package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tkrajina/gpxgo/gpx"
)

func TestMapPoint_IsMoving(t *testing.T) {
	tests := []struct {
		name     string
		distance float64
		duration time.Duration
		want     bool
	}{
		{name: "standing still", distance: 0, duration: time.Minute, want: false},
		{name: "very slow", distance: 10, duration: time.Minute, want: false},
		{name: "walking", distance: 50, duration: time.Minute, want: true},
		{name: "cycling", distance: 500, duration: time.Minute, want: true},
		{name: "no duration", distance: 0, duration: 0, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := MapPoint{Distance: tt.distance, Duration: tt.duration}
			assert.Equal(t, tt.want, p.IsMoving())
		})
	}
}

func TestGpxAsMapDataKeepsSubSecondIntervals(t *testing.T) {
	g := &gpx.GPX{}
	start := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)

	for i := range 5 {
		g.AppendPoint(&gpx.GPXPoint{
			Point:     gpx.Point{Latitude: 50 + float64(i)*0.0001, Longitude: 4},
			Timestamp: start.Add(time.Duration(i) * 1500 * time.Millisecond),
		})
	}

	data := gpxAsMapData(g, WorkoutTypeRunning, false, false)

	assert.Equal(t, 6*time.Second, data.TotalDuration)
}
