package database

import (
	"time"
)

type StatisticsItem struct {
	Kilometer     int
	Distance      float64
	TotalDistance float64
	Duration      time.Duration
	Speed         float64
	FirstPoint    *MapPoint
	LastPoint     *MapPoint
	IsBest        bool
	IsWorst       bool
}

func (si *StatisticsItem) CanHave(distance float64) bool {
	return int(si.TotalDistance+distance) < si.Kilometer*1000
}

func (si *StatisticsItem) CalcultateSpeed() {
	si.Speed = si.Distance / si.Duration.Seconds()
}

func (w *Workout) StatisticsPerKilometer() []StatisticsItem {
	var items []StatisticsItem

	nextItem := StatisticsItem{
		Kilometer:  1,
		FirstPoint: &w.Data.Points[0],
	}

	for i, p := range w.Data.Points {
		if !nextItem.CanHave(p.Distance) {
			nextItem.LastPoint = &w.Data.Points[i]
			nextItem.CalcultateSpeed()
			items = append(items, nextItem)
			nextItem = StatisticsItem{
				Kilometer:     nextItem.Kilometer + 1,
				TotalDistance: nextItem.TotalDistance,
				FirstPoint:    &w.Data.Points[i],
			}
		}

		nextItem.Distance += p.Distance
		nextItem.TotalDistance += p.Distance

		// m/s -> km/h, cut-off is speed less than 1 km/h
		if p.AverageSpeed()*3.6 >= 1.0 {
			nextItem.Duration += p.Duration
		}
	}

	nextItem.LastPoint = &w.Data.Points[len(w.Data.Points)-1]

	if nextItem.FirstPoint != nil {
		nextItem.CalcultateSpeed()
		items = append(items, nextItem)
	}

	calculateBestAndWorst(items)

	return items
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
