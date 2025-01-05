package timeago

func esLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"justo ahora", "en un rato"},
		{"hace %d segundos", "en %d segundos"},
		{"hace 1 minuto", "en 1 minuto"},
		{"hace %d minutos", "en %d minutos"},
		{"hace 1 hora", "en 1 hora"},
		{"hace %d horas", "en %d horas"},
		{"hace 1 día", "en 1 día"},
		{"hace %d días", "en %d días"},
		{"hace 1 semana", "en 1 semana"},
		{"hace %d semanas", "en %d semanas"},
		{"hace 1 mes", "en 1 mes"},
		{"hace %d meses", "en %d meses"},
		{"hace 1 año", "en 1 año"},
		{"hace %d años", "en %d años"},
	}[index]
	return res[0], res[1]
}
