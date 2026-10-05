package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// straightTrack returns n points, each 10m apart.
func straightTrack(n int) []MapPoint {
	points := make([]MapPoint, n)
	for i := range n {
		points[i] = MapPoint{Lat: 50 + float64(i)*0.0001, Lng: 4, Distance2D: 10}
	}

	points[0].Distance2D = 0

	return points
}

func TestHeatmapPointsFor_Empty(t *testing.T) {
	result := heatmapPointsFor(nil)

	assert.NotNil(t, result, "empty result must not be nil, so it is not stored as NULL")
	assert.Empty(t, result)
}

func TestHeatmapPointsFor_KeepsOnePointPerSpacing(t *testing.T) {
	// 1001 points, 10m apart = 10km
	points := straightTrack(1001)

	result := heatmapPointsFor(points)

	// One point every 50m, including first and last
	assert.Len(t, result, 201)
	assert.Equal(t, newHeatmapPoint(&points[0]), result[0])
	assert.Equal(t, newHeatmapPoint(&points[1000]), result[len(result)-1])
}

func TestHeatmapPointsFor_KeepsLastPoint(t *testing.T) {
	points := straightTrack(8) // 70m: keeps 0m, 50m and 70m

	result := heatmapPointsFor(points)

	assert.Len(t, result, 3)
	assert.Equal(t, newHeatmapPoint(&points[7]), result[2])
}

func TestHeatmapPointsFor_SkipsPointsWithoutLocation(t *testing.T) {
	points := straightTrack(3)
	points[1].Lat, points[1].Lng = 0, 0
	points = append(points, MapPoint{Distance2D: 100})

	result := heatmapPointsFor(points)

	for _, p := range result {
		assert.NotEqual(t, HeatmapPoint{0, 0}, p)
	}

	assert.Len(t, result, 2)
}

func TestHeatmapPoint_ToOrbPoint(t *testing.T) {
	p := HeatmapPoint{50.1, 4.2}.ToOrbPoint()

	assert.InDelta(t, 4.2, p.Lon(), 0)
	assert.InDelta(t, 50.1, p.Lat(), 0)
}

func TestUser_GetHeatmapPoints(t *testing.T) {
	db := createMemoryDB(t)
	u := defaultUser()
	require.NoError(t, u.Create(db))

	f1, err := gpxFS.ReadFile("sample1.gpx")
	require.NoError(t, err)

	ws, err := NewWorkout(u, WorkoutTypeAutoDetect, "", "file.gpx", f1)
	require.NoError(t, err)
	require.Len(t, ws, 1)
	require.NoError(t, ws[0].Save(db))

	expected := ws[0].Data.Details.HeatmapPoints
	require.NotEmpty(t, expected)
	assert.Less(t, len(expected), len(ws[0].Data.Details.Points))

	points, err := u.GetHeatmapPoints(db)
	require.NoError(t, err)
	assert.Equal(t, expected, points)

	other := &User{}
	other.ID = u.ID + 1
	points, err = other.GetHeatmapPoints(db)
	require.NoError(t, err)
	assert.Empty(t, points)

	// Simulate a workout processed before heatmap points existed
	detailsID := ws[0].Data.Details.ID
	require.NoError(t, db.Model(&MapDataDetails{}).Where("id = ?", detailsID).
		UpdateColumn("heatmap_points", nil).Error)

	ids, err := GetMapDataDetailsWithoutHeatmap(db)
	require.NoError(t, err)
	assert.Equal(t, []uint64{detailsID}, ids)

	points, err = u.GetHeatmapPoints(db)
	require.NoError(t, err)
	assert.Empty(t, points)

	require.NoError(t, UpdateHeatmapPointsFor(db, detailsID))

	ids, err = GetMapDataDetailsWithoutHeatmap(db)
	require.NoError(t, err)
	assert.Empty(t, ids)

	points, err = u.GetHeatmapPoints(db)
	require.NoError(t, err)
	assert.Equal(t, expected, points)
}
