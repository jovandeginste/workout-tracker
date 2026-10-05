package database

import (
	"math"

	"github.com/paulmach/orb"
	"gorm.io/gorm"
)

// heatmapPointSpacing is the minimum distance (in meters) between two
// consecutive heatmap points of a workout. At heatmap resolution, denser
// tracks add nothing visually but cost memory and bandwidth.
const heatmapPointSpacing = 50.0

// heatmapPrecision rounds coordinates to 5 decimals (~1 meter).
const heatmapPrecision = 1e5

// HeatmapPoint is a [lat, lng] pair, stored compactly as a JSON array.
type HeatmapPoint [2]float64

func (h HeatmapPoint) ToOrbPoint() *orb.Point {
	return &orb.Point{h[1], h[0]}
}

func newHeatmapPoint(p *MapPoint) HeatmapPoint {
	return HeatmapPoint{
		math.Round(p.Lat*heatmapPrecision) / heatmapPrecision,
		math.Round(p.Lng*heatmapPrecision) / heatmapPrecision,
	}
}

// heatmapPointsFor thins out the points of a track, keeping one point every
// heatmapPointSpacing meters, plus the last point. The result is never nil,
// so it is stored as "[]" instead of NULL, which marks it as calculated.
func heatmapPointsFor(points []MapPoint) []HeatmapPoint {
	result := []HeatmapPoint{}

	var (
		sinceLast  float64
		last       *MapPoint
		lastIsKept bool
	)

	for i := range points {
		p := &points[i]
		sinceLast += p.Distance2D

		if p.IsZero() {
			continue
		}

		lastIsKept = last == nil || sinceLast >= heatmapPointSpacing
		if lastIsKept {
			result = append(result, newHeatmapPoint(p))
			sinceLast = 0
		}

		last = p
	}

	if last != nil && !lastIsKept {
		result = append(result, newHeatmapPoint(last))
	}

	return result
}

// UpdateHeatmapPoints recalculates the heatmap points from the track points.
func (d *MapDataDetails) UpdateHeatmapPoints() {
	d.HeatmapPoints = heatmapPointsFor(d.Points)
}

// GetHeatmapPoints returns the heatmap points of all workouts of the user.
// Only the (small) heatmap_points column is read; the full track points are
// never loaded, keeping memory usage independent of the history size.
func (u *User) GetHeatmapPoints(db *gorm.DB) ([]HeatmapPoint, error) {
	var details []*MapDataDetails

	err := db.Model(&MapDataDetails{}).
		Select("map_data_details.id", "map_data_details.heatmap_points").
		Joins("JOIN map_data ON map_data.id = map_data_details.map_data_id").
		Joins("JOIN workouts ON workouts.id = map_data.workout_id").
		Where("workouts.user_id = ?", u.ID).
		Where("map_data_details.heatmap_points IS NOT NULL").
		Find(&details).Error
	if err != nil {
		return nil, err
	}

	var result []HeatmapPoint
	for _, d := range details {
		result = append(result, d.HeatmapPoints...)
	}

	return result, nil
}

// GetMapDataDetailsWithoutHeatmap returns the IDs of the map data details for
// which no heatmap points were calculated yet.
func GetMapDataDetailsWithoutHeatmap(db *gorm.DB) ([]uint64, error) {
	var ids []uint64

	err := db.Model(&MapDataDetails{}).
		Where("heatmap_points IS NULL").
		Pluck("id", &ids).Error

	return ids, err
}

// UpdateHeatmapPointsFor calculates and stores the heatmap points of a single
// map data details record, writing only the heatmap_points column.
func UpdateHeatmapPointsFor(db *gorm.DB, id uint64) error {
	var d MapDataDetails

	if err := db.First(&d, id).Error; err != nil {
		return err
	}

	d.UpdateHeatmapPoints()

	return db.Model(&d).Select("heatmap_points").Updates(&d).Error
}
