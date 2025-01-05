package timeago

func nlLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"recent", "binnenkort"},
		{"%d seconden geleden", "binnen %d seconden"},
		{"1 minuut geleden", "binnen 1 minuut"},
		{"%d minuten geleden", "binnen %d minuten"},
		{"1 uur geleden", "binnen 1 uur"},
		{"%d uur geleden", "binnen %d uur"},
		{"1 dag geleden", "binnen 1 dag"},
		{"%d dagen geleden", "binnen %d dagen"},
		{"1 week geleden", "binnen 1 week"},
		{"%d weken geleden", "binnen %d weken"},
		{"1 maand geleden", "binnen 1 maand"},
		{"%d maanden geleden", "binnen %d maanden"},
		{"1 jaar geleden", "binnen 1 jaar"},
		{"%d jaar geleden", "binnen %d jaar"},
	}[index]
	return res[0], res[1]
}
