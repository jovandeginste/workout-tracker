package database

import (
	"time"
)

type (
	Statistics struct {
		UserID  uint
		Buckets map[WorkoutType]map[string]Bucket
	}

	Bucket struct {
		Bucket              string `json:",omitempty"`
		WorkoutType         WorkoutType
		Workouts            int
		Distance            float64       `json:",omitempty"`
		Up                  float64       `json:",omitempty"`
		Duration            time.Duration `json:",omitempty"`
		AverageSpeed        float64       `json:",omitempty"`
		AverageSpeedNoPause float64       `json:",omitempty"`
		MaxSpeed            float64       `json:",omitempty"`
	}

	float64Record struct {
		Value float64
		Date  time.Time
		ID    uint
	}

	durationRecord struct {
		Value time.Duration
		Date  time.Time
		ID    uint
	}

	WorkoutRecord struct {
		WorkoutType         WorkoutType
		Active              bool
		AverageSpeed        float64Record
		AverageSpeedNoPause float64Record
		MaxSpeed            float64Record
		Distance            float64Record
		TotalUp             float64Record
		Duration            durationRecord
	}
)

func (b Bucket) DistanceKM() float64 {
	return b.Distance / 1000
}

func (b Bucket) AverageSpeedNoPauseKPH() float64 {
	return 3.6 * b.AverageSpeedNoPause
}

func (b Bucket) AverageSpeedKPH() float64 {
	return 3.6 * b.AverageSpeed
}

func (b Bucket) MaxSpeedKPH() float64 {
	return 3.6 * b.MaxSpeed
}
