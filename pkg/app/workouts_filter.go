package app

import (
	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type WorkoutFilters struct {
	Type     database.WorkoutType `query:"type"`
	Since    string               `query:"since"`
	OrderBy  string               `query:"order_by"`
	OrderDir string               `query:"order_dir"`

	db *gorm.DB
}

func getWorkoutsFilters(c echo.Context) (*WorkoutFilters, error) {
	filters := WorkoutFilters{}

	if err := c.Bind(&filters); err != nil {
		return nil, err
	}

	return &filters, nil
}

func (wf *WorkoutFilters) ToQuery(db *gorm.DB) *gorm.DB {
	wf.db = db

	wf.setTypeFilter()
	wf.setSinceFilter()
	wf.setOrderFilter()

	return wf.db
}

func (wf *WorkoutFilters) setTypeFilter() {
	if wf.Type == "" {
		return
	}

	wf.db = wf.db.Where(&database.Workout{Type: wf.Type})
}

func (wf *WorkoutFilters) setSinceFilter() {
	if wf.Since == "" || wf.Since == "all" {
		return
	}

	sqlDialect := wf.db.Name()
	wf.db = wf.db.Where(database.GetDateLimitExpression(sqlDialect), "-"+wf.Since)
}

func (wf *WorkoutFilters) setOrderFilter() {
	if wf.OrderBy == "" {
		return
	}

	wf.db = wf.db.Select("workouts.*").Joins("left join map_data on workouts.id = map_data.workout_id")

	dir := wf.OrderDir
	if dir == "" {
		dir = "asc"
	}

	switch wf.OrderBy {
	case "date":
		wf.db = wf.db.Order("workouts." + wf.OrderBy + " " + dir)
	case "total_distance", "total_duration", "total_weight", "total_repetitions", "total_up", "total_down",
		"average_speed_no_pause", "max_speed":
		wf.db = wf.db.Order("map_data." + wf.OrderBy + " " + dir)
	}
}
