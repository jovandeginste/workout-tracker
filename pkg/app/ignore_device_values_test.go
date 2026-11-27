package app

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIgnoreDeviceValuesCanBeSetThroughTheWebForm(t *testing.T) {
	a := configuredApp(t)

	owner := createUser(t, a, "owner", "owner", false)
	w := createAPIWorkout(t, a, owner, "workout", database.WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))

	update := func(ignore bool) {
		form := url.Values{
			"name": {"workout"},
			"type": {database.WorkoutTypeRunning.String()},
		}
		if ignore {
			form.Set("ignore_device_values", "true")
		}

		req := httptest.NewRequest(http.MethodPost, a.Reverse("workout-update", w.ID), strings.NewReader(form.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()

		c := a.echo.NewContext(req, rec)
		a.setContext(c)
		a.setCurrentUser(c, owner)
		c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), "user_info", owner)))
		c.SetPathValues(echo.PathValues{{Name: "id", Value: strconv.FormatUint(w.ID, 10)}})

		s := sessionLoadAndSave(a.sessionManager)
		require.NoError(t, s(a.workoutsUpdateHandler)(c))

		var stored database.Workout
		require.NoError(t, a.db.First(&stored, w.ID).Error)
		assert.Equal(t, ignore, stored.IgnoreDeviceValues)
		assert.True(t, stored.Dirty)
	}

	update(true)
	update(false)
}
