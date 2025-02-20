package countries

// TypeRegionCode for Typer interface
const TypeRegionCode string = "countries.RegionCode"

// TypeRegion for Typer interface
const TypeRegion string = "countries.Region"

const (
	// RegionUnknown - RegionUnknown
	RegionUnknown RegionCode = 0
	// RegionAF - RegionAF
	RegionAF RegionCode = 2
	// RegionNA - RegionNA
	RegionNA RegionCode = 3
	// RegionSA - RegionSA
	RegionSA RegionCode = 5
	// RegionOC - RegionOC
	RegionOC RegionCode = 9
	// RegionAN - RegionAN
	RegionAN RegionCode = 10
	// RegionAS - RegionAS
	RegionAS RegionCode = 142
	// RegionEU - RegionEU
	RegionEU RegionCode = 150
	// RegionNone - RegionNone
	RegionNone RegionCode = RegionCode(None)
)

const (
	// RegionAfrica       RegionCode = 2
	RegionAfrica RegionCode = 2
	// RegionNorthAmerica RegionCode = 3
	RegionNorthAmerica RegionCode = 3
	// RegionSouthAmerica RegionCode = 5
	RegionSouthAmerica RegionCode = 5
	// RegionOceania      RegionCode = 9
	RegionOceania RegionCode = 9
	// RegionAntarctica   RegionCode = 10
	RegionAntarctica RegionCode = 999
	// RegionAsia         RegionCode = 142
	RegionAsia RegionCode = 142
	// RegionEurope       RegionCode = 150
	RegionEurope RegionCode = 150
)
