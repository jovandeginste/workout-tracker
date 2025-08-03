// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// Accumulator is value accumulator.
type Accumulator struct {
	values []value // use slice over map since len(values) is relatively small
}

// NewAccumulator creates new accumulator.
func NewAccumulator() *Accumulator {
	return &Accumulator{}
}

// Collect collects value, it will either append the value when not exist or replace existing one.
func (a *Accumulator) Collect(mesgNum typedef.MesgNum, fieldNum byte, val uint32) {
	for i := range a.values {
		av := &a.values[i]
		if av.mesgNum == mesgNum && av.fieldNum == fieldNum {
			av.value = val
			av.last = val
			return
		}
	}
	a.values = append(a.values, value{
		mesgNum:  mesgNum,
		fieldNum: fieldNum,
		value:    val,
		last:     val,
	})
}

// Collect collects value, it will either append the value when not exist or replace existing one.
func (a *Accumulator) CollectValue(mesgNum typedef.MesgNum, fieldNum byte, val proto.Value) {
	switch val.Type() {
	case proto.TypeInt8:
		a.Collect(mesgNum, fieldNum, uint32(val.Int8()))
	case proto.TypeUint8:
		a.Collect(mesgNum, fieldNum, uint32(val.Uint8()))
	case proto.TypeInt16:
		a.Collect(mesgNum, fieldNum, uint32(val.Int16()))
	case proto.TypeUint16:
		a.Collect(mesgNum, fieldNum, uint32(val.Uint16()))
	case proto.TypeInt32:
		a.Collect(mesgNum, fieldNum, uint32(val.Int32()))
	case proto.TypeUint32:
		a.Collect(mesgNum, fieldNum, uint32(val.Uint32()))
	case proto.TypeInt64:
		a.Collect(mesgNum, fieldNum, uint32(val.Int64()))
	case proto.TypeUint64:
		a.Collect(mesgNum, fieldNum, uint32(val.Uint64()))
	case proto.TypeFloat32:
		a.Collect(mesgNum, fieldNum, uint32(val.Float32()))
	case proto.TypeFloat64:
		a.Collect(mesgNum, fieldNum, uint32(val.Float64()))
	case proto.TypeSliceInt8:
		vals := val.SliceInt8()
		if n := len(vals); n > 0 {
			a.Collect(mesgNum, fieldNum, uint32(vals[n-1]))
		}
	case proto.TypeSliceUint8:
		vals := val.SliceUint8()
		if n := len(vals); n > 0 {
			a.Collect(mesgNum, fieldNum, uint32(vals[n-1]))
		}
	case proto.TypeSliceInt16:
		vals := val.SliceInt16()
		if n := len(vals); n > 0 {
			a.Collect(mesgNum, fieldNum, uint32(vals[n-1]))
		}
	case proto.TypeSliceUint16:
		vals := val.SliceUint16()
		if n := len(vals); n > 0 {
			a.Collect(mesgNum, fieldNum, uint32(vals[n-1]))
		}
	case proto.TypeSliceInt32:
		vals := val.SliceInt32()
		if n := len(vals); n > 0 {
			a.Collect(mesgNum, fieldNum, uint32(vals[n-1]))
		}
	case proto.TypeSliceUint32:
		vals := val.SliceUint32()
		if n := len(vals); n > 0 {
			a.Collect(mesgNum, fieldNum, uint32(vals[n-1]))
		}
	case proto.TypeSliceInt64:
		vals := val.SliceInt64()
		if n := len(vals); n > 0 {
			a.Collect(mesgNum, fieldNum, uint32(vals[n-1]))
		}
	case proto.TypeSliceUint64:
		vals := val.SliceUint64()
		if n := len(vals); n > 0 {
			a.Collect(mesgNum, fieldNum, uint32(vals[n-1]))
		}
	case proto.TypeSliceFloat32:
		vals := val.SliceFloat32()
		if n := len(vals); n > 0 {
			a.Collect(mesgNum, fieldNum, uint32(vals[n-1]))
		}
	case proto.TypeSliceFloat64:
		vals := val.SliceFloat64()
		if n := len(vals); n > 0 {
			a.Collect(mesgNum, fieldNum, uint32(vals[n-1]))
		}
	}
}

// Accumulate calculates the accumulated value and update it accordingly. If targeted value
// does not exist, it will be collected and the original value will be returned.
func (a *Accumulator) Accumulate(mesgNum typedef.MesgNum, destFieldNum byte, val uint32, bits byte) uint32 {
	for i := range a.values {
		av := &a.values[i]
		if av.mesgNum == mesgNum && av.fieldNum == destFieldNum {
			var mask uint32 = (1 << bits) - 1
			av.value += (val - av.last) & mask
			av.last = val
			return av.value
		}
	}
	a.values = append(a.values, value{
		mesgNum:  mesgNum,
		fieldNum: destFieldNum,
		value:    val,
		last:     val,
	})
	return val
}

// Reset resets the accumulator. It retains the underlying storage for use by
// future use to reduce memory allocs.
func (a *Accumulator) Reset() { a.values = a.values[:0] }

type value struct {
	mesgNum  typedef.MesgNum
	fieldNum byte
	last     uint32
	value    uint32
}
