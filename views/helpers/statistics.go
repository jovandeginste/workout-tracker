package helpers

func StatisticSinceOptions() []string {
	return []string{
		"misc.day_7",
		"misc.month_1", "misc.month_3", "misc.month_6",
		"misc.years_1", "misc.years_2", "misc.years_5", "misc.years_10",
		"misc.forever",
	}
}

func StatisticPerOptions() []string {
	return []string{
		"misc.day",
		"misc.day_7",
		"misc.day_15",
		"misc.month",
	}
}
