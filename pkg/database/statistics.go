package database

import (
	"time"

	"gorm.io/gorm"
)

type StatConfig struct {
	Since string `query:"since"`
	Per   string `query:"per"`
}

func (sc *StatConfig) GetBucketString() string {
	switch sc.Per {
	case "year":
		return "%Y"
	default:
		return "%Y-%m"
	}
}

func (sc *StatConfig) GetSince() string {
	if sc.Since == "" {
		return "-1 year"
	}

	return sc.Since
}

type Statistics struct {
	UserID  uint
	Buckets map[string]map[WorkoutType]Bucket
}

type Bucket struct {
	Bucket              string `json:",omitempty"`
	WorkoutType         WorkoutType
	Workouts            int
	Distance            float64       `json:",omitempty"`
	Up                  float64       `json:",omitempty"`
	Duration            time.Duration `json:",omitempty"`
	AverageSpeed        float64       `json:",omitempty"`
	AverageSpeedNoPause float64       `json:",omitempty"`
	MaxSpeed            float64       `json:",omitempty"`
}

func (u *User) GetStatistics(db *gorm.DB, statConfig StatConfig) (*Statistics, error) {
	r := &Statistics{
		UserID:  u.ID,
		Buckets: map[string]map[WorkoutType]Bucket{},
	}

	rows, err := db.
		Table("workouts").
		Select(
			"count(*) as workouts",
			"type as workout_type",
			"sum(total_duration) as duration",
			"sum(total_distance) as distance",
			"sum(total_up) as up",
			"max(max_speed) as max_speed",
			"max(total_distance / (total_duration / 1000000000)) as average_speed",
			"max(total_distance / ((total_duration - pause_duration) / 1000000000)) as average_speed_no_pause",
			"strftime('"+statConfig.GetBucketString()+"', workouts.date) as bucket",
		).
		Joins("join map_data on workouts.id = map_data.workout_id").
		Where("user_id = ?", u.ID).
		Where("workouts.date > DATE(CURRENT_DATE, ?)", statConfig.GetSince()).
		Group("bucket, workout_type").Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result Bucket

	for rows.Next() {
		if err := db.ScanRows(rows, &result); err != nil {
			return nil, err
		}

		if r.Buckets[result.Bucket] == nil {
			r.Buckets[result.Bucket] = map[WorkoutType]Bucket{}
		}

		r.Buckets[result.Bucket][result.WorkoutType] = result
	}

	return r, nil
}

func (u *User) GetTotals(db *gorm.DB, t WorkoutType) (*Bucket, error) {
	if t == "" {
		t = WorkoutTypeRunning
	}

	r := &Bucket{}

	err := db.
		Table("workouts").
		Select(
			"count(*) as workouts",
			"type as workout_type",
			"sum(total_duration) as duration",
			"sum(total_distance) as distance",
			"sum(total_up) as up",
			"'all' as bucket",
		).
		Joins("join map_data on workouts.id = map_data.workout_id").
		Where("user_id = ?", u.ID).
		Where("type = ?", t).
		Scan(r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (u *User) GetRecords(db *gorm.DB, t WorkoutType) (*WorkoutRecord, error) {
	if t == "" {
		t = WorkoutTypeRunning
	}

	r := &WorkoutRecord{}

	mapping := map[*record]string{
		&r.Distance:            "max(total_distance)",
		&r.MaxSpeed:            "max(max_speed)",
		&r.TotalUp:             "max(total_up)",
		&r.AverageSpeed:        "max(total_distance / (total_duration / 1000000000))",
		&r.AverageSpeedNoPause: "max(total_distance / ((total_duration - pause_duration) / 1000000000))",
	}

	for k, v := range mapping {
		err := db.
			Table("workouts").
			Joins("join map_data on workouts.id = map_data.workout_id").
			Where("user_id = ?", u.ID).
			Where("type = ?", t).
			Select("workouts.id as id", v+" as value", "workouts.date as date").
			Scan(k).Error
		if err != nil {
			return nil, err
		}
	}

	err := db.
		Table("workouts").
		Joins("join map_data on workouts.id = map_data.workout_id").
		Where("user_id = ?", u.ID).
		Where("type = ?", t).
		Select("workouts.id as id", "max(total_duration) as value", "workouts.date as date").
		Scan(&r.Duration).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}
