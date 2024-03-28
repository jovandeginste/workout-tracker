package units

const (
	// byte ratio constants
	_          = iota
	kb float64 = 1 << (10 * iota)
	mb
	gb
	tb
	pb
	eb
	zb
	yb
)

var (
	Bi   = UnitOptionQuantity("bits")
	Data = UnitOptionQuantity("bytes")

	Byte      = NewUnit("byte", "B", Data)
	KiloByte  = NewUnit("kilobyte", "KB", Data)
	MegaByte  = NewUnit("megabyte", "MB", Data)
	GigaByte  = NewUnit("gigabyte", "GB", Data)
	TeraByte  = NewUnit("terabyte", "TB", Data)
	PetaByte  = NewUnit("petabyte", "PB", Data)
	ExaByte   = NewUnit("exabyte", "", Data)
	ZettaByte = NewUnit("zettabyte", "", Data)
	YottaByte = NewUnit("yottabyte", "", Data)

	Kibibyte = NewUnit("kibibyte", "KiB", Data, IEC)
	Mebibyte = NewUnit("mebibyte", "MiB", Data, IEC)
	Gibibyte = NewUnit("gibibyte", "GiB", Data, IEC)
	Tebibyte = NewUnit("tebibyte", "TiB", Data, IEC)
	Pebibyte = NewUnit("pebibyte", "PiB", Data, IEC)
	Exbibyte = NewUnit("exbibyte", "EiB", Data, IEC)
	Zebibyte = NewUnit("zebibyte", "ZiB", Data, IEC)
	Yobibyte = NewUnit("yobibyte", "YiB", Data, IEC)

	Bit     = NewUnit("bit", "b", Bi)
	KiloBit = Kilo(Bit)
	MegaBit = Mega(Bit)
	GigaBit = Giga(Bit)
	TeraBit = Tera(Bit)
	PetaBit = Peta(Bit)
	ExaBit  = Exa(Bit)

	Nibble = NewUnit("nibble", "", Data)
)

func init() {
	NewRatioConversion(Nibble, Bit, 4.0)
	NewRatioConversion(Byte, Bit, 8.0)

	NewRatioConversion(KiloByte, Byte, 1000)
	NewRatioConversion(MegaByte, Byte, 1000000)
	NewRatioConversion(GigaByte, Byte, 1000000000)
	NewRatioConversion(TeraByte, Byte, 1000000000000)
	NewRatioConversion(PetaByte, Byte, 1000000000000000)
	NewRatioConversion(ExaByte, Byte, 1000000000000000000)
	NewRatioConversion(ZettaByte, Byte, 1000000000000000000000)
	NewRatioConversion(YottaByte, Byte, 1000000000000000000000000)

	NewRatioConversion(Kibibyte, Byte, kb)
	NewRatioConversion(Mebibyte, Byte, mb)
	NewRatioConversion(Gibibyte, Byte, gb)
	NewRatioConversion(Tebibyte, Byte, tb)
	NewRatioConversion(Pebibyte, Byte, pb)
	NewRatioConversion(Exbibyte, Byte, eb)
	NewRatioConversion(Zebibyte, Byte, zb)
	NewRatioConversion(Yobibyte, Byte, yb)
}
