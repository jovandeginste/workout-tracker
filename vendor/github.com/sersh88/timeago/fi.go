package timeago

func fiLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"juuri äsken", "juuri nyt"},
		{"%d sekuntia sitten", "%d sekunnin päästä"},
		{"minuutti sitten", "minuutin päästä"},
		{"%d minuuttia sitten", "%d minuutin päästä"},
		{"tunti sitten", "tunnin päästä"},
		{"%d tuntia sitten", "%d tunnin päästä"},
		{"päivä sitten", "päivän päästä"},
		{"%d päivää sitten", "%d päivän päästä"},
		{"viikko sitten", "viikon päästä"},
		{"%d viikkoa sitten", "%d viikon päästä"},
		{"kuukausi sitten", "kuukauden päästä"},
		{"%d kuukautta sitten", "%d kuukauden päästä"},
		{"vuosi sitten", "vuoden päästä"},
		{"%d vuotta sitten", "%d vuoden päästä"},
	}[index]
	return res[0], res[1]
}
