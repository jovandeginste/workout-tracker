package database

// Information from:
// - https://media.hypersites.com/clients/1235/filemanager/MHC/METs.pdf
// - https://cdn-links.lww.com/permalink/mss/a/mss_43_8_2011_06_13_ainsworth_202093_sdc1.pdf

// MET returns the metabolic equivalent of task
func (w *Workout) MET() float64 {
	switch {
	case w.Type.IsDistance():
		return w.distanceMET()
	case w.Type.IsRepetition():
		return w.repititionMET()
	default:
		return 1
	}
}

// | Activity                        | Speed (km/h) | Speed (m/s) | METs |
// |---------------------------------|--------------|-------------|------|
// | Walking, slowly (stroll)        | -            | -           | 2.0  |
// | Walking, 2 mph                  | 3.2          | 0.89        | 2.5  |
// | Walking, 3 mph (20 min/mile)    | 4.8          | 1.34        | 3.3  |
// | Walking, 17 min/mile            | 5.6          | 1.56        | 3.8  |
// | Walking, 15 min/mile            | 6.4          | 1.79        | 5.0  |
// | Race walking, moderate pace     | -            | -           | 6.5  |
// | Hiking up hills                 | -            | -           | 6.9  |
// | Hiking hills, 12 lb pack        | -            | -           | 7.5  |
// | Jogging, 12 min/mile            | 8.0          | 2.23        | 8.0  |
// | Running, 10 min/mile            | 9.7          | 2.68        | 10.0 |
// | Running, 9 min/mile             | 10.8         | 3.01        | 11.0 |
// | Running, 8 min/mile             | 12.1         | 3.36        | 12.5 |
// | Running, 7 min/mile             | 13.8         | 3.85        | 14.0 |
// | Running, 6 min/mile             | 16.1         | 4.47        | 16.0 |

type metThreshold struct {
	below float64
	met   float64
}

var (
	onFootMETThresholds = []metThreshold{
		{0.89, 2.0},
		{1.34, 2.5},
		{1.56, 3.3},
		{1.79, 3.8},
		{1.90, 5.0},
		{2.10, 6.5},
		{2.68, 8.0},
		{3.01, 10.0},
		{3.36, 11.0},
		{3.85, 12.5},
		{4.47, 14.0},
	}
	cyclingMETThresholds = []metThreshold{
		{3.0, 3.0}, {4.0, 3.5}, {5.36, 5.5}, {6.47, 8.0}, {8.03, 10.0}, {8.94, 12.0},
	}
	repetitionMETThresholds = []metThreshold{{5, 3.5}, {10, 5}}
)

func metFor(value float64, thresholds []metThreshold, fallback float64) float64 {
	for _, t := range thresholds {
		if value < t.below {
			return t.met
		}
	}

	return fallback
}

func (w *Workout) distanceOnFootMET() float64 {
	return metFor(w.Data.AverageSpeedNoPause, onFootMETThresholds, 16.0)
}

func (w *Workout) distanceCyclingMET() float64 {
	return metFor(w.Data.AverageSpeedNoPause, cyclingMETThresholds, 16.0)
}

func (w *Workout) distanceMET() float64 {
	switch w.Type { //nolint:exhaustive
	case WorkoutTypeWalking, WorkoutTypeHiking, WorkoutTypeRunning:
		return w.distanceOnFootMET()
	case WorkoutTypeCycling:
		return w.distanceCyclingMET()
	default:
		return 3.5
	}
}

func (w *Workout) repititionMET() float64 {
	return metFor(w.RepetitionFrequencyPerMinute(), repetitionMETThresholds, 8.0)
}
