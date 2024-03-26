package units

var (
	Time = UnitOptionQuantity("time")

	Second      = NewUnit("second", "s", Time)
	ExaSecond   = Exa(Second)
	PetaSecond  = Peta(Second)
	TeraSecond  = Tera(Second)
	GigaSecond  = Giga(Second)
	MegaSecond  = Mega(Second)
	KiloSecond  = Kilo(Second)
	HectoSecond = Hecto(Second)
	DecaSecond  = Deca(Second)
	DeciSecond  = Deci(Second)
	CentiSecond = Centi(Second)
	MilliSecond = Milli(Second)
	MicroSecond = Micro(Second)
	NanoSecond  = Nano(Second)
	PicoSecond  = Pico(Second)
	FemtoSecond = Femto(Second)
	AttoSecond  = Atto(Second)

	Minute = NewUnit("minute", "min", Time)
	Hour   = NewUnit("hour", "hr", Time)
	Day    = NewUnit("day", "d", Time)
	Month  = NewUnit("month", "", Time)
	Year   = NewUnit("year", "yr", Time)

	Decade     = NewUnit("decade", "", Time)
	Century    = NewUnit("century", "", Time)
	Millennium = NewUnit("millennium", "", Time)

	// more esoteric time units
	PlanckTime = NewUnit("planck time", "𝑡ₚ", Time)
	Fortnight  = NewUnit("fortnight", "", Time)
	Score      = NewUnit("score", "", Time)
)

func init() {
	NewRatioConversion(Minute, Second, 60.0)
	NewRatioConversion(Hour, Second, 3600.0)
	NewRatioConversion(Day, Hour, 24.0)
	NewRatioConversion(Month, Day, 30.0)
	NewRatioConversion(Year, Day, 365.25)

	NewRatioConversion(Decade, Year, 10.0)
	NewRatioConversion(Century, Year, 100.0)
	NewRatioConversion(Millennium, Year, 1000.0)

	NewRatioConversion(PlanckTime, Second, 5.39e-44)
	NewRatioConversion(Fortnight, Day, 14)
	NewRatioConversion(Score, Year, 20.0)
}
