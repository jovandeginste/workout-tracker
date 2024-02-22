package humanize

type LocaleInfo struct {
	Code      string
	Name      string
	NameLocal string
}

var LocaleInfos = map[string]LocaleInfo{
	"af": {
		Code:      "af",
		Name:      "Afrikaans",
		NameLocal: "Afrikaans",
	},
	"ar": {
		Code:      "ar",
		Name:      "Arabic",
		NameLocal: "العربيّة",
	},
	"ar-dz": {
		Code:      "ar-dz",
		Name:      "Algerian Arabic",
		NameLocal: "العربية الجزائرية",
	},
	"ast": {
		Code:      "ast",
		Name:      "Asturian",
		NameLocal: "asturianu",
	},
	"az": {
		Code:      "az",
		Name:      "Azerbaijani",
		NameLocal: "Azərbaycanca",
	},
	"be": {
		Code:      "be",
		Name:      "Belarusian",
		NameLocal: "беларуская",
	},
	"bg": {
		Code:      "bg",
		Name:      "Bulgarian",
		NameLocal: "български",
	},
	"bn": {
		Code:      "bn",
		Name:      "Bengali",
		NameLocal: "বাংলা",
	},
	"br": {
		Code:      "br",
		Name:      "Breton",
		NameLocal: "brezhoneg",
	},
	"bs": {
		Code:      "bs",
		Name:      "Bosnian",
		NameLocal: "bosanski",
	},
	"ca": {
		Code:      "ca",
		Name:      "Catalan",
		NameLocal: "català",
	},
	"ckb": {
		Code:      "ckb",
		Name:      "Central Kurdish (Sorani)",
		NameLocal: "کوردی",
	},
	"cs": {
		Code:      "cs",
		Name:      "Czech",
		NameLocal: "česky",
	},
	"cy": {
		Code:      "cy",
		Name:      "Welsh",
		NameLocal: "Cymraeg",
	},
	"da": {
		Code:      "da",
		Name:      "Danish",
		NameLocal: "dansk",
	},
	"de": {
		Code:      "de",
		Name:      "German",
		NameLocal: "Deutsch",
	},
	"dsb": {
		Code:      "dsb",
		Name:      "Lower Sorbian",
		NameLocal: "dolnoserbski",
	},
	"el": {
		Code:      "el",
		Name:      "Greek",
		NameLocal: "Ελληνικά",
	},
	"en": {
		Code:      "en",
		Name:      "English",
		NameLocal: "English",
	},
	"en-au": {
		Code:      "en-au",
		Name:      "Australian English",
		NameLocal: "Australian English",
	},
	"en-gb": {
		Code:      "en-gb",
		Name:      "British English",
		NameLocal: "British English",
	},
	"eo": {
		Code:      "eo",
		Name:      "Esperanto",
		NameLocal: "Esperanto",
	},
	"es": {
		Code:      "es",
		Name:      "Spanish",
		NameLocal: "español",
	},
	"es-ar": {
		Code:      "es-ar",
		Name:      "Argentinian Spanish",
		NameLocal: "español de Argentina",
	},
	"es-co": {
		Code:      "es-co",
		Name:      "Colombian Spanish",
		NameLocal: "español de Colombia",
	},
	"es-mx": {
		Code:      "es-mx",
		Name:      "Mexican Spanish",
		NameLocal: "español de Mexico",
	},
	"es-ni": {
		Code:      "es-ni",
		Name:      "Nicaraguan Spanish",
		NameLocal: "español de Nicaragua",
	},
	"es-ve": {
		Code:      "es-ve",
		Name:      "Venezuelan Spanish",
		NameLocal: "español de Venezuela",
	},
	"et": {
		Code:      "et",
		Name:      "Estonian",
		NameLocal: "eesti",
	},
	"eu": {
		Code:      "eu",
		Name:      "Basque",
		NameLocal: "Basque",
	},
	"fa": {
		Code:      "fa",
		Name:      "Persian",
		NameLocal: "فارسی",
	},
	"fi": {
		Code:      "fi",
		Name:      "Finnish",
		NameLocal: "suomi",
	},
	"fr": {
		Code:      "fr",
		Name:      "French",
		NameLocal: "français",
	},
	"fy": {
		Code:      "fy",
		Name:      "Frisian",
		NameLocal: "frysk",
	},
	"ga": {
		Code:      "ga",
		Name:      "Irish",
		NameLocal: "Gaeilge",
	},
	"gd": {
		Code:      "gd",
		Name:      "Scottish Gaelic",
		NameLocal: "Gàidhlig",
	},
	"gl": {
		Code:      "gl",
		Name:      "Galician",
		NameLocal: "galego",
	},
	"he": {
		Code:      "he",
		Name:      "Hebrew",
		NameLocal: "עברית",
	},
	"hi": {
		Code:      "hi",
		Name:      "Hindi",
		NameLocal: "हिंदी",
	},
	"hr": {
		Code:      "hr",
		Name:      "Croatian",
		NameLocal: "Hrvatski",
	},
	"hsb": {
		Code:      "hsb",
		Name:      "Upper Sorbian",
		NameLocal: "hornjoserbsce",
	},
	"hu": {
		Code:      "hu",
		Name:      "Hungarian",
		NameLocal: "Magyar",
	},
	"hy": {
		Code:      "hy",
		Name:      "Armenian",
		NameLocal: "հայերեն",
	},
	"ia": {
		Code:      "ia",
		Name:      "Interlingua",
		NameLocal: "Interlingua",
	},
	"io": {
		Code:      "io",
		Name:      "Ido",
		NameLocal: "ido",
	},
	"id": {
		Code:      "id",
		Name:      "Indonesian",
		NameLocal: "Bahasa Indonesia",
	},
	"ig": {
		Code:      "ig",
		Name:      "Igbo",
		NameLocal: "Asụsụ Ìgbò",
	},
	"is": {
		Code:      "is",
		Name:      "Icelandic",
		NameLocal: "Íslenska",
	},
	"it": {
		Code:      "it",
		Name:      "Italian",
		NameLocal: "italiano",
	},
	"ja": {
		Code:      "ja",
		Name:      "Japanese",
		NameLocal: "日本語",
	},
	"ka": {
		Code:      "ka",
		Name:      "Georgian",
		NameLocal: "ქართული",
	},
	"kab": {
		Code:      "kab",
		Name:      "Kabyle",
		NameLocal: "taqbaylit",
	},
	"kk": {
		Code:      "kk",
		Name:      "Kazakh",
		NameLocal: "Қазақ",
	},
	"km": {
		Code:      "km",
		Name:      "Khmer",
		NameLocal: "Khmer",
	},
	"kn": {
		Code:      "kn",
		Name:      "Kannada",
		NameLocal: "Kannada",
	},
	"ko": {
		Code:      "ko",
		Name:      "Korean",
		NameLocal: "한국어",
	},
	"ky": {
		Code:      "ky",
		Name:      "Kyrgyz",
		NameLocal: "Кыргызча",
	},
	"lb": {
		Code:      "lb",
		Name:      "Luxembourgish",
		NameLocal: "Lëtzebuergesch",
	},
	"lt": {
		Code:      "lt",
		Name:      "Lithuanian",
		NameLocal: "Lietuviškai",
	},
	"lv": {
		Code:      "lv",
		Name:      "Latvian",
		NameLocal: "latviešu",
	},
	"mk": {
		Code:      "mk",
		Name:      "Macedonian",
		NameLocal: "Македонски",
	},
	"ml": {
		Code:      "ml",
		Name:      "Malayalam",
		NameLocal: "മലയാളം",
	},
	"mn": {
		Code:      "mn",
		Name:      "Mongolian",
		NameLocal: "Mongolian",
	},
	"mr": {
		Code:      "mr",
		Name:      "Marathi",
		NameLocal: "मराठी",
	},
	"ms": {
		Code:      "ms",
		Name:      "Malay",
		NameLocal: "Bahasa Melayu",
	},
	"my": {
		Code:      "my",
		Name:      "Burmese",
		NameLocal: "မြန်မာဘာသာ",
	},
	"nb": {
		Code:      "nb",
		Name:      "Norwegian Bokmal",
		NameLocal: "norsk (bokmål)",
	},
	"ne": {
		Code:      "ne",
		Name:      "Nepali",
		NameLocal: "नेपाली",
	},
	"nl": {
		Code:      "nl",
		Name:      "Dutch",
		NameLocal: "Nederlands",
	},
	"nn": {
		Code:      "nn",
		Name:      "Norwegian Nynorsk",
		NameLocal: "norsk (nynorsk)",
	},
	"no": {
		Code:      "no",
		Name:      "Norwegian",
		NameLocal: "norsk",
	},
	"os": {
		Code:      "os",
		Name:      "Ossetic",
		NameLocal: "Ирон",
	},
	"pa": {
		Code:      "pa",
		Name:      "Punjabi",
		NameLocal: "Punjabi",
	},
	"pl": {
		Code:      "pl",
		Name:      "Polish",
		NameLocal: "polski",
	},
	"pt": {
		Code:      "pt",
		Name:      "Portuguese",
		NameLocal: "Português",
	},
	"pt-br": {
		Code:      "pt-br",
		Name:      "Brazilian Portuguese",
		NameLocal: "Português Brasileiro",
	},
	"ro": {
		Code:      "ro",
		Name:      "Romanian",
		NameLocal: "Română",
	},
	"ru": {
		Code:      "ru",
		Name:      "Russian",
		NameLocal: "Русский",
	},
	"sk": {
		Code:      "sk",
		Name:      "Slovak",
		NameLocal: "Slovensky",
	},
	"sl": {
		Code:      "sl",
		Name:      "Slovenian",
		NameLocal: "Slovenščina",
	},
	"sq": {
		Code:      "sq",
		Name:      "Albanian",
		NameLocal: "shqip",
	},
	"sr": {
		Code:      "sr",
		Name:      "Serbian",
		NameLocal: "српски",
	},
	"sr-latn": {
		Code:      "sr-latn",
		Name:      "Serbian Latin",
		NameLocal: "srpski (latinica)",
	},
	"sv": {
		Code:      "sv",
		Name:      "Swedish",
		NameLocal: "svenska",
	},
	"sw": {
		Code:      "sw",
		Name:      "Swahili",
		NameLocal: "Kiswahili",
	},
	"ta": {
		Code:      "ta",
		Name:      "Tamil",
		NameLocal: "தமிழ்",
	},
	"te": {
		Code:      "te",
		Name:      "Telugu",
		NameLocal: "తెలుగు",
	},
	"tg": {
		Code:      "tg",
		Name:      "Tajik",
		NameLocal: "тоҷикӣ",
	},
	"th": {
		Code:      "th",
		Name:      "Thai",
		NameLocal: "ภาษาไทย",
	},
	"tk": {
		Code:      "tk",
		Name:      "Turkmen",
		NameLocal: "Türkmençe",
	},
	"tr": {
		Code:      "tr",
		Name:      "Turkish",
		NameLocal: "Türkçe",
	},
	"tt": {
		Code:      "tt",
		Name:      "Tatar",
		NameLocal: "Татарча",
	},
	"udm": {
		Code:      "udm",
		Name:      "Udmurt",
		NameLocal: "Удмурт",
	},
	"ug": {
		Code:      "ug",
		Name:      "Uyghur",
		NameLocal: "ئۇيغۇرچە",
	},
	"uk": {
		Code:      "uk",
		Name:      "Ukrainian",
		NameLocal: "Українська",
	},
	"ur": {
		Code:      "ur",
		Name:      "Urdu",
		NameLocal: "اردو",
	},
	"uz": {
		Code:      "uz",
		Name:      "Uzbek",
		NameLocal: "oʻzbek tili",
	},
	"vi": {
		Code:      "vi",
		Name:      "Vietnamese",
		NameLocal: "Tiếng Việt",
	},
	"zh-hans": {
		Code:      "zh-hans",
		Name:      "Simplified Chinese",
		NameLocal: "简体中文",
	},
	"zh-hant": {
		Code:      "zh-hant",
		Name:      "Traditional Chinese",
		NameLocal: "繁體中文",
	},
}
