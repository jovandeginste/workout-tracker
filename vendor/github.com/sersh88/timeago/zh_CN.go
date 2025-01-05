package timeago

import (
	"fmt"
)

var zhCN = []string{"秒", "分钟", "小时", "天", "周", "个月", "年"}

func zhCNLocale(diff float64, idx int) (ago string, in string) {
	if idx == 0 {
		return "刚刚", "片刻后"
	}
	var unit = zhCN[^^(idx / 2)]
	return fmt.Sprintf("%d %s前", int(diff), unit),
		fmt.Sprintf(`%d %s后`, int(diff), unit)
}
