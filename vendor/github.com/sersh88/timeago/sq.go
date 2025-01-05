package timeago

func sqLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"pak më parë", "pas pak"},
		{"para %d sekondash", "pas %d sekondash"},
		{"para një minute", "pas një minute"},
		{"para %d minutash", "pas %d minutash"},
		{"para një ore", "pas një ore"},
		{"para %d orësh", "pas %d orësh"},
		{"dje", "nesër"},
		{"para %d ditësh", "pas %d ditësh"},
		{"para një jave", "pas një jave"},
		{"para %d javësh", "pas %d javësh"},
		{"para një muaji", "pas një muaji"},
		{"para %d muajsh", "pas %d muajsh"},
		{"para një viti", "pas një viti"},
		{"para %d vjetësh", "pas %d vjetësh"},
	}[index]
	return res[0], res[1]
}
