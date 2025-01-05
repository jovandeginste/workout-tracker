package timeago

func svLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"just nu", "om en stund"},
		{"%d sekunder sedan", "om %d sekunder"},
		{"1 minut sedan", "om 1 minut"},
		{"%d minuter sedan", "om %d minuter"},
		{"1 timme sedan", "om 1 timme"},
		{"%d timmar sedan", "om %d timmar"},
		{"1 dag sedan", "om 1 dag"},
		{"%d dagar sedan", "om %d dagar"},
		{"1 vecka sedan", "om 1 vecka"},
		{"%d veckor sedan", "om %d veckor"},
		{"1 månad sedan", "om 1 månad"},
		{"%d månader sedan", "om %d månader"},
		{"1 år sedan", "om 1 år"},
		{"%d år sedan", "om %d år"},
	}[index]
	return res[0], res[1]
}
