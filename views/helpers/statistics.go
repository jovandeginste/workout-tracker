package helpers

func StatisticSinceOptions() []TranslatedKey {
	return []TranslatedKey{
		{Key: "7 day", Translation: "misc.day_7"},
		{Key: "1 month", Translation: "misc.month_1"},
		{Key: "3 months", Translation: "misc.month_3"},
		{Key: "6 months", Translation: "misc.month_6"},
		{Key: "1 year", Translation: "misc.years_1"},
		{Key: "2 years", Translation: "misc.years_2"},
		{Key: "5 years", Translation: "misc.years_5"},
		{Key: "10 years", Translation: "misc.years_10"},
		{Key: "forever", Translation: "misc.forever"},
	}
}

func StatisticPerOptions() []TranslatedKey {
	return []TranslatedKey{
		{Key: "day", Translation: "misc.day"},
		{Key: "week", Translation: "misc.day_7"},
		{Key: "month", Translation: "misc.month"},
	}
}
