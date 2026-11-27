package database

import (
	"cmp"
	"errors"
	"fmt"
	"slices"

	"gorm.io/gorm"
)

const postgresDialect = "postgres"

var ErrAnonymousUser = errors.New("no statistics available for anonymous user")

type StatConfig struct {
	Since     string   `query:"since"`
	Per       string   `query:"per"`
	Equipment []uint64 `query:"equipment"`
}

var bucketStrings = map[string]map[string]string{
	postgresDialect: {
		"year":  "YYYY",
		"week":  "YYYY-WW",
		"day":   "YYYY-MM-DD",
		"month": "YYYY-MM",
	},
	"sqlite": {
		"year":  "%Y",
		"week":  "%Y-%W",
		"day":   "%Y-%m-%d",
		"month": "%Y-%m",
	},
}

var dayBucketExpressions = map[string]map[string]string{
	postgresDialect: {
		"year":  "min(to_char(date_trunc('year', workouts.date), 'YYYY-MM-DD')) as bucket",
		"week":  "min(to_char(date_trunc('week', workouts.date), 'YYYY-MM-DD')) as bucket",
		"day":   "min(to_char(workouts.date, 'YYYY-MM-DD')) as bucket",
		"month": "min(to_char(date_trunc('month', workouts.date), 'YYYY-MM-DD')) as bucket",
	},
	"sqlite": {
		"year":  "min(strftime('%Y-01-01', workouts.date)) as bucket",
		"week":  "min(date(workouts.date, '-6 days', 'weekday 1')) as bucket",
		"day":   "min(strftime('%Y-%m-%d', workouts.date)) as bucket",
		"month": "min(strftime('%Y-%m-01', workouts.date)) as bucket",
	},
}

func bucketValue(table map[string]string, per string) string {
	if value, ok := table[per]; ok {
		return value
	}

	return table["month"]
}

func (sc *StatConfig) GetBucketString(sqlDialect string) string {
	table, ok := bucketStrings[sqlDialect]
	if !ok {
		table = bucketStrings["sqlite"]
	}

	return bucketValue(table, sc.Per)
}

func (sc *StatConfig) GetDayBucketFormatExpression(sqlDialect string) string {
	table, ok := dayBucketExpressions[sqlDialect]
	if !ok {
		table = dayBucketExpressions["sqlite"]
	}

	return bucketValue(table, sc.Per)
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

func (u *User) GetStatisticsFor(since, per string, equipment []uint64, viewer *User) (*Statistics, error) {
	return u.GetStatistics(StatConfig{
		Since:     since,
		Per:       per,
		Equipment: equipment,
	}, viewer)
}

// GetStatistics returns the statistics of the user, limited to the workouts the
// viewer may see.
func (u *User) statsBaseQuery(viewer *User) *gorm.DB {
	return VisibleWorkouts(u.db, viewer).
		Table("workouts").
		Joins("join map_data on workouts.id = map_data.workout_id").
		Where("user_id = ?", u.ID)
}

func (u *User) GetStatistics(statConfig StatConfig, viewer *User) (*Statistics, error) {
	sqlDialect := u.db.Dialector.Name()

	r := &Statistics{
		UserID:       u.ID,
		BucketFormat: statConfig.GetBucketString(sqlDialect),
		Buckets:      map[WorkoutType]Buckets{},
	}

	q := u.statsBaseQuery(viewer).
		Select(
			"count(*) as workouts",
			"workouts.type as workout_type",
			"sum(total_duration) as duration",
			"sum(total_distance) as distance",
			"sum(total_up) as up",
			"max(max_speed) as max_speed",
			"sum(total_duration * average_speed) / sum(total_duration) as average_speed",
			"sum((total_duration - pause_duration) * average_speed_no_pause) / NULLIF(sum(total_duration - pause_duration), 0) as average_speed_no_pause",
			statConfig.GetBucketFormatExpression(sqlDialect),
			statConfig.GetDayBucketFormatExpression(sqlDialect),
		)

	if statConfig.Since != "" && statConfig.Since != "forever" {
		q = q.Where(GetDateLimitExpression(sqlDialect), "-"+statConfig.GetSince())
	}

	equipment := make([]uint64, 0, len(statConfig.Equipment))
	for _, id := range statConfig.Equipment {
		if id != 0 {
			equipment = append(equipment, id)
		}
	}

	if len(equipment) > 0 {
		q = q.Where(
			"workouts.id IN (SELECT workout_id FROM workout_equipment WHERE equipment_id IN ?)",
			equipment,
		)
	}

	// Grouping by `raw_bucket` instead of `bucket` ensures that the data is grouped
	// based on the raw, unprocessed bucket values as defined in the database schema.
	// This is necessary to maintain consistency with the bucket format expressions
	// used in the SELECT clause and to avoid potential mismatches caused by
	// transformations or processing applied to `bucket`.
	// The `bucket` field is provided for frontend rendering purposes only.
	rows, err := q.Group("raw_bucket, workouts.type").Rows()
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

		result.Localize(units, &result.WorkoutType)

		r.Buckets[result.WorkoutType].Buckets[result.Bucket] = result
	}

	return r, nil
}

func GetWorkoutTypeCounts(db *gorm.DB, userID uint64, viewer *User) (map[WorkoutType]int, error) {
	var rows []struct {
		WorkoutType WorkoutType `gorm:"column:workout_type"`
		Count       int         `gorm:"column:workout_count"`
	}

	err := VisibleWorkouts(db, viewer).
		Table("workouts").
		Select("workouts.type as workout_type", "count(*) as workout_count").
		Where("user_id = ?", userID).
		Group("workouts.type").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	counts := make(map[WorkoutType]int, len(rows))
	for _, r := range rows {
		counts[r.WorkoutType] = r.Count
	}

	return counts, nil
}

func (u *User) GetHighestWorkoutType(viewer *User) (*WorkoutType, error) {
	r := ""

	err := VisibleWorkouts(u.db, viewer).
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

func (u *User) GetDefaultTotals(viewer *User) ([]*Bucket, error) {
	if u.IsAnonymous() {
		return nil, ErrAnonymousUser
	}

	types, err := u.defaultTotalsTypes(viewer)
	if err != nil {
		return nil, err
	}

	return u.GetTotalsFor(types, viewer)
}

func (u *User) defaultTotalsTypes(viewer *User) (WorkoutTypeList, error) {
	types := slices.Clone(u.Profile.TotalsShow)
	types.Sanitize()

	if len(types) == 0 {
		types = WorkoutTypeList{WorkoutTypeAutoDetect}
	}

	result := make(WorkoutTypeList, 0, len(types))

	for _, t := range types {
		if t == WorkoutTypeAutoDetect {
			ht, err := u.GetHighestWorkoutType(viewer)
			if err != nil {
				return nil, err
			}

			t = *ht

			// Keep the auto-detect label when there is no workout to detect.
			if t == "" || t == WorkoutTypeUnknown {
				t = WorkoutTypeAutoDetect
			}
		}

		if !result.Has(t) {
			result = append(result, t)
		}
	}

	return result, nil
}

func (u *User) GetTotalsFor(types WorkoutTypeList, viewer *User) ([]*Bucket, error) {
	if len(types) == 0 {
		return nil, nil
	}

	rows := []*Bucket{}

	err := u.statsBaseQuery(viewer).
		Select(
			"count(*) as workouts",
			"workouts.type as workout_type",
			"sum(total_duration) as duration",
			"sum(total_distance) as distance",
			"sum(total_up) as up",
			"'all' as bucket",
		).
		Where("workouts.type IN ?", types).
		Group("workouts.type").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	byType := map[WorkoutType]*Bucket{}
	for _, r := range rows {
		byType[r.WorkoutType] = r
	}

	result := make([]*Bucket, 0, len(types))

	for _, t := range types {
		r, ok := byType[t]
		if !ok {
			r = &Bucket{WorkoutType: t}
		}

		result = append(result, r)
	}

	return result, nil
}

func (u *User) GetTotals(t WorkoutType, viewer *User) (*Bucket, error) {
	if t == "" {
		t = WorkoutTypeRunning
	}

	r := &Bucket{}

	err := u.statsBaseQuery(viewer).
		Select(
			"count(*) as workouts",
			"max(workouts.type) as workout_type",
			"sum(total_duration) as duration",
			"sum(total_distance) as distance",
			"sum(total_up) as up",
			"'all' as bucket",
		).
		Where("workouts.type = ?", t).
		Scan(r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (u *User) GetAllRecords(viewer *User) ([]*WorkoutRecord, error) {
	if u.IsAnonymous() {
		return nil, ErrAnonymousUser
	}

	rs := []*WorkoutRecord{}

	for _, w := range DistanceWorkoutTypes() {
		r, err := u.GetRecords(w, viewer)
		if err != nil {
			return nil, err
		}

		if r != nil {
			rs = append(rs, r)
		}
	}

	// Most-used types first; keep the configured order for equal counts.
	slices.SortStableFunc(rs, func(a, b *WorkoutRecord) int {
		return cmp.Compare(b.Workouts, a.Workouts)
	})

	return rs, nil
}

func (u *User) GetRecords(t WorkoutType, viewer *User) (*WorkoutRecord, error) {
	if t == "" {
		types, err := u.defaultTotalsTypes(viewer)
		if err != nil {
			return nil, err
		}

		if len(types) > 0 {
			t = types[0]
		}
	}

	if t == "" {
		t = WorkoutTypeRunning
	}

	r := &WorkoutRecord{WorkoutType: t}

	var workouts int64

	if err := u.statsBaseQuery(viewer).
		Where("workouts.type = ?", t).
		Count(&workouts).Error; err != nil {
		return nil, err
	}

	r.Workouts = int(workouts)

	mapping := map[*Float64Record]string{
		&r.Distance:            "max(total_distance)",
		&r.MaxSpeed:            "max(max_speed)",
		&r.TotalUp:             "max(total_up)",
		&r.AverageSpeed:        "max(average_speed)",
		&r.AverageSpeedNoPause: "max(average_speed_no_pause)",
	}

	for k, v := range mapping {
		err := u.statsBaseQuery(viewer).
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

	err := VisibleWorkouts(u.db, viewer).
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
