// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

// Version is FIT Protocol Version
type Version byte

const (
	vMajorShift = 4
	vMinorMask  = 1<<vMajorShift - 1

	V1   Version = 1 << vMajorShift // V1 is Version 1.0
	V2   Version = 2 << vMajorShift // V2 is Version 2.0
	Vmax         = V2               // Vmax is an alias for the current latest version.
)

// CreateVersion creates version from major and minor value. Each value is 4 bits value (max: 15).
func CreateVersion(major, minor byte) Version {
	return Version(major<<vMajorShift | minor&vMinorMask)
}

// Major returns major value.
func (v Version) Major() byte { return byte(v >> vMajorShift) }

// Minor returns minor value.
func (v Version) Minor() byte { return byte(v & vMinorMask) }
