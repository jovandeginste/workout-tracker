package timeago

var csTimes = [][][]string{
	{{"právě teď", "právě teď"}},
	{{"před %d vteřinami", "za %d vteřiny"}, {"před %d vteřinami", "za %d vteřin"}},
	{{"před minutou", "za minutu"}},
	{{"před %d minutami", "za %d minuty"}, {"před %d minutami", "za %d minut"}},
	{{"před hodinou", "za hodinu"}},
	{{"před %d hodinami", "za %d hodiny"}, {"před %d hodinami", "za %d hodin"}},
	{{"včera", "zítra"}},
	{{"před %d dny", "za %d dny"}, {"před %d dny", "za %d dnů"}},
	{{"minulý týden", "příští týden"}},
	{{"před %d týdny", "za %d týdny"}, {"před %d týdny", "za %d týdnů"}},
	{{"minulý měsíc", "přístí měsíc"}},
	{{"před %d měsíci", "za %d měsíce"}, {"před %d měsíci", "za %d měsíců"}},
	{{"před rokem", "přístí rok"}},
	{{"před %d lety", "za %d roky"}, {"před %d lety", "za %d let"}},
}

func csLocale(number float64, index int) (ago string, in string) {
	var inflectionIndex = 0
	var isInflectionNeeded = index == 1 || index == 3 || index == 5 || index == 7 || index == 9 || index == 11 || index == 13
	if isInflectionNeeded && number >= 5 {
		inflectionIndex = 1
	}
	var res = csTimes[index][inflectionIndex]
	return res[0], res[1]
}
