package user

import "github.com/jovandeginste/workout-tracker/v2/pkg/database"

// StatisticsData holds the selected filters and preloaded data for the
// statistics page.
type StatisticsData struct {
	Since              string
	Per                string
	WorkoutTypes       []string
	SelectedEquipment  []uint64
	ChartType          string
	RouteSegments      []*database.RouteSegment
	TypeCounts         map[database.WorkoutType]int
	Equipment          []*database.Equipment
	SelectedSegment    uint64
	IncludePause       bool
	DistancePacePoints []*database.DistancePacePoint
}
