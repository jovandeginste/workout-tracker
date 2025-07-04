package database

import (
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/templatehelpers"
)

type (
	// Statistics represents the statistics for a user for a given time range and bucket size, per workout type
	Statistics struct {
		Buckets      map[WorkoutType]Buckets `json:"buckets"`      // The statistics buckets
		BucketFormat string                  `json:"bucketFormat"` // The bucket format in strftime format
		UserID       uint64                  `json:"userID"`       // The user ID
	}

	Buckets struct {
		WorkoutType      WorkoutType       `json:"workoutType"`
		LocalWorkoutType string            `json:"localWorkoutType"`
		Buckets          map[string]Bucket `json:"buckets"`
	}

	// Bucket is the consolidation of workout information for a given time bucket
	Bucket struct {
		Bucket              string        `json:"bucket,omitempty"`              // The name of the bucket
		RawBucket           string        `json:"raw_bucket,omitempty"`          // One day in the bucket (for statistic rendering)
		WorkoutType         WorkoutType   `json:"workoutType"`                   // The type of the workout
		Workouts            int           `json:"workouts"`                      // The number of workouts in the bucket
		Distance            float64       `json:"distance,omitempty"`            // The total distance in the bucket
		Up                  float64       `json:"up,omitempty"`                  // The total up elevation in the bucket
		Duration            time.Duration `json:"duration,omitempty"`            // The total duration in the bucket
		AverageSpeed        float64       `json:"averageSpeed,omitempty"`        // The average speed in the bucket
		AverageSpeedNoPause float64       `json:"averageSpeedNoPause,omitempty"` // The average speed without pause in the bucket
		MaxSpeed            float64       `json:"maxSpeed,omitempty"`            // The max speed in the bucket

		LocalDistance            string `json:"localDistance,omitempty"`            // The total distance in the bucket, localized
		LocalUp                  string `json:"localUp,omitempty"`                  // The total up elevation in the bucket, localized
		LocalDuration            string `json:"localDuration,omitempty"`            // The total duration in the bucket, localized
		LocalAverageSpeed        string `json:"localAverageSpeed,omitempty"`        // The average speed in the bucket, localized
		LocalAverageSpeedNoPause string `json:"localAverageSpeedNoPause,omitempty"` // The average speed without pause in the bucket, localized
		LocalMaxSpeed            string `json:"localMaxSpeed,omitempty"`            // The max speed in the bucket, localized

		DurationSeconds float64 `json:"durationSeconds,omitempty"` // The total duration in the bucket, in seconds
	}

	// Float64Record is a single record if the value is a float64
	Float64Record struct {
		Date  time.Time `json:"date"`  // The timestamp of the record
		Value float64   `json:"value"` // The value of the record
		ID    uint64    `json:"id"`    // The workout ID of the record
	}

	// DurationRecord is a single record if the value is a time.Duration
	DurationRecord struct {
		Date  time.Time     `json:"date"`  // The timestamp of the record
		Value time.Duration `json:"value"` // The value of the record
		ID    uint64        `json:"id"`    // The workout ID of the record
	}

	// WorkoutRecord is the collection of records for a single workout type
	WorkoutRecord struct {
		WorkoutType         WorkoutType    `json:"workoutType"`         // The type of the workout
		AverageSpeed        Float64Record  `json:"averageSpeed"`        // The record with the maximum average speed
		AverageSpeedNoPause Float64Record  `json:"averageSpeedNoPause"` // The record with the maximum average speed without pause
		MaxSpeed            Float64Record  `json:"maxSpeed"`            // The record with the maximum max speed
		Distance            Float64Record  `json:"distance"`            // The record with the maximum distance
		TotalUp             Float64Record  `json:"totalUp"`             // The record with the maximum up elevation
		Duration            DurationRecord `json:"duration"`            // The record with the maximum duration
		Active              bool           `json:"active"`              // Whether there is any data in the record
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
