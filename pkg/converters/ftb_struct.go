package converters

type FTBStruct struct {
	FitoTrack struct {
		IndoorSamples struct {
			IndoorSamples []struct {
				AbsoluteTime      string `json:"absoluteTime"`
				HeartRate         string `json:"heartRate"`
				ID                string `json:"id"`
				IntervalTriggered string `json:"intervalTriggered"`
				RelativeTime      string `json:"relativeTime"`
				AbsoluteEndTime   string `json:"absoluteEndTime"`
				Frequency         string `json:"frequency"`
				Intensity         string `json:"intensity"`
				Repetitions       string `json:"repetitions"`
				WorkoutID         string `json:"workoutId"`
			} `json:"indoorSamples"`
		} `json:"indoorSamples"`
		IndoorWorkouts struct {
			IndoorWorkouts []struct {
				AvgHeartRate      string `json:"avgHeartRate"`
				Calorie           string `json:"calorie"`
				Comment           string `json:"comment"`
				Duration          string `json:"duration"`
				Edited            string `json:"edited"`
				End               string `json:"end"`
				ID                string `json:"id"`
				IntervalSetUsedID string `json:"intervalSetUsedId"`
				MaxHeartRate      string `json:"maxHeartRate"`
				PauseDuration     string `json:"pauseDuration"`
				Start             string `json:"start"`
				AvgFrequency      string `json:"avgFrequency"`
				AvgIntensity      string `json:"avgIntensity"`
				MaxFrequency      string `json:"maxFrequency"`
				MaxIntensity      string `json:"maxIntensity"`
				Repetitions       string `json:"repetitions"`
				ExportFileName    string `json:"exportFileName"`
				WorkoutType       string `json:"workoutType"`
			} `json:"indoorWorkouts"`
		} `json:"indoorWorkouts"`
		IntervalSets any `json:"intervalSets"`
		Samples      struct {
			Samples []struct {
				AbsoluteTime      string `json:"absoluteTime"`
				HeartRate         string `json:"heartRate"`
				ID                string `json:"id"`
				IntervalTriggered string `json:"intervalTriggered"`
				RelativeTime      string `json:"relativeTime"`
				Elevation         string `json:"elevation"`
				ElevationMSL      string `json:"elevationMSL"`
				Lat               string `json:"lat"`
				Lon               string `json:"lon"`
				Pressure          string `json:"pressure"`
				Speed             string `json:"speed"`
				WorkoutID         string `json:"workoutId"`
			} `json:"samples"`
		} `json:"samples"`
		Version      string `json:"version"`
		WorkoutTypes any    `json:"workoutTypes"`
		Workouts     struct {
			Workouts []struct {
				AvgHeartRate      string `json:"avgHeartRate"`
				Calorie           string `json:"calorie"`
				Comment           string `json:"comment"`
				Duration          string `json:"duration"`
				Edited            string `json:"edited"`
				End               string `json:"end"`
				ID                string `json:"id"`
				IntervalSetUsedID string `json:"intervalSetUsedId"`
				MaxHeartRate      string `json:"maxHeartRate"`
				PauseDuration     string `json:"pauseDuration"`
				Start             string `json:"start"`
				Ascent            string `json:"ascent"`
				AvgPace           string `json:"avgPace"`
				AvgSpeed          string `json:"avgSpeed"`
				Descent           string `json:"descent"`
				Length            string `json:"length"`
				MaxElevationMSL   string `json:"maxElevationMSL"`
				MinElevationMSL   string `json:"minElevationMSL"`
				TopSpeed          string `json:"topSpeed"`
				ExportFileName    string `json:"exportFileName"`
				WorkoutType       string `json:"workoutType"`
			} `json:"workouts"`
		} `json:"workouts"`
	} `json:"fito-track"`
}
