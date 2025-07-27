package database

import (
	"errors"
	"fmt"
	"time"
)

const postgresDialect = "postgres"

var ErrAnonymousUser = errors.New("no statistics available for anonymous user")

type StatConfig struct {
	Since string `query:"since"`
	Per   string `query:"per"`
}

func (sc *StatConfig) GetBucketString(sqlDialect string) string {
	switch sqlDialect {
	case postgresDialect:
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
	default:
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
}

func (sc *StatConfig) GetDayBucketFormatExpression(sqlDialect string) string {
	switch sqlDialect {
	case postgresDialect:
		return "min(to_char(workouts.date, 'YYYY-MM-DD')) as bucket"
	default:
		return "min(strftime('%Y-%m-%d', workouts.date)) as bucket"
	}
}

func (sc *StatConfig) GetBucketFormatExpression(sqlDialect string) string {
	switch sqlDialect {
	case postgresDialect:
		return fmt.Sprintf("to_char(workouts.date, '%s') as raw_bucket", sc.GetBucketString(sqlDialect))
	default:
		return fmt.Sprintf("strftime('%s', workouts.date) as raw_bucket", sc.GetBucketString(sqlDialect))
	}
}

func GetDateLimitExpression(sqlDialect string) string {
	switch sqlDialect {
	case postgresDialect:
		return "workouts.date > CURRENT_DATE + cast(? as interval)"
	default:
		return "workouts.date > DATE(CURRENT_DATE, ?)"
	}
}

func (sc *StatConfig) GetSince() string {
	s := sc.Since
	if s == "" {
		s = "1 year"
	}

	return s
}

func (u *User) GetDefaultStatistics() (*Statistics, error) {
	return u.GetStatisticsFor("misc.years_1", "misc.month")
}

func (u *User) GetStatisticsFor(since, per string) (*Statistics, error) {
	return u.GetStatistics(StatConfig{
		Since: since,
		Per:   per,
	})
}

func (u *User) GetStatistics(statConfig StatConfig) (*Statistics, error) {
	sqlDialect := u.db.Dialector.Name()

	r := &Statistics{
		UserID:       u.ID,
		BucketFormat: statConfig.GetBucketString(sqlDialect),
		Buckets:      map[WorkoutType]Buckets{},
	}

	q := u.db.
		Table("workouts").
		Select(
			"count(*) as workouts",
			"workouts.type as workout_type",
			"sum(total_duration) as duration",
			"sum(total_distance) as distance",
			"sum(total_up) as up",
			"max(max_speed) as max_speed",
			fmt.Sprintf("avg(total_distance / COALESCE(NULLIF(NULLIF(total_duration, 0) / %d, 0), total_distance)) as average_speed", time.Second),
			fmt.Sprintf("avg(total_distance / COALESCE(NULLIF(NULLIF(total_duration - pause_duration, 0) / %d, 0), total_distance)) as average_speed_no_pause", time.Second),
			statConfig.GetBucketFormatExpression(sqlDialect),
			statConfig.GetDayBucketFormatExpression(sqlDialect),
		).
		Joins("join map_data on workouts.id = map_data.workout_id").
		Where("user_id = ?", u.ID)

	if statConfig.Since != "" && statConfig.Since != "forever" {
		q = q.Where(GetDateLimitExpression(sqlDialect), "-"+statConfig.GetSince())
	}

	// Grouping by `raw_bucket` instead of `bucket` ensures that the data is grouped
	// based on the raw, unprocessed bucket values as defined in the database schema.
	// This is necessary to maintain consistency with the bucket format expressions
	// used in the SELECT clause and to avoid potential mismatches caused by
	// transformations or processing applied to `bucket`.
	// The `bucket` field is provided for frontend rendering purposes only.
	rows, err := q.Group("raw_bucket, workout_type").Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result Bucket

	units := u.PreferredUnits()

	for rows.Next() {
		if err := u.db.ScanRows(rows, &result); err != nil {
			return nil, err
		}

		if _, ok := r.Buckets[result.WorkoutType]; !ok {
			r.Buckets[result.WorkoutType] = Buckets{
				WorkoutType:      result.WorkoutType,
				LocalWorkoutType: u.I18n(result.WorkoutType.StringT()),
				Buckets:          map[string]Bucket{},
			}
		}

		result.Localize(units)

		r.Buckets[result.WorkoutType].Buckets[result.Bucket] = result
	}

	return r, nil
}

func (u *User) GetHighestWorkoutType() (*WorkoutType, error) {
	r := ""

	err := u.db.
		Table("workouts").
		Select("workouts.type").
		Where("user_id = ?", u.ID).
		Group("workouts.type").
		Order("count(*) DESC").
		Limit(1).
		Pluck("workouts.type", &r).Error
	if err != nil {
		return nil, err
	}

	wt := AsWorkoutType(r)

	return &wt, nil
}

func (u *User) GetDefaultTotals() (*Bucket, error) {
	if u.IsAnonymous() {
		return nil, ErrAnonymousUser
	}

	t := u.Profile.TotalsShow
	if t == WorkoutTypeAutoDetect {
		ht, err := u.GetHighestWorkoutType()
		if err != nil {
			return nil, err
		}

		t = *ht
	}

	return u.GetTotals(t)
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
			"max(workouts.type) as workout_type",
			"sum(total_duration) as duration",
			"sum(total_distance) as distance",
			"sum(total_up) as up",
			"'all' as bucket",
		).
		Joins("join map_data on workouts.id = map_data.workout_id").
		Where("user_id = ?", u.ID).
		Where("workouts.type = ?", t).
		Scan(r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (u *User) GetAllRecords() ([]*WorkoutRecord, error) {
	if u.IsAnonymous() {
		return nil, ErrAnonymousUser
	}

	rs := []*WorkoutRecord{}

	for _, w := range DistanceWorkoutTypes() {
		r, err := u.GetRecords(w)
		if err != nil {
			return nil, err
		}

		if r != nil {
			rs = append(rs, r)
		}
	}

	return rs, nil
}

func (u *User) GetRecords(t WorkoutType) (*WorkoutRecord, error) {
	if t == "" {
		t = u.Profile.TotalsShow
	}

	r := &WorkoutRecord{WorkoutType: t}

	mapping := map[*Float64Record]string{
		&r.Distance:            "max(total_distance)",
		&r.MaxSpeed:            "max(max_speed)",
		&r.TotalUp:             "max(total_up)",
		&r.AverageSpeed:        fmt.Sprintf("max(total_distance / COALESCE(NULLIF(NULLIF(total_duration, 0) / %d, 0), total_distance))", time.Second),
		&r.AverageSpeedNoPause: fmt.Sprintf("max(total_distance / COALESCE(NULLIF(NULLIF(total_duration - pause_duration, 0) / %d, 0), total_distance))", time.Second),
	}

	for k, v := range mapping {
		err := u.db.
			Table("workouts").
			Joins("join map_data on workouts.id = map_data.workout_id").
			Where("user_id = ?", u.ID).
			Where("workouts.type = ?", t).
			Select("workouts.id as id", v+" as value", "workouts.date as date").
			Order(v + " DESC").
			Group("workouts.id").
			Limit(1).
			Scan(k).Error
		if err != nil {
			return nil, err
		}
	}

	err := u.db.
		Table("workouts").
		Joins("join map_data on workouts.id = map_data.workout_id").
		Where("user_id = ?", u.ID).
		Where("workouts.type = ?", t).
		Select("workouts.id as id", "max(total_duration) as value", "workouts.date as date").
		Order("max(total_duration) DESC").
		Group("workouts.id").
		Limit(1).
		Scan(&r.Duration).Error
	if err != nil {
		return nil, err
	}

	r.Active = r.Distance.Value > 0

	return r, nil
}
