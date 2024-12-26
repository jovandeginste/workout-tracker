// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"encoding/binary"
	"fmt"
	"math"
	"unicode/utf8"

	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
)

// UnmarshalValue unmarshals b into a proto.Value.
// The caller should ensure that the len(b) matches its corresponding base type's size, otherwise it might panic.
func UnmarshalValue(b []byte, arch byte, baseType basetype.BaseType, profileType profile.ProfileType, isArray bool) (Value, error) {
	switch baseType {
	case basetype.Sint8:
		if isArray {
			vals := make([]int8, 0, len(b))
			for i := range b {
				vals = append(vals, int8(b[i]))
			}
			return SliceInt8(vals), nil
		}
		return Int8(int8(b[0])), nil
	case basetype.Enum, basetype.Byte,
		basetype.Uint8, basetype.Uint8z:
		if profileType == profile.Bool { // Special Case
			if isArray {
				vals := make([]typedef.Bool, 0, len(b))
				for i := range b {
					vals = append(vals, typedef.Bool(b[i]))
				}
				return SliceBool(vals), nil
			}
			return Bool(typedef.Bool(b[0])), nil
		}
		if isArray {
			vals := make([]byte, len(b))
			copy(vals, b)
			return SliceUint8(vals), nil
		}
		return Uint8(b[0]), nil
	case basetype.Sint16:
		if isArray {
			const n = 2
			vals := make([]int16, 0, len(b)/n)
			if arch == LittleEndian {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, int16(binary.LittleEndian.Uint16(b[:n])))
				}
			} else {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, int16(binary.BigEndian.Uint16(b[:n])))
				}
			}
			return SliceInt16(vals), nil
		}
		if arch == LittleEndian {
			return Int16(int16(binary.LittleEndian.Uint16(b))), nil
		}
		return Int16(int16(binary.BigEndian.Uint16(b))), nil
	case basetype.Uint16, basetype.Uint16z:
		if isArray {
			const n = 2
			vals := make([]uint16, 0, len(b)/n)
			if arch == LittleEndian {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, binary.LittleEndian.Uint16(b[:n]))
				}
			} else {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, binary.BigEndian.Uint16(b[:n]))
				}
			}
			return SliceUint16(vals), nil
		}
		if arch == LittleEndian {
			return Uint16(binary.LittleEndian.Uint16(b)), nil
		}
		return Uint16(binary.BigEndian.Uint16(b)), nil
	case basetype.Sint32:
		if isArray {
			const n = 4
			vals := make([]int32, 0, len(b)/n)
			if arch == LittleEndian {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, int32(binary.LittleEndian.Uint32(b[:n])))
				}
			} else {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, int32(binary.BigEndian.Uint32(b[:n])))
				}
			}
			return SliceInt32(vals), nil
		}
		if arch == LittleEndian {
			return Int32(int32(binary.LittleEndian.Uint32(b))), nil
		}
		return Int32(int32(binary.BigEndian.Uint32(b))), nil
	case basetype.Uint32, basetype.Uint32z:
		if isArray {
			const n = 4
			vals := make([]uint32, 0, len(b)/n)
			if arch == LittleEndian {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, binary.LittleEndian.Uint32(b[:n]))
				}
			} else {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, binary.BigEndian.Uint32(b[:n]))
				}
			}
			return SliceUint32(vals), nil
		}
		if arch == LittleEndian {
			return Uint32(binary.LittleEndian.Uint32(b)), nil
		}
		return Uint32(binary.BigEndian.Uint32(b)), nil
	case basetype.Sint64:
		if isArray {
			const n = 8
			vals := make([]int64, 0, len(b)/n)
			if arch == LittleEndian {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, int64(binary.LittleEndian.Uint64(b[:n])))
				}
			} else {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, int64(binary.BigEndian.Uint64(b[:n])))
				}
			}
			return SliceInt64(vals), nil
		}
		if arch == LittleEndian {
			return Int64(int64(binary.LittleEndian.Uint64(b))), nil
		}
		return Int64(int64(binary.BigEndian.Uint64(b))), nil
	case basetype.Uint64, basetype.Uint64z:
		if isArray {
			const n = 8
			vals := make([]uint64, 0, len(b)/n)
			if arch == LittleEndian {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, binary.LittleEndian.Uint64(b[:n]))
				}
			} else {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, binary.BigEndian.Uint64(b[:n]))
				}
			}
			return SliceUint64(vals), nil
		}
		if arch == LittleEndian {
			return Uint64(binary.LittleEndian.Uint64(b)), nil
		}
		return Uint64(binary.BigEndian.Uint64(b)), nil
	case basetype.Float32:
		if isArray {
			const n = 4
			vals := make([]float32, 0, len(b)/n)
			if arch == LittleEndian {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, math.Float32frombits(binary.LittleEndian.Uint32(b[:n])))
				}
			} else {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, math.Float32frombits(binary.BigEndian.Uint32(b[:n])))
				}
			}
			return SliceFloat32(vals), nil
		}
		if arch == LittleEndian {
			return Float32(math.Float32frombits(binary.LittleEndian.Uint32(b))), nil
		}
		return Float32(math.Float32frombits(binary.BigEndian.Uint32(b))), nil
	case basetype.Float64:
		if isArray {
			const n = 8
			vals := make([]float64, 0, len(b)/n)
			if arch == LittleEndian {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, math.Float64frombits(binary.LittleEndian.Uint64(b[:n])))
				}
			} else {
				for ; len(b) >= n; b = b[n:] {
					vals = append(vals, math.Float64frombits(binary.BigEndian.Uint64(b[:n])))
				}
			}
			return SliceFloat64(vals), nil
		}
		if arch == LittleEndian {
			return Float64(math.Float64frombits(binary.LittleEndian.Uint64(b))), nil
		}
		return Float64(math.Float64frombits(binary.BigEndian.Uint64(b))), nil
	case basetype.String:
		if isArray {
			var size byte
			last := 0
			for i := range b {
				if b[i] == '\x00' {
					if last != i { // only if not an invalid string
						size++
					}
					last = i + 1
				}
			}
			last = 0
			vals := make([]string, 0, size)
			for i := range b {
				if b[i] == '\x00' {
					if last != i { // only if not an invalid string
						if s := utf8String(b[last:i]); len(s) > 0 {
							vals = append(vals, s)
						}
					}
					last = i + 1
				}
			}
			return SliceString(vals), nil
		}
		return String(utf8String(b)), nil
	}

	return Value{}, fmt.Errorf("type %s(%d) is not supported: %w", baseType, baseType, ErrTypeNotSupported)
}

// utf8String converts b into a valid UTF-8 string. If it encounters
// utf8.RuneError character it will discard it. It stops when all bytes
// have been successfully decoded or reach a null-terminated string '\x00'.
func utf8String(b []byte) string {
	buf := make([]byte, 0, 255)
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		if r == 0 {
			break
		}
		if r != utf8.RuneError {
			buf = utf8.AppendRune(buf, r)
		}
		b = b[size:]
	}
	return string(buf)
}
