package timeago

func frLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"à l'instant", "dans un instant"},
		{"il y a %d secondes", "dans %d secondes"},
		{"il y a 1 minute", "dans 1 minute"},
		{"il y a %d minutes", "dans %d minutes"},
		{"il y a 1 heure", "dans 1 heure"},
		{"il y a %d heures", "dans %d heures"},
		{"il y a 1 jour", "dans 1 jour"},
		{"il y a %d jours", "dans %d jours"},
		{"il y a 1 semaine", "dans 1 semaine"},
		{"il y a %d semaines", "dans %d semaines"},
		{"il y a 1 mois", "dans 1 mois"},
		{"il y a %d mois", "dans %d mois"},
		{"il y a 1 an", "dans 1 an"},
		{"il y a %d ans", "dans %d ans"},
	}[index]
	return res[0], res[1]
}
