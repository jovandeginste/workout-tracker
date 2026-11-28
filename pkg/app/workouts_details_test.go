package app

import (
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkoutShow_DetailsAreTabbed(t *testing.T) {
	a := configuredApp(t)

	owner := createUser(t, a, "owner", "owner", false)
	owner.Profile.ShowTabs = true

	w := createAPIWorkout(t, a, owner, "tabbed workout", database.WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))

	status, body := renderAs(t, a, owner, a.Reverse("workout-show", w.ID), a.workoutsShowHandler,
		echo.PathValue{Name: "id", Value: strconv.FormatUint(w.ID, 10)})
	require.Equal(t, http.StatusOK, status)

	assert.Contains(t, body, `id="workout-metadata-general"`)
	assert.Contains(t, body, `id="content-workout-metadata-general"`)
	assert.Contains(t, body, `id="workout-metadata-extra"`)
	assert.Contains(t, body, `id="content-workout-metadata-extra"`)

	general, extra, found := strings.Cut(body, `id="content-workout-metadata-extra"`)
	require.True(t, found)

	assert.Contains(t, general, "Total distance")
	assert.NotContains(t, general, "Max speed")

	assert.Contains(t, extra, "Max speed")
	assert.NotContains(t, extra, "Total distance")
}
