package timeago

func viLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"vừa xong", "một lúc"},
		{"%d giây trước", "trong %d giây"},
		{"1 phút trước", "trong 1 phút"},
		{"%d phút trước", "trong %d phút"},
		{"1 giờ trước", "trong 1 giờ"},
		{"%d giờ trước", "trong %d giờ"},
		{"1 ngày trước", "trong 1 ngày"},
		{"%d ngày trước", "trong %d ngày"},
		{"1 tuần trước", "trong 1 tuần"},
		{"%d tuần trước", "trong %d tuần"},
		{"1 tháng trước", "trong 1 tháng"},
		{"%d tháng trước", "trong %d tháng"},
		{"1 năm trước", "trong 1 năm"},
		{"%d năm trước", "trong %d năm"},
	}[index]
	return res[0], res[1]
}
