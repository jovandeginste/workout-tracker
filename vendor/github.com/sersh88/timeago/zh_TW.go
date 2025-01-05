package timeago

func zhTWLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"剛剛", "片刻後"},
		{"%d 秒前", "%d 秒後"},
		{"1 分鐘前", "1 分鐘後"},
		{"%d 分鐘前", "%d 分鐘後"},
		{"1 小時前", "1 小時後"},
		{"%d 小時前", "%d 小時後"},
		{"1 天前", "1 天後"},
		{"%d 天前", "%d 天後"},
		{"1 週前", "1 週後"},
		{"%d 週前", "%d 週後"},
		{"1 個月前", "1 個月後"},
		{"%d 個月前", "%d 個月後"},
		{"1 年前", "1 年後"},
		{"%d 年前", "%d 年後"},
	}[index]
	return res[0], res[1]
}
