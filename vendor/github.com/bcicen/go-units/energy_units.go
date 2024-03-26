package units

var (
	Energy = UnitOptionQuantity("energy")

	// metric
	Joule      = NewUnit("joule", "J", Energy)
	KiloJoule  = Kilo(Joule)
	MegaJoule  = Mega(Joule)
	GigaJoule  = Giga(Joule)
	TeraJoule  = Tera(Joule)
	PetaJoule  = Peta(Joule)
	ExaJoule   = Exa(Joule)
	ZettaJoule = Zetta(Joule)
	YottaJoule = Yotta(Joule)
	MilliJoule = Milli(Joule)
	MicroJoule = Micro(Joule)
	NanoJoule  = Nano(Joule)
	PicoJoule  = Pico(Joule)
	FemtoJoule = Femto(Joule)
	AttoJoule  = Atto(Joule)

	WattHour     = NewUnit("watt-hour", "Wh", Energy)
	KiloWattHour = Kilo(WattHour)
	GigaWattHour = Giga(WattHour)
	MegaWattHour = Mega(WattHour)
	TeraWattHour = Tera(WattHour)
	PetaWattHour = Peta(WattHour)

	// other
	ElectronVolt = NewUnit("electronvolt", "eV", Energy)
	Calorie      = NewUnit("calorie", "cal", Energy)
)

func init() {
	// Metric to Metric conversions
	NewRatioConversion(KiloJoule, Joule, 1000.0)
	NewRatioConversion(MegaJoule, Joule, 1e6)
	NewRatioConversion(GigaJoule, Joule, 1e9)
	NewRatioConversion(TeraJoule, Joule, 1e12)
	NewRatioConversion(PetaJoule, Joule, 1e15)
	NewRatioConversion(ExaJoule, Joule, 1e18)
	NewRatioConversion(ZettaJoule, Joule, 1e21)
	NewRatioConversion(YottaJoule, Joule, 1e24)
	NewRatioConversion(MilliJoule, Joule, 0.001)
	NewRatioConversion(MicroJoule, Joule, 1e-6)
	NewRatioConversion(NanoJoule, Joule, 1e-9)
	NewRatioConversion(PicoJoule, Joule, 1e-12)
	NewRatioConversion(FemtoJoule, Joule, 1e-15)
	NewRatioConversion(AttoJoule, Joule, 1e-18)

	// Non-metric to Metric conversions
	NewRatioConversion(ElectronVolt, Joule, 1.60218e-19)
	NewRatioConversion(Calorie, Joule, 4.184)
	NewRatioConversion(WattHour, Joule, 3600)
	NewRatioConversion(MegaWattHour, Joule, 3600*1e3)
	NewRatioConversion(KiloWattHour, Watt, 3600*1e6)
	NewRatioConversion(GigaWattHour, Watt, 3600*1e9)
	NewRatioConversion(TeraWattHour, Watt, 3600*1e12)
	NewRatioConversion(PetaWattHour, Watt, 3600*1e14)
}
