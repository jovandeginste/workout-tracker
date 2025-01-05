package timeago

func daLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"for et øjeblik siden", "om et øjeblik"},
		{"for %d sekunder siden", "om %d sekunder"},
		{"for 1 minut siden", "om 1 minut"},
		{"for %d minutter siden", "om %d minutter"},
		{"for 1 time siden", "om 1 time"},
		{"for %d timer siden", "om %d timer"},
		{"for 1 dag siden", "om 1 dag"},
		{"for %d dage siden", "om %d dage"},
		{"for 1 uge siden", "om 1 uge"},
		{"for %d uger siden", "om %d uger"},
		{"for 1 måned siden", "om 1 måned"},
		{"for %d måneder siden", "om %d måneder"},
		{"for 1 år siden", "om 1 år"},
		{"for %d år siden", "om %d år"},
	}[index]
	return res[0], res[1]
}
