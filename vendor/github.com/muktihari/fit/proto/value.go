// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"math"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
)

// Type is Value's type
type Type byte

const (
	TypeInvalid Type = iota
	TypeBool
	TypeInt8
	TypeUint8
	TypeInt16
	TypeUint16
	TypeInt32
	TypeUint32
	TypeInt64
	TypeUint64
	TypeFloat32
	TypeFloat64
	TypeString
	TypeSliceBool
	TypeSliceInt8
	TypeSliceUint8
	TypeSliceInt16
	TypeSliceUint16
	TypeSliceInt32
	TypeSliceUint32
	TypeSliceInt64
	TypeSliceUint64
	TypeSliceFloat32
	TypeSliceFloat64
	TypeSliceString
)

var typeStrings = [...]string{
	"Invalid",
	"Bool",
	"Int8",
	"Uint8",
	"Int16",
	"Uint16",
	"Int32",
	"Uint32",
	"Int64",
	"Uint64",
	"Float32",
	"Float64",
	"String",
	"SliceBool",
	"SliceInt8",
	"SliceUint8",
	"SliceInt16",
	"SliceUint16",
	"SliceInt32",
	"SliceUint32",
	"SliceInt64",
	"SliceUint64",
	"SliceFloat32",
	"SliceFloat64",
	"SliceString",
}

func (t Type) String() string {
	if t < Type(len(typeStrings)) {
		return typeStrings[t]
	}
	return "proto.TypeInvalid(" + strconv.Itoa(int(t)) + ")"
}

// Value is a zero alloc implementation value that hold any FIT protocol value
// (value of primitive-types or slice of primitive-types).
//
// To compare two Values of not known type, compare the results of the Any method.
// Using == on two Values is disallowed.
type Value struct {
	_   [0]func()      // disallow ==
	num uint64         // num holds either a numeric value or a slice's len.
	ptr unsafe.Pointer // ptr holds a pointer to slice's data only if it's a slice value.
	typ Type           // typ holds a Type.
}

// Return the underlying type the Value holds.
func (v Value) Type() Type { return v.typ }

// Int8 returns Value as int8, if it's not a valid int8 value, it returns basetype.Sint8Invalid (0x7F).
func (v Value) Int8() int8 {
	if v.typ != TypeInt8 {
		return basetype.Sint8Invalid
	}
	return int8(v.num)
}

// Bool returns Value as typedef.Bool, if it's not a valid typedef.Bool value, it returns typedef.BoolInvalid.
func (v Value) Bool() typedef.Bool {
	if v.typ != TypeBool {
		return typedef.BoolInvalid
	}
	return typedef.Bool(v.num)
}

// Uint8 returns Value as uint8, if it's not a valid uint8 value, it returns basetype.Uint8Invalid (0xFF).
func (v Value) Uint8() uint8 {
	if v.typ != TypeUint8 {
		return basetype.Uint8Invalid
	}
	return uint8(v.num)
}

// Uint8z returns Value as uint8, if it's not a valid uint8 value, it returns basetype.Uint8zInvalid (0).
func (v Value) Uint8z() uint8 {
	if v.typ != TypeUint8 {
		return basetype.Uint8zInvalid
	}
	return uint8(v.num)
}

// Int16 returns Value as int16, if it's not a valid int16 value, it returns basetype.Sint16Invalid (0x7FFF).
func (v Value) Int16() int16 {
	if v.typ != TypeInt16 {
		return basetype.Sint16Invalid
	}
	return int16(v.num)
}

// Uint16 returns Value as uint16, if it's not a valid uint16 value, it returns basetype.Uint16Invalid (0xFFFF).
func (v Value) Uint16() uint16 {
	if v.typ != TypeUint16 {
		return basetype.Uint16Invalid
	}
	return uint16(v.num)
}

// Uint16z returns Value as uint16, if it's not a valid uint16 value, it returns basetype.Uint16zInvalid (0).
func (v Value) Uint16z() uint16 {
	if v.typ != TypeUint16 {
		return basetype.Uint16zInvalid
	}
	return uint16(v.num)
}

// Int32 returns Value as int32, if it's not a valid int32 value, it returns basetype.Sint32Invalid (0x7FFFFFFF).
func (v Value) Int32() int32 {
	if v.typ != TypeInt32 {
		return basetype.Sint32Invalid
	}
	return int32(v.num)
}

// Uint32 returns Value as uint32, if it's not a valid uint32 value, it returns basetype.Uint32Invalid (0xFFFFFFFF).
func (v Value) Uint32() uint32 {
	if v.typ != TypeUint32 {
		return basetype.Uint32Invalid
	}
	return uint32(v.num)
}

// Uint32z returns Value as uint32, if it's not a valid uint32 value, it returns basetype.Uint32zInvalid (0).
func (v Value) Uint32z() uint32 {
	if v.typ != TypeUint32 {
		return basetype.Uint32zInvalid
	}
	return uint32(v.num)
}

// Int64 returns Value as int64, if it's not a valid int64 value, it returns basetype.Sint64Invalid (0x7FFFFFFFFFFFFFFF).
func (v Value) Int64() int64 {
	if v.typ != TypeInt64 {
		return basetype.Sint64Invalid
	}
	return int64(v.num)
}

// Uint64 returns Value as uint64, if it's not a valid uint64 value, it returns basetype.Uint64Invalid (0xFFFFFFFFFFFFFFFF).
func (v Value) Uint64() uint64 {
	if v.typ != TypeUint64 {
		return basetype.Uint64Invalid
	}
	return v.num
}

// Uint64z returns Value as uint64, if it's not a valid uint64 value, it returns basetype.Uint64Invalid (0).
func (v Value) Uint64z() uint64 {
	if v.typ != TypeUint64 {
		return basetype.Uint64zInvalid
	}
	return uint64(v.num)
}

// Float32 returns Value as float32, if it's not a valid float32 value, it returns basetype.Float32Invalid (0xFFFFFFFF) in float32 value.
func (v Value) Float32() float32 {
	if v.typ != TypeFloat32 {
		return math.Float32frombits(basetype.Float32Invalid)
	}
	return math.Float32frombits(uint32(v.num))
}

// Float64 returns Value as float64, if it's not a valid float64 value, it returns basetype.Float64Invalid (0xFFFFFFFFFFFFFFFF) in float64 value.
func (v Value) Float64() float64 {
	if v.typ != TypeFloat64 {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return math.Float64frombits(v.num)
}

// String returns Value as string, if it's not a valid string value, it returns basetype.StringInvalid.
// This should not be treated as a Go's String method, use Any() if you want to print the underlying value.
func (v Value) String() string {
	if v.typ != TypeString {
		return basetype.StringInvalid
	}
	return unsafe.String((*byte)(v.ptr), v.num)
}

// SliceBool returns Value as []typedef.Bool, if it's not a valid []typedef.Bool value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceBool() []typedef.Bool {
	if v.typ != TypeSliceBool {
		return nil
	}
	return unsafe.Slice((*typedef.Bool)(v.ptr), v.num)
}

// SliceInt8 returns Value as []int8, if it's not a valid []int8 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceInt8() []int8 {
	if v.typ != TypeSliceInt8 {
		return nil
	}
	return unsafe.Slice((*int8)(v.ptr), v.num)
}

// SliceUint8 returns Value as []uint8, if it's not a valid []uint8 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceUint8() []uint8 {
	if v.typ != TypeSliceUint8 {
		return nil
	}
	return unsafe.Slice((*uint8)(v.ptr), v.num)
}

// SliceInt16 returns Value as []int16, if it's not a valid []int16 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceInt16() []int16 {
	if v.typ != TypeSliceInt16 {
		return nil
	}
	return unsafe.Slice((*int16)(v.ptr), v.num)
}

// SliceUint16 returns Value as []uint16, if it's not a valid []uint16 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceUint16() []uint16 {
	if v.typ != TypeSliceUint16 {
		return nil
	}
	return unsafe.Slice((*uint16)(v.ptr), v.num)
}

// SliceInt32 returns Value as []int32, if it's not a valid []int32 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceInt32() []int32 {
	if v.typ != TypeSliceInt32 {
		return nil
	}
	return unsafe.Slice((*int32)(v.ptr), v.num)
}

// SliceUint32 returns Value as []uint32, if it's not a valid []uint32 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceUint32() []uint32 {
	if v.typ != TypeSliceUint32 {
		return nil
	}
	return unsafe.Slice((*uint32)(v.ptr), v.num)
}

// SliceInt64 returns Value as []int64, if it's not a valid []int64 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceInt64() []int64 {
	if v.typ != TypeSliceInt64 {
		return nil
	}
	return unsafe.Slice((*int64)(v.ptr), v.num)
}

// SliceUint64 returns Value as []uint64, if it's not a valid []uint64 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceUint64() []uint64 {
	if v.typ != TypeSliceUint64 {
		return nil
	}
	return unsafe.Slice((*uint64)(v.ptr), v.num)
}

// SliceFloat32 returns Value as []float32, if it's not a valid []float32 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceFloat32() []float32 {
	if v.typ != TypeSliceFloat32 {
		return nil
	}
	return unsafe.Slice((*float32)(v.ptr), v.num)
}

// SliceFloat64 returns Value as []float64, if it's not a valid []float64 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceFloat64() []float64 {
	if v.typ != TypeSliceFloat64 {
		return nil
	}
	return unsafe.Slice((*float64)(v.ptr), v.num)
}

// SliceString returns Value as []string, if it's not a valid []string value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceString() []string {
	if v.typ != TypeSliceString {
		return nil
	}
	return unsafe.Slice((*string)(v.ptr), v.num)
}

// Any returns Value's underlying value. If the underlying value is a slice, the caller takes ownership of that slice value,
// so Value should no longer be used after this call, except the returned value is copied and the copied value is used instead.
func (v Value) Any() any {
	switch v.typ {
	case TypeBool:
		return v.Bool()
	case TypeInt8:
		return v.Int8()
	case TypeUint8:
		return v.Uint8()
	case TypeInt16:
		return v.Int16()
	case TypeUint16:
		return v.Uint16()
	case TypeInt32:
		return v.Int32()
	case TypeUint32:
		return v.Uint32()
	case TypeInt64:
		return v.Int64()
	case TypeUint64:
		return v.Uint64()
	case TypeFloat32:
		return v.Float32()
	case TypeFloat64:
		return v.Float64()
	case TypeString:
		return v.String()
	case TypeSliceBool:
		return v.SliceBool()
	case TypeSliceInt8:
		return v.SliceInt8()
	case TypeSliceUint8:
		return v.SliceUint8()
	case TypeSliceInt16:
		return v.SliceInt16()
	case TypeSliceUint16:
		return v.SliceUint16()
	case TypeSliceInt32:
		return v.SliceInt32()
	case TypeSliceUint32:
		return v.SliceUint32()
	case TypeSliceInt64:
		return v.SliceInt64()
	case TypeSliceUint64:
		return v.SliceUint64()
	case TypeSliceFloat32:
		return v.SliceFloat32()
	case TypeSliceFloat64:
		return v.SliceFloat64()
	case TypeSliceString:
		return v.SliceString()
	}
	return nil
}

// Align checks whether Value's type is align with given basetype.
func (v Value) Align(t basetype.BaseType) bool {
	switch v.typ {
	case TypeBool, TypeSliceBool:
		return t == basetype.Enum
	case TypeInt8, TypeSliceInt8:
		return t == basetype.Sint8
	case TypeUint8, TypeSliceUint8:
		return t == basetype.Enum ||
			t == basetype.Byte ||
			t == basetype.Uint8 ||
			t == basetype.Uint8z
	case TypeInt16, TypeSliceInt16:
		return t == basetype.Sint16
	case TypeUint16, TypeSliceUint16:
		return t == basetype.Uint16 || t == basetype.Uint16z
	case TypeInt32, TypeSliceInt32:
		return t == basetype.Sint32
	case TypeUint32, TypeSliceUint32:
		return t == basetype.Uint32 || t == basetype.Uint32z
	case TypeInt64, TypeSliceInt64:
		return t == basetype.Sint64
	case TypeUint64, TypeSliceUint64:
		return t == basetype.Uint64 || t == basetype.Uint64z
	case TypeFloat32, TypeSliceFloat32:
		return t == basetype.Float32
	case TypeFloat64, TypeSliceFloat64:
		return t == basetype.Float64
	case TypeString, TypeSliceString:
		return t == basetype.String
	}
	return false
}

// Valid checks whether the Value is valid based on given basetype. This does not verify whether the Type of
// the Value aligns with the provided BaseType. For slices, even though only one element is valid, the Value will be counted a valid value.
func (v Value) Valid(t basetype.BaseType) bool {
	var invalidCount int

	switch v.typ {
	case TypeBool:
		return v.num < 2 // Only 0 (false) and 1 (true) is valid
	case TypeInt8:
		return int8(v.num) != basetype.Sint8Invalid
	case TypeUint8:
		val := uint8(v.num)
		switch t {
		case basetype.Enum:
			return val != basetype.EnumInvalid
		case basetype.Byte:
			return val != basetype.ByteInvalid
		case basetype.Uint8:
			return val != basetype.Uint8Invalid
		case basetype.Uint8z:
			return val != basetype.Uint8zInvalid
		}
		return false
	case TypeInt16:
		return int16(v.num) != basetype.Sint16Invalid
	case TypeUint16:
		if t == basetype.Uint16z {
			return uint16(v.num) != basetype.Uint16zInvalid
		}
		return uint16(v.num) != basetype.Uint16Invalid
	case TypeInt32:
		return int32(v.num) != basetype.Sint32Invalid
	case TypeUint32:
		if t == basetype.Uint32z {
			return uint32(v.num) != basetype.Uint32zInvalid
		}
		return uint32(v.num) != basetype.Uint32Invalid
	case TypeInt64:
		return int64(v.num) != basetype.Sint64Invalid
	case TypeUint64:
		if t == basetype.Uint64z {
			return v.num != basetype.Uint64zInvalid
		}
		return v.num != basetype.Uint64Invalid
	case TypeFloat32:
		return uint32(v.num) != basetype.Float32Invalid
	case TypeFloat64:
		return v.num != basetype.Float64Invalid
	case TypeString:
		s := v.String()
		return s != basetype.StringInvalid && s != "\x00"
	case TypeSliceBool:
		vals := v.SliceBool()
		for i := range vals {
			if vals[i] == typedef.BoolInvalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case TypeSliceInt8:
		vals := v.SliceInt8()
		for i := range vals {
			if vals[i] == basetype.Sint8Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case TypeSliceUint8:
		vals := v.SliceUint8()
		if t == basetype.Uint8z {
			for i := range vals {
				if vals[i] == basetype.Uint8zInvalid {
					invalidCount++
				}
			}
		} else {
			for i := range vals {
				if vals[i] == basetype.Uint8Invalid {
					invalidCount++
				}
			}
		}
		return invalidCount != len(vals)
	case TypeSliceInt16:
		vals := v.SliceInt16()
		for i := range vals {
			if vals[i] == basetype.Sint16Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case TypeSliceUint16:
		vals := v.SliceUint16()
		if t == basetype.Uint16z {
			for i := range vals {
				if vals[i] == basetype.Uint16zInvalid {
					invalidCount++
				}
			}
		} else {
			for i := range vals {
				if vals[i] == basetype.Uint16Invalid {
					invalidCount++
				}
			}
		}
		return invalidCount != len(vals)
	case TypeSliceInt32:
		vals := v.SliceInt32()
		for i := range vals {
			if vals[i] == basetype.Sint32Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case TypeSliceUint32:
		vals := v.SliceUint32()
		if t == basetype.Uint32z {
			for i := range vals {
				if vals[i] == basetype.Uint32zInvalid {
					invalidCount++
				}
			}
		} else {
			for i := range vals {
				if vals[i] == basetype.Uint32Invalid {
					invalidCount++
				}
			}
		}
		return invalidCount != len(vals)
	case TypeSliceInt64:
		vals := v.SliceInt64()
		for i := range vals {
			if vals[i] == basetype.Sint64Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case TypeSliceUint64:
		vals := v.SliceUint64()
		if t == basetype.Uint64z {
			for i := range vals {
				if vals[i] == basetype.Uint64zInvalid {
					invalidCount++
				}
			}
		} else {
			for i := range vals {
				if vals[i] == basetype.Uint64Invalid {
					invalidCount++
				}
			}
		}
		return invalidCount != len(vals)
	case TypeSliceFloat32:
		vals := v.SliceFloat32()
		for i := range vals {
			if math.Float32bits(vals[i]) == basetype.Float32Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case TypeSliceFloat64:
		vals := v.SliceFloat64()
		for i := range vals {
			if math.Float64bits(vals[i]) == basetype.Float64Invalid {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	case TypeSliceString:
		vals := v.SliceString()
		for i := range vals {
			if vals[i] == basetype.StringInvalid || vals[i] == "\x00" {
				invalidCount++
			}
		}
		return invalidCount != len(vals)
	}
	return false
}

// Bool converts typedef.Bool as Value. If v > 1, it will be treated as typedef.BoolInvalid.
func Bool(v typedef.Bool) Value {
	num := uint64(v)
	if v > 1 {
		num = uint64(typedef.BoolInvalid)
	}
	return Value{num: num, typ: TypeBool}
}

// Int8 converts int8 as Value.
func Int8(v int8) Value {
	return Value{num: uint64(v), typ: TypeInt8}
}

// Uint8 converts uint8 as Value.
func Uint8(v uint8) Value {
	return Value{num: uint64(v), typ: TypeUint8}
}

// Int16 converts int16 as Value.
func Int16(v int16) Value {
	return Value{num: uint64(v), typ: TypeInt16}
}

// Uint16 converts uint16 as Value.
func Uint16(v uint16) Value {
	return Value{num: uint64(v), typ: TypeUint16}
}

// Int32 converts int32 as Value.
func Int32(v int32) Value {
	return Value{num: uint64(v), typ: TypeInt32}
}

// Uint32 converts uint32 as Value.
func Uint32(v uint32) Value {
	return Value{num: uint64(v), typ: TypeUint32}
}

// Int64 converts int64 as Value.
func Int64(v int64) Value {
	return Value{num: uint64(v), typ: TypeInt64}
}

// Uint64 converts uint64 as Value.
func Uint64(v uint64) Value {
	return Value{num: v, typ: TypeUint64}
}

// Float32 converts float32 as Value.
func Float32(v float32) Value {
	return Value{num: uint64(math.Float32bits(v)), typ: TypeFloat32}
}

// Float64 converts float64 as Value.
func Float64(v float64) Value {
	return Value{num: math.Float64bits(v), typ: TypeFloat64}
}

// String converts string as Value.
func String(v string) Value {
	return Value{num: uint64(len(v)), typ: TypeString, ptr: unsafe.Pointer(unsafe.StringData(v))}
}

// SliceBool converts []typedef.Bool as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceBool(s []typedef.Bool) Value {
	return Value{num: uint64(len(s)), typ: TypeSliceBool, ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceInt8 converts []int8 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceInt8[S []E, E ~int8](s S) Value {
	return Value{num: uint64(len(s)), typ: TypeSliceInt8, ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceUint8 converts []uint8 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceUint8[S []E, E ~uint8](s S) Value {
	return Value{num: uint64(len(s)), typ: TypeSliceUint8, ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceInt16 converts []int16 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceInt16[S []E, E ~int16](s S) Value {
	return Value{num: uint64(len(s)), typ: TypeSliceInt16, ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceUint16 converts []uint16 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceUint16[S []E, E ~uint16](s S) Value {
	return Value{num: uint64(len(s)), typ: TypeSliceUint16, ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceInt32 converts []int32 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceInt32[S []E, E ~int32](s S) Value {
	return Value{num: uint64(len(s)), typ: TypeSliceInt32, ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceUint32 converts []uint32 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceUint32[S []E, E ~uint32](s S) Value {
	return Value{num: uint64(len(s)), typ: TypeSliceUint32, ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceInt64 converts []int64 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceInt64[S []E, E ~int64](s S) Value {
	return Value{num: uint64(len(s)), typ: TypeSliceInt64, ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceUint64 converts []uint64 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceUint64[S []E, E ~uint64](s S) Value {
	return Value{num: uint64(len(s)), typ: TypeSliceUint64, ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceFloat32 converts []float32 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceFloat32[S []E, E ~float32](s S) Value {
	return Value{num: uint64(len(s)), typ: TypeSliceFloat32, ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceFloat64 converts []float64 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceFloat64[S []E, E ~float64](s S) Value {
	return Value{num: uint64(len(s)), typ: TypeSliceFloat64, ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceString converts []string as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceString[S []E, E ~string](s S) Value {
	return Value{num: uint64(len(s)), typ: TypeSliceString, ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// Any converts any value into Value. If the given v is not a primitive-type value or
// a slice of primitive-type, it will determine its types using reflection.
//
// It works with important caveats:
//   - If v is not a supported value such as int, uint, []int, []uint, []any, etc. Value with TypeInvalid will be returned.
//   - If v is a slice, this will take ownership of v, and the caller should not use v after this call.
func Any(v any) Value {
	switch val := v.(type) { // Fast path
	case int, uint, []int, []uint, []any: // Fast return on invalid value
		return Value{}
	case Value:
		return val
	case bool:
		if val {
			return Bool(typedef.BoolTrue)
		}
		return Bool(typedef.BoolFalse)
	case typedef.Bool:
		return Bool(val)
	case int8:
		return Int8(val)
	case uint8:
		return Uint8(val)
	case int16:
		return Int16(val)
	case uint16:
		return Uint16(val)
	case int32:
		return Int32(val)
	case uint32:
		return Uint32(val)
	case int64:
		return Int64(val)
	case uint64:
		return Uint64(val)
	case float32:
		return Float32(val)
	case float64:
		return Float64(val)
	case string:
		return String(val)
	case []bool:
		// Casting a bool to uint8 is not guaranteed to work correctly since bool and uint8 may not
		// have the same memory layout. It's safer to allocate memory explicitly for the conversion.
		bools := make([]typedef.Bool, len(val))
		for i := range val {
			if val[i] {
				bools[i] = typedef.BoolTrue
			} else {
				bools[i] = typedef.BoolFalse
			}
		}
		return SliceBool(bools)
	case []typedef.Bool:
		return SliceBool(val)
	case []int8:
		return SliceInt8(val)
	case []uint8:
		return SliceUint8(val)
	case []int16:
		return SliceInt16(val)
	case []uint16:
		return SliceUint16(val)
	case []int32:
		return SliceInt32(val)
	case []uint32:
		return SliceUint32(val)
	case []int64:
		return SliceInt64(val)
	case []uint64:
		return SliceUint64(val)
	case []float32:
		return SliceFloat32(val)
	case []float64:
		return SliceFloat64(val)
	case []string:
		return SliceString(val)
	}

	// Fallback to reflection.
	rv := reflect.Indirect(reflect.ValueOf(v))
	switch rv.Kind() {
	case reflect.Bool:
		if rv.Bool() {
			return Bool(typedef.BoolTrue)
		}
		return Bool(typedef.BoolFalse)
	case reflect.Int8:
		return Int8(int8(rv.Int()))
	case reflect.Uint8:
		return Uint8(uint8(rv.Uint()))
	case reflect.Int16:
		return Int16(int16(rv.Int()))
	case reflect.Uint16:
		return Uint16(uint16(rv.Uint()))
	case reflect.Int32:
		return Int32(int32(rv.Int()))
	case reflect.Uint32:
		return Uint32(uint32(rv.Uint()))
	case reflect.Int64:
		return Int64(int64(rv.Int()))
	case reflect.Uint64:
		return Uint64(uint64(rv.Uint()))
	case reflect.Float32:
		return Float32(float32(rv.Float()))
	case reflect.Float64:
		return Float64(float64(rv.Float()))
	case reflect.String:
		return String(rv.String())
	case reflect.Slice:
		ptr := rv.UnsafePointer()
		switch rv.Type().Elem().Kind() {
		case reflect.Bool:
			vals := unsafe.Slice((*bool)(ptr), rv.Len())
			bools := make([]typedef.Bool, len(vals)) // See: case []bool for details on why we must allocate.
			for i := range vals {
				if vals[i] {
					bools[i] = typedef.BoolTrue
				} else {
					bools[i] = typedef.BoolFalse
				}
			}
			return SliceBool(bools)
		case reflect.Int8:
			return SliceInt8(unsafe.Slice((*int8)(ptr), rv.Len()))
		case reflect.Uint8:
			return SliceUint8(unsafe.Slice((*uint8)(ptr), rv.Len()))
		case reflect.Int16:
			return SliceInt16(unsafe.Slice((*int16)(ptr), rv.Len()))
		case reflect.Uint16:
			return SliceUint16(unsafe.Slice((*uint16)(ptr), rv.Len()))
		case reflect.Int32:
			return SliceInt32(unsafe.Slice((*int32)(ptr), rv.Len()))
		case reflect.Uint32:
			return SliceUint32(unsafe.Slice((*uint32)(ptr), rv.Len()))
		case reflect.Int64:
			return SliceInt64(unsafe.Slice((*int64)(ptr), rv.Len()))
		case reflect.Uint64:
			return SliceUint64(unsafe.Slice((*uint64)(ptr), rv.Len()))
		case reflect.Float32:
			return SliceFloat32(unsafe.Slice((*float32)(ptr), rv.Len()))
		case reflect.Float64:
			return SliceFloat64(unsafe.Slice((*float64)(ptr), rv.Len()))
		case reflect.String:
			return SliceString(unsafe.Slice((*string)(ptr), rv.Len()))
		}
	}

	return Value{}
}

var sizes = [...]int{
	TypeInvalid: 0,
	TypeBool:    1,
	TypeInt8:    int(basetype.Sint8.Size()),
	TypeUint8:   int(basetype.Uint8.Size()),
	TypeInt16:   int(basetype.Sint16.Size()),
	TypeUint16:  int(basetype.Uint16.Size()),
	TypeInt32:   int(basetype.Sint32.Size()),
	TypeUint32:  int(basetype.Uint32.Size()),
	TypeInt64:   int(basetype.Sint64.Size()),
	TypeUint64:  int(basetype.Uint64.Size()),
	TypeFloat32: int(basetype.Float32.Size()),
	TypeFloat64: int(basetype.Float64.Size()),
	TypeString:  int(basetype.String.Size()),
}

// Sizeof returns the size of val in bytes. For every string in Value, if the last index of the string is not '\x00', size += 1.
func Sizeof(val Value) int {
	switch val.typ {
	case TypeString:
		s := val.String()
		n := len(s)
		if n == 0 || s[n-1] != '\x00' {
			n += 1
		}
		return n * sizes[TypeString]
	case TypeSliceBool:
		return int(val.num) * sizes[TypeBool]
	case TypeSliceInt8:
		return int(val.num) * sizes[TypeInt8]
	case TypeSliceUint8:
		return int(val.num) * sizes[TypeUint8]
	case TypeSliceInt16:
		return int(val.num) * sizes[TypeInt16]
	case TypeSliceUint16:
		return int(val.num) * sizes[TypeUint16]
	case TypeSliceInt32:
		return int(val.num) * sizes[TypeInt32]
	case TypeSliceUint32:
		return int(val.num) * sizes[TypeUint32]
	case TypeSliceInt64:
		return int(val.num) * sizes[TypeInt64]
	case TypeSliceUint64:
		return int(val.num) * sizes[TypeUint64]
	case TypeSliceFloat32:
		return int(val.num) * sizes[TypeFloat32]
	case TypeSliceFloat64:
		return int(val.num) * sizes[TypeFloat64]
	case TypeSliceString:
		vs := val.SliceString()
		var size int
		for i := range vs {
			n := len(vs[i])
			if n == 0 || vs[i][n-1] != '\x00' {
				n += 1 // utf-8 null terminated string
			}
			size += n
		}
		if size == 0 {
			return 1 * sizes[TypeString] // utf-8 null terminated string
		}
		return size * sizes[TypeString]
	default:
		return sizes[val.typ]
	}
}
