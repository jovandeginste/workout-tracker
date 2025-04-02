package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/templatehelpers"
)

type BreakdownItem struct {
	FirstPoint    *MapPoint     `json:"firstPoint"`    // First GPS point in this item
	LastPoint     *MapPoint     `json:"lastPoint"`     // Last GPS point in this item
	UnitName      string        `json:"unitName"`      // Unit name
	UnitCount     float64       `json:"unitCount"`     // Count of the unit per item
	Counter       int           `json:"counter"`       // Counter of this item in the list of items
	Distance      float64       `json:"distance"`      // Distance in this item
	TotalDistance float64       `json:"totalDistance"` // Total distance in all items up to and including this item
	Duration      time.Duration `json:"duration"`      // Duration in this item
	TotalDuration time.Duration `json:"totalDuration"` // Total duration in all items up to and including this item
	Speed         float64       `json:"speed"`         // Speed in this item
	IsBest        bool          `json:"isBest"`        // Whether this item is the best of the list
	IsWorst       bool          `json:"isWorst"`       // Whether this item is the worst of the list

	LocalTotalDistance string `json:"localTotalDistance,omitempty"` // Total distance in all items up to and including this item
	LocalDistance      string `json:"localDistance,omitempty"`      // The total distance in the bucket, localized
	LocalAverageSpeed  string `json:"localAverageSpeed,omitempty"`  // The average speed in the bucket, localized
	LocalElevation     string `json:"localElevation,omitempty"`     // The starting elevation in the bucket, localized
	LocalHeartRate     string `json:"localHeartRate,omitempty"`     // The starting heart rate in the bucket, localized
	LocalCadence       string `json:"localCadence,omitempty"`       // The starting cadence in the bucket, localized

	TotalDurationSeconds float64 `json:"totalDurationSeconds,omitempty"` // The total duration in the bucket, in seconds
}

func (bi *BreakdownItem) Localize(units *UserPreferredUnits) {
	bi.LocalTotalDistance = templatehelpers.HumanDistanceFor(units.Distance())(bi.TotalDistance)
	bi.TotalDurationSeconds = bi.TotalDuration.Seconds()

	bi.LocalDistance = templatehelpers.HumanDistanceFor(units.Distance())(bi.Distance)
	bi.LocalAverageSpeed = templatehelpers.HumanSpeedFor(units.Distance())(bi.Speed)

	bi.LocalElevation = templatehelpers.HumanElevationFor(units.Elevation())(bi.FirstPoint.ExtraMetrics.Get("elevation"))
	bi.LocalHeartRate = fmt.Sprintf("%.0f", bi.FirstPoint.ExtraMetrics.Get("heart-rate"))
	bi.LocalCadence = fmt.Sprintf("%.0f", bi.FirstPoint.ExtraMetrics.Get("cadence"))
}

func (bi *BreakdownItem) createNext(fp *MapPoint) BreakdownItem {
	return BreakdownItem{
		UnitCount:     bi.UnitCount,
		UnitName:      bi.UnitName,
		Counter:       bi.Counter + 1,
		TotalDistance: bi.TotalDistance,
		TotalDuration: bi.TotalDuration,
		FirstPoint:    fp,
	}
}

func (bi *BreakdownItem) canHave(count float64, unit string, fp *MapPoint) bool {
	switch unit {
	case "distance":
		return bi.canHaveDistance(fp.Distance, float64(bi.Counter)*count)
	case "duration":
		return bi.canHaveDuration(fp.Duration, time.Duration(float64(bi.Counter)*count))
	}

	return true
}

func (bi *BreakdownItem) canHaveDistance(distance, next float64) bool {
	return bi.TotalDistance+distance < next
}

func (bi *BreakdownItem) canHaveDuration(duration, next time.Duration) bool {
	return bi.TotalDuration+duration < next
}

func (bi *BreakdownItem) CalcultateSpeed() {
	if bi.Duration.Seconds() == 0 {
		bi.Speed = 0
		return
	}

	bi.Speed = bi.Distance / bi.Duration.Seconds()
}

func calculateBestAndWorst(items []BreakdownItem) {
	if len(items) == 0 {
		return
	}

	worst := 0
	best := 0

	for i := range items {
		if items[i].Speed < items[worst].Speed {
			worst = i
		}

		if items[i].Speed > items[best].Speed {
			best = i
		}
	}

	items[worst].IsWorst = true
	items[best].IsBest = true
}

func (w *Workout) statisticsWithUnit(count float64, unit string) []BreakdownItem {
	if w.Data.Details == nil ||
		len(w.Data.Details.Points) == 0 {
		return nil
	}

	var items []BreakdownItem

	nextItem := BreakdownItem{
		UnitCount:  count,
		UnitName:   unit,
		Counter:    1,
		FirstPoint: &w.Data.Details.Points[0],
	}

	for i, p := range w.Data.Details.Points {
		if !nextItem.canHave(count, unit, &w.Data.Details.Points[i]) {
			nextItem.LastPoint = &w.Data.Details.Points[i]
			nextItem.CalcultateSpeed()
			items = append(items, nextItem)
			nextItem = nextItem.createNext(&w.Data.Details.Points[i])
		}

		nextItem.Distance += p.Distance
		nextItem.TotalDistance += p.Distance

		// m/s -> km/h, cut-off is speed less than 1 km/h
		if p.AverageSpeed()*3.6 >= 1.0 {
			nextItem.Duration += p.Duration
			nextItem.TotalDuration += p.Duration
		}
	}

	nextItem.LastPoint = &w.Data.Details.Points[len(w.Data.Details.Points)-1]

	if nextItem.FirstPoint != nil {
		nextItem.CalcultateSpeed()
		items = append(items, nextItem)
	}

	calculateBestAndWorst(items)

	return items
}

type WorkoutBreakdown struct {
	Unit  string          `json:"unit"`
	Items []BreakdownItem `json:"items"`
}

func (w *Workout) StatisticsPer(count float64, unit string) (WorkoutBreakdown, error) {
	wb := WorkoutBreakdown{Unit: unit}

	switch unit {
	case "m":
		wb.Items = w.statisticsWithUnit(count, "distance")
	case "km":
		wb.Items = w.statisticsWithUnit(count*templatehelpers.MeterPerKM, "distance")
	case "mi":
		wb.Items = w.statisticsWithUnit(count*templatehelpers.MeterPerMile, "distance")
	case "sec":
		wb.Items = w.statisticsWithUnit(count*float64(time.Second), "duration")
	case "min":
		wb.Items = w.statisticsWithUnit(count*float64(time.Minute), "duration")
	case "hour":
		wb.Items = w.statisticsWithUnit(count*float64(time.Hour), "duration")
	default:
		return wb, fmt.Errorf("unknown unit: %s", unit)
	}

	if len(wb.Items) == 0 {
		return wb, errors.New("no data")
	}

	return wb, nil
}
