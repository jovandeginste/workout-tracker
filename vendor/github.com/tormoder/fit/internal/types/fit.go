package types

import "fmt"

type Kind byte

const (
	NativeFit Kind = 0x00 // Standard 	-> Fit base type/alias
	TimeUTC   Kind = 0x01 // Time UTC 	-> time.Time
	TimeLocal Kind = 0x02 // Time Local 	-> time.Time with Location
	Lat       Kind = 0x03 // Latitude 	-> fit.Latitude
	Lng       Kind = 0x04 // Longitude 	-> fit.Longitude
)

func (k Kind) String() string {
	if k > 0x04 {
		return fmt.Sprintf("unknown kind (%d)", k)
	}
	return kname[k]
}

var kname = [...]string{
	"NativeFit",
	"TimeUTC",
	"TimeLocal",
	"Lat",
	"Lng",
}

func Make(kind Kind, array bool) Fit {
	var f Fit
	if array {
		f = f.setArray()
	}
	if kind == NativeFit {
		return f
	}
	f = f.setKind(kind)
	switch kind {
	case TimeUTC, TimeLocal:
		f = f.setBase(BaseUint32)
	case Lat, Lng:
		f = f.setBase(BaseSint32)
	}
	return f
}

func MakeNative(b Base, array bool) Fit {
	var f Fit
	f = f.setBase(b)
	if array {
		f = f.setArray()
	}
	return f
}

// Bit-packing layout:
//
//	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
//	| bF | bE | bD | bC | bB | bA | b9 | b8 | b7 | b6 | b5 | b4 | b3 | b2 | b1 | b0 |
//	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
//	\_________________________________/\______________/\__/\________________________/
//
//	               Unused                    Kind      Array       Base type
type Fit uint16

func (f Fit) setKind(k Kind) Fit {
	return Fit(uint16(f) | uint16(k)<<6)
}

func (f Fit) setArray() Fit {
	f = f | 0x20
	return f
}

func (f Fit) setBase(b Base) Fit {
	return Fit(uint16(f) | uint16(b&0x1F))
}

func (f Fit) Kind() Kind {
	return Kind((f & 0x1C0) >> 6)
}

func (f Fit) Array() bool {
	return (f&0x20)>>5 == 1
}

func (f Fit) BaseType() Base {
	return decompress(byte(f))
}

func (f Fit) Valid() bool {
	return int(f.Kind()) < len(fgotype) && f.BaseType().Known()
}

func (f Fit) GoInvalidValue() string {
	if !f.Valid() {
		return "invalid type: " + f.String()
	}
	if f.Array() {
		return "nil"
	}
	if f.Kind() == NativeFit {
		return f.BaseType().GoInvalidValue()
	}
	return fgoinvalid[f.Kind()]
}

func (f Fit) GoType() string {
	if !f.Valid() {
		return "invalid type: " + f.String()
	}
	var gt string
	if f.Kind() == NativeFit {
		gt = f.BaseType().GoType()
	} else {
		gt = fgotype[f.Kind()]
	}
	if f.Array() {
		return "[]" + gt
	}
	return gt
}

func (f Fit) String() string {
	return fmt.Sprintf(
		"kind: %v |  base type: %v | array: %t",
		f.Kind(), f.BaseType(), f.Array(),
	)
}

func (f Fit) ValueString() string {
	return fmt.Sprintf("types.Fit(%d)", f)
}

var fgoinvalid = [...]string{
	"",
	"timeBase",
	"timeBase",
	"NewLatitudeInvalid()",
	"NewLongitudeInvalid()",
}

var fgotype = [...]string{
	"",
	"time.Time",
	"time.Time",
	"Latitude",
	"Longitude",
}
