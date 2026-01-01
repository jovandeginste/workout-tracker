// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"fmt"
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
type Value struct {
	_   [0]func()      // disallow ==
	num uint64         // num holds either a numeric value or a slice's len + type identifier (5 msb).
	ptr unsafe.Pointer // ptr holds either a pointer to type identifier or a pointer to slice's data.
}

// Type identifier only for numeric values.
// For slices, we use 5 most significant bits (msb) of Value's num for type identifier.
//
// The use of pointer arithmetic in "(Value) Type" method depends on:
//   - The pointer addresses MUST be incremented by 1 byte.
//   - The order MUST identical with constants Types order.
var (
	memptr     [12]byte // we only need the addresses
	startAddr  = uintptr(unsafe.Pointer(&memptr[TypeInvalid]))
	ptrBool    = unsafe.Pointer(&memptr[TypeBool])
	ptrInt8    = unsafe.Pointer(&memptr[TypeInt8])
	ptrUint8   = unsafe.Pointer(&memptr[TypeUint8])
	ptrInt16   = unsafe.Pointer(&memptr[TypeInt16])
	ptrUint16  = unsafe.Pointer(&memptr[TypeUint16])
	ptrInt32   = unsafe.Pointer(&memptr[TypeInt32])
	ptrUint32  = unsafe.Pointer(&memptr[TypeUint32])
	ptrInt64   = unsafe.Pointer(&memptr[TypeInt64])
	ptrUint64  = unsafe.Pointer(&memptr[TypeUint64])
	ptrFloat32 = unsafe.Pointer(&memptr[TypeFloat32])
	ptrFloat64 = unsafe.Pointer(&memptr[TypeFloat64])
	endAddr    = uintptr(ptrFloat64)
)

const (
	vbits  = 5                       // reserved bits for slice's type identifier (max value: 31, large enough to hold a Type)
	vshift = 64 - vbits              // type identifier bits shifter only for slices.
	vmask  = math.MaxUint64 >> vbits // mask for retrieving slice's len from Value's num, no slice's len exceed this value.
)

// Return the underlying type the Value holds.
func (v Value) Type() Type {
	if p := uintptr(v.ptr); p >= startAddr && p <= endAddr {
		return Type(p - startAddr)
	}
	return Type(v.num >> vshift)
}

var sizes = [...]int{
	TypeInvalid: 0,
	TypeBool:    1,
	TypeInt8:    1,
	TypeUint8:   1,
	TypeInt16:   2,
	TypeUint16:  2,
	TypeInt32:   4,
	TypeUint32:  4,
	TypeInt64:   8,
	TypeUint64:  8,
	TypeFloat32: 4,
	TypeFloat64: 8,
	TypeString:  1,
}

// Size returns the size of Value in binary from. For every string in Value,
// if the last index of the string is not '\x00', size += 1.
func (v Value) Size() int {
	switch typ := v.Type(); typ {
	case TypeString:
		s := v.String()
		n := len(s)
		if n == 0 || s[n-1] != '\x00' {
			n += 1
		}
		return n * sizes[TypeString]
	case TypeSliceBool:
		return int(v.num&vmask) * sizes[TypeBool]
	case TypeSliceInt8:
		return int(v.num&vmask) * sizes[TypeInt8]
	case TypeSliceUint8:
		return int(v.num&vmask) * sizes[TypeUint8]
	case TypeSliceInt16:
		return int(v.num&vmask) * sizes[TypeInt16]
	case TypeSliceUint16:
		return int(v.num&vmask) * sizes[TypeUint16]
	case TypeSliceInt32:
		return int(v.num&vmask) * sizes[TypeInt32]
	case TypeSliceUint32:
		return int(v.num&vmask) * sizes[TypeUint32]
	case TypeSliceInt64:
		return int(v.num&vmask) * sizes[TypeInt64]
	case TypeSliceUint64:
		return int(v.num&vmask) * sizes[TypeUint64]
	case TypeSliceFloat32:
		return int(v.num&vmask) * sizes[TypeFloat32]
	case TypeSliceFloat64:
		return int(v.num&vmask) * sizes[TypeFloat64]
	case TypeSliceString:
		vs := v.SliceString()
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
		return sizes[typ]
	}
}

// Bool returns Value as typedef.Bool, if it's not a valid typedef.Bool value, it returns typedef.BoolInvalid.
func (v Value) Bool() typedef.Bool {
	if v.ptr != ptrBool {
		return typedef.BoolInvalid
	}
	return typedef.Bool(v.num)
}

// Int8 returns Value as int8, if it's not a valid int8 value, it returns basetype.Sint8Invalid (0x7F).
func (v Value) Int8() int8 {
	if v.ptr != ptrInt8 {
		return basetype.Sint8Invalid
	}
	return int8(v.num)
}

// Uint8 returns Value as uint8, if it's not a valid uint8 value, it returns basetype.Uint8Invalid (0xFF).
func (v Value) Uint8() uint8 {
	if v.ptr != ptrUint8 {
		return basetype.Uint8Invalid
	}
	return uint8(v.num)
}

// Uint8z returns Value as uint8, if it's not a valid uint8 value, it returns basetype.Uint8zInvalid (0).
func (v Value) Uint8z() uint8 {
	if v.ptr != ptrUint8 {
		return basetype.Uint8zInvalid
	}
	return uint8(v.num)
}

// Int16 returns Value as int16, if it's not a valid int16 value, it returns basetype.Sint16Invalid (0x7FFF).
func (v Value) Int16() int16 {
	if v.ptr != ptrInt16 {
		return basetype.Sint16Invalid
	}
	return int16(v.num)
}

// Uint16 returns Value as uint16, if it's not a valid uint16 value, it returns basetype.Uint16Invalid (0xFFFF).
func (v Value) Uint16() uint16 {
	if v.ptr != ptrUint16 {
		return basetype.Uint16Invalid
	}
	return uint16(v.num)
}

// Uint16z returns Value as uint16, if it's not a valid uint16 value, it returns basetype.Uint16zInvalid (0).
func (v Value) Uint16z() uint16 {
	if v.ptr != ptrUint16 {
		return basetype.Uint16zInvalid
	}
	return uint16(v.num)
}

// Int32 returns Value as int32, if it's not a valid int32 value, it returns basetype.Sint32Invalid (0x7FFFFFFF).
func (v Value) Int32() int32 {
	if v.ptr != ptrInt32 {
		return basetype.Sint32Invalid
	}
	return int32(v.num)
}

// Uint32 returns Value as uint32, if it's not a valid uint32 value, it returns basetype.Uint32Invalid (0xFFFFFFFF).
func (v Value) Uint32() uint32 {
	if v.ptr != ptrUint32 {
		return basetype.Uint32Invalid
	}
	return uint32(v.num)
}

// Uint32z returns Value as uint32, if it's not a valid uint32 value, it returns basetype.Uint32zInvalid (0).
func (v Value) Uint32z() uint32 {
	if v.ptr != ptrUint32 {
		return basetype.Uint32zInvalid
	}
	return uint32(v.num)
}

// Int64 returns Value as int64, if it's not a valid int64 value, it returns basetype.Sint64Invalid (0x7FFFFFFFFFFFFFFF).
func (v Value) Int64() int64 {
	if v.ptr != ptrInt64 {
		return basetype.Sint64Invalid
	}
	return int64(v.num)
}

// Uint64 returns Value as uint64, if it's not a valid uint64 value, it returns basetype.Uint64Invalid (0xFFFFFFFFFFFFFFFF).
func (v Value) Uint64() uint64 {
	if v.ptr != ptrUint64 {
		return basetype.Uint64Invalid
	}
	return v.num
}

// Uint64z returns Value as uint64, if it's not a valid uint64 value, it returns basetype.Uint64zInvalid (0).
func (v Value) Uint64z() uint64 {
	if v.ptr != ptrUint64 {
		return basetype.Uint64zInvalid
	}
	return uint64(v.num)
}

// Float32 returns Value as float32, if it's not a valid float32 value, it returns basetype.Float32Invalid (0xFFFFFFFF) in float32 value.
func (v Value) Float32() float32 {
	if v.ptr != ptrFloat32 {
		return math.Float32frombits(basetype.Float32Invalid)
	}
	return math.Float32frombits(uint32(v.num))
}

// Float64 returns Value as float64, if it's not a valid float64 value, it returns basetype.Float64Invalid (0xFFFFFFFFFFFFFFFF) in float64 value.
func (v Value) Float64() float64 {
	if v.ptr != ptrFloat64 {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return math.Float64frombits(v.num)
}

// String returns Value as string, if it's not a valid string value, it returns basetype.StringInvalid.
// This should not be treated as a Go's String method, use Any() if you want to print the underlying value.
func (v Value) String() string {
	if v.Type() != TypeString {
		return basetype.StringInvalid
	}
	return unsafe.String((*byte)(v.ptr), v.num&vmask)
}

var _ fmt.Formatter = (*Value)(nil)

// Format controls how Value is formatted when using fmt. It overrides the String method, as the String method
// is used to return string value, rather than the Value formatted as a string.
func (v Value) Format(p fmt.State, verb rune) {
	switch {
	case v.num == 0 && v.ptr == nil:
		fmt.Fprintf(p, "<invalid proto.Value>")
	case verb != 'v':
		fmt.Fprintf(p, fmt.FormatString(p, verb), v.Any())
	case p.Flag('#'):
		fmt.Fprintf(p, "%#v", v.Any())
	default:
		fmt.Fprintf(p, "%v", v.Any())
	}
}

// SliceBool returns Value as []typedef.Bool, if it's not a valid []typedef.Bool value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceBool() []typedef.Bool {
	if v.Type() != TypeSliceBool {
		return nil
	}
	return unsafe.Slice((*typedef.Bool)(v.ptr), v.num&vmask)
}

// SliceInt8 returns Value as []int8, if it's not a valid []int8 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceInt8() []int8 {
	if v.Type() != TypeSliceInt8 {
		return nil
	}
	return unsafe.Slice((*int8)(v.ptr), v.num&vmask)
}

// SliceUint8 returns Value as []uint8, if it's not a valid []uint8 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceUint8() []uint8 {
	if v.Type() != TypeSliceUint8 {
		return nil
	}
	return unsafe.Slice((*uint8)(v.ptr), v.num&vmask)
}

// SliceInt16 returns Value as []int16, if it's not a valid []int16 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceInt16() []int16 {
	if v.Type() != TypeSliceInt16 {
		return nil
	}
	return unsafe.Slice((*int16)(v.ptr), v.num&vmask)
}

// SliceUint16 returns Value as []uint16, if it's not a valid []uint16 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceUint16() []uint16 {
	if v.Type() != TypeSliceUint16 {
		return nil
	}
	return unsafe.Slice((*uint16)(v.ptr), v.num&vmask)
}

// SliceInt32 returns Value as []int32, if it's not a valid []int32 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceInt32() []int32 {
	if v.Type() != TypeSliceInt32 {
		return nil
	}
	return unsafe.Slice((*int32)(v.ptr), v.num&vmask)
}

// SliceUint32 returns Value as []uint32, if it's not a valid []uint32 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceUint32() []uint32 {
	if v.Type() != TypeSliceUint32 {
		return nil
	}
	return unsafe.Slice((*uint32)(v.ptr), v.num&vmask)
}

// SliceInt64 returns Value as []int64, if it's not a valid []int64 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceInt64() []int64 {
	if v.Type() != TypeSliceInt64 {
		return nil
	}
	return unsafe.Slice((*int64)(v.ptr), v.num&vmask)
}

// SliceUint64 returns Value as []uint64, if it's not a valid []uint64 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceUint64() []uint64 {
	if v.Type() != TypeSliceUint64 {
		return nil
	}
	return unsafe.Slice((*uint64)(v.ptr), v.num&vmask)
}

// SliceFloat32 returns Value as []float32, if it's not a valid []float32 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceFloat32() []float32 {
	if v.Type() != TypeSliceFloat32 {
		return nil
	}
	return unsafe.Slice((*float32)(v.ptr), v.num&vmask)
}

// SliceFloat64 returns Value as []float64, if it's not a valid []float64 value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceFloat64() []float64 {
	if v.Type() != TypeSliceFloat64 {
		return nil
	}
	return unsafe.Slice((*float64)(v.ptr), v.num&vmask)
}

// SliceString returns Value as []string, if it's not a valid []string value, it returns nil.
// The caller takes ownership of the returned value, so Value should no longer be used after this call,
// except the returned value is copied and the copied value is used instead.
func (v Value) SliceString() []string {
	if v.Type() != TypeSliceString {
		return nil
	}
	return unsafe.Slice((*string)(v.ptr), v.num&vmask)
}

// Any returns Value's underlying value. If the underlying value is a slice, the caller takes ownership of that slice value,
// so Value should no longer be used after this call, except the returned value is copied and the copied value is used instead.
func (v Value) Any() any {
	switch v.Type() {
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
	switch v.Type() {
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
	switch v.Type() {
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
			if vals[i] != typedef.BoolInvalid {
				return true
			}
		}
		return false
	case TypeSliceInt8:
		vals := v.SliceInt8()
		for i := range vals {
			if vals[i] != basetype.Sint8Invalid {
				return true
			}
		}
		return false
	case TypeSliceUint8:
		vals := v.SliceUint8()
		if t == basetype.Uint8z {
			for i := range vals {
				if vals[i] != basetype.Uint8zInvalid {
					return true
				}
			}
		} else {
			for i := range vals {
				if vals[i] != basetype.Uint8Invalid {
					return true
				}
			}
		}
		return false
	case TypeSliceInt16:
		vals := v.SliceInt16()
		for i := range vals {
			if vals[i] != basetype.Sint16Invalid {
				return true
			}
		}
		return false
	case TypeSliceUint16:
		vals := v.SliceUint16()
		if t == basetype.Uint16z {
			for i := range vals {
				if vals[i] != basetype.Uint16zInvalid {
					return true
				}
			}
		} else {
			for i := range vals {
				if vals[i] != basetype.Uint16Invalid {
					return true
				}
			}
		}
		return false
	case TypeSliceInt32:
		vals := v.SliceInt32()
		for i := range vals {
			if vals[i] != basetype.Sint32Invalid {
				return true
			}
		}
		return false
	case TypeSliceUint32:
		vals := v.SliceUint32()
		if t == basetype.Uint32z {
			for i := range vals {
				if vals[i] != basetype.Uint32zInvalid {
					return true
				}
			}
		} else {
			for i := range vals {
				if vals[i] != basetype.Uint32Invalid {
					return true
				}
			}
		}
		return false
	case TypeSliceInt64:
		vals := v.SliceInt64()
		for i := range vals {
			if vals[i] != basetype.Sint64Invalid {
				return true
			}
		}
		return false
	case TypeSliceUint64:
		vals := v.SliceUint64()
		if t == basetype.Uint64z {
			for i := range vals {
				if vals[i] != basetype.Uint64zInvalid {
					return true
				}
			}
		} else {
			for i := range vals {
				if vals[i] != basetype.Uint64Invalid {
					return true
				}
			}
		}
		return false
	case TypeSliceFloat32:
		vals := v.SliceFloat32()
		for i := range vals {
			if math.Float32bits(vals[i]) != basetype.Float32Invalid {
				return true
			}
		}
		return false
	case TypeSliceFloat64:
		vals := v.SliceFloat64()
		for i := range vals {
			if math.Float64bits(vals[i]) != basetype.Float64Invalid {
				return true
			}
		}
		return false
	case TypeSliceString:
		vals := v.SliceString()
		for i := range vals {
			if vals[i] != basetype.StringInvalid && vals[i] != "\x00" {
				return true
			}
		}
		return false
	}
	return false
}

// Bool converts typedef.Bool as Value. If v > 1, it will be treated as typedef.BoolInvalid.
func Bool(v typedef.Bool) Value {
	num := uint64(v)
	if v > 1 {
		num = uint64(typedef.BoolInvalid)
	}
	return Value{num: num, ptr: ptrBool}
}

// Int8 converts int8 as Value.
func Int8(v int8) Value {
	return Value{num: uint64(v), ptr: ptrInt8}
}

// Uint8 converts uint8 as Value.
func Uint8(v uint8) Value {
	return Value{num: uint64(v), ptr: ptrUint8}
}

// Int16 converts int16 as Value.
func Int16(v int16) Value {
	return Value{num: uint64(v), ptr: ptrInt16}
}

// Uint16 converts uint16 as Value.
func Uint16(v uint16) Value {
	return Value{num: uint64(v), ptr: ptrUint16}
}

// Int32 converts int32 as Value.
func Int32(v int32) Value {
	return Value{num: uint64(v), ptr: ptrInt32}
}

// Uint32 converts uint32 as Value.
func Uint32(v uint32) Value {
	return Value{num: uint64(v), ptr: ptrUint32}
}

// Int64 converts int64 as Value.
func Int64(v int64) Value {
	return Value{num: uint64(v), ptr: ptrInt64}
}

// Uint64 converts uint64 as Value.
func Uint64(v uint64) Value {
	return Value{num: v, ptr: ptrUint64}
}

// Float32 converts float32 as Value.
func Float32(v float32) Value {
	return Value{num: uint64(math.Float32bits(v)), ptr: ptrFloat32}
}

// Float64 converts float64 as Value.
func Float64(v float64) Value {
	return Value{num: math.Float64bits(v), ptr: ptrFloat64}
}

// String converts string as Value.
func String(v string) Value {
	return Value{num: uint64(TypeString)<<vshift | uint64(len(v)), ptr: unsafe.Pointer(unsafe.StringData(v))}
}

// SliceBool converts []typedef.Bool as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceBool(s []typedef.Bool) Value {
	return Value{num: uint64(TypeSliceBool)<<vshift | uint64(len(s)), ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceInt8 converts []int8 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceInt8[S []E, E ~int8](s S) Value {
	return Value{num: uint64(TypeSliceInt8)<<vshift | uint64(len(s)), ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceUint8 converts []uint8 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceUint8[S []E, E ~uint8](s S) Value {
	return Value{num: uint64(TypeSliceUint8)<<vshift | uint64(len(s)), ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceInt16 converts []int16 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceInt16[S []E, E ~int16](s S) Value {
	return Value{num: uint64(TypeSliceInt16)<<vshift | uint64(len(s)), ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceUint16 converts []uint16 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceUint16[S []E, E ~uint16](s S) Value {
	return Value{num: uint64(TypeSliceUint16)<<vshift | uint64(len(s)), ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceInt32 converts []int32 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceInt32[S []E, E ~int32](s S) Value {
	return Value{num: uint64(TypeSliceInt32)<<vshift | uint64(len(s)), ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceUint32 converts []uint32 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceUint32[S []E, E ~uint32](s S) Value {
	return Value{num: uint64(TypeSliceUint32)<<vshift | uint64(len(s)), ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceInt64 converts []int64 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceInt64[S []E, E ~int64](s S) Value {
	return Value{num: uint64(TypeSliceInt64)<<vshift | uint64(len(s)), ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceUint64 converts []uint64 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceUint64[S []E, E ~uint64](s S) Value {
	return Value{num: uint64(TypeSliceUint64)<<vshift | uint64(len(s)), ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceFloat32 converts []float32 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceFloat32[S []E, E ~float32](s S) Value {
	return Value{num: uint64(TypeSliceFloat32)<<vshift | uint64(len(s)), ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceFloat64 converts []float64 as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceFloat64[S []E, E ~float64](s S) Value {
	return Value{num: uint64(TypeSliceFloat64)<<vshift | uint64(len(s)), ptr: unsafe.Pointer(unsafe.SliceData(s))}
}

// SliceString converts []string as Value. This takes ownership of s, and the caller should not use s after this call.
func SliceString[S []E, E ~string](s S) Value {
	return Value{num: uint64(TypeSliceString)<<vshift | uint64(len(s)), ptr: unsafe.Pointer(unsafe.SliceData(s))}
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
