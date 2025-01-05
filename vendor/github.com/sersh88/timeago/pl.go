package timeago

func plLocale(diff float64, index int) (ago string, in string) {
	// 0-13 alternately: single unit of time,
	// genitive plural form for all other numbers excluding cases below:
	// 14-20: nominative plural form for the numbers 2,3,4
	// and all other numbers higher than 21 which end in 2,3,4
	var l = [][]string{
		{"w tej chwili", "za chwilę"},
		{"%d sekund temu", "za %d sekund"},
		{"1 minutę temu", "za 1 minutę"},
		{"%d minut temu", "za %d minut"},
		{"1 godzinę temu", "za 1 godzinę"},
		{"%d godzin temu", "za %d godzin"},
		{"1 dzień temu", "za 1 dzień"}, // {"wczoraj", "jutro"},
		{"%d dni temu", "za %d dni"},
		{"1 tydzień temu", "za 1 tydzień"},
		{"%d tygodni temu", "za %d tygodni"},
		{"1 miesiąc temu", "za 1 miesiąc"},
		{"%d miesięcy temu", "za %d miesięcy"},
		{"1 rok temu", "za 1 rok"},
		{"%d lat temu", "za %d lat"},
		{"%d sekundy temu", "za %d sekundy"},
		{"%d minuty temu", "za %d minuty"},
		{"%d godziny temu", "za %d godziny"},
		{"%d dni temu", "za %d dni"},
		{"%d tygodnie temu", "za %d tygodnie"},
		{"%d miesiące temu", "za %d miesiące"},
		{"%d lata temu", "za %d lata"},
	}
	// to determine which plural form must be used check the last 2 digits
	// and calculate new index value to get the nominative form (14-20)
	// for all other cases use index value as it is (0-13)
	var number = int64(diff)
	idx := index
	if index&1 != 0 {
		if number%10 > 4 || number%10 < 2 || 1 == ^^(number/10)%10 {
			idx = index
		} else {
			idx = index + 1
			idx = idx/2 + 13
		}
	}
	return l[idx][0], l[idx][1]
}
