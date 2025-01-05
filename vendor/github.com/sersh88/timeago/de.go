package timeago

func deLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"gerade eben", "vor einer Weile"},
		{"vor %d Sekunden", "in %d Sekunden"},
		{"vor 1 Minute", "in 1 Minute"},
		{"vor %d Minuten", "in %d Minuten"},
		{"vor 1 Stunde", "in 1 Stunde"},
		{"vor %d Stunden", "in %d Stunden"},
		{"vor 1 Tag", "in 1 Tag"},
		{"vor %d Tagen", "in %d Tagen"},
		{"vor 1 Woche", "in 1 Woche"},
		{"vor %d Wochen", "in %d Wochen"},
		{"vor 1 Monat", "in 1 Monat"},
		{"vor %d Monaten", "in %d Monaten"},
		{"vor 1 Jahr", "in 1 Jahr"},
		{"vor %d Jahren", "in %d Jahren"},
	}[index]
	return res[0], res[1]
}
