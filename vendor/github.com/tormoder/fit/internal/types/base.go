// Package types provides the set of base types defined by the FIT protocol.
// Users of this package must validate base types parsed from a raw byte using
// the Known method before calling any other methods (except String).
package types

import (
	"fmt"
	"math"
)

const (
	pkg           = "types"
	typeNumMask   = 0x1F
	multiByteFlag = 0x80
)

type Base byte

// Base types for fit data.
// The base 5 bits increase by 1 for each definition with all multi-byte number types having the MSB set.
const (
	BaseEnum    Base = 0x00
	BaseSint8   Base = 0x01 // 2's complement format
	BaseUint8   Base = 0x02
	BaseSint16  Base = 0x83 // 2's complement format
	BaseUint16  Base = 0x84
	BaseSint32  Base = 0x85 // 2's complement format
	BaseUint32  Base = 0x86
	BaseString  Base = 0x07 // Null terminated string encoded in UTF-8
	BaseFloat32 Base = 0x88
	BaseFloat64 Base = 0x89
	BaseUint8z  Base = 0x0A
	BaseUint16z Base = 0x8B
	BaseUint32z Base = 0x8C
	BaseByte    Base = 0x0D // Array of bytes. Field is invalid if all bytes are invalid
	BaseSint64  Base = 0x8E // 2's complement format
	BaseUint64  Base = 0x8F
	BaseUint64z Base = 0x90
)

// Internal compressed representation of certain base types.
// With this, we can fit all base types in 5 bits as opposed to 8 bits.
// Base types should be decompressed before use.
const (
	compressedSint16  byte = 0x03
	compressedUint16  byte = 0x04
	compressedSint32  byte = 0x05
	compressedUint32  byte = 0x06
	compressedFloat32 byte = 0x08
	compressedFloat64 byte = 0x09
	compressedUint16z byte = 0x0B
	compressedUint32z byte = 0x0C
	compressedSint64  byte = 0x0E
	compressedUint64  byte = 0x0F
	compressedUint64z byte = 0x10
)

func decompress(b byte) Base {
	b = b & typeNumMask
	switch b {
	case compressedSint16:
		return BaseSint16
	case compressedUint16:
		return BaseUint16
	case compressedSint32:
		return BaseSint32
	case compressedUint32:
		return BaseUint32
	case compressedFloat32:
		return BaseFloat32
	case compressedFloat64:
		return BaseFloat64
	case compressedUint16z:
		return BaseUint16z
	case compressedUint32z:
		return BaseUint32z
	case compressedSint64:
		return BaseSint64
	case compressedUint64:
		return BaseUint64
	case compressedUint64z:
		return BaseUint64z
	default:
		return Base(b)
	}
}

func (t Base) index() byte {
	return byte(t) & typeNumMask
}

func (t Base) multibyte() bool {
	return (byte(t) & multiByteFlag) == multiByteFlag
}

func (t Base) Float() bool {
	return !t.Integer() && t.Signed()
}

func (t Base) GoInvalidValue() string {
	return binvalid[t.index()]
}

func (t Base) GoType() string {
	return bgotype[t.index()]
}

func (t Base) Integer() bool {
	return binteger[t.index()]
}

func (t Base) Known() bool {
	return int(t.index()) < len(bname) && (t.multibyte() == (t.Size() > 1))
}

func (t Base) PkgString() string {
	return pkg + "." + t.String()
}

func (t Base) Signed() bool {
	return bsigned[t.index()]
}

func (t Base) Size() int {
	return bsize[t.index()]
}

func (t Base) String() string {
	if t.Known() {
		return bname[t.index()]
	}
	return fmt.Sprintf("unknown (0x%X)", byte(t))
}

func (t Base) Invalid() interface{} {
	if t.Known() {
		return goinvalid[t.index()]
	}
	return fmt.Sprintf("unknown (0x%X)", byte(t))
}

var bsize = [...]int{
	1,
	1,
	1,
	2,
	2,
	4,
	4,
	1,
	4,
	8,
	1,
	2,
	4,
	1,
	8,
	8,
	8,
}

var bname = [...]string{
	"BaseEnum",
	"BaseSint8",
	"BaseUint8",
	"BaseSint16",
	"BaseUint16",
	"BaseSint32",
	"BaseUint32",
	"BaseString",
	"BaseFloat32",
	"BaseFloat64",
	"BaseUint8z",
	"BaseUint16z",
	"BaseUint32z",
	"BaseByte",
	"BaseSint64",
	"BaseUint64",
	"BaseUint64z",
}

var binteger = [...]bool{
	false,
	true,
	true,
	true,
	true,
	true,
	true,
	false,
	false,
	false,
	true,
	true,
	true,
	false,
	true,
	true,
	true,
}

var bsigned = [...]bool{
	false,
	true,
	false,
	true,
	false,
	true,
	false,
	false,
	true,
	true,
	false,
	false,
	false,
	false,
	true,
	false,
	false,
}

var bgotype = [...]string{
	"byte",
	"int8",
	"uint8",
	"int16",
	"uint16",
	"int32",
	"uint32",
	"string",
	"float32",
	"float64",
	"uint8",
	"uint16",
	"uint32",
	"byte",
	"int64",
	"uint64",
	"uint64",
}

var binvalid = [...]string{
	"0xFF",
	"0x7F",
	"0xFF",
	"0x7FFF",
	"0xFFFF",
	"0x7FFFFFFF",
	"0xFFFFFFFF",
	"\"\"",
	"0xFFFFFFFF",
	"0xFFFFFFFFFFFFFFFF",
	"0x00",
	"0x0000",
	"0x00000000",
	"0xFF",
	"0x7FFFFFFFFFFFFFFF",
	"0xFFFFFFFFFFFFFFFF",
	"0x0000000000000000",
}

var goinvalid = [...]interface{}{
	byte(0xFF),
	int8(0x7F),
	uint8(0xFF),
	int16(0x7FFF),
	uint16(0xFFFF),
	int32(0x7FFFFFFF),
	uint32(0xFFFFFFFF),
	string(""),
	math.Float32frombits(0xFFFFFFFF),
	math.Float64frombits(0xFFFFFFFFFFFFFFFF),
	uint8(0x00),
	uint16(0x0000),
	uint32(0x00000000),
	byte(0xFF),
	int64(0x7FFFFFFFFFFFFFFF),
	uint64(0xFFFFFFFFFFFFFFFF),
	uint64(0x0000000000000000),
}

func BaseFromString(s string) (Base, error) {
	t, found := baseStringToType[s]
	if !found {
		return 0xFF, fmt.Errorf("no base type found for string: %q", s)
	}
	return t, nil
}

var baseStringToType = map[string]Base{
	"enum":    BaseEnum,
	"sint8":   BaseSint8,
	"uint8":   BaseUint8,
	"sint16":  BaseSint16,
	"uint16":  BaseUint16,
	"sint32":  BaseSint32,
	"uint32":  BaseUint32,
	"string":  BaseString,
	"float32": BaseFloat32,
	"float64": BaseFloat64,
	"uint8z":  BaseUint8z,
	"uint16z": BaseUint16z,
	"uint32z": BaseUint32z,
	"byte":    BaseByte,
	"sint64":  BaseSint64,
	"uint64":  BaseUint64,
	"uint64z": BaseUint64z,

	// Typo in SDK 20.14:
	"unit8": BaseUint8,
}
