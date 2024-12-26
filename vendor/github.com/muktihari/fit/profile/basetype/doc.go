// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package basetype defines the base of all types used in FIT. It is manually coded since Profile.xlsx does not provided
// same level details as it is defined on https://developer.garmin.com/fit/protocol.
//
// This package refer to SDK -> Profile.xlsx (Sheet Name: "Types", Type Name (Column A): "fit_base_type") and is segregated since it's a special type.
// When code generation is generating types from Profile.xlsx and encounter type name "fit_base_type",
// that type name should not be generated to avoid unnecessary redundancy and confusion.
package basetype
