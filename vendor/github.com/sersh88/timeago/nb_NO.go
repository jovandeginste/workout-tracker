package timeago

func nbLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"akkurat nå", "om litt"},
		{"%d sekunder siden", "om %d sekunder"},
		{"1 minutt siden", "om 1 minutt"},
		{"%d minutter siden", "om %d minutter"},
		{"1 time siden", "om 1 time"},
		{"%d timer siden", "om %d timer"},
		{"1 dag siden", "om 1 dag"},
		{"%d dager siden", "om %d dager"},
		{"1 uke siden", "om 1 uke"},
		{"%d uker siden", "om %d uker"},
		{"1 måned siden", "om 1 måned"},
		{"%d måneder siden", "om %d måneder"},
		{"1 år siden", "om 1 år"},
		{"%d år siden", "om %d år"},
	}[index]
	return res[0], res[1]
}
