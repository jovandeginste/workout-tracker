package timeago

func itLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"poco fa", "fra poco"},
		{"%d secondi fa", "fra %d secondi"},
		{"un minuto fa", "fra un minuto"},
		{"%d minuti fa", "fra %d minuti"},
		{"un'ora fa", "fra un'ora"},
		{"%d ore fa", "fra %d ore"},
		{"un giorno fa", "fra un giorno"},
		{"%d giorni fa", "fra %d giorni"},
		{"una settimana fa", "fra una settimana"},
		{"%d settimane fa", "fra %d settimane"},
		{"un mese fa", "fra un mese"},
		{"%d mesi fa", "fra %d mesi"},
		{"un anno fa", "fra un anno"},
		{"%d anni fa", "fra %d anni"},
	}[index]
	return res[0], res[1]
}
