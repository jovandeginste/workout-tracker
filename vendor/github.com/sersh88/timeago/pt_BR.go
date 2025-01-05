package timeago

func ptLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"agora mesmo", "agora"},
		{"há %d segundos", "em %d segundos"},
		{"há um minuto", "em um minuto"},
		{"há %d minutos", "em %d minutos"},
		{"há uma hora", "em uma hora"},
		{"há %d horas", "em %d horas"},
		{"há um dia", "em um dia"},
		{"há %d dias", "em %d dias"},
		{"há uma semana", "em uma semana"},
		{"há %d semanas", "em %d semanas"},
		{"há um mês", "em um mês"},
		{"há %d meses", "em %d meses"},
		{"há um ano", "em um ano"},
		{"há %d anos", "em %d anos"},
	}[index]
	return res[0], res[1]
}
