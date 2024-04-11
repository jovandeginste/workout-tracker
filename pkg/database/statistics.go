package database

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

func (u *User) GetDefaultStatistics() (*Statistics, error) {
	return u.GetStatistics(StatConfig{
		Since: "-1 year",
		Per:   "month",
	})
}

func (u *User) GetStatistics(statConfig StatConfig) (*Statistics, error) {
	r := &Statistics{
		UserID:  u.ID,
		Buckets: map[WorkoutType]map[string]Bucket{},
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
		&r.AverageSpeed:        "max(total_distance / (total_duration / 1000000000))",
		&r.AverageSpeedNoPause: "max(total_distance / ((total_duration - pause_duration) / 1000000000))",
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
