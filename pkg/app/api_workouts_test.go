package app

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/converters"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createAPIWorkout(t *testing.T, a *App, u *database.User, name string, wt database.WorkoutType, date time.Time) *database.Workout {
	t.Helper()

	w := &database.Workout{
		UserID: u.ID,
		Date:   date,
		Name:   name,
		Type:   wt,
	}
	require.NoError(t, a.db.Create(w).Error)

	md := &database.MapData{
		WorkoutID: w.ID,
		WorkoutData: converters.WorkoutData{
			TotalDistance: 1000,
			TotalDuration: time.Hour,
		},
	}
	require.NoError(t, a.db.Create(md).Error)

	return w
}

func apiWorkoutRequest(t *testing.T, ts *httptest.Server, method, path, body string) (int, []byte) {
	t.Helper()

	req, err := http.NewRequest(method, ts.URL+path, strings.NewReader(body))
	require.NoError(t, err)

	req.Header.Set("Authorization", "Bearer my-api-key")

	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	return res.StatusCode, b
}

func TestAPI_WorkoutUpdate(t *testing.T) {
	a := configuredApp(t)
	u := defaultAPIUser(a.db)
	date := time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC)
	w := createAPIWorkout(t, a, u, "Original", database.WorkoutTypeRunning, date)

	ts := httptest.NewServer(a.echo)
	defer ts.Close()

	t.Run("updates only the provided metadata", func(t *testing.T) {
		status, b := apiWorkoutRequest(t, ts, http.MethodPatch, a.Reverse("api-workout-update", w.ID), `{"name":"Renamed","type":"cycling"}`)
		require.Equal(t, http.StatusOK, status)

		var resp struct {
			Results database.Workout
			Errors  []string
		}
		require.NoError(t, json.Unmarshal(b, &resp))
		assert.Empty(t, resp.Errors)
		assert.NotContains(t, string(b), "my-api-key")
		assert.NotContains(t, string(b), "my-password")

		assert.Equal(t, "Renamed", resp.Results.Name)
		assert.Equal(t, database.WorkoutTypeCycling, resp.Results.Type)
		assert.True(t, date.Equal(resp.Results.Date))
		require.NotNil(t, resp.Results.Data)
		assert.Equal(t, 1000.0, resp.Results.Data.TotalDistance)

		var stored database.Workout
		require.NoError(t, a.db.First(&stored, w.ID).Error)
		assert.Equal(t, "Renamed", stored.Name)
		assert.Equal(t, database.WorkoutTypeCycling, stored.Type)
	})

	t.Run("rejects an invalid body", func(t *testing.T) {
		status, b := apiWorkoutRequest(t, ts, http.MethodPatch, a.Reverse("api-workout-update", w.ID), `{"name":`)

		assert.Equal(t, http.StatusBadRequest, status)
		assert.Contains(t, string(b), "errors")
	})

	t.Run("does not update another user's workout", func(t *testing.T) {
		other := defaultUser(a.db)
		other.Username = "other-user"
		other.Name = "Other"
		require.NoError(t, other.Save(a.db))

		otherWorkout := createAPIWorkout(t, a, other, "Other workout", database.WorkoutTypeWalking, date.Add(time.Hour))

		status, _ := apiWorkoutRequest(t, ts, http.MethodPatch, a.Reverse("api-workout-update", otherWorkout.ID), `{"name":"Hijacked"}`)
		assert.Equal(t, http.StatusBadRequest, status)

		var stored database.Workout
		require.NoError(t, a.db.First(&stored, otherWorkout.ID).Error)
		assert.Equal(t, "Other workout", stored.Name)
	})
}

func TestAPI_WorkoutDelete(t *testing.T) {
	a := configuredApp(t)
	u := defaultAPIUser(a.db)
	w := createAPIWorkout(t, a, u, "To delete", database.WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))

	other := defaultUser(a.db)
	other.Username = "other-user"
	other.Name = "Other"
	require.NoError(t, other.Save(a.db))

	otherWorkout := createAPIWorkout(t, a, other, "Not mine", database.WorkoutTypeWalking, time.Date(2025, 1, 1, 11, 0, 0, 0, time.UTC))

	ts := httptest.NewServer(a.echo)
	defer ts.Close()

	t.Run("does not delete another user's workout", func(t *testing.T) {
		status, _ := apiWorkoutRequest(t, ts, http.MethodDelete, a.Reverse("api-workout-delete", otherWorkout.ID), "")
		assert.Equal(t, http.StatusBadRequest, status)

		var count int64
		require.NoError(t, a.db.Model(&database.Workout{}).Where("id = ?", otherWorkout.ID).Count(&count).Error)
		assert.Equal(t, int64(1), count)
	})

	t.Run("deletes the workout", func(t *testing.T) {
		status, b := apiWorkoutRequest(t, ts, http.MethodDelete, a.Reverse("api-workout-delete", w.ID), "")
		require.Equal(t, http.StatusOK, status)

		var resp struct {
			Results database.Workout
			Errors  []string
		}
		require.NoError(t, json.Unmarshal(b, &resp))
		assert.Empty(t, resp.Errors)
		assert.NotContains(t, string(b), "my-api-key")
		assert.NotContains(t, string(b), "my-password")
		assert.Equal(t, w.ID, resp.Results.ID)

		var count int64
		require.NoError(t, a.db.Model(&database.Workout{}).Where("id = ?", w.ID).Count(&count).Error)
		assert.Equal(t, int64(0), count)

		require.NoError(t, a.db.Model(&database.MapData{}).Where("workout_id = ?", w.ID).Count(&count).Error)
		assert.Equal(t, int64(0), count)
	})

	t.Run("returns an error for an unknown workout", func(t *testing.T) {
		status, _ := apiWorkoutRequest(t, ts, http.MethodDelete, a.Reverse("api-workout-delete", 99999), "")
		assert.Equal(t, http.StatusBadRequest, status)
	})
}
