package timeago

func ocLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"fa un moment", "d'aquí un moment"},
		{"fa %d segondas", "d'aquí %d segondas"},
		{"fa 1 minuta", "d'aquí 1 minuta"},
		{"fa %d minutas", "d'aquí %d minutas"},
		{"fa 1 ora", "d'aquí 1 ora"},
		{"fa %d oras", "d'aquí %d oras"},
		{"fa 1 jorn", "d'aquí 1 jorn"},
		{"fa %d jorns", "d'aquí %d jorns"},
		{"fa 1 setmana", "d'aquí 1 setmana"},
		{"fa %d setmanas", "d'aquí %d setmanas"},
		{"fa 1 mes", "d'aquí 1 mes"},
		{"fa %d meses", "d'aquí %d meses"},
		{"fa 1 an", "d'aquí 1 an"},
		{"fa %d ans", "d'aquí %d ans"},
	}[index]
	return res[0], res[1]
}
