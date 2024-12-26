// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

// Version is FIT Protocol Version
type Version byte

const (
	ErrProtocolVersionNotSupported = errorString("protocol version not supported")

	MajorVersionShift = 4
	MajorVersionMask  = 0x0F << MajorVersionShift
	MinorVersionMask  = 0x0F

	V1   Version = 1 << MajorVersionShift // V1 is Version 1.0
	V2   Version = 2 << MajorVersionShift // V2 is Version 2.0
	Vmax         = V2                     // Vmax is an alias for the current latest version.
)

// CreateVersion creates version from major and minor value, it can only create version up < Vmax.
func CreateVersion(major, minor byte) (Version, bool) {
	version := Version((major << MajorVersionShift) | minor)
	if version > Vmax {
		return 0, false
	}
	return version, true
}

// Validate checks whether given version is a valid version.
func Validate(version Version) error {
	if VersionMajor(version) > VersionMajor(Vmax) {
		return ErrProtocolVersionNotSupported
	}
	return nil
}

// VersionMajor returns major value of given version
func VersionMajor(version Version) byte {
	return byte(version >> MajorVersionShift)
}

// VersionMinor returns minor value of given version
func VersionMinor(version Version) byte {
	return byte(version & MinorVersionMask)
}
