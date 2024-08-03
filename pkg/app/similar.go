package app

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/labstack/echo/v4"
)

type workoutSimilarityReport struct {
	res []WorkoutSimilarity
	l   sync.Mutex
	wg  sync.WaitGroup
	ch  chan int
}

func (w *workoutSimilarityReport) add(s WorkoutSimilarity) {
	w.l.Lock()
	w.res = append(w.res, s)
	w.l.Unlock()
}

func (w *workoutSimilarityReport) inc() {
	w.wg.Add(1)
	w.ch <- 1
}

func (w *workoutSimilarityReport) dec() {
	<-w.ch
	w.wg.Done()
}

type WorkoutSimilarity struct {
	WorkoutID  uint    `json:"workout_id"` // The ID of the other workout
	Similarity float64 `json:"similarity"` // The similarity score, as a fraction between 0 and 1, where 1 is identical
}

// apiWorkoutSimilarHandler returns a list of all other workouts with a similarity score
// @Summary      Break down a workdown per units
// @Param        id      path       int     true  "Workout ID"
// @Produce      json
// @Success      200  {object}  APIResponse{result=[]WorkoutSimilarity}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /workouts/{id}/similar [get]
func (a *App) apiWorkoutSimilarHandler(c echo.Context) error {
	resp := APIResponse{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return a.renderAPIError(c, resp, err)
	}

	w, err := a.getCurrentUser(c).GetWorkout(a.db, id)
	if err != nil {
		resp.Errors = append(resp.Errors, err.Error())
	}

	workouts, err := a.getCurrentUser(c).GetWorkouts(a.db.Preload("Data.Details"))
	if err != nil {
		resp.Errors = append(resp.Errors, err.Error())
	}

	rep := workoutSimilarityReport{
		ch: make(chan int, 4),
	}

	// Calculate similarity for each other workouts:
	for _, j := range workouts {
		a.calculateSimilarity(w, j, &rep)
	}

	rep.wg.Wait()

	resp.Results = rep.res

	return c.JSON(http.StatusOK, resp)
}

func (a *App) calculateSimilarity(w, j *database.Workout, rep *workoutSimilarityReport) {
	if w.ID == j.ID {
		return
	}

	if !j.HasDetails() {
		return
	}

	a.logger.Debug(fmt.Sprintf("Comparing %d and %d", w.ID, j.ID))

	rep.inc()

	go func() {
		defer rep.dec()

		s := w.Data.Details.SimilarityTo(j.Data.Details)

		rep.add(WorkoutSimilarity{WorkoutID: j.ID, Similarity: s})
	}()
}
