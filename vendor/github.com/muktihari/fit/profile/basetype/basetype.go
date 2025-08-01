// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basetype

import (
	"math"
	"strconv"
)

// BaseType is the base of all types used in FIT.
//
// Bits layout:
//   - 7  : Endian Ability (0: for single byte data; 1: if base type has endianness)
//   - 5–6: Reserved
//   - 0–4: Base Type Number
type BaseType byte

const (
	// BaseTypeNumMask is used to get the Base Type Number.
	//
	// Example: (Sint16 & BaseTypeNumMask) -> 3.
	BaseTypeNumMask = 0b00011111

	// EndianAbilityMask is used to get the Endian Ability.
	//
	// Example: (Sint32 & EndianAbilityMask == EndianAbilityMask) -> true
	EndianAbilityMask = 0b10000000
)

const (
	Enum    BaseType = 0x00
	Sint8   BaseType = 0x01 // 2’s complement format
	Uint8   BaseType = 0x02
	Sint16  BaseType = 0x83 // 2’s complement format
	Uint16  BaseType = 0x84
	Sint32  BaseType = 0x85 // 2’s complement format
	Uint32  BaseType = 0x86
	String  BaseType = 0x07 // Null terminated string encoded in UTF-8 format: 0x00
	Float32 BaseType = 0x88
	Float64 BaseType = 0x89
	Uint8z  BaseType = 0x0A
	Uint16z BaseType = 0x8B
	Uint32z BaseType = 0x8C
	Byte    BaseType = 0x0D // Array of bytes. Field is invalid if all bytes are invalid.
	Sint64  BaseType = 0x8E // 2’s complement format
	Uint64  BaseType = 0x8F
	Uint64z BaseType = 0x90
)

const (
	EnumInvalid    byte   = math.MaxUint8  // 0xFF
	Sint8Invalid   int8   = math.MaxInt8   // 0x7F
	Uint8Invalid   uint8  = math.MaxUint8  // 0xFF
	Sint16Invalid  int16  = math.MaxInt16  // 0x7FFF
	Uint16Invalid  uint16 = math.MaxUint16 // 0xFFFF
	Sint32Invalid  int32  = math.MaxInt32  // 0x7FFFFFFF
	Uint32Invalid  uint32 = math.MaxUint32 // 0xFFFFFFFF
	StringInvalid  string = ""             // We use empty string to represent an invalid string in Go. However, it will be converted automatically into an utf8 null-terminated string "\x00" by the Value Marshaler.
	Float32Invalid uint32 = math.MaxUint32 // 0xFFFFFFFF. math.Float32frombits(0xFFFFFFFF) produces float64 NaN which is uncomparable. Can only check in its integer form e.g. math.Float32bits(float32value) == Float32Invalid.
	Float64Invalid uint64 = math.MaxUint64 // 0xFFFFFFFFFFFFFFFF. math.Float64frombits(0xFFFFFFFFFFFFFFFF) produces float64 NaN which is uncomparable. Can only check in its integer form e.g. math.Float64bits(float64value) == Float64Invalid.
	Uint8zInvalid  uint8  = 0              // 0x00
	Uint16zInvalid uint16 = 0              // 0x0000
	Uint32zInvalid uint32 = 0              // 0x00000000
	ByteInvalid    byte   = math.MaxUint8  // 0xFF
	Sint64Invalid  int64  = math.MaxInt64  // 0x7FFFFFFFFFFFFFFF
	Uint64Invalid  uint64 = math.MaxUint64 // 0xFFFFFFFFFFFFFFFF
	Uint64zInvalid uint64 = 0              // 0x0000000000000000
)

// FromString convert given s into BaseType, if not valid 255 will be returned.
func FromString(s string) BaseType {
	switch s {
	case "enum":
		return Enum
	case "sint8":
		return Sint8
	case "uint8":
		return Uint8
	case "sint16":
		return Sint16
	case "uint16":
		return Uint16
	case "sint32":
		return Sint32
	case "uint32":
		return Uint32
	case "string":
		return String
	case "float32":
		return Float32
	case "float64":
		return Float64
	case "uint8z":
		return Uint8z
	case "uint16z":
		return Uint16z
	case "uint32z":
		return Uint32z
	case "byte":
		return Byte
	case "sint64":
		return Sint64
	case "uint64":
		return Uint64
	case "uint64z":
		return Uint64z
	}
	return 255
}

// String returns string representation of t.
func (t BaseType) String() string {
	switch t {
	case Enum:
		return "enum"
	case Sint8:
		return "sint8"
	case Uint8:
		return "uint8"
	case Sint16:
		return "sint16"
	case Uint16:
		return "uint16"
	case Sint32:
		return "sint32"
	case Uint32:
		return "uint32"
	case String:
		return "string"
	case Float32:
		return "float32"
	case Float64:
		return "float64"
	case Uint8z:
		return "uint8z"
	case Uint16z:
		return "uint16z"
	case Uint32z:
		return "uint32z"
	case Byte:
		return "byte"
	case Sint64:
		return "sint64"
	case Uint64:
		return "uint64"
	case Uint64z:
		return "uint64z"
	}
	return "invalid(" + strconv.Itoa(int(t)) + ")"
}

var sizes = [256]byte{
	Enum:    1,
	Sint8:   1,
	Uint8:   1,
	Sint16:  2,
	Uint16:  2,
	Sint32:  4,
	Uint32:  4,
	String:  1,
	Float32: 4,
	Float64: 8,
	Uint8z:  1,
	Uint16z: 2,
	Uint32z: 4,
	Byte:    1,
	Sint64:  8,
	Uint64:  8,
	Uint64z: 8,
}

// Size returns how many bytes it needs in binary form. If BaseType is invalid, zero will be returned.
func (t BaseType) Size() byte {
	return sizes[t] // PERF: use array to optimize speed since this method is frequently used.
}

// Valid checks whether BaseType is valid or not.
func (t BaseType) Valid() bool {
	return sizes[t] > 0
}

// GoType returns go equivalent type in string.
func (t BaseType) GoType() string {
	switch t {
	case Enum:
		return "byte"
	case Sint8:
		return "int8"
	case Uint8, Uint8z:
		return "uint8"
	case Sint16:
		return "int16"
	case Uint16, Uint16z:
		return "uint16"
	case Sint32:
		return "int32"
	case Uint32, Uint32z:
		return "uint32"
	case String:
		return "string"
	case Float32:
		return "float32"
	case Float64:
		return "float64"
	case Byte:
		return "byte"
	case Sint64:
		return "int64"
	case Uint64, Uint64z:
		return "uint64"
	}
	return "invalid(" + strconv.Itoa(int(t)) + ")"
}

// Invalid returns invalid value of t. e.g. Byte is 255 (its highest value).
func (t BaseType) Invalid() any {
	switch t {
	case Enum:
		return EnumInvalid
	case Sint8:
		return Sint8Invalid
	case Uint8:
		return Uint8Invalid
	case Sint16:
		return Sint16Invalid
	case Uint16:
		return Uint16Invalid
	case Sint32:
		return Sint32Invalid
	case Uint32:
		return Uint32Invalid
	case String:
		return StringInvalid
	case Float32:
		return math.Float32frombits(Float32Invalid)
	case Float64:
		return math.Float64frombits(Float64Invalid)
	case Uint8z:
		return Uint8zInvalid
	case Uint16z:
		return Uint16zInvalid
	case Uint32z:
		return Uint32zInvalid
	case Byte:
		return ByteInvalid
	case Sint64:
		return Sint64Invalid
	case Uint64:
		return Uint64Invalid
	case Uint64z:
		return Uint64zInvalid
	}
	return nil
}

// List returns all constants.
func List() []BaseType {
	return []BaseType{
		Enum,
		Sint8,
		Uint8,
		Sint16,
		Uint16,
		Sint32,
		Uint32,
		String,
		Float32,
		Float64,
		Uint8z,
		Uint16z,
		Uint32z,
		Byte,
		Sint64,
		Uint64,
		Uint64z,
	}
}
