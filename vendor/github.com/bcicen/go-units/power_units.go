package units

var (
	Power = UnitOptionQuantity("power")

	// metric
	Watt      = NewUnit("watt", "W", Power)
	KiloWatt  = Kilo(Watt)
	MegaWatt  = Mega(Watt)
	GigaWatt  = Giga(Watt)
	TeraWatt  = Tera(Watt)
	PetaWatt  = Peta(Watt)
	ExaWatt   = Exa(Watt)
	ZettaWatt = Zetta(Watt)
	YottaWatt = Yotta(Watt)
)

func init() {
	NewRatioConversion(KiloWatt, Watt, 1000.0)
	NewRatioConversion(MegaWatt, Watt, 1e6)
	NewRatioConversion(GigaWatt, Watt, 1e9)
	NewRatioConversion(TeraWatt, Watt, 1e12)
	NewRatioConversion(PetaWatt, Watt, 1e15)
	NewRatioConversion(ExaWatt, Watt, 1e18)
	NewRatioConversion(ZettaWatt, Watt, 1e21)
	NewRatioConversion(YottaWatt, Watt, 1e24)
}
