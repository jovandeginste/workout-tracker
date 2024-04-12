package database

import (
	"fmt"
	"time"
)

type StatConfig struct {
	Since string `query:"since"`
	Per   string `query:"per"`
}

func (sc *StatConfig) GetBucketString() string {
	switch sc.Per {
	case "year":
		return "%Y"
	case "week":
		return "%Y-%W"
	case "day":
		return "%Y-%m-%d"
	default:
		return "%Y-%m"
	}
}

func (sc *StatConfig) GetPostgresBucketString() string {
	switch sc.Per {
	case "year":
		return "YYYY"
	case "week":
		return "YYYY-WW"
	case "day":
		return "YYYY-MM-DD"
	default:
		return "YYYY-MM"
	}
}

func (sc *StatConfig) GetSince() string {
	if sc.Since == "" {
		return "-1 year"
	}

	return sc.Since
}

func (u *User) GetDefaultStatistics() (*Statistics, error) {
	return u.GetStatisticsFor("-1 year", "month")
}

func (u *User) GetStatisticsFor(since, per string) (*Statistics, error) {
	return u.GetStatistics(StatConfig{
		Since: since,
		Per:   per,
	})
}

func (u *User) GetStatistics(statConfig StatConfig) (*Statistics, error) {
	r := &Statistics{
		UserID:       u.ID,
		BucketFormat: statConfig.GetBucketString(),
		Buckets:      map[WorkoutType]map[string]Bucket{},
	}

	databaseType := u.db.Dialector.Name()

	bucketFormat := "strftime('" + statConfig.GetBucketString() + "', workouts.date) as bucket"
	if databaseType == "postgres" {
		bucketFormat = "to_char(workouts.date, '" + statConfig.GetPostgresBucketString() + "') as bucket"
	}

	dateLimit := "workouts.date > DATE(CURRENT_DATE, ?)"
	if databaseType == "postgres" {
		dateLimit = "workouts.date > CURRENT_DATE + cast(? as interval)"
	}

	rows, err := u.db.
		Table("workouts").
		Select(
			"count(*) as workouts",
			"type as workout_type",
			"sum(total_duration) as duration",
			"sum(total_distance) as distance",
			"sum(total_up) as up",
			"max(max_speed) as max_speed",
			fmt.Sprintf("avg(total_distance / (total_duration / %d)) as average_speed", time.Second),
			fmt.Sprintf("avg(total_distance / ((total_duration - pause_duration) / %d)) as average_speed_no_pause", time.Second),
			bucketFormat,
		).
		Joins("join map_data on workouts.id = map_data.workout_id").
		Where("user_id = ?", u.ID).
		Where(dateLimit, statConfig.GetSince()).
		Group("bucket, workout_type").Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result Bucket

	for rows.Next() {
		if err := u.db.ScanRows(rows, &result); err != nil {
			return nil, err
		}

		if r.Buckets[result.WorkoutType] == nil {
			r.Buckets[result.WorkoutType] = map[string]Bucket{}
		}

		r.Buckets[result.WorkoutType][result.Bucket] = result
	}

	return r, nil
}

func (u *User) GetDefaultTotals() (*Bucket, error) {
	return u.GetTotals(u.Profile.TotalsShow)
}

func (u *User) GetTotals(t WorkoutType) (*Bucket, error) {
	if t == "" {
		t = WorkoutTypeRunning
	}

	r := &Bucket{}

	err := u.db.
		Table("workouts").
		Select(
			"count(*) as workouts",
			"max(type) as workout_type",
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

func (u *User) GetAllRecords() ([]*WorkoutRecord, error) {
	rs := []*WorkoutRecord{}

	for _, w := range DistanceWorkoutTypes() {
		r, err := u.GetRecords(w)
		if err != nil {
			return nil, err
		}

		rs = append(rs, r)
	}

	return rs, nil
}

func (u *User) GetRecords(t WorkoutType) (*WorkoutRecord, error) {
	if t == "" {
		t = WorkoutTypeRunning
	}

	r := &WorkoutRecord{WorkoutType: t}

	mapping := map[*float64Record]string{
		&r.Distance:            "max(total_distance)",
		&r.MaxSpeed:            "max(max_speed)",
		&r.TotalUp:             "max(total_up)",
		&r.AverageSpeed:        fmt.Sprintf("avg(total_distance / (total_duration / %d))", time.Second),
		&r.AverageSpeedNoPause: fmt.Sprintf("avg(total_distance / ((total_duration - pause_duration) / %d))", time.Second),
	}

	for k, v := range mapping {
		err := u.db.
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

	err := u.db.
		Table("workouts").
		Joins("join map_data on workouts.id = map_data.workout_id").
		Where("user_id = ?", u.ID).
		Where("type = ?", t).
		Select("workouts.id as id", "max(total_duration) as value", "workouts.date as date").
		Scan(&r.Duration).Error
	if err != nil {
		return nil, err
	}

	r.Active = r.Distance.Value > 0

	return r, nil
}
