// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package factory contains predefined messages based on messages in [Profile.xlsx] provided in the Official FIT SDK.
// This package handles message and field creation, as well as registering manufacturer specific messages to be used for
// certain manufacturer's generated FIT file. You can always generate your own custom SDK using cmd/fitgen, but for those
// who prefer using this SDK as it is, what this package provides might be sufficient.
package factory
