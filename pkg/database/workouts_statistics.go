package database

import (
	"time"

	"github.com/bcicen/go-units"
	"github.com/jovandeginste/workout-tracker/pkg/templatehelpers"
)

type StatisticsItem struct {
	Unit          string
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

func (si *StatisticsItem) HumanSpeed(u templatehelpers.Units) float64 {
	switch u { //nolint:exhaustive
	case templatehelpers.ImperialUnits:
		return si.SpeedMPH()
	default:
		return si.SpeedKPH()
	}
}

func (si *StatisticsItem) SpeedMPH() float64 {
	value := units.NewValue(si.SpeedKPH(), units.KiloMeter)

	f := value.MustConvert(units.Mile)

	return f.Float()
}

func (si *StatisticsItem) SpeedKPH() float64 {
	return 3.6 * si.Speed
}

func (si *StatisticsItem) createNext(fp *MapPoint) StatisticsItem {
	return StatisticsItem{
		Unit:          si.Unit,
		Counter:       si.Counter + 1,
		TotalDistance: si.TotalDistance,
		TotalDuration: si.TotalDuration,
		FirstPoint:    fp,
	}
}

func (si *StatisticsItem) canHave(unit string, fp *MapPoint) bool {
	switch unit {
	case "km":
		return si.canHaveDistance(fp.Distance)
	case "min":
		return si.canHaveDuration(fp.Duration)
	}

	return true
}

func (si *StatisticsItem) canHaveDistance(distance float64) bool {
	return int(si.TotalDistance+distance) < si.Counter*1000
}

func (si *StatisticsItem) canHaveDuration(duration time.Duration) bool {
	return si.Duration+duration < time.Minute
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

func (w *Workout) StatisticsPerKilometer() []StatisticsItem {
	return w.statisticsWithUnit("km")
}

func (w *Workout) StatisticsPerMinute() []StatisticsItem {
	return w.statisticsWithUnit("min")
}

func (w *Workout) statisticsWithUnit(unit string) []StatisticsItem {
	if len(w.Data.Details.Points) == 0 {
		return nil
	}

	var items []StatisticsItem

	nextItem := StatisticsItem{
		Unit:       unit,
		Counter:    1,
		FirstPoint: &w.Data.Details.Points[0],
	}

	for i, p := range w.Data.Details.Points {
		if !nextItem.canHave(unit, &w.Data.Details.Points[i]) {
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
