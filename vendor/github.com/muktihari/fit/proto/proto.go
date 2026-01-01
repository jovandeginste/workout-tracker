// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"fmt"

	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
)

type errorString string

func (e errorString) Error() string { return string(e) }

// NOTE: The term "Global FIT Profile" refers to the definition provided in the Profile.xlsx.

const ( // header is 1 byte ->	 0bxxxxxxxx
	MesgDefinitionMask         = 0b01000000 // Mask for determining if the message type is a message definition.
	MesgNormalHeaderMask       = 0b00000000 // Mask for determining if the message type is a normal message data .
	MesgCompressedHeaderMask   = 0b10000000 // Mask for determining if the message type is a compressed timestamp message data.
	LocalMesgNumMask           = 0b00001111 // Mask for mapping normal message data to the message definition.
	CompressedLocalMesgNumMask = 0b01100000 // Mask for mapping compressed timestamp message data to the message definition. Used with CompressedBitShift.
	CompressedTimeMask         = 0b00011111 // Mask for measuring time offset value from header. Compressed timestamp is using 5 least significant bits (lsb) of header
	DevDataMask                = 0b00100000 // Mask for determining if a message contains developer fields.

	CompressedBitShift = 5 // Used for right-shifting the 5 least significant bits (lsb) of compressed time.

	LittleEndian = 0
	BigEndian    = 1

	DefaultFileHeaderSize byte   = 14     // The preferred size is 14
	DataTypeFIT           string = ".FIT" // FIT is a constant string ".FIT"

	// Field Num for timestamp across all defined messages in the profile.
	//
	// All timestamp fields should have num 253, except for these two messages that have different num
	// due to the messages may have been added before the rule was put in place:
	//   - Course Point's Timestamp num: 1
	//   - Set's Timestamp num: 254
	//
	// However, we don't see Official SDK implementations making exception regarding this, such as when decoding
	// and encoding compressed timestamp, only messages that have timestamp num 253 are eligible for that feature.
	//
	// Ref: https://forums.garmin.com/developer/fit-sdk/f/discussion/311106/why-has-course_point-message-a-timestamp-field-1-and-not-253
	FieldNumTimestamp = 253
)

// FIT represents a structure for FIT Files.
type FIT struct {
	FileHeader FileHeader // File Header contains either 12 or 14 bytes
	Messages   []Message  // Messages.
	CRC        uint16     // Cyclic Redundancy Check 16-bit value to ensure the integrity of the messages.
}

// FileHeader is a FIT's FileHeader with either 12 bytes size without CRC or a 14 bytes size with CRC, while 14 bytes size is the preferred size.
type FileHeader struct {
	Size            byte    // File header size either 12 (legacy) or 14.
	ProtocolVersion Version // The FIT Protocol version which is being used to encode the FIT file.
	ProfileVersion  uint16  // The FIT Profile Version (associated with data defined in Global FIT Profile).
	DataSize        uint32  // The size of the messages in bytes (this field will be automatically updated by the encoder)
	DataType        string  // ".FIT" (a string constant)
	CRC             uint16  // Cyclic Redundancy Check 16-bit value to ensure the integrity of the file header. (this field will be automatically updated by the encoder)
}

// MessageDefinition is the definition of the upcoming data messages.
type MessageDefinition struct {
	Header                    byte                       // The message definition header with mask 0b01000000.
	Reserved                  byte                       // Currently undetermined; the default value is 0.
	Architecture              byte                       // The Byte Order to be used to decode the values of both this message definition and the upcoming message. (0: Little-Endian, 1: Big-Endian)
	MesgNum                   typedef.MesgNum            // Global Message Number defined by factory (retrieved from Profile.xslx). (endianness of this 2 Byte value is defined in the Architecture byte)
	FieldDefinitions          []FieldDefinition          // List of the field definition
	DeveloperFieldDefinitions []DeveloperFieldDefinition // List of the developer field definition (only if Developer Data Flag is set in Header)
}

// FieldDefinition is the definition of the upcoming field within the message's structure.
type FieldDefinition struct {
	Num      byte              // The field definition number
	Size     byte              // The size of the upcoming value
	BaseType basetype.BaseType // The type of the upcoming value to be represented
}

// FieldDefinition is the definition of the upcoming developer field within the message's structure.
type DeveloperFieldDefinition struct { // 3 bits
	Num                byte // Maps to `field_definition_number` of a `field_description` message.
	Size               byte // Size (in bytes) of the specified FIT message’s field
	DeveloperDataIndex byte // Maps to `developer_data_index` of a `developer_data_id` and a `field_description` messages.
}

// Message is a FIT protocol message containing the data defined in the Message Definition
type Message struct {
	Header          byte             // Message Header serves to distinguish whether the message is a Normal Data or a Compressed Timestamp Data. Unlike MessageDefinition, Message's Header should not contain Developer Data Flag.
	Num             typedef.MesgNum  // Global Message Number defined in Global FIT Profile, except number within range 0xFF00 - 0xFFFE are manufacturer specific number.
	Fields          []Field          // List of Field
	DeveloperFields []DeveloperField // List of DeveloperField
}

// FieldByNum returns a pointer to the Field in a Message, if not found return nil.
func (m *Message) FieldByNum(num byte) *Field {
	for i := range m.Fields {
		if m.Fields[i].Num == num {
			return &m.Fields[i]
		}
	}
	return nil
}

// FieldValueByNum returns the value of the Field in a Messsage, if not found return invalid value.
func (m *Message) FieldValueByNum(num byte) Value {
	for i := range m.Fields {
		if m.Fields[i].Num == num {
			return m.Fields[i].Value
		}
	}
	return Value{}
}

// RemoveFieldByNum removes Field in a Message by num.
func (m *Message) RemoveFieldByNum(num byte) {
	for i := range m.Fields {
		if m.Fields[i].Num == num {
			m.Fields = append(m.Fields[:i], m.Fields[i+1:]...)
			return
		}
	}
}

// FieldBase acts as a fundamental representation of a field as defined in the Global FIT Profile.
// The value of this representation should not be altered, except in the case of an unknown field.
type FieldBase struct {
	Name       string              // Defined in the Global FIT profile for the specified FIT message, otherwise its a manufaturer specific name (defined by manufacturer).
	Num        byte                // Defined in the Global FIT profile for the specified FIT message, otherwise its a manufaturer specific number (defined by manufacturer). (255 == invalid)
	Type       profile.ProfileType // Type is defined type that serves as an abstraction layer above base types (primitive-types), e.g. DateTime is a time representation in uint32.
	BaseType   basetype.BaseType   // BaseType is the base of the ProfileType. E.g. profile.DateTime -> basetype.Uint32.
	Array      bool                // Flag whether the value of this field is an array
	Accumulate bool                // Flag to indicate if the value of the field is accumulable.
	Scale      float64             // A scale or offset specified in the FIT profile for binary fields (sint/uint etc.) only. the binary quantity is divided by the scale factor and then the offset is subtracted. (default: 1)
	Offset     float64             // A scale or offset specified in the FIT profile for binary fields (sint/uint etc.) only. the binary quantity is divided by the scale factor and then the offset is subtracted. (default: 0)
	Units      string              // Units of the value, such as m (meter), m/s (meter per second), s (second), etc.
	Components []Component         // List of components
	SubFields  []SubField          // List of sub-fields
}

// Field represents the full representation of a field, as specified in the Global FIT Profile.
type Field struct {
	// PERF: Embedding the struct as a pointer to avoid runtime duffcopy when creating a field since FieldBase should not be altered.
	*FieldBase

	// Value holds any primitive-type single value (or slice value) in a form of proto.Value.
	// We use proto.Value instead of interface{} because interface{} is heap-allocated, whereas proto.Value is not.
	Value Value

	// A flag to detect whether this field is generated through component expansion.
	IsExpandedField bool
}

var _ fmt.Formatter = (*Field)(nil)

// Format controls how Field is formatted when using fmt. Instead of only printing
// FieldBase's pointer, it also include the value, making it easier to debug.
func (f Field) Format(p fmt.State, verb rune) {
	switch {
	case verb != 'v':
		fmt.Fprintf(p, "%%!%c(%T=%v)", verb, f, f)
	case p.Flag('+'):
		fmt.Fprintf(p, "{FieldBase:(%p)(%+v) Value:%+v IsExpandedField:%t}",
			f.FieldBase, f.FieldBase, f.Value, f.IsExpandedField)
	case p.Flag('#'):
		fmt.Fprintf(p, "{FieldBase:(%p)(%#v), Value:%#v, IsExpandedField:%t}",
			f.FieldBase, f.FieldBase, f.Value, f.IsExpandedField)
	default: // %v
		fmt.Fprintf(p, "{(%p)(%v) %v %t}",
			f.FieldBase, f.FieldBase, f.Value, f.IsExpandedField)
	}
}

// WithValue returns a Field containing v value.
func (f Field) WithValue(v any) Field {
	f.Value = Any(v)
	return f
}

// SubFieldSubstitution returns any sub-field that can substitute the properties interpretation of the parent Field (Dynamic Field).
func (f *Field) SubFieldSubstitution(mesgRef *Message) *SubField {
	for i := range f.SubFields {
		subField := &f.SubFields[i]
		for j := range subField.Maps {
			smap := &subField.Maps[j]
			fieldRef := mesgRef.FieldByNum(smap.RefFieldNum)
			if fieldRef == nil {
				continue
			}
			if fieldRef.isValueEqualTo(smap.RefFieldValue) {
				return subField
			}
		}
	}
	return nil
}

// isValueEqualTo compare if Value == SubField's Map RefFieldValue.
// The FIT documentation on dynamic fields says: reference fields must be of integer type, floating point reference values are not supported.
func (f *Field) isValueEqualTo(refFieldValue int64) bool {
	fieldValue, ok := convertToInt64(f.Value)
	if !ok {
		return ok
	}
	return fieldValue == refFieldValue
}

// convertToInt64 converts any integer value of val to int64, if val is non-integer value return false.
func convertToInt64(val Value) (int64, bool) {
	switch val.Type() {
	case TypeInt8,
		TypeUint8,
		TypeInt16,
		TypeUint16,
		TypeInt32,
		TypeUint32,
		TypeInt64,
		TypeUint64:
		return int64(val.num), true
	}
	return 0, false
}

// DeveloperField is a way to add custom data fields to existing messages. Developer Data Fields can be added
// to any message at runtime by providing a self-describing FieldDefinition messages prior to that message.
// The combination of the DeveloperDataIndex and FieldDefinitionNumber create a unique id for each FieldDescription.
// Developer Data Fields are also used by the Connect IQ FIT Contributor library, allowing Connect IQ apps
// and data fields to include custom data in FIT Activity files during the recording of activities.
//
// NOTE: If DeveloperField contains a valid NativeMesgNum and NativeFieldNum, the value should be treated as
// native value (scale, offset, etc shall apply). [Added since protocol version 2.0]
type DeveloperField struct {
	Num                byte // Maps to `field_definition_number` of a `field_description` message.
	DeveloperDataIndex byte // Maps to `developer_data_index` of a `developer_data_id` and a `field_description` messages.
	Value              Value
}

// Component is a way of compressing one or more fields into a bit field expressed in a single containing field.
// The component can be expanded as a main Field in a Message or to update the value of the destination main Field.
type Component struct {
	FieldNum   byte
	Accumulate bool
	Bits       byte // bit value max 32
	Scale      float64
	Offset     float64
}

// SubField is a dynamic interpretation of the main Field in a Message when the SubFieldMap mapping match. See SubFieldMap's docs.
type SubField struct {
	Name       string
	Type       profile.ProfileType
	Scale      float64
	Offset     float64
	Units      string
	Maps       []SubFieldMap
	Components []Component
}

// SubFieldMap is the mapping between SubField and the corresponding main Field in a Message.
// When any Field in a Message has Field.Num == RefFieldNum and Field.Value == RefFieldValue, then the SubField containing
// this mapping can be interpreted as the main Field's properties (name, scale, type etc.)
type SubFieldMap struct {
	RefFieldNum   byte
	RefFieldValue int64
}

// LocalMesgNum extracts LocalMesgNum from message header.
func LocalMesgNum(header byte) byte {
	if (header & MesgCompressedHeaderMask) == MesgCompressedHeaderMask {
		return (header & CompressedLocalMesgNumMask) >> CompressedBitShift
	}
	return header & LocalMesgNumMask
}

const (
	errNilMesg            = errorString("mesg is nil")
	errValueSizeExceed255 = errorString("value's size exceed 255")
)

// NewMessageDefinition returns a new MessageDefinition based on the given Message or an error if one occurs.
// This will set Reserved and Architecture with 0 value. It's up to the caller to change the returning value.
//
// This serves as a testing helper and is for documentation purposes only.
func NewMessageDefinition(mesg *Message) (*MessageDefinition, error) {
	if mesg == nil {
		return nil, errNilMesg
	}

	const maxValueSize = 255

	mesgDef := &MessageDefinition{
		Header:           MesgDefinitionMask,
		Reserved:         0,
		Architecture:     LittleEndian,
		MesgNum:          mesg.Num,
		FieldDefinitions: make([]FieldDefinition, 0, len(mesg.Fields)),
	}

	for i := range mesg.Fields {
		size := mesg.Fields[i].Value.Size()
		if size > maxValueSize {
			return nil, fmt.Errorf("Fields[%d].Value's size should be <= %d: %w",
				i, maxValueSize, errValueSizeExceed255)
		}
		mesgDef.FieldDefinitions = append(mesgDef.FieldDefinitions, FieldDefinition{
			Num:      mesg.Fields[i].Num,
			Size:     byte(size),
			BaseType: mesg.Fields[i].BaseType,
		})
	}

	if len(mesg.DeveloperFields) == 0 {
		return mesgDef, nil
	}

	mesgDef.DeveloperFieldDefinitions = make([]DeveloperFieldDefinition, 0, len(mesg.DeveloperFields))
	mesgDef.Header |= DevDataMask
	for i := range mesg.DeveloperFields {
		size := mesg.DeveloperFields[i].Value.Size()
		if size > maxValueSize {
			return nil, fmt.Errorf("Fields[%d].Value's size should be <= %d: %w",
				i, maxValueSize, errValueSizeExceed255)
		}
		mesgDef.DeveloperFieldDefinitions = append(mesgDef.DeveloperFieldDefinitions, DeveloperFieldDefinition{
			Num:                mesg.DeveloperFields[i].Num,
			Size:               byte(size),
			DeveloperDataIndex: mesg.DeveloperFields[i].DeveloperDataIndex,
		})
	}

	return mesgDef, nil
}
