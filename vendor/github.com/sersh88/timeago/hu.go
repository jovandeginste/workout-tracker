package timeago

func huLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"éppen most", "éppen most"},
		{"%d másodperce", "%d másodpercen belül"},
		{"1 perce", "1 percen belül"},
		{"%d perce", "%d percen belül"},
		{"1 órája", "1 órán belül"},
		{"%d órája", "%d órán belül"},
		{"1 napja", "1 napon belül"},
		{"%d napja", "%d napon belül"},
		{"1 hete", "1 héten belül"},
		{"%d hete", "%d héten belül"},
		{"1 hónapja", "1 hónapon belül"},
		{"%d hónapja", "%d hónapon belül"},
		{"1 éve", "1 éven belül"},
		{"%d éve", "%d éven belül"},
	}[index]
	return res[0], res[1]
}
