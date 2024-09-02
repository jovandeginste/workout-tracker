package app

import (
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/jovandeginste/workout-tracker/internal/database"
	"github.com/jovandeginste/workout-tracker/internal/pkg/geocoder"
	"github.com/labstack/echo/v4"
)

const (
	htmlDateFormat = "2006-01-02T15:04"
)

func uploadedFile(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Read all from r into a bytes slice
	content, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}

	return content, nil
}

type ManualWorkout struct {
	Name            *string               `form:"name"`
	Date            *string               `form:"date"`
	Location        *string               `form:"location"`
	DurationHours   *int                  `form:"duration_hours"`
	DurationMinutes *int                  `form:"duration_minutes"`
	DurationSeconds *int                  `form:"duration_seconds"`
	Distance        *float64              `form:"distance"`
	Repetitions     *int                  `form:"repetitions"`
	Weight          *float64              `form:"weight"`
	Notes           *string               `form:"notes"`
	Type            *database.WorkoutType `form:"type"`

	units *database.UserPreferredUnits
}

func (m *ManualWorkout) ToDate() *time.Time {
	if m.Date == nil {
		return nil
	}

	d, err := time.Parse(htmlDateFormat, *m.Date)
	if err != nil {
		return nil
	}

	return &d
}

func (m *ManualWorkout) ToDistance() *float64 {
	if m.Distance == nil || *m.Distance == 0 {
		return nil
	}

	d := m.units.DistanceToDatabase(*m.Distance)

	return &d
}

func (m *ManualWorkout) ToDuration() *time.Duration {
	var totalDuration time.Duration

	if m.DurationHours != nil {
		totalDuration += time.Duration(*m.DurationHours) * time.Hour
	}

	if m.DurationMinutes != nil {
		totalDuration += time.Duration(*m.DurationMinutes) * time.Minute
	}

	if m.DurationSeconds != nil {
		totalDuration += time.Duration(*m.DurationSeconds) * time.Second
	}

	if totalDuration == 0 {
		return nil
	}

	return &totalDuration
}

func setIfNotNil[T any](dst *T, src *T) {
	if src == nil {
		return
	}

	*dst = *src
}

func (m *ManualWorkout) Update(w *database.Workout) {
	if w.Data == nil {
		w.Data = &database.MapData{}
	}

	dDate := m.ToDate()

	setIfNotNil(&w.Name, m.Name)
	setIfNotNil(&w.Notes, m.Notes)
	setIfNotNil(&w.Date, &dDate)
	setIfNotNil(&w.Type, m.Type)

	setIfNotNil(&w.Data.AddressString, m.Location)
	setIfNotNil(&w.Data.TotalDistance, m.ToDistance())
	setIfNotNil(&w.Data.TotalDuration, m.ToDuration())
	setIfNotNil(&w.Data.TotalRepetitions, m.Repetitions)
	setIfNotNil(&w.Data.TotalWeight, m.Weight)

	a, err := geocoder.Find(*m.Location)
	if err != nil {
		w.Data.Address = nil
		return
	}

	setIfNotNil(&w.Data.Address, &a)

	w.Data.UpdateAddress()
}

func (a *App) addWorkout(c echo.Context) error {
	if strings.HasPrefix(c.Request().Header.Get(echo.HeaderContentType), echo.MIMEMultipartForm) {
		return a.addWorkoutFromFile(c)
	}

	d := &ManualWorkout{units: a.getCurrentUser(c).PreferredUnits()}
	if err := c.Bind(d); err != nil {
		return a.redirectWithError(c, "/workouts", err)
	}

	workout := &database.Workout{}
	d.Update(workout)

	workout.User = a.getCurrentUser(c)
	workout.UserID = a.getCurrentUser(c).ID
	workout.Data.Creator = "web-interface"

	var equipmentIDS struct {
		EquipmentIDs []uint `form:"equipment"`
	}

	if err := c.Bind(&equipmentIDS); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-edit", c.Param("id")), err)
	}

	equipment, err := database.GetEquipmentByIDs(a.db, a.getCurrentUser(c).ID, equipmentIDS.EquipmentIDs)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-edit", c.Param("id")), err)
	}

	if err := workout.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-edit", c.Param("id")), err)
	}

	if err := a.db.Model(&workout).Association("Equipment").Replace(equipment); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-show", c.Param("id")), err)
	}

	a.setNotice(c, "The workout '%s' has been created.", workout.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("workouts"))
}

func (a *App) workoutsUpdateHandler(c echo.Context) error {
	workout, err := a.getWorkout(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-show", c.Param("id")), err)
	}

	d := &ManualWorkout{units: a.getCurrentUser(c).PreferredUnits()}
	if err := c.Bind(d); err != nil {
		return a.redirectWithError(c, "/workouts", err)
	}

	d.Update(workout)

	var equipmentIDS struct {
		EquipmentIDs []uint `form:"equipment"`
	}

	if err := c.Bind(&equipmentIDS); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-edit", c.Param("id")), err)
	}

	equipment, err := database.GetEquipmentByIDs(a.db, a.getCurrentUser(c).ID, equipmentIDS.EquipmentIDs)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-edit", c.Param("id")), err)
	}

	if err := workout.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-edit", c.Param("id")), err)
	}

	if err := a.db.Model(&workout).Association("Equipment").Replace(equipment); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-show", c.Param("id")), err)
	}

	a.setNotice(c, "The workout '%s' has been updated.", workout.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("workout-show", c.Param("id")))
}

func (a *App) addWorkoutFromFile(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["file"]

	msg := []string{}
	errMsg := []string{}

	for _, file := range files {
		content, parseErr := uploadedFile(file)
		if parseErr != nil {
			errMsg = append(errMsg, parseErr.Error())
			continue
		}

		notes := c.FormValue("notes")
		workoutType := database.WorkoutType(c.FormValue("type"))

		w, addErr := a.getCurrentUser(c).AddWorkout(a.db, workoutType, notes, file.Filename, content)
		if addErr != nil {
			errMsg = append(errMsg, addErr.Error())
			continue
		}

		msg = append(msg, w.Name)
	}

	if len(errMsg) > 0 {
		a.setError(c, "Encountered %d problems while adding workouts: %s", len(errMsg), strings.Join(errMsg, "; "))
	}

	if len(msg) > 0 {
		a.setNotice(c, "Added %d new workout(s): %s", len(msg), strings.Join(msg, "; "))
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("workouts"))
}
