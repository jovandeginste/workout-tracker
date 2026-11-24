package helpers

func StatisticSinceOptions() []TranslatedKey {
	return []TranslatedKey{
		{Key: "7 day", Translation: "misc.day-7"},
		{Key: "1 month", Translation: "misc.month-1"},
		{Key: "3 months", Translation: "misc.month-3"},
		{Key: "6 months", Translation: "misc.month-6"},
		{Key: "1 year", Translation: "misc.years-1"},
		{Key: "2 years", Translation: "misc.years-2"},
		{Key: "5 years", Translation: "misc.years-5"},
		{Key: "10 years", Translation: "misc.years-10"},
		{Key: "forever", Translation: "misc.forever"},
	}
}

func StatisticPerOptions() []TranslatedKey {
	return []TranslatedKey{
		{Key: "day", Translation: "misc.day"},
		{Key: "week", Translation: "misc.day-7"},
		{Key: "month", Translation: "misc.month"},
		{Key: "year", Translation: "misc.year"},
	}
}
