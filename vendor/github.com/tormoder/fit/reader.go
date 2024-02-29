package fit

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
	"sort"
	"time"

	"github.com/tormoder/fit/dyncrc16"
	"github.com/tormoder/fit/internal/types"
)

var (
	le = binary.LittleEndian
	be = binary.BigEndian
)

type decoder struct {
	r     io.Reader
	bytes struct {
		limit int
		n     int
		buf   [4096]byte
		i, j  int
	}

	crc     dyncrc16.Hash16
	tmp     [255 * 3]byte
	defmsgs [maxLocalMesgs]*defmsg

	timestamp      uint32
	lastTimeOffset int32

	opts  decodeOptions
	debug bool

	unknownFields   map[unknownField]int
	unknownMessages map[MesgNum]int

	h    Header
	file *File
}

// CheckIntegrity verifies the FIT header and file CRC. Only the header CRC is
// verified if headerOnly is true.
func CheckIntegrity(r io.Reader, headerOnly bool) error {
	var d decoder
	return d.decode(r, headerOnly, false, true)
}

// DecodeHeader returns the FIT file header without decoding the entire FIT
// file.
func DecodeHeader(r io.Reader) (Header, error) {
	var d decoder
	if err := d.decode(r, true, false, false); err != nil {
		return Header{}, err
	}
	return d.h, nil
}

// DecodeHeaderAndFileID returns the FIT file header and FileId message without
// decoding the entire FIT file. The FileId message must be present in all FIT
// files.
func DecodeHeaderAndFileID(r io.Reader) (Header, FileIdMsg, error) {
	var d decoder
	if err := d.decode(r, false, true, false); err != nil {
		return Header{}, FileIdMsg{}, err
	}
	return d.h, d.file.FileId, nil
}

// Decode reads a FIT file from r and returns it as a *File.
// If error is non-nil, all data decoded before the error was
// encountered is also returned.
func Decode(r io.Reader, opts ...DecodeOption) (*File, error) {
	var d decoder
	for _, opt := range opts {
		opt(&d.opts)
	}
	err := d.decode(r, false, false, false)
	return d.file, err
}

// DecodeChained reads chained FIT files from r until an error is encountered
// or no more data is available. If error is non-nil, all data decoded before
// the error was encountered is also returned for the last file read.
func DecodeChained(r io.Reader, opts ...DecodeOption) ([]*File, error) {
	var fitFiles []*File
	var i int
	for {
		var d decoder
		for _, opt := range opts {
			opt(&d.opts)
		}
		err := d.decode(r, false, false, false)
		if err != nil {
			if d.h.Size == 0 && i != 0 {
				// Header not read, and not first file:
				// EOF, no more data.
				return fitFiles, nil
			}
			if d.file != nil {
				fitFiles = append(fitFiles, d.file)
			}
			return fitFiles, fmt.Errorf("error parsing chained fit: file #%d: %v", i+1, err)
		}
		fitFiles = append(fitFiles, d.file)
		i++
	}
}

func (d *decoder) decode(r io.Reader, headerOnly, fileIDOnly, crcOnly bool) error {
	if d.opts.logger != nil {
		d.debug = true
	}

	d.r = r
	d.crc = dyncrc16.New()

	err := d.decodeHeader()
	if err != nil {
		return fmt.Errorf("error decoding header: %v", err)
	}

	d.file = new(File)
	d.file.Header = d.h
	d.bytes.limit = int(d.h.DataSize)

	if d.debug {
		d.opts.logger.Println("header decoded:", d.h)
	}

	if headerOnly {
		return nil
	}

	if crcOnly {
		_, err = io.CopyN(d.crc, d.r, int64(d.h.DataSize))
		if err != nil {
			return fmt.Errorf("error parsing data: %v", err)
		}
		return d.checkCRC()
	}

	if d.opts.unknownFields {
		d.unknownFields = make(map[unknownField]int)
		defer d.handleUnknownFields()
	}
	if d.opts.unknownMessages {
		d.unknownMessages = make(map[MesgNum]int)
		defer d.handleUnknownMessages()
	}

	err = d.parseFileIdMsg()
	if err != nil {
		return fmt.Errorf("error parsing file id message: %v", err)
	}
	if fileIDOnly {
		return nil
	}

	err = d.file.init()
	if err != nil {
		return err
	}

	err = d.decodeFileData()
	if err != nil {
		return err
	}

	// Check invariant pre-read CRC:
	if !crcOnly && d.bytes.n != d.bytes.limit {
		fatalErr := fmt.Sprintf("internal decoder error: pre-crc check: data size limit is %d, but n is %d", d.bytes.limit, d.bytes.n)
		panic(fatalErr)
	}

	return d.checkCRC()
}

func (d *decoder) decodeFileData() error {
	for d.bytes.n < d.bytes.limit {
		var (
			b   byte
			dm  *defmsg
			msg reflect.Value
			err error
		)

		b, err = d.readByte()
		if err != nil {
			return fmt.Errorf("error parsing record header: %v", err)
		}

		switch {
		case (b & compressedHeaderMask) == compressedHeaderMask:
			msg, err = d.parseDataMessage(b, true)
			if err != nil {
				return fmt.Errorf("parsing compressed timestamp message: %v", err)
			}
			if msg.IsValid() {
				d.file.add(msg)
			}
		case (b & mesgDefinitionMask) == mesgDefinitionMask:
			dm, err = d.parseDefinitionMessage(b)
			if err != nil {
				return fmt.Errorf("parsing definition message: %v", err)
			}
			d.defmsgs[dm.localMsgType] = dm
		case (b & mesgDefinitionMask) == mesgHeaderMask:
			msg, err = d.parseDataMessage(b, false)
			if err != nil {
				return fmt.Errorf("parsing data message: %v", err)
			}
			if msg.IsValid() {
				d.file.add(msg)
			}
		default:
			return fmt.Errorf("unknown record header, got: %#x", b)
		}
	}

	return nil
}

func (d *decoder) checkCRC() error {
	if d.debug {
		d.opts.logger.Printf("expecting crc value: 0x%x", d.crc.Sum16())
	}
	if _, err := io.ReadFull(d.r, d.tmp[:bytesForCRC]); err != nil {
		err = noEOF(err)
		return fmt.Errorf("error parsing file CRC: %v", err)
	}
	d.crc.Write(d.tmp[:bytesForCRC])
	d.file.CRC = le.Uint16(d.tmp[:bytesForCRC])
	if d.debug {
		d.opts.logger.Printf("read crc value: 0x%x", d.file.CRC)
	}
	if d.crc.Sum16() != 0x0000 {
		return IntegrityError("file checksum failed")
	}

	return nil
}

func (d *decoder) fill() error {
	if d.bytes.i != d.bytes.j {
		panic("internal decoder error: fill called when unread bytes exist")
	}
	if d.bytes.n == d.bytes.limit {
		return FormatError("fit file requested data beyond data size listed in header")
	}

	d.bytes.i, d.bytes.j = 0, 0
	end := len(d.bytes.buf)
	max := d.bytes.limit - d.bytes.n
	if max < end {
		end = max
	}

	n, err := d.r.Read(d.bytes.buf[d.bytes.i:end])
	d.bytes.j += n
	if n > 0 {
		err = nil
		d.crc.Write(d.bytes.buf[d.bytes.i:d.bytes.j])
	}

	return err
}

func (d *decoder) readByte() (byte, error) {
	for d.bytes.i == d.bytes.j {
		if err := d.fill(); err != nil {
			err = noEOF(err)
			return 0, err
		}
	}
	x := d.bytes.buf[d.bytes.i]
	d.bytes.i++
	d.bytes.n++
	return x, nil
}

func (d *decoder) skipByte() error {
	for d.bytes.i == d.bytes.j {
		if err := d.fill(); err != nil {
			err = noEOF(err)
			return err
		}
	}
	d.bytes.i++
	d.bytes.n++
	return nil
}

func (d *decoder) readFull(p []byte) error {
	for {
		n := copy(p, d.bytes.buf[d.bytes.i:d.bytes.j])
		p = p[n:]
		d.bytes.i += n
		d.bytes.n += n
		if len(p) == 0 {
			break
		}
		if err := d.fill(); err != nil {
			err = noEOF(err)
			return err
		}
	}
	return nil
}

type defmsg struct {
	localMsgType      uint8
	arch              binary.ByteOrder
	globalMsgNum      MesgNum
	fields            byte
	fieldDefs         []fieldDef
	devDataFieldDescs []devDataFieldDesc
}

func (dm defmsg) String() string {
	return fmt.Sprintf(
		"local: %d | global: %v | arch: %v | fields: %d | data dev fields: %d",
		dm.localMsgType, dm.globalMsgNum, dm.arch, dm.fields, len(dm.devDataFieldDescs),
	)
}

type fieldDef struct {
	num   byte
	size  byte
	btype types.Base
}

func (fd fieldDef) String() string {
	return fmt.Sprintf("num: %d | size: %d | btype: %v", fd.num, fd.size, fd.btype)
}

type devDataFieldDesc struct {
	fieldNum     byte
	size         byte
	devDataIndex byte
}

func (ddfd devDataFieldDesc) String() string {
	return fmt.Sprintf("field number: %d | size: %d | developer data index: %d", ddfd.fieldNum, ddfd.size, ddfd.devDataIndex)
}

func (d *decoder) parseFileIdMsg() error {
	b, err := d.readByte()
	if err != nil {
		return fmt.Errorf("error parsing record header: %v", err)
	}

	if !((b & mesgDefinitionMask) == mesgDefinitionMask) {
		return fmt.Errorf("expected record header byte for definition message, got %#x - %8b", b, b)
	}

	dm, err := d.parseDefinitionMessage(b)
	if err != nil {
		return fmt.Errorf("error parsing definition message: %v", err)
	}
	if dm.globalMsgNum != MesgNumFileId {
		return fmt.Errorf("parsed definition message was not for file_id (was %v)", dm.globalMsgNum)
	}
	d.defmsgs[dm.localMsgType] = dm

	b, err = d.readByte()
	if err != nil {
		return fmt.Errorf("error parsing record header: %v", err)
	}

	if !((b & mesgHeaderMask) == mesgHeaderMask) {
		return fmt.Errorf("expected record header byte for data message, got %#x - %8b", b, b)
	}
	msg, err := d.parseDataMessage(b, false)
	if err != nil {
		return fmt.Errorf("error reading data message:  %v", err)
	}

	_, ok := msg.Interface().(FileIdMsg)
	if !ok {
		return errors.New("parsed message was not of type file_id")
	}

	if d.debug {
		d.opts.logger.Println("parsed file_id message:", msg)
	}

	d.file.add(msg)

	return nil
}

func (d *decoder) parseDefinitionMessage(recordHeader byte) (*defmsg, error) {
	dm := defmsg{}
	dm.localMsgType = recordHeader & localMesgNumMask
	if dm.localMsgType > localMesgNumMask {
		if d.debug {
			d.opts.logger.Printf("illegal local message number: %d\n", dm.localMsgType)
		}
		return nil, FormatError("illegal local message number")
	}

	// Next byte reserved.
	if err := d.skipByte(); err != nil {
		return nil, err
	}

	arch, err := d.readByte()
	if err != nil {
		return nil, err
	}

	switch arch {
	case littleEndian:
		dm.arch = le
	case bigEndian:
		dm.arch = be
	default:
		return nil, fmt.Errorf("unknown arch: %#x", arch)
	}

	if err = d.readFull(d.tmp[:2]); err != nil {
		return nil, fmt.Errorf("error parsing global message number: %v", err)
	}
	dm.globalMsgNum = MesgNum(dm.arch.Uint16(d.tmp[:2]))
	if dm.globalMsgNum == MesgNumInvalid {
		return nil, FormatError("global message number was set invalid")
	}

	dm.fields, err = d.readByte()
	if err != nil {
		return nil, err
	}
	if dm.fields == 0 {
		if d.debug {
			d.opts.logger.Println("parseDefinitionMessage: warning: 0 fields")
			d.opts.logger.Println("parseDefinitionMessage: message:", dm)
		}
		return &dm, nil
	}

	if err = d.readFull(d.tmp[0 : 3*uint16(dm.fields)]); err != nil {
		return nil, fmt.Errorf("error parsing fields: %v", err)
	}

	dm.fieldDefs = make([]fieldDef, dm.fields)
	for i, fd := range dm.fieldDefs {
		fd.num = d.tmp[i*3]
		fd.size = d.tmp[(i*3)+1]
		fd.btype = types.Base(d.tmp[(i*3)+2])
		if err = d.validateFieldDef(dm.globalMsgNum, fd); err != nil {
			if d.debug {
				d.opts.logger.Println("illegal definition message:", dm)
			}
			return nil, fmt.Errorf("validating %v failed: %v", dm.globalMsgNum, err)
		}
		dm.fieldDefs[i] = fd
	}

	if recordHeader&devDataMask == devDataMask {
		numDevFields, err := d.readByte()
		if err != nil {
			return nil, fmt.Errorf("error reading number of developer data fields: %w", err)
		}

		if err = d.readFull(d.tmp[0 : 3*uint16(numDevFields)]); err != nil {
			return nil, fmt.Errorf("error reading developer data field description data: %w", err)
		}

		dm.devDataFieldDescs = make([]devDataFieldDesc, numDevFields)
		for i, ddfd := range dm.devDataFieldDescs {
			ddfd.fieldNum = d.tmp[i*3]
			ddfd.size = d.tmp[(i*3)+1]
			ddfd.devDataIndex = d.tmp[(i*3)+2]
			dm.devDataFieldDescs[i] = ddfd
		}
	}

	if d.debug {
		d.opts.logger.Println("definition message parsed:", dm)
	}

	return &dm, nil
}

func (d *decoder) validateFieldDef(gmsgnum MesgNum, dfield fieldDef) error {
	if !dfield.btype.Known() {
		return fmt.Errorf("field %d: unknown base type: %v", dfield.num, dfield.btype)
	}

	var pfield *field
	pfound := false
	if knownMsgNums[gmsgnum] {
		pfield, pfound = getField(gmsgnum, dfield.num)
	}

	if dfield.btype == types.BaseString {
		if !pfound {
			return nil
		}
		if pfield.t.BaseType() == dfield.btype {
			return nil
		}
		return fmt.Errorf(
			"field %d: field base type is string, but profile lists it as %v, not compatible",
			dfield.num, pfield.t.BaseType())
	}

	// Verify that field definition size is not less than field definition
	// base type size.
	if int(dfield.size) < dfield.btype.Size() {
		return fmt.Errorf(
			"field %d: size (%d) is less than base type size (%d)",
			dfield.num, dfield.size, dfield.btype.Size())
	}

	if !pfound {
		return nil
	}

	// Profile field.
	if !pfield.t.Array() {
		// Profile field not an array. Verify that the field size is
		// not greater than the profile base type size. A smaller size
		// is allowed due to dynamic fields.
		switch {
		case int(dfield.size) > pfield.t.BaseType().Size():
			return fmt.Errorf(
				"field %d: size %d for %v as base type in definition message is greater than size %d for %v as base type from profile",
				dfield.num, dfield.size, dfield.btype, pfield.t.BaseType().Size(), pfield.t.BaseType())

		case int(dfield.size) <= pfield.t.BaseType().Size() && dfield.btype != pfield.t.BaseType():
			// Size is less or equal, but we can only allow
			// "compatible" types that will not panic when setting
			// fields using reflection.
			switch {
			case pfield.t.BaseType().Signed() != dfield.btype.Signed():
				fallthrough
			case dfield.btype.Float() && !pfield.t.BaseType().Float():
				fallthrough
			case pfield.t.BaseType() == types.BaseString && dfield.btype != types.BaseString:
				return fmt.Errorf(
					"field %d: type %v is not compatible with profile type %v",
					dfield.num, dfield.btype, pfield.t.BaseType())
			}
		}

		return nil
	}

	// Profile field is an array.
	switch {
	case (int(dfield.size) % dfield.btype.Size()) != 0:
		return fmt.Errorf(
			"field %d: array, but size (%d) is not a multiple of base type %v size (%d)",
			dfield.num, dfield.size, dfield.btype, dfield.btype.Size())
	case dfield.btype != pfield.t.BaseType():
		// Require correct base type if an array. I have not seen a
		// dynamic field that is an array and have a smaller base type
		// for array elements. Maybe allow equal sized compatible types
		// later if needed (like for non-array fields).
		return fmt.Errorf(
			"field %d: array, but definition (%v) and profile (%v) base types differ",
			dfield.num, dfield.btype, pfield.t.BaseType())
	default:
		return nil
	}
}

func (d *decoder) parseDataMessage(recordHeader byte, compressed bool) (reflect.Value, error) {
	var localMsgNum byte
	if compressed {
		localMsgNum = (recordHeader & compressedLocalMesgNumMask) >> 5
	} else {
		localMsgNum = recordHeader & localMesgNumMask
	}

	dm := d.defmsgs[localMsgNum]
	if dm == nil {
		return reflect.Value{}, fmt.Errorf(
			"missing data definition message for local message number %d",
			localMsgNum)
	}

	var msgv reflect.Value
	knownMsg := knownMsgNums[dm.globalMsgNum]
	if knownMsg {
		msgv = getMesgAllInvalid(dm.globalMsgNum)
	} else if d.opts.unknownMessages {
		d.unknownMessages[dm.globalMsgNum]++
	}

	if !compressed {
		return d.parseDataFields(dm, knownMsg, msgv)
	}

	// Data message has compressed timestamp header.
	if d.timestamp == 0 {
		if d.debug {
			d.opts.logger.Println(
				"warning: parsing compressed timestamp",
				"header, but have no previous reference",
				"time, skipping setting timestamp for message",
			)
		}
		return d.parseDataFields(dm, knownMsg, msgv)
	}

	timeOffset := int32(recordHeader & compressedTimeMask)
	d.timestamp += uint32((timeOffset - d.lastTimeOffset) & int32(compressedTimeMask))
	d.lastTimeOffset = timeOffset

	fieldTimestamp, found := getField(dm.globalMsgNum, fieldNumTimeStamp)
	if found {
		fieldval := msgv.Field(fieldTimestamp.sindex)
		t := decodeDateTime(d.timestamp)
		fieldval.Set(reflect.ValueOf(t))
		return d.parseDataFields(dm, knownMsg, msgv)
	}

	if d.debug {
		d.opts.logger.Println(
			"warning: parsing message with compressed timestamp header,",
			"but did not find timestamp field in message of type", dm.globalMsgNum)
	}

	return d.parseDataFields(dm, knownMsg, msgv)
}

func (d *decoder) parseDataFields(dm *defmsg, knownMsg bool, msgv reflect.Value) (reflect.Value, error) {
	for i, dfield := range dm.fieldDefs {
		dsize := int(dfield.size)
		padding := 0

		pfield, pfound := getField(dm.globalMsgNum, dfield.num)
		if pfound {
			if pfield.t.BaseType() != types.BaseString && !pfield.t.Array() {
				padding = pfield.t.BaseType().Size() - dsize
			}
		} else if d.opts.unknownFields {
			d.unknownFields[unknownField{dm.globalMsgNum, dfield.num}]++
		}

		err := d.readFull(d.tmp[0:dsize])
		if err != nil {
			return reflect.Value{}, fmt.Errorf(
				"error parsing data message: %v (field %d [%v] for [%v])",
				err, i, dfield, dm)
		}

		if padding != 0 {
			if dm.arch == le {
				for j := dsize; j < pfield.t.BaseType().Size(); j++ {
					d.tmp[j] = 0x00
				}
			} else {
				for j := 0; j < pfield.t.BaseType().Size(); j++ {
					d.tmp[j], d.tmp[j+padding] = 0x00, d.tmp[j]
				}
			}
		}

		if !knownMsg || !pfound {
			continue
		}

		fieldv := msgv.Field(pfield.sindex)

		switch pfield.t.Kind() {
		case types.NativeFit:
			if !pfield.t.Array() {
				err = d.parseFitField(dm, dfield, fieldv)
			} else {
				err = d.parseFitFieldArray(dm, dfield, fieldv)
			}
			if err == nil {
				continue
			}
			return reflect.Value{}, fmt.Errorf("error parsing data message: %v", err)
		case types.TimeUTC, types.TimeLocal:
			d.parseTimeStamp(dm, fieldv, pfield)
		case types.Lat:
			i32 := dm.arch.Uint32(d.tmp[:types.BaseSint32.Size()])
			lat := NewLatitude(int32(i32))
			fieldv.Set(reflect.ValueOf(lat))
		case types.Lng:
			i32 := dm.arch.Uint32(d.tmp[:types.BaseSint32.Size()])
			lng := NewLongitude(int32(i32))
			fieldv.Set(reflect.ValueOf(lng))
		default:
			panic("parseDataFields: unreachable: unknown kind")
		}
	}

	for i, ddfd := range dm.devDataFieldDescs {
		err := d.readFull(d.tmp[0:int(ddfd.size)])
		if err != nil {
			return reflect.Value{}, fmt.Errorf("error parsing data developer message: %v (field %d [%v] for [%v])", err, i, ddfd, dm)
		}
	}

	if knownMsg && !msgv.IsValid() {
		panic("internal decoder error: parse data fields: known message, but not (reflect) valid")
	}

	return msgv, nil
}

func (d *decoder) parseFitField(dm *defmsg, dfield fieldDef, fieldv reflect.Value) error {
	dsize := int(dfield.size)
	switch dfield.btype {
	case types.BaseByte, types.BaseEnum, types.BaseUint8, types.BaseUint8z:
		fieldv.SetUint(uint64(d.tmp[0]))
	case types.BaseSint8:
		fieldv.SetInt(int64(d.tmp[0]))
	case types.BaseSint16:
		i16 := int64(dm.arch.Uint16(d.tmp[:dsize]))
		fieldv.SetInt(i16)
	case types.BaseUint16, types.BaseUint16z:
		u16 := uint64(dm.arch.Uint16(d.tmp[:dsize]))
		fieldv.SetUint(u16)
	case types.BaseSint32:
		i32 := int64(dm.arch.Uint32(d.tmp[:dsize]))
		fieldv.SetInt(i32)
	case types.BaseUint32, types.BaseUint32z:
		u32 := uint64(dm.arch.Uint32(d.tmp[:dsize]))
		fieldv.SetUint(u32)
	case types.BaseFloat32:
		bits := dm.arch.Uint32(d.tmp[:dsize])
		f32 := float64(math.Float32frombits(bits))
		fieldv.SetFloat(f32)
	case types.BaseFloat64:
		bits := dm.arch.Uint64(d.tmp[:dsize])
		f64 := math.Float64frombits(bits)
		fieldv.SetFloat(f64)
	case types.BaseString:
		for j := 0; j < dsize; j++ {
			if d.tmp[j] == 0x00 {
				if j > 0 {
					fieldv.SetString(string(d.tmp[:j]))
				}
				break
			}
			if j == dsize-1 {
				fieldv.SetString(string(d.tmp[:j]))
			}
		}
	default:
		return fmt.Errorf(
			"unknown base type %d for field %v in definition message %v",
			dfield.btype, dfield, dm)
	}

	return nil
}

func (d *decoder) parseFitFieldArray(dm *defmsg, dfield fieldDef, fieldv reflect.Value) error {
	dbt := dfield.btype
	dsize := int(dfield.size)

	if dbt == types.BaseByte {
		byteArray := make([]byte, dsize)
		copy(byteArray, d.tmp[:dsize])
		fieldv.SetBytes(byteArray)
		return nil
	}

	slicev := reflect.MakeSlice(
		fieldv.Type(),
		dsize/dbt.Size(),
		dsize/dbt.Size(),
	)

	switch dbt {
	case types.BaseUint8, types.BaseUint8z, types.BaseEnum:
		for j := 0; j < dsize; j++ {
			slicev.Index(j).SetUint(uint64(d.tmp[j]))
		}
	case types.BaseSint8:
		for j := 0; j < dsize; j++ {
			slicev.Index(j).SetInt(int64(d.tmp[j]))
		}
	case types.BaseSint16:
		for j, k := 0, 0; j < dsize; j, k = j+dbt.Size(), k+1 {
			i16 := int64(dm.arch.Uint16(d.tmp[j : j+dbt.Size()]))
			slicev.Index(k).SetInt(i16)
		}
	case types.BaseUint16, types.BaseUint16z:
		for j, k := 0, 0; j < dsize; j, k = j+dbt.Size(), k+1 {
			ui16 := uint64(dm.arch.Uint16(d.tmp[j : j+dbt.Size()]))
			slicev.Index(k).SetUint(ui16)
		}
	case types.BaseSint32:
		for j, k := 0, 0; j < dsize; j, k = j+dbt.Size(), k+1 {
			i32 := int64(dm.arch.Uint32(d.tmp[j : j+dbt.Size()]))
			slicev.Index(k).SetInt(i32)
		}
	case types.BaseUint32, types.BaseUint32z:
		for j, k := 0, 0; j < dsize; j, k = j+dbt.Size(), k+1 {
			ui32 := uint64(dm.arch.Uint32(d.tmp[j : j+dbt.Size()]))
			slicev.Index(k).SetUint(ui32)
		}
	case types.BaseFloat32:
		for j, k := 0, 0; j < dsize; j, k = j+dbt.Size(), k+1 {
			bits := dm.arch.Uint32(d.tmp[j : j+dbt.Size()])
			f32 := float64(math.Float32frombits(bits))
			slicev.Index(k).SetFloat(f32)
		}
	case types.BaseFloat64:
		for j, k := 0, 0; j < dsize; j, k = j+dbt.Size(), k+1 {
			bits := dm.arch.Uint64(d.tmp[j : j+dbt.Size()])
			f64 := math.Float64frombits(bits)
			slicev.Index(k).SetFloat(f64)
		}
	case types.BaseString:
		if dfield.size == 0 {
			return nil
		}
		var strings []string
		j, k := 0, 0
		for {
			if d.tmp[j+k] == 0x00 {
				if k == 0 {
					break
				}
				strings = append(strings, string(d.tmp[j:j+k]))
				j = j + k + 1
				if j >= dsize {
					break
				}
				k = 0
			} else {
				k++
				if j+k >= dsize {
					// We have not seen a 0x00 terminator,
					// but there's no room for one.
					// Take the string we have and exit loop.
					strings = append(strings, string(d.tmp[j:dsize]))
					break
				}
			}
		}
		fieldv.Set(reflect.ValueOf(strings))
		return nil // We don't want the Set after the switch.
	default:
		return fmt.Errorf(
			"unknown base type %d for field %v in definition message %v",
			dbt, dfield, dm)
	}

	fieldv.Set(slicev)
	return nil
}

func (d *decoder) parseTimeStamp(dm *defmsg, fieldv reflect.Value, pfield *field) {
	u32 := dm.arch.Uint32(d.tmp[:types.BaseUint32.Size()])
	if u32 == 0xFFFFFFFF {
		return
	}
	if u32 < systemTimeMarker && d.debug {
		d.opts.logger.Println("parsing time: seconds from device power on")
	}

	if pfield.t.Kind() == types.TimeUTC {
		if pfield.num == fieldNumTimeStamp {
			d.timestamp = u32
			d.lastTimeOffset = int32(d.timestamp & uint32(compressedTimeMask))
		}
		t := decodeDateTime(u32)
		fieldv.Set(reflect.ValueOf(t))
		return
	}

	// Local timestamp.
	//
	// Use a custom timezone with the calculated offset to indicate that it
	// is not UTC.
	//
	// Also see the SetLocalTimeZone function in the timeutil subpackage.
	// For now not used due to an external dependency.
	var local time.Time
	switch {
	case d.timestamp == 0, d.timestamp < systemTimeMarker:
		// No time reference.
		// Set local with zero offset.
		d.timestamp = u32
		tzone := time.FixedZone(localZoneName, 0)
		local = decodeDateTime(u32)
		local = local.In(tzone)
	default:
		local = decodeDateTime(u32)
		utc := decodeDateTime(d.timestamp)
		offsetDur := local.Sub(utc)
		tzone := time.FixedZone(localZoneName, int(offsetDur.Seconds()))
		local = utc.In(tzone)
	}
	fieldv.Set(reflect.ValueOf(local))
}

func noEOF(err error) error {
	if errors.Is(err, io.EOF) {
		return io.ErrUnexpectedEOF
	}
	return err
}

func (d *decoder) handleUnknownFields() {
	d.file.UnknownFields = make([]UnknownField, 0, len(d.unknownFields))
	for field, count := range d.unknownFields {
		d.file.UnknownFields = append(d.file.UnknownFields, UnknownField{
			MesgNum:  field.mesgNum,
			FieldNum: field.fieldNum,
			Count:    count,
		})
	}
	sort.Sort(unknownFieldSlice(d.file.UnknownFields))
}

func (d *decoder) handleUnknownMessages() {
	d.file.UnknownMessages = make([]UnknownMessage, 0, len(d.unknownMessages))
	for mesgNum, count := range d.unknownMessages {
		d.file.UnknownMessages = append(d.file.UnknownMessages, UnknownMessage{
			MesgNum: mesgNum,
			Count:   count,
		})
	}
	sort.Sort(unknownMessageSlice(d.file.UnknownMessages))
}
