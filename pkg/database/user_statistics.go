package database

import (
	"time"

	"gorm.io/gorm"
)

type UserStatistics struct {
	Total struct {
		Workouts int
		Distance float64
		Duration time.Duration
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
	}

	return us
}
