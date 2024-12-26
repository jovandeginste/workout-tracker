// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package profile defines all the types and messages generated from Profile.xlsx using fitgen.
// Additionally, supplementary items related to "Global Profile" that are not declared in the
// Profile.xlsx itself but are provided by Garmin or its affiliates website, such as common file
// types, are added manually.
//
// The Version constant indicates the version to which this profile refers, represented by the
// value corresponding to the Official FIT SDK's version.
//
// The ProfileType defines an abstraction layer of types used in the FIT file above the base types
// (primitive-types) such as sint, uint, etc. Here is an example to help understanding it better:
//   - Type DateTime is a time representation decoded in uint32 format in the FIT binary proto.
//     The value of uint32 is a number counted since FIT Epoch (time since 31 Dec 1989 00:00:000 UTC).
//
// Using an abstraction layer like the profile type allows time to be stored in binary files as
// compact as uint32 values. This means that when we encounter a field with the DateTime type,
// we can decode it into time.Time{}.
package profile
