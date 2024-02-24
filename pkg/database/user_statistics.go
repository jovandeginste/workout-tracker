package database

import (
	"time"

	"gorm.io/gorm"
)

type (
	record struct {
		Value float64
		Date  time.Time
		ID    uint
	}

	Totals struct {
		Workouts            int
		Distance            float64
		Up                  float64
		Duration            time.Duration
		AverageSpeed        float64
		AverageSpeedNoPause float64
		MaxSpeed            float64
	}
	UserStatistics struct {
		Total    Totals
		PerYear  map[int]*Totals
		PerMonth map[int]map[int]*Totals
		Records  struct {
			Active              bool
			AverageSpeed        record
			AverageSpeedNoPause record
			MaxSpeed            record
			Distance            record
			TotalUp             record
			Duration            struct {
				Value time.Duration
				Date  time.Time
				ID    uint
			}
		}
	}
)

func (r *record) CheckAndSwap(value float64, id uint, date *time.Time) {
	if r.Value < value {
		r.Value = value
		r.Date = *date
		r.ID = id
	}
}

func (us *UserStatistics) Add(w *Workout) {
	us.Total.Workouts++
	us.Total.Distance += w.Data.TotalDistance
	us.Total.Duration += w.Data.TotalDuration
	us.Total.Up += w.Data.TotalUp
	us.Total.AverageSpeed += w.Data.AverageSpeed()
	us.Total.AverageSpeedNoPause += w.Data.AverageSpeedNoPause()
	us.Total.MaxSpeed += w.Data.MaxSpeed

	d := w.Date
	year := d.Year()
	month := int(d.Month())

	us.AddYear(year, w)
	us.AddMonth(year, month, w)
}

func (us *UserStatistics) AddMonth(year, month int, w *Workout) {
	if _, ok := us.PerMonth[year]; !ok {
		us.PerMonth[year] = map[int]*Totals{}
	}

	entry, ok := us.PerMonth[year][month]
	if !ok {
		us.PerMonth[year][month] = &Totals{
			Workouts:            1,
			Distance:            w.Data.TotalDistance,
			Up:                  w.Data.TotalUp,
			Duration:            w.Data.TotalDuration,
			AverageSpeed:        w.Data.AverageSpeed(),
			AverageSpeedNoPause: w.Data.AverageSpeedNoPause(),
			MaxSpeed:            w.Data.MaxSpeed,
		}

		return
	}

	entry.Workouts++
	entry.Distance += w.Data.TotalDistance
	entry.Duration += w.Data.TotalDuration
	entry.Up += w.Data.TotalUp
	entry.AverageSpeed += w.Data.AverageSpeed()
	entry.AverageSpeedNoPause += w.Data.AverageSpeedNoPause()
	entry.MaxSpeed += w.Data.MaxSpeed
}

func (us *UserStatistics) AddYear(year int, w *Workout) {
	entry, ok := us.PerYear[year]
	if !ok {
		us.PerYear[year] = &Totals{
			Workouts:            1,
			Distance:            w.Data.TotalDistance,
			Up:                  w.Data.TotalUp,
			Duration:            w.Data.TotalDuration,
			AverageSpeed:        w.Data.AverageSpeed(),
			AverageSpeedNoPause: w.Data.AverageSpeedNoPause(),
			MaxSpeed:            w.Data.MaxSpeed,
		}

		return
	}

	entry.Workouts++
	entry.Distance += w.Data.TotalDistance
	entry.Duration += w.Data.TotalDuration
	entry.Up += w.Data.TotalUp
	entry.AverageSpeed += w.Data.AverageSpeed()
	entry.AverageSpeedNoPause += w.Data.AverageSpeedNoPause()
	entry.MaxSpeed += w.Data.MaxSpeed
}

func (u *User) Statistics(db *gorm.DB) (*UserStatistics, error) {
	us := &UserStatistics{}

	us.PerYear = map[int]*Totals{}
	us.PerMonth = map[int]map[int]*Totals{}

	workouts, err := u.GetWorkouts(db)
	if err != nil {
		return nil, err
	}

	for _, w := range workouts {
		if !w.Type.IsDistance() {
			continue
		}

		us.Records.Active = true
		us.Add(w)

		us.Records.AverageSpeedNoPause.CheckAndSwap(
			w.Data.AverageSpeedNoPause(),
			w.ID,
			w.Date,
		)

		us.Records.AverageSpeed.CheckAndSwap(
			w.Data.AverageSpeed(),
			w.ID,
			w.Date,
		)

		us.Records.MaxSpeed.CheckAndSwap(
			w.Data.MaxSpeed,
			w.ID,
			w.Date,
		)

		us.Records.Distance.CheckAndSwap(
			w.Data.TotalDistance,
			w.ID,
			w.Date,
		)

		us.Records.TotalUp.CheckAndSwap(
			w.Data.TotalUp,
			w.ID,
			w.Date,
		)

		if w.Data.TotalDuration > us.Records.Duration.Value {
			us.Records.Duration.Value = w.Data.TotalDuration
			us.Records.Duration.ID = w.ID
			us.Records.Duration.Date = *w.Date
		}
	}

	return us, nil
}
