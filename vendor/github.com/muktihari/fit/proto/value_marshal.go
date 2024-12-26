// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"encoding/binary"
	"fmt"
	"math"
)

const ErrTypeNotSupported = errorString("type is not supported")

// MarshalAppend appends the FIT format encoding of Value to b. Returning the result.
// If arch is 0, marshal in Little-Endian, otherwise marshal in Big-Endian.
func (v Value) MarshalAppend(b []byte, arch byte) ([]byte, error) {
	// NOTE: The size of the resulting bytes should align with Sizeof.
	switch v.typ {
	case TypeBool:
		val := v.Bool()
		if val > 1 {
			b = append(b, 255)
		} else {
			b = append(b, byte(val))
		}
		return b, nil
	case TypeInt8:
		return append(b, uint8(v.num)), nil
	case TypeUint8:
		return append(b, uint8(v.num)), nil
	case TypeInt16:
		if arch == LittleEndian {
			b = binary.LittleEndian.AppendUint16(b, uint16(v.num))
		} else {
			b = binary.BigEndian.AppendUint16(b, uint16(v.num))
		}
		return b, nil
	case TypeUint16:
		if arch == LittleEndian {
			b = binary.LittleEndian.AppendUint16(b, uint16(v.num))
		} else {
			b = binary.BigEndian.AppendUint16(b, uint16(v.num))
		}
		return b, nil
	case TypeInt32:
		if arch == LittleEndian {
			b = binary.LittleEndian.AppendUint32(b, uint32(v.num))
		} else {
			b = binary.BigEndian.AppendUint32(b, uint32(v.num))
		}
		return b, nil
	case TypeUint32:
		if arch == LittleEndian {
			b = binary.LittleEndian.AppendUint32(b, uint32(v.num))
		} else {
			b = binary.BigEndian.AppendUint32(b, uint32(v.num))
		}
		return b, nil
	case TypeInt64:
		if arch == LittleEndian {
			b = binary.LittleEndian.AppendUint64(b, uint64(v.num))
		} else {
			b = binary.BigEndian.AppendUint64(b, uint64(v.num))
		}
		return b, nil
	case TypeUint64:
		if arch == LittleEndian {
			b = binary.LittleEndian.AppendUint64(b, v.num)
		} else {
			b = binary.BigEndian.AppendUint64(b, v.num)
		}
		return b, nil
	case TypeFloat32:
		if arch == LittleEndian {
			b = binary.LittleEndian.AppendUint32(b, uint32(v.num))
		} else {
			b = binary.BigEndian.AppendUint32(b, uint32(v.num))
		}
		return b, nil
	case TypeFloat64:
		if arch == LittleEndian {
			b = binary.LittleEndian.AppendUint64(b, v.num)
		} else {
			b = binary.BigEndian.AppendUint64(b, v.num)
		}
		return b, nil
	case TypeString:
		val := v.String()
		n := len(val)
		b = append(b, val...)
		if n == 0 || val[n-1] != 0 {
			b = append(b, 0) // add utf-8 null-terminated string
		}
		return b, nil
	case TypeSliceBool:
		vals := v.SliceBool()
		for i := range vals {
			if vals[i] > 1 {
				b = append(b, 255)
			} else {
				b = append(b, byte(vals[i]))
			}
		}
		return b, nil
	case TypeSliceInt8:
		vals := v.SliceInt8()
		for i := range vals {
			b = append(b, uint8(vals[i]))
		}
		return b, nil
	case TypeSliceUint8:
		return append(b, v.SliceUint8()...), nil
	case TypeSliceInt16:
		vals := v.SliceInt16()
		if arch == LittleEndian {
			for i := range vals {
				b = binary.LittleEndian.AppendUint16(b, uint16(vals[i]))
			}
		} else {
			for i := range vals {
				b = binary.BigEndian.AppendUint16(b, uint16(vals[i]))
			}
		}
		return b, nil
	case TypeSliceUint16:
		vals := v.SliceUint16()
		if arch == LittleEndian {
			for i := range vals {
				b = binary.LittleEndian.AppendUint16(b, vals[i])
			}
		} else {
			for i := range vals {
				b = binary.BigEndian.AppendUint16(b, vals[i])
			}
		}
		return b, nil
	case TypeSliceInt32:
		vals := v.SliceInt32()
		if arch == LittleEndian {
			for i := range vals {
				b = binary.LittleEndian.AppendUint32(b, uint32(vals[i]))
			}
		} else {
			for i := range vals {
				b = binary.BigEndian.AppendUint32(b, uint32(vals[i]))
			}
		}
		return b, nil
	case TypeSliceUint32:
		vals := v.SliceUint32()
		if arch == LittleEndian {
			for i := range vals {
				b = binary.LittleEndian.AppendUint32(b, vals[i])
			}
		} else {
			for i := range vals {
				b = binary.BigEndian.AppendUint32(b, vals[i])
			}
		}
		return b, nil
	case TypeSliceInt64:
		vals := v.SliceInt64()
		if arch == LittleEndian {
			for i := range vals {
				b = binary.LittleEndian.AppendUint64(b, uint64(vals[i]))
			}
		} else {
			for i := range vals {
				b = binary.BigEndian.AppendUint64(b, uint64(vals[i]))
			}
		}
		return b, nil
	case TypeSliceUint64:
		vals := v.SliceUint64()
		if arch == LittleEndian {
			for i := range vals {
				b = binary.LittleEndian.AppendUint64(b, vals[i])
			}
		} else {
			for i := range vals {
				b = binary.BigEndian.AppendUint64(b, vals[i])
			}
		}
		return b, nil
	case TypeSliceFloat32:
		vals := v.SliceFloat32()
		if arch == LittleEndian {
			for i := range vals {
				b = binary.LittleEndian.AppendUint32(b, math.Float32bits(vals[i]))
			}
		} else {
			for i := range vals {
				b = binary.BigEndian.AppendUint32(b, math.Float32bits(vals[i]))
			}
		}
		return b, nil
	case TypeSliceFloat64:
		vals := v.SliceFloat64()
		if arch == LittleEndian {
			for i := range vals {
				b = binary.LittleEndian.AppendUint64(b, math.Float64bits(vals[i]))
			}
		} else {
			for i := range vals {
				b = binary.BigEndian.AppendUint64(b, math.Float64bits(vals[i]))
			}
		}
		return b, nil
	case TypeSliceString:
		vals := v.SliceString()
		for i := range vals {
			n := len(vals[i])
			b = append(b, vals[i]...)
			if n == 0 || vals[i][n-1] != 0 {
				b = append(b, 0) // add utf-8 null-terminated string
			}
		}
		if len(vals) == 0 {
			b = append(b, 0) // add utf-8 null-terminated string
		}
		return b, nil
	default:
		return b, fmt.Errorf("type Value(%T) is not supported: %w", v.Type(), ErrTypeNotSupported)
	}
}
