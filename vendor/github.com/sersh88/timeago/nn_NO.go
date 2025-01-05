package timeago

func nnLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"nett no", "om litt"},
		{"%d sekund sidan", "om %d sekund"},
		{"1 minutt sidan", "om 1 minutt"},
		{"%d minutt sidan", "om %d minutt"},
		{"1 time sidan", "om 1 time"},
		{"%d timar sidan", "om %d timar"},
		{"1 dag sidan", "om 1 dag"},
		{"%d dagar sidan", "om %d dagar"},
		{"1 veke sidan", "om 1 veke"},
		{"%d veker sidan", "om %d veker"},
		{"1 månad sidan", "om 1 månad"},
		{"%d månadar sidan", "om %d månadar"},
		{"1 år sidan", "om 1 år"},
		{"%d år sidan", "om %d år"},
	}[index]
	return res[0], res[1]
}
