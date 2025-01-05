package timeago

func idLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"baru saja", "sebentar"},
		{"%d detik yang lalu", "dalam %d detik"},
		{"1 menit yang lalu", "dalam 1 menit"},
		{"%d menit yang lalu", "dalam %d menit"},
		{"1 jam yang lalu", "dalam 1 jam"},
		{"%d jam yang lalu", "dalam %d jam"},
		{"1 hari yang lalu", "dalam 1 hari"},
		{"%d hari yang lalu", "dalam %d hari"},
		{"1 minggu yang lalu", "dalam 1 minggu"},
		{"%d minggu yang lalu", "dalam %d minggu"},
		{"1 bulan yang lalu", "dalam 1 bulan"},
		{"%d bulan yang lalu", "dalam %d bulan"},
		{"1 tahun yang lalu", "dalam 1 tahun"},
		{"%d tahun yang lalu", "dalam %d tahun"},
	}[index]
	return res[0], res[1]
}
