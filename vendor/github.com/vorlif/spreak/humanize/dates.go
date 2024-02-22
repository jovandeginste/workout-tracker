package humanize

var (
	weekdays = map[int]string{
		0: "Monday",
		1: "Tuesday",
		2: "Wednesday",
		3: "Thursday",
		4: "Friday",
		5: "Saturday",
		6: "Sunday",
	}

	weekdaysAbbr = map[int]string{
		0: "Mon",
		1: "Tue",
		2: "Wed",
		3: "Thu",
		4: "Fri",
		5: "Sat",
		6: "Sun",
	}

	months = map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	months3 = map[int]string{
		1:  "jan",
		2:  "feb",
		3:  "mar",
		4:  "apr",
		5:  "may",
		6:  "jun",
		7:  "jul",
		8:  "aug",
		9:  "sep",
		10: "oct",
		11: "nov",
		12: "dec",
	}

	monthsAp = map[int]gettextEntry{ // month names in Associated Press style
		1:  {"abbrev. month", "Jan.", ""},
		2:  {"abbrev. month", "Feb.", ""},
		3:  {"abbrev. month", "March", ""},
		4:  {"abbrev. month", "April", ""},
		5:  {"abbrev. month", "May", ""},
		6:  {"abbrev. month", "June", ""},
		7:  {"abbrev. month", "July", ""},
		8:  {"abbrev. month", "Aug.", ""},
		9:  {"abbrev. month", "Sept.", ""},
		10: {"abbrev. month", "Oct.", ""},
		11: {"abbrev. month", "Nov.", ""},
		12: {"abbrev. month", "Dec.", ""},
	}

	monthsAlt = map[int]gettextEntry{ // required for long date representation by some locales
		1:  {"alt. month", "January", ""},
		2:  {"alt. month", "February", ""},
		3:  {"alt. month", "March", ""},
		4:  {"alt. month", "April", ""},
		5:  {"alt. month", "May", ""},
		6:  {"alt. month", "June", ""},
		7:  {"alt. month", "July", ""},
		8:  {"alt. month", "August", ""},
		9:  {"alt. month", "September", ""},
		10: {"alt. month", "October", ""},
		11: {"alt. month", "November", ""},
		12: {"alt. month", "December", ""},
	}
)
