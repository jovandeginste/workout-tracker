package database

import (
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/templatehelpers"
)

type (
	// Statistics represents the statistics for a user for a given time range and bucket size, per workout type
	Statistics struct {
		Buckets      map[WorkoutType]Buckets // The statistics buckets
		BucketFormat string                  // The bucket format in strftime format
		UserID       uint                    // The user ID
	}

	Buckets struct {
		WorkoutType      WorkoutType
		LocalWorkoutType string
		Buckets          map[string]Bucket
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

		LocalDistance            string `json:",omitempty"` // The total distance in the bucket, localized
		LocalUp                  string `json:",omitempty"` // The total up elevation in the bucket, localized
		LocalDuration            string `json:",omitempty"` // The total duration in the bucket, localized
		LocalAverageSpeed        string `json:",omitempty"` // The average speed in the bucket, localized
		LocalAverageSpeedNoPause string `json:",omitempty"` // The average speed without pause in the bucket, localized
		LocalMaxSpeed            string `json:",omitempty"` // The max speed in the bucket, localized

		DurationSeconds float64 `json:",omitempty"` // The total duration in the bucket, in seconds
	}

	// Float64Record is a single record if the value is a float64
	Float64Record struct {
		Date  time.Time // The timestamp of the record
		Value float64   // The value of the record
		ID    uint      // The workout ID of the record
	}

	// DurationRecord is a single record if the value is a time.Duration
	DurationRecord struct {
		Date  time.Time     // The timestamp of the record
		Value time.Duration // The value of the record
		ID    uint          // The workout ID of the record
	}

	// WorkoutRecord is the collection of records for a single workout type
	WorkoutRecord struct {
		WorkoutType         WorkoutType    // The type of the workout
		AverageSpeed        Float64Record  // The record with the maximum average speed
		AverageSpeedNoPause Float64Record  // The record with the maximum average speed without pause
		MaxSpeed            Float64Record  // The record with the maximum max speed
		Distance            Float64Record  // The record with the maximum distance
		TotalUp             Float64Record  // The record with the maximum up elevation
		Duration            DurationRecord // The record with the maximum duration
		Active              bool           // Whether there is any data in the record
	}
)

func (b *Bucket) Localize(units *UserPreferredUnits) {
	b.LocalDistance = templatehelpers.HumanDistanceFor(units.Distance())(b.Distance)
	b.LocalUp = templatehelpers.HumanElevationFor(units.Elevation())(b.Up)
	b.LocalDuration = templatehelpers.HumanDuration(b.Duration)
	b.LocalAverageSpeed = templatehelpers.HumanSpeedFor(units.Distance())(b.AverageSpeed)
	b.LocalAverageSpeedNoPause = templatehelpers.HumanSpeedFor(units.Distance())(b.AverageSpeedNoPause)
	b.LocalMaxSpeed = templatehelpers.HumanSpeedFor(units.Distance())(b.MaxSpeed)
	b.DurationSeconds = b.Duration.Seconds()
}
