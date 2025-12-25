package app

import (
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/invopop/ctxi18n/i18n"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/pkg/geocoder"
	"github.com/jovandeginste/workout-tracker/v2/pkg/templatehelpers"
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
	Name            *string               `form:"name" json:"name"`
	Date            *string               `form:"date" json:"date"`
	Timezone        *string               `form:"timezone" json:"timezone"`
	Location        *string               `form:"location" json:"location"`
	DurationHours   *int                  `form:"duration_hours" json:"duration_hours"`
	DurationMinutes *int                  `form:"duration_minutes" json:"duration_minutes"`
	DurationSeconds *int                  `form:"duration_seconds" json:"duration_seconds"`
	Distance        *float64              `form:"distance" json:"distance"`
	Repetitions     *int                  `form:"repetitions" json:"repetitions"`
	Weight          *float64              `form:"weight" json:"weight"`
	Notes           *string               `form:"notes" json:"notes"`
	Type            *database.WorkoutType `form:"type" json:"type"`
	CustomType      *string               `form:"custom_type" json:"custom_type"`

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

	if m.Timezone == nil {
		return &d
	}

	// Handle timezone offset
	tzLoc, err := time.LoadLocation(*m.Timezone)
	if err == nil {
		d = d.In(tzLoc)
	}

	_, zoneOffset := d.Zone()
	d = d.Add(-time.Duration(zoneOffset) * time.Second)

	// handle DST transitions
	if d.IsDST() {
		d = d.Add(1 * time.Hour)
	}

	return &d
}

func (m *ManualWorkout) ToWeight() *float64 {
	if m.Weight == nil || *m.Weight == 0 {
		return nil
	}

	d := templatehelpers.WeightToDatabase(*m.Weight, m.units.Weight())

	return &d
}

func (m *ManualWorkout) ToDistance() *float64 {
	if m.Distance == nil || *m.Distance == 0 {
		return nil
	}

	d := templatehelpers.DistanceToDatabase(*m.Distance, m.units.Distance())

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
	setIfNotNil(&w.Date, dDate)
	setIfNotNil(&w.Type, m.Type)
	setIfNotNil(&w.CustomType, m.CustomType)

	setIfNotNil(&w.Data.AddressString, m.Location)
	setIfNotNil(&w.Data.TotalDistance, m.ToDistance())
	setIfNotNil(&w.Data.TotalDuration, m.ToDuration())
	setIfNotNil(&w.Data.TotalRepetitions, m.Repetitions)
	setIfNotNil(&w.Data.TotalWeight, m.ToWeight())

	if m.Location != nil && w.FullAddress() != *m.Location {
		a, err := geocoder.Find(*m.Location)
		if err != nil {
			w.Data.Address = nil
			return
		}

		w.Data.Address = a
		w.Data.AddressString = database.GetAddressString(a)
	}

	w.Data.UpdateExtraMetrics()
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
		EquipmentIDs []uint64 `form:"equipment"`
	}

	if err := c.Bind(&equipmentIDS); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-add"), err)
	}

	equipment, err := database.GetEquipmentByIDs(a.db, a.getCurrentUser(c).ID, equipmentIDS.EquipmentIDs)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-add"), err)
	}

	if err := workout.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-add"), err)
	}

	if err := a.db.Model(&workout).Association("Equipment").Replace(equipment); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-show", workout.ID), err)
	}

	a.worker.Submit(database.NewUpdateMapDataAddressTask(workout.Data.ID))

	a.addNoticeT(c, "translation.The_workout_s_has_been_created", workout.Name)

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
		EquipmentIDs []uint64 `form:"equipment"`
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

	a.worker.Submit(database.NewUpdateMapDataAddressTask(workout.Data.ID))

	a.addNoticeT(c, "translation.The_workout_s_has_been_updated", workout.Name)

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

		ws, addErr := a.getCurrentUser(c).AddWorkout(a.db, workoutType, notes, file.Filename, content)
		if len(addErr) > 0 {
			for _, e := range addErr {
				errMsg = append(errMsg, e.Error())
			}
			continue
		}

		for _, w := range ws {
			msg = append(msg, w.Name)
			a.worker.Submit(database.NewUpdateMapDataAddressTask(w.Data.ID))
		}
	}

	if len(errMsg) > 0 {
		a.addErrorN(c, "alerts.workouts_added", len(errMsg), i18n.M{"count": len(errMsg), "list": strings.Join(errMsg, "; ")})
	}

	if len(msg) > 0 {
		a.addNoticeN(c, "notices.workouts_added", len(msg), i18n.M{"count": len(msg), "list": strings.Join(msg, "; ")})
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("workouts"))
}
