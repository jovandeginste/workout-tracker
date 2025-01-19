package helpers

func StatisticSinceOptions() []string {
	return []string{
		"7 days",
		"1 month", "3 months", "6 months",
		"1 year", "2 years", "5 years", "10 years",
		"forever",
	}
}

func StatisticPerOptions() []string {
	return []string{
		"day",
		"7 days",
		"15 days",
		"month",
	}
}
