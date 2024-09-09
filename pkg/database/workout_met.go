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

func (w *Workout) distanceOnFootMET() float64 {
	s := w.Data.AverageSpeedNoPause() // meters per second

	switch {
	case s < 0.89:
		return 2.0
	case s < 1.34:
		return 2.5
	case s < 1.56:
		return 3.3
	case s < 1.79:
		return 3.8
	case s < 1.90:
		return 5.0
	case s < 2.10:
		return 6.5
	case s < 2.23:
		return 6.5
	case s < 2.68:
		return 8.0
	case s < 3.01:
		return 10.0
	case s < 3.36:
		return 11.0
	case s < 3.85:
		return 12.5
	case s < 4.47:
		return 14.0
	default:
		return 16.0
	}
}

// | Activity                         | Speed (km/h) | Speed (m/s) | METs |
// |----------------------------------|--------------|-------------|------|
// | Stationary cycling, 50 watts     | -            | -           | 3.0  |
// | Bicycling, leisurely             | -            | -           | 3.5  |
// | Stationary cycling, 100 watts    | -            | -           | 5.5  |
// | Bicycling, 12-13 mph             | 19.3         | 5.36        | 8.0  |
// | Bicycling, 14-15 mph             | 23.3         | 6.47        | 10.0 |
// | Bicycling, 16-19 mph             | 28.9         | 8.03        | 12.0 |
// | Bicycling, 20+ mph               | 32.2+        | 8.94+       | 16.0 |

func (w *Workout) distanceCyclingMET() float64 {
	s := w.Data.AverageSpeedNoPause() // meters per second

	switch {
	case s < 3.0:
		return 3.0
	case s < 4.0:
		return 3.5
	case s < 5.36:
		return 5.5
	case s < 6.47:
		return 8.0
	case s < 8.03:
		return 10.0
	case s < 8.94:
		return 12.0
	default:
		return 16.0
	}
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
	freq := w.RepetitionFrequencyPerMinute()

	switch {
	case freq < 5:
		return 3.5
	case freq < 10:
		return 5
	default:
		return 8.0
	}
}
