// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

// Bool is a special type to accommodate nullable/unset boolean value.
// 0 is false and 1 is true, other value should be treated as invalid 255.
type Bool uint8

const (
	BoolFalse   Bool = 0   // false
	BoolTrue    Bool = 1   // true
	BoolInvalid Bool = 255 // invalid or unset
)

func (b Bool) Uint8() uint8 { return uint8(b) }

func (b Bool) String() string {
	switch b {
	case BoolTrue:
		return "true"
	case BoolFalse:
		return "false"
	}
	return "BoolInvalid(" + strconv.Itoa(int(b)) + ")"
}

// BoolFromString parse string into Bool constant it represents, return BoolInvalid if not found.
func BoolFromString(s string) Bool {
	switch s {
	case "true":
		return BoolTrue
	case "false":
		return BoolFalse
	}
	return BoolInvalid
}

// BoolFromBool parse bool value into Bool.
func BoolFromBool(t bool) Bool {
	if t {
		return BoolTrue
	}
	return BoolFalse
}
