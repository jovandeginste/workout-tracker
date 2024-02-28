package database

import (
	"time"

	"gorm.io/gorm"
)

type (
	record struct {
		Value float64
		Date  time.Time
		ID    uint
	}

	WorkoutStatistics map[string]*WorkoutStatistic

	WorkoutStatistic struct {
		WorkoutType WorkoutType
		Total       Totals
		PerYear     map[int]*Totals
		PerMonth    map[int]map[int]*Totals
		Records     WorkoutRecord
	}

	Totals struct {
		WorkoutType         WorkoutType
		Workouts            int
		Distance            float64
		Up                  float64
		Duration            time.Duration
		AverageSpeed        float64
		AverageSpeedNoPause float64
		MaxSpeed            float64
	}

	WorkoutRecord struct {
		WorkoutType         WorkoutType
		Active              bool
		AverageSpeed        record
		AverageSpeedNoPause record
		MaxSpeed            record
		Distance            record
		TotalUp             record
		Duration            struct {
			Value time.Duration
			Date  time.Time
			ID    uint
		}
	}
)

func (t *Totals) AverageSpeedKPH() float64 {
	return 3.6 * t.AverageSpeed
}

func (t *Totals) AverageSpeedNoPauseKPH() float64 {
	return 3.6 * t.AverageSpeedNoPause
}

func (t *Totals) MaxSpeedKPH() float64 {
	return 3.6 * t.MaxSpeed
}

func NewWorkoutStatistic(t WorkoutType) *WorkoutStatistic {
	return &WorkoutStatistic{
		WorkoutType: t,
		Records:     WorkoutRecord{WorkoutType: t},
		Total:       Totals{WorkoutType: t},
		PerYear:     map[int]*Totals{},
		PerMonth:    map[int]map[int]*Totals{},
	}
}

func (r *record) CheckAndSwap(value float64, id uint, date *time.Time) {
	if r.Value < value {
		r.Value = value
		r.Date = *date
		r.ID = id
	}
}

func (us *WorkoutStatistic) Add(w *Workout) {
	us.Total.Workouts++
	us.Total.Distance += w.Data.TotalDistance
	us.Total.Duration += w.Data.TotalDuration
	us.Total.Up += w.Data.TotalUp
	us.Total.AverageSpeed += w.Data.AverageSpeed()
	us.Total.AverageSpeedNoPause += w.Data.AverageSpeedNoPause()
	us.Total.MaxSpeed += w.Data.MaxSpeed

	d := w.Date
	year := d.Year()
	month := int(d.Month())

	us.AddYear(us.WorkoutType, year, w)
	us.AddMonth(us.WorkoutType, year, month, w)
}

func NewTotal(t WorkoutType, d *MapData) *Totals {
	return &Totals{
		Workouts:            1,
		WorkoutType:         t,
		Distance:            d.TotalDistance,
		Up:                  d.TotalUp,
		Duration:            d.TotalDuration,
		AverageSpeed:        d.AverageSpeed(),
		AverageSpeedNoPause: d.AverageSpeedNoPause(),
		MaxSpeed:            d.MaxSpeed,
	}
}

func calcNewAvg(oldValue, newValue float64, newCounter int) float64 {
	c := float64(newCounter)
	return ((oldValue * (c - 1)) + newValue) / c
}

func (t *Totals) Add(d *MapData) {
	t.Workouts++
	t.Distance += d.TotalDistance
	t.Duration += d.TotalDuration
	t.Up += d.TotalUp

	t.AverageSpeed = calcNewAvg(t.AverageSpeed, d.AverageSpeed(), t.Workouts)
	t.AverageSpeedNoPause = calcNewAvg(t.AverageSpeedNoPause, d.AverageSpeedNoPause(), t.Workouts)

	if d.MaxSpeed > t.MaxSpeed {
		t.MaxSpeed = d.MaxSpeed
	}
}

func (us *WorkoutStatistic) AddMonth(t WorkoutType, year, month int, w *Workout) {
	if _, ok := us.PerMonth[year]; !ok {
		us.PerMonth[year] = map[int]*Totals{}
	}

	entry, ok := us.PerMonth[year][month]
	if !ok {
		us.PerMonth[year][month] = NewTotal(t, w.Data)

		return
	}

	entry.Add(w.Data)
}

func (us *WorkoutStatistic) AddYear(t WorkoutType, year int, w *Workout) {
	entry, ok := us.PerYear[year]
	if !ok {
		us.PerYear[year] = NewTotal(t, w.Data)

		return
	}

	entry.Add(w.Data)
}

func (u *User) Statistics(db *gorm.DB) (WorkoutStatistics, error) {
	us := WorkoutStatistics{}

	workouts, err := u.GetWorkouts(db)
	if err != nil {
		return nil, err
	}

	for _, w := range workouts {
		if !w.Type.IsDistance() {
			continue
		}

		s, ok := us[w.Type.String()]
		if !ok {
			s = NewWorkoutStatistic(w.Type)
			us[w.Type.String()] = s
		}

		s.Records.Active = true
		s.Add(w)

		s.Records.CheckAndSwap(w)
	}

	return us, nil
}

func (wr *WorkoutRecord) CheckAndSwap(w *Workout) {
	wr.AverageSpeedNoPause.CheckAndSwap(w.Data.AverageSpeedNoPause(), w.ID, w.Date)
	wr.AverageSpeed.CheckAndSwap(w.Data.AverageSpeed(), w.ID, w.Date)
	wr.MaxSpeed.CheckAndSwap(w.Data.MaxSpeed, w.ID, w.Date)
	wr.Distance.CheckAndSwap(w.Data.TotalDistance, w.ID, w.Date)
	wr.TotalUp.CheckAndSwap(w.Data.TotalUp, w.ID, w.Date)

	if w.Data.TotalDuration > wr.Duration.Value {
		wr.Duration.Value = w.Data.TotalDuration
		wr.Duration.ID = w.ID
		wr.Duration.Date = *w.Date
	}
}
