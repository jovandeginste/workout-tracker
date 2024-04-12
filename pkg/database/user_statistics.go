package database

import (
	"time"
)

type (
	Statistics struct {
		UserID       uint
		BucketFormat string
		Buckets      map[WorkoutType]map[string]Bucket
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
