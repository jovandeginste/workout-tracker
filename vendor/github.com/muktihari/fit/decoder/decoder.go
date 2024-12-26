// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"sync"

	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/kit/hash"
	"github.com/muktihari/fit/kit/hash/crc16"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

type errorString string

func (e errorString) Error() string { return string(e) }

const (
	// ErrNotFITFile will be returned if the first byte of every FIT sequence does not match
	// with FIT FileHeader's Size specification (either 12 or 14), byte 8-12 are not ".FIT",
	// or byte 4-8 are all zero (FileHeader's DataSize == 0).
	ErrNotFITFile = errorString("not a FIT file")

	// ErrCRCChecksumMismatch will be returned if the CRC checksum does not match with
	// the CRC in the file, whether during FileHeader or messages integrity checks.
	ErrCRCChecksumMismatch = errorString("crc checksum mismatch")

	// ErrMesgDefMissing will be returned if message definition for the incoming message data is missing.
	ErrMesgDefMissing = errorString("message definition missing") // NOTE: Kept exported since it's used by RawDecoder

	errInvalidBaseType = errorString("invalid basetype")
)

// Decoder is FIT file decoder. See New() for details.
type Decoder struct {
	readBuffer  *readBuffer // read from io.Reader with buffer without extra copying.
	n           int64       // n is a read bytes counter, always moving forward, do not reset (except on full reset).
	factory     Factory
	accumulator *Accumulator
	crc16       hash.Hash16
	err         error // Any error occurs during process.

	fieldsArray          [255]proto.Field
	developerFieldsArray [255]proto.DeveloperField

	options options

	once           sync.Once // It is used to invoke decodeFileHeader exactly once. Must be reassigned on init/reset.
	cur            uint32    // The current byte position relative to bytes of the messages, reset on next chained FIT file.
	timestamp      uint32    // Active timestamp
	lastTimeOffset byte      // Last time offset

	// FIT File Representation
	fileHeader proto.FileHeader
	messages   []proto.Message
	crc        uint16

	// FileId Message is a special message that must present in a FIT file.
	fileId *mesgdef.FileId

	// Message Definition Lookup
	localMessageDefinitions      [proto.LocalMesgNumMask + 1]*proto.MessageDefinition // message definition for upcoming message data
	localMessageDefinitionsArray [proto.LocalMesgNumMask + 1]proto.MessageDefinition  // PERF: backing array for message definition

	// Developer Data Lookup
	developerDataIndexes []uint8
	fieldDescriptions    []*mesgdef.FieldDescription
}

// Factory defines a contract that any Factory containing these method can be used by the Decoder.
type Factory interface {
	// CreateField creates new field based on defined messages in the factory.
	// If not found, it returns new field with "unknown" name.
	CreateField(mesgNum typedef.MesgNum, num byte) proto.Field
}

type options struct {
	factory               Factory
	logWriter             io.Writer
	mesgListeners         []MesgListener    // Each listener will received every decoded message.
	mesgDefListeners      []MesgDefListener // Each listener will received every decoded message definition.
	readBufferSize        int
	shouldChecksum        bool
	broadcastOnly         bool
	shouldExpandComponent bool
	broadcastMesgCopy     bool
}

func defaultOptions() options {
	return options{
		factory:               factory.StandardFactory(),
		logWriter:             nil,
		readBufferSize:        defaultReadBufferSize,
		shouldChecksum:        true,
		broadcastOnly:         false,
		shouldExpandComponent: true,
		broadcastMesgCopy:     false,
	}
}

// Option is Decoder's option.
type Option func(o *options)

// WithFactory sets custom factory.
func WithFactory(factory Factory) Option {
	return func(o *options) {
		if factory != nil {
			o.factory = factory
		}
	}
}

// WithMesgListener adds listeners to the listener pool, where each listener will receive
// every message as soon as it is decoded. The listeners will be appended not replaced.
// If users need to reset use Reset().
func WithMesgListener(listeners ...MesgListener) Option {
	return func(o *options) {
		o.mesgListeners = append(o.mesgListeners, listeners...)
	}
}

// WithMesgDefListener adds listeners to the listener pool, where each listener will receive
// every message definition as soon as it is decoded. The listeners will be appended not replaced.
// If users need to reset use Reset().
func WithMesgDefListener(listeners ...MesgDefListener) Option {
	return func(o *options) {
		o.mesgDefListeners = append(o.mesgDefListeners, listeners...)
	}
}

// WithBroadcastOnly directs the Decoder to only broadcast the messages without retaining them, reducing memory usage when
// it's not going to be used anyway. This option is intended to be used with WithMesgListener and
// When this option is specified, the Decode will return a FIT with empty messages.
func WithBroadcastOnly() Option {
	return func(o *options) { o.broadcastOnly = true }
}

// WithBroadcastMesgCopy directs the Decoder to copy the mesg before passing it to listeners
// (it was the default behavior on version <= v0.14.0).
func WithBroadcastMesgCopy() Option {
	return func(o *options) { o.broadcastMesgCopy = true }
}

// WithIgnoreChecksum directs the Decoder to not checking data integrity (CRC Checksum).
func WithIgnoreChecksum() Option {
	return func(o *options) { o.shouldChecksum = false }
}

// WithNoComponentExpansion directs the Decoder to not expand the components.
func WithNoComponentExpansion() Option {
	return func(o *options) { o.shouldExpandComponent = false }
}

// WithLogWriter specifies where the log messages will be written to. By default, the Decoder do not write any log if
// log writer is not specified. The Decoder will only write log messages when it encountered a bad encoded FIT file such as:
//   - Field Definition's Size (or Developer Field Definition's Size) is zero.
//   - Field Definition's Size (or Developer Field Definition's Size) is less than basetype's Size.
//     e.g. Size 1 bytes but having basetype uint32 (4 bytes).
//   - Field Definition's Size is more than basetype's Size but field.Array is false.
//   - Encountering a Developer Field without prior DeveloperDataId or FieldDescription Message.
func WithLogWriter(w io.Writer) Option {
	return func(o *options) { o.logWriter = w }
}

// WithReadBufferSize directs the Decoder to use this buffer size for reading from io.Reader instead of default 4096.
func WithReadBufferSize(size int) Option {
	return func(o *options) { o.readBufferSize = size }
}

// New returns a FIT File Decoder to decode given r.
//
// The FIT protocol allows for multiple FIT files to be chained together in a single FIT file.
// Each FIT file in the chain must be a properly formatted FIT file (header, data records, CRC).
//
// To decode a chained FIT file containing multiple FIT data, invoke Decode() or DecodeWithContext()
// method multiple times. For convenience, we can wrap it with the Next() method as follows (optional):
//
//	for dec.Next() {
//	   fit, err := dec.Decode()
//	}
//
// Note: Decoder already implements efficient io.Reader buffering, so there's no need to wrap 'r' using *bufio.Reader;
// doing so will only reduce performance.
func New(r io.Reader, opts ...Option) *Decoder {
	d := &Decoder{
		readBuffer:  new(readBuffer),
		accumulator: NewAccumulator(),
		crc16:       crc16.New(nil),
	}
	d.Reset(r, opts...)
	return d
}

// Reset resets the Decoder to read its input from r, clear any error and
// reset previous options to default options so any options needs to be inputed again.
// It is similar to New() but it retains the underlying storage for use by
// future decode to reduce memory allocs.
func (d *Decoder) Reset(r io.Reader, opts ...Option) {
	d.reset()
	d.n = 0 // Must reset bytes counter since it's a full reset.

	// Reuse listeners' slices
	for i := range d.options.mesgListeners {
		d.options.mesgListeners[i] = nil // avoid memory leaks
	}
	mesgListeners := d.options.mesgListeners[:0]
	for i := range d.options.mesgDefListeners {
		d.options.mesgDefListeners[i] = nil // avoid memory leaks
	}
	mesgDefListeners := d.options.mesgDefListeners[:0]

	d.options = defaultOptions()
	d.options.mesgListeners = mesgListeners
	d.options.mesgDefListeners = mesgDefListeners
	for i := range opts {
		opts[i](&d.options)
	}

	d.readBuffer.Reset(r, d.options.readBufferSize)
	d.factory = d.options.factory
}

func (d *Decoder) reset() {
	d.accumulator.Reset()
	d.crc16.Reset()
	d.once = sync.Once{}
	d.cur = 0
	d.timestamp = 0
	d.lastTimeOffset = 0
	d.err = nil
	d.fileHeader = proto.FileHeader{}
	d.messages = nil
	d.crc = 0
	d.fileId = nil
}

// releaseTemporaryObjects releases objects that being created during a single decode process
// by stops referencing those objects so it can be garbage-collected on next GC cycle.
func (d *Decoder) releaseTemporaryObjects() {
	d.localMessageDefinitions = [proto.LocalMesgNumMask + 1]*proto.MessageDefinition{}
	d.fieldsArray = [255]proto.Field{}
	d.developerFieldsArray = [255]proto.DeveloperField{}
	d.fileId = nil
	d.messages = nil
	d.developerDataIndexes = d.developerDataIndexes[:0]
	for i := range d.fieldDescriptions {
		d.fieldDescriptions[i] = nil
	}
	d.fieldDescriptions = d.fieldDescriptions[:0]
}

// CheckIntegrity checks all FIT sequences of given reader are valid determined by these following checks:
//  1. Has valid FileHeader's size and bytes 8–11 of the FileHeader is “.FIT”
//  2. FileHeader's DataSize > 0
//  3. CRC checksum of messages should match with File's CRC value.
//
// It returns the number of sequences completed and any error encountered. The number of sequences completed can help recovering
// valid FIT sequences in a chained FIT that contains invalid or corrupted data.
//
// After invoking this method, the underlying reader should be reset afterward as the reader has been fully read.
// If the underlying reader implements io.Seeker, we can do reader.Seek(0, io.SeekStart).
func (d *Decoder) CheckIntegrity() (seq int, err error) {
	if d.err != nil {
		return 0, d.err
	}

	shouldChecksum := d.options.shouldChecksum
	d.options.shouldChecksum = true // Must checksum

	for {
		// Check File Header Integrity
		pos := d.n
		if err = d.decodeFileHeaderOnce(); err != nil {
			if pos != 0 && pos == d.n && err == io.EOF {
				// When EOF error occurs exactly after a sequence has been completed,
				// make the error as nil, it means we have reached the desirable EOF.
				err = nil
			}
			break
		}
		// Read bytes acquired by messages to calculate crc checksum of its contents
		if err = d.discardMessages(); err != nil {
			break
		}
		if err = d.decodeCRC(); err != nil {
			break
		}
		d.once = sync.Once{}
		d.cur = 0
		seq++
	}

	if err != nil { // When there is an error, wrap it with informative message before return.
		err = fmt.Errorf("byte pos: %d: %w", d.n, err)
	}

	// Reset used variables so that the decoder can be reused by the same reader.
	d.reset()
	d.n = 0 // Must reset bytes counter
	d.options.shouldChecksum = shouldChecksum

	return seq, err
}

// discardMessages efficiently discards bytes used by messages.
func (d *Decoder) discardMessages() (err error) {
	for d.cur < d.fileHeader.DataSize {
		size := int(d.fileHeader.DataSize - d.cur)
		if size > reservedbuf {
			size = reservedbuf
		}
		if _, err = d.readN(size); err != nil { // Discard bytes
			return err
		}
	}
	return nil
}

// PeekFileHeader decodes only up to FileHeader (first 12-14 bytes) without decoding the whole reader.
//
// If we choose to continue, Decode picks up where this left then continue decoding next messages instead of starting from zero.
func (d *Decoder) PeekFileHeader() (*proto.FileHeader, error) {
	if d.err != nil {
		return nil, d.err
	}
	if d.err = d.decodeFileHeaderOnce(); d.err != nil {
		return nil, d.err
	}
	return &d.fileHeader, nil
}

// PeekFileId decodes only up to FileId message without decoding the whole reader.
// The FileId is expected to be present as the first message; however, we don't validate this,
// as it's an edge case that occurs when a FIT file is poorly encoded.
//
// If we choose to continue, Decode picks up where this left then continue decoding next messages instead of starting from zero.
func (d *Decoder) PeekFileId() (*mesgdef.FileId, error) {
	if d.err != nil {
		return nil, d.err
	}
	if d.err = d.decodeFileHeaderOnce(); d.err != nil {
		return nil, d.err
	}
	for d.fileId == nil {
		if d.err = d.decodeMessage(); d.err != nil {
			return nil, d.err
		}
	}
	return d.fileId, nil
}

// Next checks whether next bytes are still a valid FIT File sequence. Return false when invalid or reach EOF.
func (d *Decoder) Next() bool {
	if d.err != nil {
		return false
	}
	if d.n == 0 {
		return true
	}
	return d.decodeFileHeaderOnce() == nil
}

// Decode method decodes `r` into FIT data. One invocation will produce one valid FIT data or
// an error if it occurs. To decode a chained FIT file containing multiple FIT data, invoke this
// method multiple times. For convenience, we can wrap it with the Next() method as follows (optional):
//
//	for dec.Next() {
//	     fit, err := dec.Decode()
//	     if err != nil {
//	         return err
//	     }
//	}
func (d *Decoder) Decode() (*proto.FIT, error) {
	if d.err != nil {
		return nil, d.err
	}
	if d.err = d.decodeFileHeaderOnce(); d.err != nil {
		return nil, d.err
	}
	defer d.releaseTemporaryObjects()
	if d.err = d.decodeMessages(); d.err != nil {
		return nil, d.err
	}
	if d.err = d.decodeCRC(); d.err != nil {
		return nil, d.err
	}
	fit := &proto.FIT{
		FileHeader: d.fileHeader,
		Messages:   d.messages,
		CRC:        d.crc,
	}
	d.reset()
	return fit, nil
}

// Discard discards a single FIT file sequence and returns any error encountered. This method directs the Decoder to
// point to the byte sequence of the next valid FIT file sequence, discarding the current FIT file sequence.
//
// Example: - A chained FIT file consist of Activity, Course, Workout and Settings. And we only want to decode Course.
//
//	for dec.Next() {
//		fileId, err := dec.PeekFileId()
//		if err != nil {
//			return err
//		}
//		if fileId.Type != typedef.FileCourse {
//		    if err := dec.Discard(); err != nil {
//		    	return err
//		    }
//		    continue
//		}
//		fit, err := dec.Decode()
//		if err != nil {
//			return err
//		}
//	 }
func (d *Decoder) Discard() error {
	if d.err != nil {
		return d.err
	}

	optionsShouldChecksum := d.options.shouldChecksum
	d.options.shouldChecksum = false
	defer func() { d.options.shouldChecksum = optionsShouldChecksum }()

	if d.err = d.decodeFileHeaderOnce(); d.err != nil {
		return d.err
	}
	if d.err = d.discardMessages(); d.err != nil {
		return d.err
	}
	if _, d.err = d.readN(2); d.err != nil { // Discard File CRC
		return d.err
	}
	d.reset()
	return d.err
}

// decodeFileHeaderOnce invokes decodeFileHeader exactly once.
func (d *Decoder) decodeFileHeaderOnce() error {
	d.once.Do(func() { d.err = d.decodeFileHeader() })
	return d.err
}

// decodeFileHeader is only invoked through decodeFileHeaderOnce.
func (d *Decoder) decodeFileHeader() error {
	b, err := d.readBuffer.ReadN(1)
	if err != nil {
		return err
	}
	d.n += 1
	size := b[0]

	if size != 12 && size != 14 { // current spec is either 12 or 14
		return fmt.Errorf("file header size [%d] is invalid: %w", size, ErrNotFITFile)
	}
	_, _ = d.crc16.Write(b)

	rem := int(size - 1)
	b, err = d.readBuffer.ReadN(rem)
	if err != nil {
		return err
	}
	d.n += int64(rem)

	// PERF: Neither string(b[7:11]) nor assigning proto.DataTypeFIT constant to a variable escape to the heap.
	if string(b[7:11]) != proto.DataTypeFIT {
		return ErrNotFITFile
	}

	d.fileHeader = proto.FileHeader{
		Size:            size,
		ProtocolVersion: proto.Version(b[0]),
		ProfileVersion:  binary.LittleEndian.Uint16(b[1:3]),
		DataSize:        binary.LittleEndian.Uint32(b[3:7]),
		DataType:        proto.DataTypeFIT,
	}

	if err := proto.Validate(d.fileHeader.ProtocolVersion); err != nil {
		return err
	}

	if d.fileHeader.DataSize == 0 {
		return fmt.Errorf("invalid data size: %w", ErrNotFITFile)
	}

	if size == 14 {
		d.fileHeader.CRC = binary.LittleEndian.Uint16(b[11:13])
	}

	if d.fileHeader.CRC == 0x0000 || !d.options.shouldChecksum { // do not need to check file header's crc integrity.
		d.crc16.Reset()
		return nil
	}

	_, _ = d.crc16.Write(b[:len(b)-2])
	if d.crc16.Sum16() != d.fileHeader.CRC { // check file header integrity
		return fmt.Errorf("expected file header's crc: %d, got: %d: %w", d.fileHeader.CRC, d.crc16.Sum16(), ErrCRCChecksumMismatch)
	}

	d.crc16.Reset() // this hash will be re-used for calculating data integrity.

	return nil
}

func (d *Decoder) decodeMessages() (err error) {
	for d.cur < d.fileHeader.DataSize {
		if err = d.decodeMessage(); err != nil {
			return fmt.Errorf("decodeMessage [byte pos: %d]: %w", d.n, err)
		}
	}
	return nil
}

func (d *Decoder) decodeMessage() error {
	b, err := d.readN(1)
	if err != nil {
		return err
	}
	header := b[0]

	// NOTE: Compressed Timestamp Header Bit 5-6 is the local message type.
	// Bit 6 overlap with MesgDefinitionMask; It's a message definition only if Bit 7 is zero.
	if (header & (proto.MesgCompressedHeaderMask | proto.MesgDefinitionMask)) == proto.MesgDefinitionMask {
		return d.decodeMessageDefinition(header)
	}
	return d.decodeMessageData(header)
}

func (d *Decoder) decodeMessageDefinition(header byte) error {
	b, err := d.readN(5)
	if err != nil {
		return err
	}

	localMesgNum := header & proto.LocalMesgNumMask
	mesgDef := d.localMessageDefinitions[localMesgNum]
	if mesgDef == nil {
		// PERF: Use backing array to avoid object creation. On init, allocate slices
		// with max cap for more deterministic performance by avoiding re-allocation.
		mesgDef = &d.localMessageDefinitionsArray[localMesgNum]
		if mesgDef.FieldDefinitions == nil && mesgDef.DeveloperFieldDefinitions == nil {
			mesgDef.FieldDefinitions = make([]proto.FieldDefinition, 0, 255)
			mesgDef.DeveloperFieldDefinitions = make([]proto.DeveloperFieldDefinition, 0, 255)
		}
	}

	mesgDef.Header = header
	mesgDef.Reserved = b[0]
	mesgDef.Architecture = b[1]
	if mesgDef.Architecture == proto.LittleEndian {
		mesgDef.MesgNum = typedef.MesgNum(binary.LittleEndian.Uint16(b[2:4]))
	} else {
		mesgDef.MesgNum = typedef.MesgNum(binary.BigEndian.Uint16(b[2:4]))
	}

	n := int(b[4])
	b, err = d.readN(n * 3) // 3 byte per field
	if err != nil {
		return err
	}

	mesgDef.FieldDefinitions = mesgDef.FieldDefinitions[:0]
	var baseType basetype.BaseType
	for ; len(b) >= 3; b = b[3:] {
		baseType = basetype.BaseType(b[2])
		if !baseType.Valid() {
			return fmt.Errorf("message definition number: %s(%d): fields[%d].BaseType: %s: %w",
				mesgDef.MesgNum, mesgDef.MesgNum, len(mesgDef.FieldDefinitions), baseType, errInvalidBaseType)
		}
		mesgDef.FieldDefinitions = append(mesgDef.FieldDefinitions,
			proto.FieldDefinition{
				Num:      b[0],
				Size:     b[1],
				BaseType: baseType,
			})
	}

	mesgDef.DeveloperFieldDefinitions = mesgDef.DeveloperFieldDefinitions[:0]
	if (header & proto.DevDataMask) == proto.DevDataMask {
		b, err = d.readN(1)
		if err != nil {
			return err
		}

		n = int(b[0])
		b, err = d.readN(n * 3) // 3 byte per field
		if err != nil {
			return err
		}

		for ; len(b) >= 3; b = b[3:] {
			mesgDef.DeveloperFieldDefinitions = append(mesgDef.DeveloperFieldDefinitions,
				proto.DeveloperFieldDefinition{
					Num:                b[0],
					Size:               b[1],
					DeveloperDataIndex: b[2],
				})
		}
	}

	d.localMessageDefinitions[localMesgNum] = mesgDef

	if len(d.options.mesgDefListeners) > 0 {
		// Clone since we don't have control of the object lifecycle outside Decoder.
		mesgDef := *mesgDef
		mesgDef.FieldDefinitions = sliceutil.Clone(mesgDef.FieldDefinitions)
		mesgDef.DeveloperFieldDefinitions = sliceutil.Clone(mesgDef.DeveloperFieldDefinitions)
		for i := range d.options.mesgDefListeners {
			d.options.mesgDefListeners[i].OnMesgDef(mesgDef) // blocking or non-blocking depends on listeners' implementation.
		}
	}

	return nil
}

func (d *Decoder) decodeMessageData(header byte) (err error) {
	localMesgNum := header
	if (header & proto.MesgCompressedHeaderMask) == proto.MesgCompressedHeaderMask {
		localMesgNum = (header & proto.CompressedLocalMesgNumMask) >> proto.CompressedBitShift
	}
	mesgDef := d.localMessageDefinitions[localMesgNum&proto.LocalMesgNumMask] // bounds check eliminated due to the mask
	if mesgDef == nil {
		return ErrMesgDefMissing
	}

	mesg := proto.Message{Num: mesgDef.MesgNum}
	mesg.Header = header
	mesg.Fields = d.fieldsArray[:0]

	if (header & proto.MesgCompressedHeaderMask) == proto.MesgCompressedHeaderMask { // Compressed Timestamp Message Data
		timeOffset := header & proto.CompressedTimeMask
		d.timestamp += uint32((timeOffset - d.lastTimeOffset) & proto.CompressedTimeMask)
		d.lastTimeOffset = timeOffset

		timestampField := d.factory.CreateField(mesgDef.MesgNum, proto.FieldNumTimestamp)
		if timestampField.Name == factory.NameUnknown {
			timestampField.BaseType = basetype.Uint32
			timestampField.Type = profile.DateTime
		}
		timestampField.Value = proto.Uint32(d.timestamp)

		mesg.Fields = append(mesg.Fields, timestampField) // add timestamp field
	}

	if err = d.decodeFields(mesgDef, &mesg); err != nil {
		return err
	}

	// FileId Message
	if d.fileId == nil && mesg.Num == mesgnum.FileId {
		d.fileId = mesgdef.NewFileId(&mesg)
	}

	// Prerequisites for decoding developer fields
	switch mesg.Num {
	case mesgnum.DeveloperDataId:
		// These messages must occur before any related field description messages are written to the proto.
		d.developerDataIndexes = append(d.developerDataIndexes,
			mesg.FieldValueByNum(fieldnum.DeveloperDataIdDeveloperDataIndex).Uint8())
	case mesgnum.FieldDescription:
		// These messages must occur in the file before any related developer data is written to the proto.
		d.fieldDescriptions = append(d.fieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	}

	if len(mesgDef.DeveloperFieldDefinitions) != 0 {
		mesg.DeveloperFields = d.developerFieldsArray[:0]
		if err = d.decodeDeveloperFields(mesgDef, &mesg); err != nil {
			return err
		}
	}

	if !d.options.broadcastOnly || d.options.broadcastMesgCopy {
		mesg.Fields = sliceutil.Clone(mesg.Fields)
		mesg.DeveloperFields = sliceutil.Clone(mesg.DeveloperFields)
	}

	if !d.options.broadcastOnly {
		d.messages = append(d.messages, mesg)
	}

	for i := range d.options.mesgListeners {
		d.options.mesgListeners[i].OnMesg(mesg) // blocking or non-blocking depends on listeners' implementation.
	}

	return nil
}

func (d *Decoder) decodeFields(mesgDef *proto.MessageDefinition, mesg *proto.Message) (err error) {
	for i := range mesgDef.FieldDefinitions {
		fieldDef := &mesgDef.FieldDefinitions[i]

		// We enforce field.Array for string type to match the value defined in factory for all non-unknown fields.
		var overrideStringArray bool
		field := d.factory.CreateField(mesgDef.MesgNum, fieldDef.Num)
		if field.Name == factory.NameUnknown {
			// Assign fieldDef's type for unknown field so later we can encode it as per its original value.
			field.BaseType = fieldDef.BaseType
			field.Type = profile.ProfileTypeFromBaseType(field.BaseType)
			// Check if the size corresponds to an array.
			field.Array = fieldDef.Size > field.BaseType.Size() && fieldDef.Size%field.BaseType.Size() == 0
			// Fallback to FIT Protocol's string rule: decoder will determine it by counting the utf8 null-terminated string.
			overrideStringArray = field.BaseType == basetype.String
		}

		var (
			baseType    = field.BaseType
			profileType = field.Type
			array       = field.Array
		)

		// Gracefully handle poorly encoded FIT file.
		if fieldDef.Size == 0 {
			d.logField(mesg, fieldDef, "Size is zero. Skip")
			continue
		} else if fieldDef.Size < baseType.Size() {
			baseType = basetype.Byte
			profileType = profile.Byte
			array = fieldDef.Size > baseType.Size() && fieldDef.Size&baseType.Size() == 0
			d.logField(mesg, fieldDef, "Size is less than expected. Fallback: decode as byte(s) and convert the value")
		} else if fieldDef.Size > baseType.Size() && !field.Array && baseType != basetype.String {
			d.logField(mesg, fieldDef, "field.Array is false. Fallback: retrieve first array's value only")
		}

		field.Value, err = d.readValue(fieldDef.Size, mesgDef.Architecture, baseType, profileType, array, overrideStringArray)
		if err != nil {
			return err
		}

		if baseType != field.BaseType { // Convert value
			field.Value = convertBytesToValue(field.Value, field.BaseType)
		}

		if field.Num == proto.FieldNumTimestamp && field.Value.Type() == proto.TypeUint32 {
			timestamp := field.Value.Uint32()
			d.timestamp = timestamp
			d.lastTimeOffset = byte(timestamp & proto.CompressedTimeMask)
		}

		if field.Accumulate && d.options.shouldExpandComponent {
			d.collectAccumulableValues(mesg.Num, field.Num, field.Value)
		}

		mesg.Fields = append(mesg.Fields, field)
	}

	if !d.options.shouldExpandComponent {
		return nil
	}

	// Now that all fields has been decoded, we need to expand all components and accumulate the accumulable values.
	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if subField := field.SubFieldSubtitution(mesg); subField != nil {
			// Expand sub-field components as the main field components
			d.expandComponents(mesg, field.Value, field.BaseType, subField.Components)
			continue
		}
		// No sub-field can interpret as main field, expand main field components
		d.expandComponents(mesg, field.Value, field.BaseType, field.Components)
	}

	return nil
}

// collectAccumulableValues collects the field values to be used in component expansion.
func (d *Decoder) collectAccumulableValues(mesgNum typedef.MesgNum, fieldNum byte, val proto.Value) {
	switch val.Type() {
	case proto.TypeInt8:
		d.accumulator.Collect(mesgNum, fieldNum, uint32(val.Int8()))
	case proto.TypeUint8:
		d.accumulator.Collect(mesgNum, fieldNum, uint32(val.Uint8()))
	case proto.TypeInt16:
		d.accumulator.Collect(mesgNum, fieldNum, uint32(val.Int16()))
	case proto.TypeUint16:
		d.accumulator.Collect(mesgNum, fieldNum, uint32(val.Uint16()))
	case proto.TypeInt32:
		d.accumulator.Collect(mesgNum, fieldNum, uint32(val.Int32()))
	case proto.TypeUint32:
		d.accumulator.Collect(mesgNum, fieldNum, uint32(val.Uint32()))
	case proto.TypeInt64:
		d.accumulator.Collect(mesgNum, fieldNum, uint32(val.Int64()))
	case proto.TypeUint64:
		d.accumulator.Collect(mesgNum, fieldNum, uint32(val.Uint64()))
	case proto.TypeFloat32:
		d.accumulator.Collect(mesgNum, fieldNum, uint32(val.Float32()))
	case proto.TypeFloat64:
		d.accumulator.Collect(mesgNum, fieldNum, uint32(val.Float64()))
	case proto.TypeSliceInt8:
		vals := val.SliceInt8()
		for i := range vals {
			d.accumulator.Collect(mesgNum, fieldNum, uint32(vals[i]))
		}
	case proto.TypeSliceUint8:
		vals := val.SliceUint8()
		for i := range vals {
			d.accumulator.Collect(mesgNum, fieldNum, uint32(vals[i]))
		}
	case proto.TypeSliceInt16:
		vals := val.SliceInt16()
		for i := range vals {
			d.accumulator.Collect(mesgNum, fieldNum, uint32(vals[i]))
		}
	case proto.TypeSliceUint16:
		vals := val.SliceUint16()
		for i := range vals {
			d.accumulator.Collect(mesgNum, fieldNum, uint32(vals[i]))
		}
	case proto.TypeSliceInt32:
		vals := val.SliceInt32()
		for i := range vals {
			d.accumulator.Collect(mesgNum, fieldNum, uint32(vals[i]))
		}
	case proto.TypeSliceUint32:
		vals := val.SliceUint32()
		for i := range vals {
			d.accumulator.Collect(mesgNum, fieldNum, uint32(vals[i]))
		}
	case proto.TypeSliceInt64:
		vals := val.SliceInt64()
		for i := range vals {
			d.accumulator.Collect(mesgNum, fieldNum, uint32(vals[i]))
		}
	case proto.TypeSliceUint64:
		vals := val.SliceUint64()
		for i := range vals {
			d.accumulator.Collect(mesgNum, fieldNum, uint32(vals[i]))
		}
	case proto.TypeSliceFloat32:
		vals := val.SliceFloat32()
		for i := range vals {
			d.accumulator.Collect(mesgNum, fieldNum, uint32(vals[i]))
		}
	case proto.TypeSliceFloat64:
		vals := val.SliceFloat64()
		for i := range vals {
			d.accumulator.Collect(mesgNum, fieldNum, uint32(vals[i]))
		}
	}
}

func (d *Decoder) expandComponents(mesg *proto.Message, containingValue proto.Value, baseType basetype.BaseType, components []proto.Component) {
	if len(components) == 0 {
		return
	}

	if !containingValue.Valid(baseType) {
		return
	}

	vbits, ok := makeBits(containingValue)
	if !ok {
		return
	}

	var componentField proto.Field
	for i := range components {
		component := &components[i]

		componentField = d.factory.CreateField(mesg.Num, component.FieldNum)
		componentField.IsExpandedField = true

		// A component can only have max 32 bits value.
		// If a field has only one component, expand it even if its value is zero
		// e.g. speed (0) -> enhanced_speed (0).
		val := vbits.Pull(component.Bits)
		if val == 0 && len(components) > 1 {
			break
		}

		if component.Accumulate {
			val = d.accumulator.Accumulate(mesg.Num, component.FieldNum, val, component.Bits)
		}

		componentScaled := scaleoffset.Apply(val, component.Scale, component.Offset)
		val = uint32(scaleoffset.Discard(componentScaled, componentField.Scale, componentField.Offset))
		value := convertUint32ToValue(val, componentField.BaseType)

		// All components fields are appended, so it makes more sense to search from the last order.
		// Our goal is to create new or update existing expanded field. However, there is an edge case
		// where the decoded field with the same field number as the expanded field candidate already exist
		// before the expansion begin. For such cases, that field will be updated instead of creating new
		// expanded field. If its value is an array, value will be appended, otherwise, value will be replaced.
		var fieldRef *proto.Field
		for j := len(mesg.Fields) - 1; j >= 0; j-- {
			field := &mesg.Fields[j]
			if field.Num == component.FieldNum {
				fieldRef = field
				break
			}
		}

		var shouldAppend bool
		if fieldRef == nil {
			fieldRef = &componentField
			shouldAppend = true
		}

		// Some of expanded field's values are in the form of slice, e.g.:
		// - Hr: event_timestamp FROM event_timestamp_12 (approx. up to 120 bits)
		// - RawBbi: time, quality, gap FROM data (approx. up to 240 bits)
		if fieldRef.Array {
			fieldRef.Value = valueAppend(fieldRef.Value, value)
		} else {
			fieldRef.Value = value
		}

		if shouldAppend {
			mesg.Fields = append(mesg.Fields, componentField)
		}

		// The destination field (componentField) can itself contain components requiring expansion.
		// e.g. compressed_speed_distance -> (speed, distance), speed -> enhanced_speed.
		//
		// NOTE: We pass the 32 bits component's value to ensure we only expand this value.
		if subField := componentField.SubFieldSubtitution(mesg); subField != nil {
			d.expandComponents(mesg, value, componentField.BaseType, subField.Components)
		} else {
			d.expandComponents(mesg, value, componentField.BaseType, componentField.Components)
		}
	}
}

func (d *Decoder) decodeDeveloperFields(mesgDef *proto.MessageDefinition, mesg *proto.Message) error {
	for i := range mesgDef.DeveloperFieldDefinitions {
		devFieldDef := &mesgDef.DeveloperFieldDefinitions[i]

		var ok bool
		for _, developerDataIndex := range d.developerDataIndexes {
			if developerDataIndex == devFieldDef.DeveloperDataIndex {
				ok = true
				break
			}
		}

		if !ok {
			// NOTE: Currently, we allow missing DeveloperDataId message,
			// we only use FieldDescription messages to decode developer data.
			if d.options.logWriter != nil {
				fmt.Fprintf(d.options.logWriter,
					"mesg.Num: %d, developerFields[%d].Num: %d: missing developer data id with developer data index '%d'",
					mesg.Num, i, devFieldDef.Num, devFieldDef.DeveloperDataIndex)
			}
		}

		// Find the FieldDescription that refers to this DeveloperField.
		// The combination of the Developer Data Index and Field Definition Number
		// create a unique id for each Field Description.
		var fieldDesc *mesgdef.FieldDescription
		for _, f := range d.fieldDescriptions {
			if f.DeveloperDataIndex != devFieldDef.DeveloperDataIndex {
				continue
			}
			if f.FieldDefinitionNumber != devFieldDef.Num {
				continue
			}
			fieldDesc = f
			break
		}

		if fieldDesc == nil {
			if d.options.logWriter != nil {
				fmt.Fprintf(d.options.logWriter, "mesg.Num: %d, developerFields[%d].Num: %d: Can't interpret developer field, "+
					"no field description mesg found. Just read acquired bytes (%d) and move forward. [byte pos: %d]\n",
					mesg.Num, i, devFieldDef.Num, devFieldDef.Size, d.n)
			}
			if _, err := d.readN(int(devFieldDef.Size)); err != nil {
				return fmt.Errorf("no field description found, unable to read acquired bytes: %w", err)
			}
			continue
		}

		if !fieldDesc.FitBaseTypeId.Valid() {
			return fmt.Errorf("fieldDescription.FitBaseTypeId: %s: %w",
				fieldDesc.FitBaseTypeId, errInvalidBaseType)
		}

		var isArray bool
		baseType := fieldDesc.FitBaseTypeId
		profileType := profile.ProfileTypeFromBaseType(baseType)

		// Gracefully handle poorly encoded FIT file.
		if devFieldDef.Size == 0 {
			d.logDeveloperField(mesg, devFieldDef, fieldDesc.FitBaseTypeId, "Size is zero. Skip")
			continue
		} else if devFieldDef.Size < fieldDesc.FitBaseTypeId.Size() {
			baseType = basetype.Byte
			profileType = profile.Byte
			d.logDeveloperField(mesg, devFieldDef, fieldDesc.FitBaseTypeId,
				"Size is less than expected. Fallback: decode as byte(s) and convert the value")
		}

		if devFieldDef.Size > baseType.Size() && devFieldDef.Size%baseType.Size() == 0 {
			isArray = true
		}

		// NOTE: It seems there is no standard on utilizing Array field to handle []string in developer fields.
		// Discussion: https://forums.garmin.com/developer/fit-sdk/f/discussion/355554/how-to-determine-developer-field-s-value-type-is-a-string-or-string
		overrideStringArray := fieldDesc.FitBaseTypeId == basetype.String
		val, err := d.readValue(devFieldDef.Size, mesgDef.Architecture, baseType, profileType, isArray, overrideStringArray)
		if err != nil {
			return err
		}

		if baseType != fieldDesc.FitBaseTypeId { // Convert value
			val = convertBytesToValue(val, fieldDesc.FitBaseTypeId)
		}

		// NOTE: Decoder will not attempt to validate native data when both NativeMesgNum and NativeFieldNum are valid.
		// Users need to handle this themselves due to the limited context available.
		mesg.DeveloperFields = append(mesg.DeveloperFields,
			proto.DeveloperField{
				Num:                devFieldDef.Num,
				DeveloperDataIndex: devFieldDef.DeveloperDataIndex,
				Value:              val,
			})
	}
	return nil
}

func (d *Decoder) decodeCRC() error {
	b, err := d.readBuffer.ReadN(2)
	if err != nil {
		return err
	}
	d.n += 2
	d.crc = binary.LittleEndian.Uint16(b)
	if d.options.shouldChecksum && d.crc16.Sum16() != d.crc { // check data integrity
		return fmt.Errorf("expected crc %d, got: %d: %w", d.crc, d.crc16.Sum16(), ErrCRCChecksumMismatch)
	}
	d.crc16.Reset()
	return nil
}

func (d *Decoder) readN(n int) ([]byte, error) {
	b, err := d.readBuffer.ReadN(n)
	if err != nil {
		return nil, err
	}
	d.n, d.cur = d.n+int64(n), d.cur+uint32(n)
	if d.options.shouldChecksum {
		_, _ = d.crc16.Write(b)
	}
	return b, nil
}

// readValue reads message value bytes from reader and convert it into its corresponding type. Size should not be zero.
func (d *Decoder) readValue(size byte, arch byte, baseType basetype.BaseType, profileType profile.ProfileType, isArray, overrideStringArray bool) (val proto.Value, err error) {
	b, err := d.readN(int(size))
	if err != nil {
		return val, err
	}
	if overrideStringArray && baseType == basetype.String {
		isArray = strcount(b) > 1
	}
	return proto.UnmarshalValue(b, arch, baseType, profileType, isArray)
}

const logFieldTemplate = "mesg.Num: %q, %s.Num: %d, size: %d, type: %q (size: %d). %s. [bytes pos: %d]\n"

// logField logs field related issues only if logWriter is not nil.
func (d *Decoder) logField(m *proto.Message, fd *proto.FieldDefinition, msg string) {
	if d.options.logWriter == nil {
		return
	}
	fmt.Fprintf(d.options.logWriter, logFieldTemplate, m.Num, "field", fd.Num, fd.Size, fd.BaseType, fd.BaseType.Size(), msg, d.n)
}

// logDeveloperField logs developerField related issues only if logWriter is not nil.
func (d *Decoder) logDeveloperField(m *proto.Message, dfd *proto.DeveloperFieldDefinition, bt basetype.BaseType, msg string) {
	if d.options.logWriter == nil {
		return
	}
	fmt.Fprintf(d.options.logWriter, logFieldTemplate, m.Num, "developerField", dfd.Num, dfd.Size, bt, bt.Size(), msg, d.n)
}

// DecodeWithContext is similar to Decode but with respect to context propagation.
func (d *Decoder) DecodeWithContext(ctx context.Context) (*proto.FIT, error) {
	if d.err != nil {
		return nil, d.err
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if d.err = checkContext(ctx); d.err != nil {
		return nil, d.err
	}
	if d.err = d.decodeFileHeaderOnce(); d.err != nil {
		return nil, d.err
	}
	defer d.releaseTemporaryObjects()
	if d.err = d.decodeMessagesWithContext(ctx); d.err != nil {
		return nil, d.err
	}
	if d.err = checkContext(ctx); d.err != nil {
		return nil, d.err
	}
	if d.err = d.decodeCRC(); d.err != nil {
		return nil, d.err
	}
	fit := &proto.FIT{
		FileHeader: d.fileHeader,
		Messages:   d.messages,
		CRC:        d.crc,
	}
	d.reset()
	return fit, nil
}

func checkContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (d *Decoder) decodeMessagesWithContext(ctx context.Context) (err error) {
	for d.cur < d.fileHeader.DataSize {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err = d.decodeMessage(); err != nil {
				return fmt.Errorf("decodeMessage [byte pos: %d]: %w", d.n, err)
			}
		}
	}
	return nil
}

// strcount counts how many valid string in b.
// This should align with the logic in proto.UnmarshalValue.
func strcount(b []byte) (size byte) {
	last := 0
	for i := range b {
		if b[i] == '\x00' {
			if last != i { // only if not an invalid string
				size++
			}
			last = i + 1
		}
	}
	return size
}

// convertUint32ToValue val into proto.Value of targeted baseType.
// If targeted baseType is not supported, it returns proto.Value{}.
func convertUint32ToValue(val uint32, baseType basetype.BaseType) proto.Value {
	switch baseType {
	case basetype.Sint8:
		return proto.Int8(int8(val))
	case basetype.Enum, basetype.Uint8, basetype.Uint8z:
		return proto.Uint8(uint8(val))
	case basetype.Sint16:
		return proto.Int16(int16(val))
	case basetype.Uint16, basetype.Uint16z:
		return proto.Uint16(uint16(val))
	case basetype.Sint32:
		return proto.Int32(int32(val))
	case basetype.Uint32, basetype.Uint32z:
		return proto.Uint32(uint32(val))
	case basetype.Sint64:
		return proto.Int64(int64(val))
	case basetype.Uint64, basetype.Uint64z:
		return proto.Uint64(uint64(val))
	case basetype.Float32:
		return proto.Float32(float32(val))
	case basetype.Float64:
		return proto.Float64(float64(val))
	}
	return proto.Value{}
}

// convertBytesToValue converts val in the form of byte or []byte into target baseType.
// This is used for casting value of bad encoded FIT files.
func convertBytesToValue(val proto.Value, baseType basetype.BaseType) proto.Value {
	var value uint64
	switch val.Type() {
	case proto.TypeUint8:
		value = uint64(val.Uint8())
	case proto.TypeSliceUint8:
		b := val.SliceUint8()
		for i := range b {
			value |= uint64(b[i]) << (i * 8)
		}
	}
	switch baseType {
	case basetype.Sint8:
		return proto.Int8(int8(value))
	case basetype.Enum, basetype.Uint8, basetype.Uint8z:
		return proto.Uint8(uint8(value))
	case basetype.Sint16:
		return proto.Int16(int16(value))
	case basetype.Uint16, basetype.Uint16z:
		return proto.Uint16(uint16(value))
	case basetype.Sint32:
		return proto.Int32(int32(value))
	case basetype.Uint32, basetype.Uint32z:
		return proto.Uint32(uint32(value))
	case basetype.Sint64:
		return proto.Int64(int64(value))
	case basetype.Uint64, basetype.Uint64z:
		return proto.Uint64(value)
	case basetype.Float32:
		return proto.Float32(float32(value))
	case basetype.Float64:
		return proto.Float64(float64(value))
	}
	return val
}

// valueAppend appends elem into slice. Elem must be has type element of
// slice's element. Otherwise, undefined behavior.
func valueAppend(slice proto.Value, elem proto.Value) proto.Value {
	switch elem.Type() {
	case proto.TypeInt8:
		return proto.SliceInt8(append(slice.SliceInt8(), elem.Int8()))
	case proto.TypeUint8:
		return proto.SliceUint8(append(slice.SliceUint8(), elem.Uint8()))
	case proto.TypeInt16:
		return proto.SliceInt16(append(slice.SliceInt16(), elem.Int16()))
	case proto.TypeUint16:
		return proto.SliceUint16(append(slice.SliceUint16(), elem.Uint16()))
	case proto.TypeInt32:
		return proto.SliceInt32(append(slice.SliceInt32(), elem.Int32()))
	case proto.TypeUint32:
		return proto.SliceUint32(append(slice.SliceUint32(), elem.Uint32()))
	case proto.TypeInt64:
		return proto.SliceInt64(append(slice.SliceInt64(), elem.Int64()))
	case proto.TypeUint64:
		return proto.SliceUint64(append(slice.SliceUint64(), elem.Uint64()))
	case proto.TypeFloat32:
		return proto.SliceFloat32(append(slice.SliceFloat32(), elem.Float32()))
	case proto.TypeFloat64:
		return proto.SliceFloat64(append(slice.SliceFloat64(), elem.Float64()))
	}
	return slice
}
