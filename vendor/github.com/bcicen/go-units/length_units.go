package units

var (
	Length = UnitOptionQuantity("length")

	// metric
	Meter      = NewUnit("meter", "m", Length, SI, UnitOptionAliases("metre"))
	ExaMeter   = Exa(Meter)
	PetaMeter  = Peta(Meter)
	TeraMeter  = Tera(Meter)
	GigaMeter  = Giga(Meter)
	MegaMeter  = Mega(Meter)
	KiloMeter  = Kilo(Meter)
	HectoMeter = Hecto(Meter)
	DecaMeter  = Deca(Meter)
	DeciMeter  = Deci(Meter)
	CentiMeter = Centi(Meter)
	MilliMeter = Milli(Meter)
	MicroMeter = Micro(Meter)
	NanoMeter  = Nano(Meter)
	PicoMeter  = Pico(Meter)
	FemtoMeter = Femto(Meter)
	AttoMeter  = Atto(Meter)

	Angstrom = NewUnit("angstrom", "Å", Length, BI, UnitOptionPlural("angstroms"))
	Inch     = NewUnit("inch", "in", Length, BI, UnitOptionPlural("inches"))
	Foot     = NewUnit("foot", "ft", Length, BI, UnitOptionPlural("feet"))
	Yard     = NewUnit("yard", "yd", Length, BI)
	Mile     = NewUnit("mile", "mi", Length, BI)
	League   = NewUnit("league", "lea", Length, BI)
	Furlong  = NewUnit("furlong", "fur", Length, BI)
)

func init() {
	NewRatioConversion(Angstrom, Meter, 0.0000000001)
	NewRatioConversion(Inch, Meter, 0.0254)
	NewRatioConversion(Foot, Meter, 0.3048)
	NewRatioConversion(Yard, Meter, 0.9144)
	NewRatioConversion(Mile, Meter, 1609.344)
	NewRatioConversion(League, Meter, 4828.032)
	NewRatioConversion(Furlong, Meter, 201.168)
}
