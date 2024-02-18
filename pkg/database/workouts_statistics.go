package database

import (
	"time"
)

type StatisticsItem struct {
	Kilometer int
	Distance  float64
	Duration  time.Duration
	Speed     float64
	Point     *MapPoint
	IsBest    bool
	IsWorst   bool
}

func (si *StatisticsItem) CanHave(distance float64) bool {
	return si.Distance+distance < 1000
}

func (si *StatisticsItem) CalcultateSpeed() {
	si.Speed = si.Distance / si.Duration.Seconds()
}

func (w *Workout) StatisticsPerKilometer() []StatisticsItem {
	var items []StatisticsItem

	nextItem := StatisticsItem{Kilometer: 1}

	for i, p := range w.Data.Points {
		if nextItem.Point == nil {
			nextItem.Point = &w.Data.Points[i]
		}

		if !nextItem.CanHave(p.Distance) {
			nextItem.CalcultateSpeed()
			items = append(items, nextItem)
			nextItem = StatisticsItem{
				Kilometer: nextItem.Kilometer + 1,
			}

			continue
		}

		nextItem.Distance += p.Distance
		nextItem.Duration += p.Duration
	}

	if nextItem.Point != nil {
		nextItem.CalcultateSpeed()
		items = append(items, nextItem)
	}

	calculateBestAndWorst(items)

	return items
}

func calculateBestAndWorst(items []StatisticsItem) {
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
