package workouts

import (
	"testing"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestSegmentToMap(t *testing.T) {
	segment := database.Segment{
		Type:     database.SlopeKindLift,
		Index:    2,
		StartIdx: 10,
		EndIdx:   20,
		Start:    database.MapPoint{Lat: 1.5, Lng: 2.5},
		End:      database.MapPoint{Lat: 3.5, Lng: 4.5},
		Gain:     100,
		Length:   500,
		AvgSlope: 0.2,
		Duration: 90 * time.Second,
	}

	assert.Equal(t, mapSegment{
		Kind:     "lift",
		Index:    2,
		StartIdx: 10,
		EndIdx:   20,
		Start:    []float64{1.5, 2.5},
		End:      []float64{3.5, 4.5},
		Gain:     100,
		Length:   500,
		AvgSlope: 0.2,
		Duration: 90,
	}, segmentToMap(segment))
}
