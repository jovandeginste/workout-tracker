package database

import (
	"fmt"
	"time"
)

type StatisticsItem struct {
	UnitCount     float64
	UnitName      string
	Counter       int
	Distance      float64
	TotalDistance float64
	Duration      time.Duration
	TotalDuration time.Duration
	Speed         float64
	FirstPoint    *MapPoint
	LastPoint     *MapPoint
	IsBest        bool
	IsWorst       bool
}

func (si *StatisticsItem) SpeedKPH() float64 {
	return 3.6 * si.Speed
}

func (si *StatisticsItem) createNext(fp *MapPoint) StatisticsItem {
	return StatisticsItem{
		UnitCount:     si.UnitCount,
		UnitName:      si.UnitName,
		Counter:       si.Counter + 1,
		TotalDistance: si.TotalDistance,
		TotalDuration: si.TotalDuration,
		FirstPoint:    fp,
	}
}

func (si *StatisticsItem) canHave(count float64, unit string, fp *MapPoint) bool {
	switch unit {
	case "distance":
		return si.canHaveDistance(fp.Distance, float64(si.Counter)*count)
	case "duration":
		return si.canHaveDuration(fp.Duration, time.Duration(float64(si.Counter)*count))
	}

	return true
}

func (si *StatisticsItem) canHaveDistance(distance, next float64) bool {
	return si.TotalDistance+distance < next
}

func (si *StatisticsItem) canHaveDuration(duration, next time.Duration) bool {
	return si.TotalDuration+duration < next
}

func (si *StatisticsItem) CalcultateSpeed() {
	si.Speed = si.Distance / si.Duration.Seconds()
}

func calculateBestAndWorst(items []StatisticsItem) {
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

func (w *Workout) statisticsWithUnit(count float64, unit string) []StatisticsItem {
	if len(w.Data.Details.Points) == 0 {
		return nil
	}

	var items []StatisticsItem

	nextItem := StatisticsItem{
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
	Unit  string
	Items []StatisticsItem
}

func (w *Workout) StatisticsPer(count float64, unit string) (WorkoutBreakdown, error) {
	wb := WorkoutBreakdown{Unit: unit}

	switch unit {
	case "m":
		wb.Items = w.statisticsWithUnit(count, "distance")
	case "km":
		wb.Items = w.statisticsWithUnit(count*1000, "distance")
	case "mi":
		wb.Items = w.statisticsWithUnit(count*1609.344, "distance")
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
		return wb, fmt.Errorf("no data")
	}

	return wb, nil
}
