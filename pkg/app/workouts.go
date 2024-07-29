package app

import (
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/labstack/echo/v4"
)

const (
	htmlDateFormat     = "2006-01-02T15:04"
	htmlDurationFormat = "15:04"
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
	Name        string               `form:"name"`
	Date        string               `form:"date"`
	Duration    string               `form:"duration"`
	Repetitions int                  `form:"repetitions"`
	Weight      float64              `form:"weight"`
	Notes       string               `form:"notes"`
	Type        database.WorkoutType `form:"type"`
}

func (m *ManualWorkout) ToDate() time.Time {
	d, err := time.Parse(htmlDateFormat, m.Date)
	if err != nil {
		return time.Time{}
	}

	return d
}

func (m *ManualWorkout) ToDuration() time.Duration {
	d, err := time.Parse(htmlDurationFormat, m.Duration)
	if err != nil {
		return 0
	}

	return time.Duration(d.Hour())*time.Hour + time.Duration(d.Minute())*time.Minute
}

func (a *App) addWorkout(c echo.Context) error {
	if strings.HasPrefix(c.Request().Header.Get(echo.HeaderContentType), echo.MIMEMultipartForm) {
		return a.addWorkoutFromFile(c)
	}

	d := &ManualWorkout{}

	if err := c.Bind(d); err != nil {
		return a.redirectWithError(c, "/workouts", err)
	}

	dDate := d.ToDate()
	dDuration := d.ToDuration()

	w := database.Workout{
		Name:   d.Name,
		Notes:  d.Notes,
		Date:   &dDate,
		Type:   d.Type,
		User:   a.getCurrentUser(c),
		UserID: a.getCurrentUser(c).ID,
		Data: &database.MapData{
			TotalDuration:    dDuration,
			TotalRepetitions: d.Repetitions,
			TotalWeight:      d.Weight,
			Creator:          "web-interface",
		},
	}

	if err := w.Save(a.db); err != nil {
		return a.redirectWithError(c, "/workouts", err)
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("workouts"))
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
