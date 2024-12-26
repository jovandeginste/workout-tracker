// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package semicircles

import (
	"math"

	"github.com/muktihari/fit/profile/basetype"
)

const (
	piRadians        = 1 << 31 // 2^31; 31 bit representation
	conversionFactor = 180.0 / piRadians
)

// ToDegrees converts semicircles into degrees value. If semicircles equals to
// basetype.Sint32Invalid, float64 invalid value will be returned.
func ToDegrees(semicircles int32) float64 {
	if semicircles == basetype.Sint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(semicircles) * conversionFactor
}

// ToSemicircles converts degrees into semicircles value. If degrees equals to
// either IsNaN, IsInf or in integer form it is equals to float64 invalid value
// in integer form, basetype.Sint32Invalid will be returned.
func ToSemicircles(degrees float64) int32 {
	if math.Float64bits(degrees) == basetype.Float64Invalid ||
		math.IsNaN(degrees) || math.IsInf(degrees, 0) {
		return basetype.Sint32Invalid
	}
	return int32(degrees / conversionFactor)
}
