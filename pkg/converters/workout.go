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
		Name                string        `json:"name"`                                // The name of the workout
		Type                string        `json:"type"`                                // The type of the workout
		Start               time.Time     `json:"start"`                               // The start time of the workout
		Stop                time.Time     `json:"stop"`                                // The stop time of the workout
		TotalDistance       float64       `json:"totalDistance"`                       // The total distance of the workout
		TotalDistance2D     float64       `json:"totalDistance2D"`                     // The total 2D distance of the workout
		TotalDuration       time.Duration `json:"totalDuration"`                       // The total duration of the workout
		MaxSpeed            float64       `json:"maxSpeed"`                            // The maximum speed of the workout
		MaxCadence          float64       `json:"maxCadence"`                          // The maximum cadence of the workout
		AverageSpeed        float64       `json:"averageSpeed"`                        // The average speed of the workout
		AverageSpeedNoPause float64       `json:"averageSpeedNoPause"`                 // The average speed of the workout without pausing
		AverageCadence      float64       `json:"averageCadence"`                      // The average cadence of the workout
		PauseDuration       time.Duration `json:"pauseDuration"`                       // The total pause duration of the workout
		MinElevation        float64       `json:"minElevation"`                        // The minimum elevation of the workout
		MaxElevation        float64       `json:"maxElevation"`                        // The maximum elevation of the workout
		TotalUp             float64       `json:"totalUp"`                             // The total distance up of the workout
		TotalDown           float64       `json:"totalDown"`                           // The total distance down of the workout
		TotalRepetitions    int           `json:"totalRepetitions"`                    // The number of repetitions of the workout
		TotalWeight         float64       `json:"totalWeight"`                         // The weight of the workout
		TotalCalories       float64       `json:"totalCalories"`                       // The total calories of the workout
		ExtraMetrics        []string      `gorm:"serializer:json" json:"extraMetrics"` // Extra metrics available
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
