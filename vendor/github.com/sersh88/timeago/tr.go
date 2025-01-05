package timeago

func trLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"az önce", "şimdi"},
		{"%d saniye önce", "%d saniye içinde"},
		{"1 dakika önce", "1 dakika içinde"},
		{"%d dakika önce", "%d dakika içinde"},
		{"1 saat önce", "1 saat içinde"},
		{"%d saat önce", "%d saat içinde"},
		{"1 gün önce", "1 gün içinde"},
		{"%d gün önce", "%d gün içinde"},
		{"1 hafta önce", "1 hafta içinde"},
		{"%d hafta önce", "%d hafta içinde"},
		{"1 ay önce", "1 ay içinde"},
		{"%d ay önce", "%d ay içinde"},
		{"1 yıl önce", "1 yıl içinde"},
		{"%d yıl önce", "%d yıl içinde"},
	}[index]
	return res[0], res[1]
}
