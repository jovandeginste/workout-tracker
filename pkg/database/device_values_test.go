package database

import (
	"encoding/xml"
	"os"
	"testing"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/converters"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tkrajina/gpxgo/gpx"
)

type deviceTestPoint struct {
	latitude  float64
	longitude float64
	elevation float64
	distance  *float64
	speed     *float64
}

func deviceValue(v float64) *float64 {
	return &v
}

func deviceTestGPX(t *testing.T, points []deviceTestPoint) *gpx.GPX {
	t.Helper()

	g := &gpx.GPX{Creator: "garmin"}
	start := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)

	for i, p := range points {
		point := gpx.GPXPoint{
			Point: gpx.Point{
				Latitude:  p.latitude,
				Longitude: p.longitude,
				Elevation: *gpx.NewNullableFloat64(p.elevation),
			},
			Timestamp: start.Add(time.Duration(i) * 10 * time.Second),
		}

		if p.distance != nil {
			point.Extensions.Nodes = append(point.Extensions.Nodes, gpx.ExtensionNode{
				XMLName: xml.Name{Local: "distance"},
				Data:    cast.ToString(*p.distance),
			})
		}

		if p.speed != nil {
			point.Extensions.Nodes = append(point.Extensions.Nodes, gpx.ExtensionNode{
				XMLName: xml.Name{Local: "speed"},
				Data:    cast.ToString(*p.speed),
			})
		}

		g.AppendPoint(&point)
	}

	return g
}

func TestGpxAsMapDataUsesDeviceDistance(t *testing.T) {
	g := deviceTestGPX(t, []deviceTestPoint{
		{latitude: 50.000, longitude: 4, elevation: 100, distance: deviceValue(0)},
		{latitude: 50.001, longitude: 4, elevation: 110, distance: deviceValue(50)},
		{latitude: 50.002, longitude: 4, elevation: 120, distance: deviceValue(120)},
	})

	converters.AddGPXExtension(g, converters.ExtensionTotalDistance, 120)
	converters.AddGPXExtension(g, converters.ExtensionTotalAscent, 12)
	converters.AddGPXExtension(g, converters.ExtensionTotalDescent, 3)
	converters.AddGPXExtension(g, converters.ExtensionTotalTimerTime, 15)
	converters.AddGPXExtension(g, converters.ExtensionAverageCadence, 80)
	converters.AddGPXExtension(g, converters.ExtensionMaxCadence, 95)
	converters.AddGPXExtension(g, converters.ExtensionAveragePower, 200)
	converters.AddGPXExtension(g, converters.ExtensionMaxPower, 300)

	data := gpxAsMapData(g, WorkoutTypeCycling, false, false)

	require.Len(t, data.Details.Points, 3)

	assert.InDelta(t, 50, data.Details.Points[1].Distance, 0.001)
	assert.InDelta(t, 70, data.Details.Points[2].Distance, 0.001)
	assert.InDelta(t, 120, data.Details.Points[2].TotalDistance, 0.001)

	assert.InDelta(t, 120, data.TotalDistance, 0.001)
	assert.InDelta(t, 120, data.TotalDistance2D, 0.001)
	assert.Equal(t, 20*time.Second, data.TotalDuration)

	assert.InDelta(t, 12, data.TotalUp, 0.001)
	assert.InDelta(t, 3, data.TotalDown, 0.001)

	// The device timer ran for 15 of the 20 seconds.
	assert.Equal(t, 5*time.Second, data.PauseDuration)

	assert.InDelta(t, 80, data.AverageCadence, 0.001)
	assert.InDelta(t, 95, data.MaxCadence, 0.001)
	assert.InDelta(t, 200, data.AveragePower, 0.001)
	assert.InDelta(t, 300, data.MaxPower, 0.001)
}

func TestGpxAsMapDataIgnoreDeviceValuesRecalculates(t *testing.T) {
	g := deviceTestGPX(t, []deviceTestPoint{
		{latitude: 50.000, longitude: 4, elevation: 100, distance: deviceValue(0)},
		{latitude: 50.001, longitude: 4, elevation: 110, distance: deviceValue(50)},
		{latitude: 50.002, longitude: 4, elevation: 120, distance: deviceValue(120)},
	})

	converters.AddGPXExtension(g, converters.ExtensionTotalDistance, 120)
	converters.AddGPXExtension(g, converters.ExtensionTotalAscent, 12)
	converters.AddGPXExtension(g, converters.ExtensionTotalDescent, 3)
	converters.AddGPXExtension(g, converters.ExtensionTotalTimerTime, 15)

	data := gpxAsMapData(g, WorkoutTypeCycling, false, true)

	// The track points are roughly 111 m apart, so the calculated distance is
	// larger than the device distance.
	assert.Greater(t, data.TotalDistance, 200.0)
	assert.Greater(t, data.Details.Points[1].Distance, 100.0)
	assert.InDelta(t, 20, data.TotalUp, 0.001)
	assert.Equal(t, time.Duration(0), data.PauseDuration)
}

func TestGpxAsMapDataIgnoreDeviceValuesRecalculatesMaxSpeed(t *testing.T) {
	g := deviceTestGPX(t, []deviceTestPoint{
		{latitude: 50.0000, longitude: 4, elevation: 100, speed: deviceValue(5)},
		{latitude: 50.0001, longitude: 4, elevation: 100, speed: deviceValue(5)},
		{latitude: 50.0100, longitude: 4, elevation: 100, speed: deviceValue(5)},
		{latitude: 50.0001, longitude: 4, elevation: 100, speed: deviceValue(5)},
	})

	data := gpxAsMapData(g, WorkoutTypeCycling, false, true)

	assert.Greater(t, data.MaxSpeed, 5.0)
}

func TestGpxAsMapDataDeviceDistanceFallsBackWhenMissing(t *testing.T) {
	g := deviceTestGPX(t, []deviceTestPoint{
		{latitude: 50.000, longitude: 4, elevation: 100, distance: deviceValue(0)},
		{latitude: 50.001, longitude: 4, elevation: 100, distance: deviceValue(100)},
		{latitude: 50.002, longitude: 4, elevation: 100},
		{latitude: 50.003, longitude: 4, elevation: 100, distance: deviceValue(150)},
	})

	data := gpxAsMapData(g, WorkoutTypeCycling, false, false)

	require.Len(t, data.Details.Points, 4)

	assert.InDelta(t, 100, data.Details.Points[1].Distance, 0.001)
	assert.Greater(t, data.Details.Points[2].Distance, 100.0)
	assert.Greater(t, data.Details.Points[3].Distance, 100.0)

	for _, p := range data.Details.Points {
		assert.GreaterOrEqual(t, p.Distance, 0.0)
	}
}

func TestGpxAsMapDataDeviceDistanceHandlesReset(t *testing.T) {
	g := deviceTestGPX(t, []deviceTestPoint{
		{latitude: 50.000, longitude: 4, elevation: 100, distance: deviceValue(0)},
		{latitude: 50.001, longitude: 4, elevation: 100, distance: deviceValue(100)},
		{latitude: 50.002, longitude: 4, elevation: 100, distance: deviceValue(20)},
		{latitude: 50.003, longitude: 4, elevation: 100, distance: deviceValue(120)},
	})

	data := gpxAsMapData(g, WorkoutTypeCycling, false, false)

	require.Len(t, data.Details.Points, 4)

	assert.InDelta(t, 100, data.Details.Points[1].Distance, 0.001)
	assert.Greater(t, data.Details.Points[2].Distance, 100.0)
	assert.InDelta(t, 100, data.Details.Points[3].Distance, 0.001)
}

func TestGpxAsMapDataDeviceValuesRespectIgnoreElevation(t *testing.T) {
	g := deviceTestGPX(t, []deviceTestPoint{
		{latitude: 50.000, longitude: 4, elevation: 100, distance: deviceValue(0)},
		{latitude: 50.001, longitude: 4, elevation: 120, distance: deviceValue(100)},
	})

	converters.AddGPXExtension(g, converters.ExtensionTotalAscent, 20)
	converters.AddGPXExtension(g, converters.ExtensionTotalDescent, 0)

	data := gpxAsMapData(g, WorkoutTypeCycling, true, false)

	assert.InDelta(t, 0, data.TotalUp, 0.001)
	assert.InDelta(t, 0, data.TotalDown, 0.001)
}

func TestGpxAsMapDataMaxSpeedIgnoresGpsTeleport(t *testing.T) {
	g := deviceTestGPX(t, []deviceTestPoint{
		{latitude: 50.0000, longitude: 4, elevation: 100, speed: deviceValue(5)},
		{latitude: 50.0001, longitude: 4, elevation: 100, speed: deviceValue(5)},
		{latitude: 50.0100, longitude: 4, elevation: 100, speed: deviceValue(5)},
		{latitude: 50.0001, longitude: 4, elevation: 100, speed: deviceValue(5)},
	})

	data := gpxAsMapData(g, WorkoutTypeCycling, false, false)

	// The geometry contains a jump of more than a kilometre between two
	// consecutive points, but the device reported 5 m/s.
	assert.InDelta(t, 5, data.MaxSpeed, 0.001)
}

func TestGpxAsMapDataMaxSpeedFallsBackToGeometry(t *testing.T) {
	g := deviceTestGPX(t, []deviceTestPoint{
		{latitude: 50.0000, longitude: 4, elevation: 100},
		{latitude: 50.0010, longitude: 4, elevation: 100},
		{latitude: 50.0020, longitude: 4, elevation: 100},
		{latitude: 50.0030, longitude: 4, elevation: 100},
	})

	data := gpxAsMapData(g, WorkoutTypeCycling, false, false)

	assert.Greater(t, data.MaxSpeed, 0.0)
}

func TestGpxAsMapDataUsesDeviceValuesFromFitSample(t *testing.T) {
	content, err := os.ReadFile("testdata/issues/issue-0658-15718428901-activity.fit")
	require.NoError(t, err)

	g, err := converters.ParseFit(content)
	require.NoError(t, err)

	data := gpxAsMapData(g, WorkoutTypeCycling, false, false)

	assert.InDelta(t, 52185.44, data.TotalDistance, 0.01)
	assert.InDelta(t, 52185.44, data.TotalDistance2D, 0.01)
	assert.InDelta(t, 262, data.TotalUp, 0.001)
	assert.InDelta(t, 242, data.TotalDown, 0.001)

	// The device timer ran for 15771.708 seconds, which leaves the rest of the
	// elapsed time as pause.
	elapsed := data.TotalDuration - data.PauseDuration
	assert.InDelta(t, 15771.708, elapsed.Seconds(), 1)
}
