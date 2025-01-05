package timeago

var secArray = [6]float64{60, // 60 seconds in 1 min
	60,           // 60 mins in 1 hour
	24,           // 24 hours in 1 day
	7,            // 7 days in 1 week
	365 / 7 / 12, // 4.345238095238096 weeks in 1 month
	12,           // 12 months in 1 year
}

type LocaleFunction func(number float64, index int) (ago string, in string)

var localeFunctions = map[string]LocaleFunction{
	"ar":      arLocale,
	"be":      beLocale,
	"bg":      bgLocale,
	"bn_IN":   bnLocale,
	"bn":      bnLocale,
	"ca":      caLocale,
	"cs":      csLocale,
	"da":      daLocale,
	"de":      deLocale,
	"el":      elLocale,
	"en":      enLocale,
	"es":      esLocale,
	"eu":      euLocale,
	"fa":      faLocale,
	"fi":      fiLocale,
	"fr":      frLocale,
	"gl":      glLocale,
	"he":      heLocale,
	"hi_IN":   hiLocale,
	"hi":      hiLocale,
	"hu":      huLocale,
	"id_ID":   idLocale,
	"id":      idLocale,
	"it":      itLocale,
	"ja":      jaLocale,
	"ka":      kaLocale,
	"ko":      koLocale,
	"ml":      mlLocale,
	"my":      myLocale,
	"nb_NO":   nbLocale,
	"nb":      nbLocale,
	"nl":      nlLocale,
	"nn_NO":   nnLocale,
	"nn":      nnLocale,
	"oc":      ocLocale,
	"pl":      plLocale,
	"pt_BR":   ptLocale,
	"pt":      ptLocale,
	"ro":      roLocale,
	"ru":      ruLocale,
	"sq":      sqLocale,
	"sr":      srLocale,
	"sv":      svLocale,
	"ta":      taLocale,
	"th":      thLocale,
	"tk":      tkLocale,
	"tr":      trLocale,
	"uk":      ukLocale,
	"vi":      viLocale,
	"zh_CN":   zhCNLocale,
	"zh_Hans": zhCNLocale,
	"zh":      zhCNLocale,
	"zh_TW":   zhTWLocale,
	"zh_Hant": zhCNLocale,
}
