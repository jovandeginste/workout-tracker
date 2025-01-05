package timeago

func euLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"orain", "denbora bat barru"},
		{"duela %d segundu", "%d segundu barru"},
		{"duela minutu 1", "minutu 1 barru"},
		{"duela %d minutu", "%d minutu barru"},
		{"duela ordu 1", "ordu 1 barru"},
		{"duela %d ordu", "%d ordu barru"},
		{"duela egun 1", "egun 1 barru"},
		{"duela %d egun", "%d egun barru"},
		{"duela aste 1", "aste 1 barru"},
		{"duela %d aste", "%d aste barru"},
		{"duela hillabete 1", "hillabete 1 barru"},
		{"duela %d hillabete", "%d hillabete barru"},
		{"duela urte 1", "urte 1 barru"},
		{"duela %d urte", "%d urte barru"},
	}[index]
	return res[0], res[1]
}
