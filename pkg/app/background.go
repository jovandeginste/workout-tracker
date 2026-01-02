package app

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime/debug"
	"slices"
	"strings"
	"time"

	"github.com/alitto/pond/v2"
	"github.com/jovandeginste/workout-tracker/v2/pkg/converters"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"gorm.io/gorm"
)

var (
	ErrWorker          = errors.New("worker error")
	ErrNothingImported = errors.New("nothing imported")
)

const (
	FileAddDelay            = -1 * time.Minute
	workerWorkoutsBatchSize = 10
)

func (a *App) BackgroundWorker() {
	a.workerPool = pond.NewPool(10)
	a.workerPoolGeo = pond.NewPool(1)

	a.Logger().Info("Background worker loop initialized", "delay_seconds", a.Config.WorkerDelaySeconds)
	for {
		a.bgLoop()
		time.Sleep(time.Duration(a.Config.WorkerDelaySeconds) * time.Second)
	}
}

func (a *App) bgLoop() {
	l := a.logger.With("module", "worker")

	defer func() {
		if r := recover(); r != nil {
			l.Error(fmt.Sprintf("Panic in bgLoop: %#v", r))
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	l.Info("Worker started...")

	if a.workerPool.WaitingTasks() > 0 {
		l.With("size", a.workerPool.WaitingTasks()).Warn("Waiting for current updater to finish")
	} else {
		a.updateWorkouts(l.With("update", "workouts"))
		a.updateRouteSegments(l.With("update", "route_segments"))
		a.autoImports(l.With("update", "imports"))
	}

	if a.workerPoolGeo.WaitingTasks() > 0 {
		l.With("size", a.workerPoolGeo.WaitingTasks()).Warn("Waiting for current updater to finish")
	} else {
		a.updateAddresses(l.With("update", "addresses"))
	}

	l.Info("Worker finished...")
}

func (a *App) autoImports(l *slog.Logger) {
	var uID []uint64

	q := a.db.Model(&database.User{}).Pluck("ID", &uID)
	if err := q.Error; err != nil {
		l.Error(ErrWorker.Error() + ": " + err.Error())
	}

	for i := range uID {
		a.workerPool.Go(func() {
			if err := a.autoImportForUser(l, uID[i]); err != nil {
				l.Error(ErrWorker.Error() + ": " + err.Error())
			}
		})
	}
}

func (a *App) autoImportForUser(l *slog.Logger, userID uint64) error {
	u, err := database.GetUserByID(a.db, userID)
	if err != nil {
		return err
	}

	ok, err := u.Profile.CanImportFromDirectory()
	if err != nil {
		return fmt.Errorf("could not use auto-import dir %v for user %v: %w", u.Profile.AutoImportDirectory, u.Username, err)
	}

	if !ok {
		return nil
	}

	userLogger := l.With("user_id", userID).With("user", u.Username)
	userLogger.Info("Importing from '" + u.Profile.AutoImportDirectory + "'")

	// parse all files in the directory, non-recusive
	files, err := filepath.Glob(filepath.Join(u.Profile.AutoImportDirectory, "*"))
	if err != nil {
		return err
	}

	for _, path := range files {
		pathLogger := userLogger.With("path", path)

		if err := a.importForUser(pathLogger, u, path); err != nil {
			pathLogger.Error("Could not import: " + err.Error())
		}
	}

	return nil
}

func (a *App) importForUser(logger *slog.Logger, u *database.User, path string) error {
	// Get file info for path
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !fileCanBeImported(path, info) {
		return nil
	}

	if importErr := a.importFile(logger, u, path); importErr != nil {
		logger.Error("Could not import: " + importErr.Error())
		return moveImportFile(logger, u.Profile.AutoImportDirectory, path, "failed")
	}

	return moveImportFile(logger, u.Profile.AutoImportDirectory, path, "done")
}

func moveImportFile(logger *slog.Logger, dir, path, statusDir string) error {
	destDir := filepath.Join(dir, statusDir)

	destPath := filepath.Join(destDir, filepath.Base(path))
	logger.Info("Moving file", "src", path, "dst", destPath)

	// If destDir does not exist, create it
	if _, err := os.Stat(destDir); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(destDir, 0o755); err != nil {
			return err
		}
	}

	if err := os.Rename(path, destPath); err != nil {
		return err
	}

	logger.Info("Files moved", "destination", destDir)

	return nil
}

func (a *App) importFile(logger *slog.Logger, u *database.User, path string) error {
	logger.Info("Importing path")

	dat, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	w, addErr := u.AddWorkout(a.db, database.WorkoutTypeAutoDetect, "", path, dat)
	if len(addErr) > 0 {
		return addErr[0]
	}

	if w == nil {
		return ErrNothingImported
	}

	logger.Info("Finished import.")

	return nil
}

func fileCanBeImported(p string, i os.FileInfo) bool {
	if i.IsDir() {
		return false
	}

	// If file was changed within the last minute, don't import it
	if i.ModTime().After(time.Now().Add(FileAddDelay)) {
		return false
	}

	return slices.Contains(converters.SupportedFileTypes, strings.ToLower(filepath.Ext(p)))
}

// For the route segment, re-match against all workouts, marking the segments as clean after matching.
func (a *App) rematchRouteSegmentToWorkouts(rs *database.RouteSegment, l *slog.Logger) error {
	rs.RouteSegmentMatches = []*database.RouteSegmentMatch{}

	var workoutsBatch []*database.Workout
	qw := a.db.Preload("Data.Details").Preload("User").Model(&database.Workout{}).
		FindInBatches(&workoutsBatch, workerWorkoutsBatchSize, func(wtx *gorm.DB, batchNo int) error {
			l.With("batch_no", batchNo).
				With("workouts_batch_size", len(workoutsBatch)).
				Debug("rematchRouteSegmentsToWorkouts start")

				// Match this batch of workouts against the current route segment
			newMatches := rs.FindMatches(workoutsBatch)
			rs.RouteSegmentMatches = append(rs.RouteSegmentMatches, newMatches...)
			l.With("route_segment_id", rs.ID).
				With("new_matches", len(newMatches)).
				With("total_matches", len(rs.RouteSegmentMatches)).
				Debug("Updating route segments")

			l.With("batch_no", batchNo).
				With("workouts_batch_size", len(workoutsBatch)).
				Debug("rematchRouteSegmentsToWorkouts done")

			return nil
		})

	if qw.Error != nil {
		return qw.Error
	}

	// Mark route segment as non-dirty and save matches
	rs.Dirty = false

	return rs.Save(a.db)
}

func (a *App) updateRouteSegments(l *slog.Logger) {
	var rss []*database.RouteSegment

	r := a.db.Preload("RouteSegmentMatches").Model(&database.RouteSegment{}).
		Where(&database.RouteSegment{Dirty: true}).
		FindInBatches(&rss, workerWorkoutsBatchSize, func(rtx *gorm.DB, batchNo int) error {
			for i := range rss {
				rs := rss[i]

				a.workerPool.Go(func() {
					rl := l.With("route_segment_id", rs.ID)
					rl.Info("Updating route segment")

					if err := a.rematchRouteSegmentToWorkouts(rs, rl); err != nil {
						rl.Error("Error during matching", "error", err)
					}
				})
			}

			return nil
		})

	if r.Error != nil {
		l.Error("Error during batch query", "error", r.Error)
	}
}

func (a *App) updateWorkouts(l *slog.Logger) {
	var ws []*database.Workout

	r := a.db.Preload("GPX").Preload("Data.Details").Preload("User").Model(&database.Workout{}).
		Where(&database.Workout{Dirty: true}).
		FindInBatches(&ws, workerWorkoutsBatchSize, func(wtx *gorm.DB, batchNo int) error {
			for i := range ws {
				w := ws[i]

				a.workerPool.Go(func() {
					wl := l.With("workout_id", w.ID)
					wl.Info("Updating workout")

					if err := w.UpdateData(a.db); err != nil {
						wl.Error("Error during data update", "error", err)
					}
				})
			}

			return nil
		})

	if r.Error != nil {
		l.Error("Error during batch query", "error", r.Error)
	}
}

func (a *App) updateAddresses(l *slog.Logger) {
	var mds []*database.MapData

	r := a.db.Model(&database.MapData{}).
		Where("center IS NOT NULL").Where("address_string", "").
		FindInBatches(&mds, workerWorkoutsBatchSize, func(wtx *gorm.DB, batchNo int) error {
			for i := range mds {
				md := mds[i]

				a.workerPoolGeo.Go(func() {
					wl := l.With("workout_id", md.WorkoutID).With("map_data_id", md.ID)
					wl.Info("Updating address")

					md.UpdateAddress()
					if err := md.Save(a.db); err != nil {
						wl.Error("Error during address update", "error", err)
					}
				})
			}

			return nil
		})

	if r.Error != nil {
		l.Error("Error during batch query", "error", r.Error)
	}
}
