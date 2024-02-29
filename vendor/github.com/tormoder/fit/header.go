package fit

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/tormoder/fit/dyncrc16"
)

// Header represents a FIT file header.
type Header struct {
	Size            byte
	ProtocolVersion byte
	ProfileVersion  uint16
	DataSize        uint32
	DataType        [4]byte
	CRC             uint16
}

var (
	errNotFit     = FormatError("header data type was not '.FIT'")
	errHeaderSize = FormatError("illegal header size")
	errHdrCRC     = IntegrityError("header checksum failed")
	errReadSize   = ioError{op: "read size", err: io.ErrUnexpectedEOF}
	errReadData   = ioError{op: "read data", err: io.ErrUnexpectedEOF}
)

// NewHeader creates a new Header with required info.
func NewHeader(v ProtocolVersion, crc bool) Header {
	h := Header{
		ProtocolVersion: v.Version(),
		ProfileVersion:  ProfileVersion,
		Size:            headerSizeNoCRC,
	}
	if crc {
		h.Size = headerSizeCRC
	}
	copy(h.DataType[:], fitDataTypeString[:])
	return h
}

func (d *decoder) decodeHeader() error {
	err := binary.Read(d.r, le, &d.h.Size)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return errReadSize
		}
		return ioError{"reading size", err}
	}
	if d.h.Size != headerSizeCRC && d.h.Size != headerSizeNoCRC {
		return errHeaderSize
	}

	_, err = io.ReadFull(d.r, d.tmp[:d.h.Size-1])
	if err != nil {
		if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
			return errReadData
		}
		return ioError{"reading data", err}
	}

	if err = checkProtocolVersion(d.tmp[0]); err != nil {
		return err
	}
	d.h.ProtocolVersion = d.tmp[0]
	d.h.ProfileVersion = le.Uint16(d.tmp[1:3])
	d.h.DataSize = le.Uint32(d.tmp[3:7])

	if string(d.tmp[7:11]) != fitDataTypeString {
		return errNotFit
	}
	copy(d.h.DataType[:], d.tmp[7:11])

	d.crc.Write([]byte{d.h.Size})
	d.crc.Write(d.tmp[:d.h.Size-1])

	if d.h.Size == headerSizeNoCRC {
		return nil
	}

	d.h.CRC = le.Uint16(d.tmp[11 : d.h.Size-1])
	if d.h.CRC == 0x0000 {
		return nil
	}

	if d.crc.Sum16() != 0x0000 {
		return errHdrCRC
	}

	return nil
}

func checkProtocolVersion(b byte) error {
	fileProtoVer := ProtocolVersion(b)
	if fileProtoVer.Major() > CurrentProtocolVersion().Major() {
		err := fmt.Sprintf(
			"protocol version %v not supported by sdk protocol version %v",
			fileProtoVer,
			CurrentProtocolVersion(),
		)
		return NotSupportedError(err)
	}
	return nil
}

// CheckIntegrity verifies the FIT header CRC.
func (h Header) CheckIntegrity() error {
	if err := checkProtocolVersion(h.ProtocolVersion); err != nil {
		return err
	}
	if string(h.DataType[:len(h.DataType)]) != fitDataTypeString {
		return errNotFit
	}
	if h.Size == headerSizeNoCRC {
		return nil
	}
	if h.CRC == 0 {
		return nil
	}

	crc := dyncrc16.New()
	bh := make([]byte, h.Size)
	bh[0] = h.Size
	bh[1] = h.ProtocolVersion
	le.PutUint16(bh[2:4], h.ProfileVersion)
	le.PutUint32(bh[4:8], h.DataSize)
	copy(bh[8:12], h.DataType[:])
	le.PutUint16(bh[12:14], h.CRC)

	if crc.Sum16() != 0x0000 {
		return errHdrCRC
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (h Header) MarshalJSON() ([]byte, error) {
	e := newJSONEncodeState()
	e.open()
	e.writeFieldName("Size")
	e.writeUint(uint64(h.Size))
	e.c()
	e.writeFieldName("ProtocolVersion")
	e.writeUint(uint64(h.ProtocolVersion))
	e.c()
	e.writeFieldName("ProfileVersion")
	e.writeUint(uint64(h.ProfileVersion))
	e.c()
	e.writeFieldName("DataSize")
	e.writeUint(uint64(h.DataSize))
	e.c()
	e.writeFieldName("DataType")
	e.writeStringBytes(h.DataType[:4])
	e.c()
	e.writeFieldName("CRC")
	e.writeUint(uint64(h.CRC))
	e.close()
	return e.Bytes(), nil
}

func (h Header) String() string {
	return fmt.Sprintf(
		"size: %d | protover: %d | profver: %d | dsize: %d | dtype: %s | crc: 0x%x",
		h.Size, h.ProtocolVersion, h.ProfileVersion, h.DataSize, string(h.DataType[:]), h.CRC,
	)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (h Header) MarshalBinary() ([]byte, error) {
	buf := &bytes.Buffer{}

	err := binary.Write(buf, binary.LittleEndian, h.Size)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, h.ProtocolVersion)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, h.ProfileVersion)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, h.DataSize)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, h.DataType)
	if err != nil {
		return nil, err
	}

	h.CRC = dyncrc16.Checksum(buf.Bytes())

	if h.Size == headerSizeCRC {
		err = binary.Write(buf, binary.LittleEndian, h.CRC)
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}
