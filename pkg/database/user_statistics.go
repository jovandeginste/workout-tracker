package database

import (
	"time"

	"gorm.io/gorm"
)

type record struct {
	Value float64
	Date  time.Time
	ID    uint
}

type UserStatistics struct {
	Total struct {
		Workouts int
		Distance float64
		Up       float64
		Duration time.Duration
	}
	Records struct {
		AverageSpeed record
		MaxSpeed     record
		Distance     record
		TotalUp      record
		Duration     struct {
			Value time.Duration
			Date  time.Time
			ID    uint
		}
	}
}

func (r *record) CheckAndSwap(value float64, id uint, date *time.Time) {
	if r.Value < value {
		r.Value = value
		r.Date = *date
		r.ID = id
	}
}

func (u *User) Statistics(db *gorm.DB) *UserStatistics {
	us := &UserStatistics{}

	workouts, err := u.GetWorkouts(db)
	if err != nil {
		return us
	}

	us.Total.Workouts = len(workouts)

	for _, w := range workouts {
		us.Total.Distance += w.Data.TotalDistance
		us.Total.Duration += w.Data.TotalDuration
		us.Total.Up += w.Data.TotalUp

		us.Records.AverageSpeed.CheckAndSwap(
			w.Data.AverageSpeed,
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

	return us
}
