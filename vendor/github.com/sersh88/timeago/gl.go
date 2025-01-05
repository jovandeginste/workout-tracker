package timeago

func glLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"xusto agora", "daquí a un pouco"},
		{"hai %d segundos", "en %d segundos"},
		{"hai 1 minuto", "nun minuto"},
		{"hai %d minutos", "en %d minutos"},
		{"hai 1 hora", "nunha hora"},
		{"hai %d horas", "en %d horas"},
		{"hai 1 día", "nun día"},
		{"hai %d días", "en %d días"},
		{"hai 1 semana", "nunha semana"},
		{"hai %d semanas", "en %d semanas"},
		{"hai 1 mes", "nun mes"},
		{"hai %d meses", "en %d meses"},
		{"hai 1 ano", "nun ano"},
		{"hai %d anos", "en %d anos"},
	}[index]
	return res[0], res[1]
}
