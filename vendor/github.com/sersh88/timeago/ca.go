package timeago

var caTimeTypes = [][]string{
	{"fa un moment", "d'aquí un moment"},
	{"fa %d segons", "d'aquí %d segons"},
	{"fa 1 minut", "d'aquí 1 minut"},
	{"fa %d minuts", "d'aquí %d minuts"},
	{"fa 1 hora", "d'aquí 1 hora"},
	{"fa %d hores", "d'aquí %d hores"},
	{"fa 1 dia", "d'aquí 1 dia"},
	{"fa %d dies", "d'aquí %d dies"},
	{"fa 1 setmana", "d'aquí 1 setmana"},
	{"fa %d setmanes", "d'aquí %d setmanes"},
	{"fa 1 mes", "d'aquí 1 mes"},
	{"fa %d mesos", "d'aquí %d mesos"},
	{"fa 1 any", "d'aquí 1 any"},
	{"fa %d anys", "d'aquí %d anys"},
}

func caLocale(_ float64, idx int) (ago string, in string) {
	var res = caTimeTypes[idx]
	return res[0], res[1]
}
