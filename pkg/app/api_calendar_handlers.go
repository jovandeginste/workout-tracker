package app

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/jovandeginste/workout-tracker/v2/views/workouts"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const calTS = "2006-01-02T15:04:05"

type Event struct {
	Title string    `json:"title"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	URL   string    `json:"url"`
}

type calendarQueryParams struct {
	Start    *string `query:"start"`
	End      *string `query:"end"`
	TimeZone *string `query:"timeZone"` // eg Europe/Brussels
	tz       *time.Location
}

func (cqp *calendarQueryParams) SetTZ() {
	cqp.tz = time.UTC
	if cqp.TimeZone == nil {
		return
	}

	tz, err := time.LoadLocation(*cqp.TimeZone)
	if err != nil {
		return
	}

	cqp.tz = tz
}

func (cqp calendarQueryParams) SetStart(db *gorm.DB) *gorm.DB {
	if cqp.Start == nil {
		return db
	}

	start, err := time.ParseInLocation(calTS, *cqp.Start, cqp.tz)
	if err != nil {
		return db
	}

	return db.Where("workouts.date >= ?", start)
}

func (cqp calendarQueryParams) SetEnd(db *gorm.DB) *gorm.DB {
	if cqp.End == nil {
		return db
	}

	end, err := time.ParseInLocation(calTS, *cqp.End, cqp.tz)
	if err != nil {
		return db
	}

	return db.Where("workouts.date <= ?", end)
}

// apiCalendar returns the calendar events of all workouts of the current user
// @Summary      List the calendar events of all workouts of the current user
// @Param        start query    string false "Start date of the calendar view"
// @Param        end query    string false "End date of the calendar view"
// @Produce      json
// @Success      200  {object}  APIResponse{results=[]Event}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /workouts/coordinates [get]
func (a *App) apiCalendar(c echo.Context) error {
	resp := APIResponse{}
	events := []Event{}
	queryParams := calendarQueryParams{}

	if err := c.Bind(&queryParams); err != nil {
		return a.renderAPIError(c, resp, err)
	}

	u := a.getCurrentUser(c)
	db := a.db.Preload("Data").Preload("Data.Details")

	queryParams.SetTZ()
	db = queryParams.SetStart(db)
	db = queryParams.SetEnd(db)

	wos, err := u.GetWorkouts(db)
	if err != nil {
		return a.renderAPIError(c, resp, err)
	}

	for _, w := range wos {
		buf := templ.GetBuffer()
		defer templ.ReleaseBuffer(buf)

		t := workouts.EventTitle(w, u.PreferredUnits())
		if err := t.Render(c.Request().Context(), buf); err != nil {
			return a.renderAPIError(c, resp, err)
		}

		d := buf.String()
		// Remove all newlines and surrounding whitespace
		d = htmlConcatenizer.ReplaceAllString(d, "")

		events = append(events, Event{
			Title: d,
			Start: w.GetDate().In(queryParams.tz),
			End:   w.GetEnd().In(queryParams.tz),
			URL:   a.echo.Reverse("workout-show", w.ID),
		})
	}

	resp.Results = events

	return c.JSON(http.StatusOK, resp)
}
