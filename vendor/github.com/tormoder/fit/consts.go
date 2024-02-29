package fit

import "fmt"

// ProfileVersion is the current supported profile version of the FIT SDK.
const ProfileVersion uint16 = ((ProfileMajorVersion * 100) + ProfileMinorVersion)

var currentProtocolVersion = V20

// CurrentProtocolVersion returns the current supported FIT protocol version.
func CurrentProtocolVersion() ProtocolVersion {
	return currentProtocolVersion
}

// ProtocolVersion represents the FIT protocol version.
type ProtocolVersion byte

// FIT protocol versions.
const (
	V10 ProtocolVersion = 0x10
	V20 ProtocolVersion = 0x20
)

// Version returns the full FIT protocol version encoded as a single byte.
func (p ProtocolVersion) Version() byte {
	return byte(p)
}

// Major returns the major FIT protocol version.
func (p ProtocolVersion) Major() byte {
	return byte(p&protocolVersionMajorMask) >> protocolVersionMajorShift
}

// Minor returns the minor FIT protocol version.
func (p ProtocolVersion) Minor() byte {
	return byte(p & protocolVersionMinorMask)
}

func (p ProtocolVersion) String() string {
	return fmt.Sprintf("%d.%d", p.Major(), p.Minor())
}

const (
	protocolVersionMajorShift = 4
	protocolVersionMajorMask  = 0x0F << protocolVersionMajorShift
	protocolVersionMinorMask  = 0x0F
)

const (
	headerTypeMask             byte = 0xF0
	compressedHeaderMask       byte = 0x80
	compressedTimeMask         byte = 0x1F
	compressedLocalMesgNumMask byte = 0x60

	mesgDefinitionMask byte = 0x40
	devDataMask        byte = 0x20
	mesgHeaderMask     byte = 0x00
	localMesgNumMask   byte = 0x0F

	maxLocalMesgs byte = localMesgNumMask + 1

	littleEndian byte = 0x00
	bigEndian    byte = 0x01

	bytesForCRC     byte = 2
	headerSizeCRC   byte = 14
	headerSizeNoCRC byte = headerSizeCRC - bytesForCRC

	fitDataTypeString string = ".FIT"

	fieldNumTimeStamp byte = 253
)
