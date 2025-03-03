package converters

import (
	"time"

	"github.com/tkrajina/gpxgo/gpx"
)

type (
	Workout struct {
		Data     WorkoutData
		Content  []byte
		GPX      *gpx.GPX
		FileType string
	}

	WorkoutData struct {
		Name                string        // The name of the workout
		Type                string        // The type of the workout
		Start               time.Time     // The start time of the workout
		Stop                time.Time     // The stop time of the workout
		TotalDistance       float64       // The total distance of the workout
		TotalDuration       time.Duration // The total duration of the workout
		MaxSpeed            float64       // The maximum speed of the workout
		AverageSpeed        float64       // The average speed of the workout
		AverageSpeedNoPause float64       // The average speed of the workout without pausing
		PauseDuration       time.Duration // The total pause duration of the workout
		MinElevation        float64       // The minimum elevation of the workout
		MaxElevation        float64       // The maximum elevation of the workout
		TotalUp             float64       // The total distance up of the workout
		TotalDown           float64       // The total distance down of the workout
		TotalRepetitions    int           // The number of repetitions of the workout
		TotalWeight         float64       // The weight of the workout
		ExtraMetrics        []string      `gorm:"serializer:json"` // Extra metrcis available
	}
)

func (w *Workout) IsGPXBAsed() bool {
	return w.GPX != nil
}

func (w *Workout) FixName(basename string) {
	if w.Data.Name != "" {
		return
	}

	if !w.IsGPXBAsed() {
		w.Data.Name = basename
		return
	}

	if w.GPX.Name != "" {
		// We have a name
		w.Data.Name = w.GPX.Name
		return
	}

	if len(w.GPX.Tracks) > 0 && w.GPX.Tracks[0].Name != "" {
		// Copy the name of the first track
		w.Data.Name = w.GPX.Tracks[0].Name
		return
	}

	// Use the filename
	w.Data.Name = basename
}

func (w *Workout) Filename() string {
	if w.FileType == "" {
		return w.Data.Name
	}

	return w.Data.Name + "." + w.FileType
}
