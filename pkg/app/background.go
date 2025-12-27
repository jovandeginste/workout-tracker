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
	FileAddDelay                 = -1 * time.Minute
	workerRouteSegmentsBatchSize = 10
	workerWorkoutsBatchSize      = 10
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
		a.updateWorkouts(l)
		a.updateRouteSegments(l)
		a.autoImports(l)
	}

	if a.workerPoolGeo.WaitingTasks() > 0 {
		l.With("size", a.workerPoolGeo.WaitingTasks()).Warn("Waiting for current updater to finish")
	} else {
		a.updateAddresses(l)
	}

	l.Info("Worker finished...")
}

func (a *App) autoImports(l *slog.Logger) {
	var uID []uint64

	q := a.db.Model(&database.User{}).Pluck("ID", &uID)
	if err := q.Error; err != nil {
		l.Error(ErrWorker.Error() + ": " + err.Error())
	}

	for _, i := range uID {
		a.workerPool.Submit(func() {
			if err := a.autoImportForUser(l, i); err != nil {
				l.Error(ErrWorker.Error() + ": " + err.Error())
			}
		})
	}
}

func (a *App) autoImportForUser(l *slog.Logger, userID uint64) error {
	l = l.With("user_id", userID)

	u, err := database.GetUserByID(a.db, userID)
	if err != nil {
		return err
	}

	userLogger := l.With("user", u.Username)

	ok, err := u.Profile.CanImportFromDirectory()
	if err != nil {
		return fmt.Errorf("could not use auto-import dir %v for user %v: %w", u.Profile.AutoImportDirectory, u.Username, err)
	}

	if !ok {
		return nil
	}

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

	logger.Info("Moving to '" + destDir + "'")

	// If destDir does not exist, create it
	if _, err := os.Stat(destDir); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(destDir, 0o755); err != nil {
			return err
		}
	}

	if err := os.Rename(path, filepath.Join(destDir, filepath.Base(path))); err != nil {
		return err
	}

	logger.Info("Moved to '" + destDir + "'")

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

// For the given set of route segments, re-match against all workouts, marking the segments as clean after matching.
func (a *App) rematchRouteSegmentsToWorkouts(routeSegments []*database.RouteSegment, l *slog.Logger) error {
	if len(routeSegments) == 0 {
		l.Debug("rematchRouteSegmentsToWorkouts: no segments provided")
		return nil
	}

	// Reset matches for each segment
	for _, rs := range routeSegments {
		rs.RouteSegmentMatches = []*database.RouteSegmentMatch{}
	}

	var workoutsBatch []*database.Workout
	qw := a.db.Preload("Data.Details").Preload("User").Model(&database.Workout{}).FindInBatches(&workoutsBatch, workerRouteSegmentsBatchSize, func(wtx *gorm.DB, batchNo int) error {
		l.With("batch_no", batchNo).
			With("route_segments_batch_size", len(routeSegments)).
			With("workouts_batch_size", len(workoutsBatch)).
			Debug("rematchRouteSegmentsToWorkouts start")

		// Match this batch of workouts against the current batch of route segments
		for _, rs := range routeSegments {
			newMatches := rs.FindMatches(workoutsBatch)
			rs.RouteSegmentMatches = append(rs.RouteSegmentMatches, newMatches...)
			l.With("route_segment_id", rs.ID).
				With("new_matches", len(newMatches)).
				With("total_matches", len(rs.RouteSegmentMatches)).
				Debug("Updating route segments")
		}

		l.With("batch_no", batchNo).
			With("route_segments_batch_size", len(routeSegments)).
			With("workouts_batch_size", len(workoutsBatch)).
			Debug("rematchRouteSegmentsToWorkouts done")

		return nil
	})

	if qw.Error != nil {
		l.Error("Worker error fetching workouts: " + qw.Error.Error())
	}

	// Mark all route segments as non-dirty and save matches
	var errs error

	for _, rs := range routeSegments {
		rs.Dirty = false
		if err := rs.Save(a.db); err != nil {
			errs = errors.Join(errs, err)
			l.Error("Worker error saving route segment: " + err.Error())
		}
	}

	return errs
}

func (a *App) updateRouteSegments(l *slog.Logger) {
	var routeSegmentsBatch []*database.RouteSegment

	// Fetch next batch of dirty segments
	q := a.db.Preload("RouteSegmentMatches").Model(&database.RouteSegment{}).Where(&database.RouteSegment{Dirty: true}).Limit(workerRouteSegmentsBatchSize).Find(&routeSegmentsBatch)
	l.With("route_segments_batch_size", len(routeSegmentsBatch)).
		Info("updateRouteSegments batch")

	if err := q.Error; err != nil {
		l.Error("Worker error", "error", err)
	}

	err := a.rematchRouteSegmentsToWorkouts(routeSegmentsBatch, l)
	if err != nil {
		l.Error("Worker error during matching", "error", err)
	}
}

func (a *App) updateWorkouts(l *slog.Logger) {
	var wID []uint64

	db := a.db.Preload("Data.Details").Preload("User")

	q := db.Model(&database.Workout{}).Where(&database.Workout{Dirty: true}).Limit(workerWorkoutsBatchSize).Pluck("ID", &wID)
	if err := q.Error; err != nil {
		l.Error("Worker error", "error", err)
	}

	for _, i := range wID {
		l.With("workout_id", i).Info("Updating workout")

		a.workerPool.Submit(func() {
			if err := a.updateWorkout(i); err != nil {
				l.Error("Worker error", "error", err)
			}
		})
	}
}

func (a *App) updateWorkout(i uint64) error {
	w, err := database.GetWorkoutDetails(a.db, i)
	if err != nil {
		return err
	}

	return w.UpdateData(a.db)
}

func (a *App) updateAddresses(l *slog.Logger) {
	var mID []uint64

	q := a.db.Model(&database.MapData{}).
		Where("center IS NOT NULL").Where("address_string", "").
		Limit(workerWorkoutsBatchSize).Pluck("ID", &mID)
	if err := q.Error; err != nil {
		l.Error("Worker error", "error", err)
	}

	for _, i := range mID {
		a.workerPoolGeo.Submit(func() {
			wl := l.With("map_data_id", i)
			wl.Info("Updating workout address")

			if err := a.updateAddress(i); err != nil {
				wl.Error("Worker error", "error", err)
			}

			wl.Info("Workout address updated")
		})
	}
}

func (a *App) updateAddress(id uint64) error {
	var m database.MapData

	if err := a.db.First(&m, id).Error; err != nil {
		return err
	}

	m.UpdateAddress()

	return m.Save(a.db)
}
