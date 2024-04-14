package database

import (
	"time"
)

type (
	// Statistics represents the statistics for a user for a given time range and bucket size, per workout type
	Statistics struct {
		UserID       uint                              // The user ID
		BucketFormat string                            // The bucket format in strftime format
		Buckets      map[WorkoutType]map[string]Bucket // The statistics buckets
	}

	// Bucket is the consolidation of workout information for a given time bucket
	Bucket struct {
		Bucket              string        `json:",omitempty"` // The name of the bucket
		WorkoutType         WorkoutType   // The type of the workout
		Workouts            int           // The number of workouts in the bucket
		Distance            float64       `json:",omitempty"` // The total distance in the bucket
		Up                  float64       `json:",omitempty"` // The total up elevation in the bucket
		Duration            time.Duration `json:",omitempty"` // The total duration in the bucket
		AverageSpeed        float64       `json:",omitempty"` // The average speed in the bucket
		AverageSpeedNoPause float64       `json:",omitempty"` // The average speed without pause in the bucket
		MaxSpeed            float64       `json:",omitempty"` // The max speed in the bucket
	}

	// float64Record is a single record if the value is a float64
	float64Record struct {
		Value float64   // The value of the record
		Date  time.Time // The timestamp of the record
		ID    uint      // The workout ID of the record
	}

	// durationRecord is a single record if the value is a time.Duration
	durationRecord struct {
		Value time.Duration // The value of the record
		Date  time.Time     // The timestamp of the record
		ID    uint          // The workout ID of the record
	}

	// WorkoutRecord is the collection of records for a single workout type
	WorkoutRecord struct {
		WorkoutType         WorkoutType    // The type of the workout
		Active              bool           // Whether there is any data in the record
		AverageSpeed        float64Record  // The record with the maximum average speed
		AverageSpeedNoPause float64Record  // The record with the maximum average speed without pause
		MaxSpeed            float64Record  // The record with the maximum max speed
		Distance            float64Record  // The record with the maximum distance
		TotalUp             float64Record  // The record with the maximum up elevation
		Duration            durationRecord // The record with the maximum duration
	}
)
