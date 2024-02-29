package fit

import (
	"math"
	"strconv"
)

const (
	sint32Invalid = 0x7FFFFFFF
	stringInvalid = "Invalid"
	precision     = 5 // 1.1 m
)

var (
	semiToDegFactor = 180 / math.Pow(2, 31)
	degToSemiFactor = math.Pow(2, 31) / 180
)

// Latitude represents the geographical coordinate latitude.
type Latitude struct {
	semicircles int32
}

// NewLatitude returns a new latitude from a semicircle. If semicircles is
// outside the range of a latitude, (math.MinInt32/2, math.MaxInt32/2) then an
// invalid latitude is returned.
func NewLatitude(semicircles int32) Latitude {
	if semicircles == sint32Invalid {
		return NewLatitudeInvalid()
	}
	if semicircles < math.MinInt32/2 || semicircles > math.MaxInt32/2 {
		return NewLatitudeInvalid()
	}
	return Latitude{semicircles: semicircles}
}

// NewLatitudeDegrees returns a new latitude from a degree. If degrees is
// outside the range of a latitude (+/- 90°) then an invalid latitude is
// returned.
func NewLatitudeDegrees(degrees float64) Latitude {
	if degrees >= 90 || degrees <= -90 {
		return NewLatitudeInvalid()
	}
	return Latitude{semicircles: int32(degrees * degToSemiFactor)}
}

// NewLatitudeInvalid returns an invalid latitude. The underlying storage is
// set to the invalid value of the FIT base type (sint32) used to represent a
// latitude.
func NewLatitudeInvalid() Latitude {
	return Latitude{semicircles: sint32Invalid}
}

// Semicircles returns l in semicircles.
func (l Latitude) Semicircles() int32 {
	return l.semicircles
}

// Degrees returns l in degrees. If l is invalid then NaN is returned.
func (l Latitude) Degrees() float64 {
	if l.semicircles == sint32Invalid {
		return math.NaN()
	}
	return float64(l.semicircles) * semiToDegFactor
}

// Invalid reports whether l represents an invalid latitude.
func (l Latitude) Invalid() bool {
	return l.semicircles == sint32Invalid
}

// String returns a string representation of l in degrees with 5 decimal
// places. If l is invalid then the string "Invalid" is returned.
func (l Latitude) String() string {
	if l.semicircles == sint32Invalid {
		return stringInvalid
	}
	return strconv.FormatFloat(l.Degrees(), 'f', precision, 32)
}

// Longitude represents the geographical coordinate longitude.
type Longitude struct {
	semicircles int32
}

// NewLongitude returns a new longitude from a semicircle.
func NewLongitude(semicircles int32) Longitude {
	return Longitude{semicircles: semicircles}
}

// NewLongitudeDegrees returns a new longitude from a degree. If degrees is
// outside the range of a longitude (+/- 180°) then an invalid longitude is
// returned.
func NewLongitudeDegrees(degrees float64) Longitude {
	if degrees >= 180 || degrees <= -180 {
		return Longitude{semicircles: sint32Invalid}
	}
	return Longitude{semicircles: int32(degrees * degToSemiFactor)}
}

// NewLongitudeInvalid returns an invalid longitude. The underlying storage is
// set to the invalid value of the FIT base type (sint32) used to represent a
// longitude.
func NewLongitudeInvalid() Longitude {
	return Longitude{semicircles: sint32Invalid}
}

// Semicircles returns l in semicircles.
func (l Longitude) Semicircles() int32 {
	return l.semicircles
}

// Degrees returns l in degrees. If l is invalid then NaN is returned.
func (l Longitude) Degrees() float64 {
	if l.semicircles == sint32Invalid {
		return math.NaN()
	}
	return float64(l.semicircles) * semiToDegFactor
}

// Invalid reports whether l represents an invalid longitude.
func (l Longitude) Invalid() bool {
	return l.semicircles == sint32Invalid
}

// String returns a string representation of l in degrees with 5 decimal
// places. If l is invalid then the string "Invalid" is returned.
func (l Longitude) String() string {
	if l.semicircles == sint32Invalid {
		return stringInvalid
	}
	return strconv.FormatFloat(l.Degrees(), 'f', precision, 32)
}
