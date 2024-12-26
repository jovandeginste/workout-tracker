// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"github.com/muktihari/fit/proto"
)

// bits is 2048 bits value implementation, large enough to hold proto.Value in its integer form.
// This bits value enable us to do bitwise operation and it's used for component expansion as
// Field's Value requiring expansion can hold up to 255 bytes (2040 bits) data, this is obviously
// way more bits than Go's primitive value can handle.
//
// In Profile.xlsx v21.141, the biggest value for component expansion is "raw_bbi" message for
// having 240 bits on "data" and the second is "hr" message for having 120 bits on "event_timestamp_12".
type bits struct {
	// We use array to avoid memory allocation. It's simple to maintain and more deterministic.
	store [32]uint64
}

// makeBits creates 2048 bits value from proto.Value.
func makeBits(value proto.Value) (v bits, ok bool) {
	switch value.Type() {
	case proto.TypeInt8:
		v.store[0] = uint64(value.Int8())
	case proto.TypeUint8:
		v.store[0] = uint64(value.Uint8())
	case proto.TypeInt16:
		v.store[0] = uint64(value.Int16())
	case proto.TypeUint16:
		v.store[0] = uint64(value.Uint16())
	case proto.TypeInt32:
		v.store[0] = uint64(value.Int32())
	case proto.TypeUint32:
		v.store[0] = uint64(value.Uint32())
	case proto.TypeInt64:
		v.store[0] = uint64(value.Int64())
	case proto.TypeUint64:
		v.store[0] = value.Uint64()
	case proto.TypeFloat32:
		v.store[0] = uint64(value.Float32())
	case proto.TypeFloat64:
		v.store[0] = uint64(value.Float64())
	case proto.TypeSliceInt8:
		v.store = storeFromSlice(value.SliceInt8(), 1)
	case proto.TypeSliceUint8:
		v.store = storeFromSlice(value.SliceUint8(), 1)
	case proto.TypeSliceInt16:
		v.store = storeFromSlice(value.SliceInt16(), 2)
	case proto.TypeSliceUint16:
		v.store = storeFromSlice(value.SliceUint16(), 2)
	case proto.TypeSliceInt32:
		v.store = storeFromSlice(value.SliceInt32(), 4)
	case proto.TypeSliceUint32:
		v.store = storeFromSlice(value.SliceUint32(), 4)
	case proto.TypeSliceInt64:
		v.store = storeFromSlice(value.SliceInt64(), 8)
	case proto.TypeSliceUint64:
		v.store = storeFromSlice(value.SliceUint64(), 8)
	case proto.TypeSliceFloat32:
		v.store = storeFromSlice(value.SliceFloat32(), 4)
	case proto.TypeSliceFloat64:
		v.store = storeFromSlice(value.SliceFloat64(), 8)
	default:
		return v, false
	}
	return v, true
}

type numeric interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

// storeFromSlice creates value store from given s (slice of supported numeric type).
func storeFromSlice[S []E, E numeric](s S, bitsize uint8) (store [32]uint64) {
	var index, pos uint8
	for len(s) > 0 && index < 32 {
		store[index] |= uint64(s[0]) << (pos * 8)
		pos += bitsize
		if pos == 8 {
			index, pos = index+1, 0
		}
		s = s[1:]
	}
	return store
}

// Pull retrieves a value of the specified bit size from the value store and
// the value store will be updated accordingly.
func (v *bits) Pull(bitsize byte) (val uint32) {
	mask := uint64(1)<<bitsize - 1  // e.g. (1 << 8) - 1     = 255
	val = uint32(v.store[0] & mask) // e.g. 0x27010E08 & 255 = 0x08
	v.store[0] >>= bitsize          // e.g. 0x27010E08 >> 8  = 0x27010E

	for i := 1; i < len(v.store); i++ {
		if v.store[i] == 0 {
			continue
		}
		// e.g. 128 bits Layout Before: 0x0000_0000_0000_FFFF_0000_0000_2701_0E08
		hi := v.store[i] & mask    // e.g. 0x0000_0000_0000_FFFF & 0xFF                  = 0x0000_0000_0000_00FF
		lo := hi << (64 - bitsize) // e,g. 0x0000_0000_0000_00FF << (64 - 8)             = 0xFF00_0000_0000_0000
		v.store[i-1] |= lo         // e.g. 0x0000_0000_0027_010E | 0xFF00_0000_0000_0000 = 0xFF00_0000_0027_010E
		v.store[i] >>= bitsize     // e.g. 0x0000_0000_0000_FFFF >> 8                    = 0x0000_0000_0000_00FF
		// e.g. 128 bits Layout After:  0x0000_0000_0000_00FF_FF00_0000_0027_010E
	}

	return val
}
