package converters

import "time"

type (
	FitoTrackBackup struct {
		Version string `xml:"version"`

		IndoorWorkouts indoorWorkouts `xml:"indoorWorkouts"`
		Workouts       workouts       `xml:"workouts"`

		// IntervalSets any `xml:"intervalSets"`
		// WorkoutTypes any `xml:"workoutTypes"`
	}

	indoorWorkouts struct {
		IndoorWorkouts []indoorWorkout `xml:"indoorWorkouts"`
	}
	indoorWorkout struct {
		AvgFrequency float64 `xml:"avgFrequency"`
		AvgIntensity float64 `xml:"avgIntensity"`
		MaxFrequency float64 `xml:"maxFrequency"`
		MaxIntensity float64 `xml:"maxIntensity"`
		Repetitions  int     `xml:"repetitions"`
		workoutCommon
	}

	workouts struct {
		Workouts []workout `xml:"workouts"`
	}
	workout struct {
		Ascent          string  `xml:"ascent"`
		AvgPace         float64 `xml:"avgPace"`
		AvgSpeed        float64 `xml:"avgSpeed"`
		Descent         float64 `xml:"descent"`
		Length          int64   `xml:"length"`
		MaxElevationMSL float64 `xml:"maxElevationMSL"`
		MinElevationMSL float64 `xml:"minElevationMSL"`
		TopSpeed        float64 `xml:"topSpeed"`
		workoutCommon
	}

	workoutCommon struct {
		AvgHeartRate      float64 `xml:"avgHeartRate"`
		Calorie           float64 `xml:"calorie"`
		Comment           string  `xml:"comment"`
		Duration          int64   `xml:"duration"`
		Edited            string  `xml:"edited"`
		End               int64   `xml:"end"`
		ID                string  `xml:"id"`
		IntervalSetUsedID string  `xml:"intervalSetUsedId"`
		MaxHeartRate      float64 `xml:"maxHeartRate"`
		Start             int64   `xml:"start"`
		PauseDuration     int64   `xml:"pauseDuration"`
		ExportFileName    string  `xml:"exportFileName"`
		WorkoutType       string  `xml:"workoutType"`
	}
)

func (iw *indoorWorkout) StartTime() time.Time {
	return time.Unix(iw.Start/1000, 1000000*(iw.Start%1000))
}

func (iw *indoorWorkout) EndTime() time.Time {
	return time.Unix(iw.End/1000, 1000000*(iw.End%1000))
}
