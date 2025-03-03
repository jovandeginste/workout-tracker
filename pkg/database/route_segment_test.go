package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouteSegment_Parse(t *testing.T) {
	{
		rs, err := NewRouteSegment("", "meer.gpx", []byte(meer))
		assert.NoError(t, err)
		assert.NotNil(t, rs)
		assert.Greater(t, rs.TotalDistance, 1800.0)
	}

	{
		rs, err := NewRouteSegment("", "finsepiste.gpx", []byte(finsepiste))
		assert.NoError(t, err)
		assert.NotNil(t, rs)
		assert.Greater(t, rs.TotalDistance, 900.0)
	}
}

func TestRouteSegment_FindMatches(t *testing.T) {
	rs, err := NewRouteSegment("", "finsepiste.gpx", []byte(finsepiste))
	assert.NoError(t, err)

	w1, err := NewWorkout(AnonymousUser(), WorkoutTypeAutoDetect, "", "match.gpx", []byte(track))
	assert.NoError(t, err)
	assert.Len(t, w1, 1)

	w1_1 := w1[0]
	assert.True(t, w1_1.Type.IsLocation())
	assert.True(t, w1_1.HasTracks())

	w2, err := NewWorkout(AnonymousUser(), WorkoutTypeAutoDetect, "", "nomatch.gpx", []byte(GpxSample1))
	assert.NoError(t, err)
	assert.Len(t, w2, 1)

	w2_1 := w2[0]
	assert.True(t, w2_1.Type.IsLocation())
	assert.True(t, w2_1.HasTracks())

	workouts := []*Workout{w1_1, w2_1}
	matches := rs.FindMatches(workouts)

	if !assert.Len(t, matches, 1) {
		return
	}

	assert.Len(t, matches[0].Workout.Data.Details.Points, 158)
}

func TestRouteSegment_StartingPoints_NoMatch(t *testing.T) {
	rs, err := NewRouteSegment("", "finsepiste.gpx", []byte(finsepiste))
	assert.NoError(t, err)

	w, err := NewWorkout(AnonymousUser(), WorkoutTypeAutoDetect, "", "nomatch.gpx", []byte(GpxSample1))
	assert.NoError(t, err)
	assert.Len(t, w, 1)

	w1 := w[0]
	sp := rs.StartingPoints(w1.Data.Details.Points)
	assert.Empty(t, sp)
}

func TestRouteSegment_StartingPoints_Match(t *testing.T) {
	rs, err := NewRouteSegment("", "finsepiste.gpx", []byte(finsepiste))
	assert.NoError(t, err)

	w, err := NewWorkout(AnonymousUser(), WorkoutTypeAutoDetect, "", "match.gpx", []byte(track))
	assert.NoError(t, err)
	assert.Len(t, w, 1)

	w1 := w[0]
	sp := rs.StartingPoints(w1.Data.Details.Points)
	assert.NotEmpty(t, sp)

	for _, p := range sp {
		assert.Less(t, rs.Points[0].DistanceTo(&w1.Data.Details.Points[p]), MaxDeltaMeter)
	}
}

func TestRouteSegment_StartingPoints_MatchSegment(t *testing.T) {
	rs, err := NewRouteSegment("", "finsepiste.gpx", []byte(finsepiste))
	assert.NoError(t, err)

	w, err := NewWorkout(AnonymousUser(), WorkoutTypeAutoDetect, "", "match.gpx", []byte(track))
	assert.NoError(t, err)
	assert.Len(t, w, 1)

	w1 := w[0]
	sp := rs.StartingPoints(w1.Data.Details.Points)
	assert.NotEmpty(t, sp)

	{
		last, ok := rs.MatchSegment(w1, 3, true)
		assert.Zero(t, last)
		assert.False(t, ok)
	}

	{
		last, ok := rs.MatchSegment(w1, 4, true)
		assert.NotZero(t, last)
		assert.True(t, ok)
	}
}

func TestRouteSegment_Match(t *testing.T) {
	rs, err := NewRouteSegment("", "finsepiste.gpx", []byte(finsepiste))
	assert.NoError(t, err)

	w, err := NewWorkout(AnonymousUser(), WorkoutTypeAutoDetect, "", "match.gpx", []byte(track))
	assert.NoError(t, err)
	assert.Len(t, w, 1)

	w1 := w[0]
	rsm := rs.Match(w1)
	if !assert.NotNil(t, rsm) {
		return
	}

	assert.Greater(t, rsm.Distance, 900.0)
	assert.True(t, rsm.MatchesDistance(rs.TotalDistance))
}
