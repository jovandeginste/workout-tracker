package fit

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"time"
	"unicode/utf8"

	"github.com/tormoder/fit/dyncrc16"
	"github.com/tormoder/fit/internal/types"
)

type encoder struct {
	w    io.Writer
	arch binary.ByteOrder
}

func encodeString(str string, size byte) ([]byte, error) {
	length := len(str)
	if length > int(size)-1 {
		length = int(size) - 1
	}

	bstr := make([]byte, size)
	copy(bstr, str[:length])
	if !utf8.Valid(bstr) {
		return nil, fmt.Errorf("can't encode %+v as UTF-8 string", str)
	}
	return bstr, nil
}

func (e *encoder) encodeValue(value interface{}, f *field) error {
	switch f.t.Kind() {
	case types.TimeUTC:
		t := value.(time.Time)
		u32 := encodeTime(t)
		binary.Write(e.w, e.arch, u32)
	case types.TimeLocal:
		t := value.(time.Time)
		_, offs := t.Zone()
		u32 := uint32(int64(encodeTime(t)) + int64(offs))
		binary.Write(e.w, e.arch, u32)
	case types.Lat:
		lat := value.(Latitude)
		binary.Write(e.w, e.arch, lat.semicircles)
	case types.Lng:
		lng := value.(Longitude)
		binary.Write(e.w, e.arch, lng.semicircles)
	case types.NativeFit:
		if f.t.BaseType() == types.BaseString {
			str, ok := value.(string)
			if !ok {
				return fmt.Errorf("not a string: %+v", value)
			}

			var err error
			value, err = encodeString(str, f.length)
			if err != nil {
				return fmt.Errorf("can't encode %+v as UTF-8 string: %w", value, err)
			}
		}
		binary.Write(e.w, e.arch, value)
	default:
		return fmt.Errorf("unknown Fit type %+v", f.t)
	}

	return nil
}

func (e *encoder) writeField(value reflect.Value, f *field) error {
	if !f.t.Array() {
		return e.encodeValue(value.Interface(), f)
	}

	if f.t.BaseType() == types.BaseString {
		return fmt.Errorf("can't encode array of strings")
	}

	invalid := f.t.BaseType().Invalid()
	max := byte(value.Len())
	if max > f.length {
		max = f.length
	}
	for i := byte(0); i < max; i++ {
		elem := value.Index(int(i))
		err := e.encodeValue(elem.Interface(), f)
		if err != nil {
			return err
		}
	}
	for i := max; i < f.length; i++ {
		err := e.encodeValue(invalid, f)
		if err != nil {
			return err
		}
	}

	return nil
}

type encodeMesgDef struct {
	globalMesgNum MesgNum
	localMesgNum  byte
	fields        []*field
}

func (e *encoder) writeMesg(mesg reflect.Value, def *encodeMesgDef) error {
	hdr := def.localMesgNum & localMesgNumMask
	err := binary.Write(e.w, e.arch, hdr)
	if err != nil {
		return err
	}

	for _, f := range def.fields {
		value := mesg.Field(f.sindex)

		err := e.writeField(value, f)
		if err != nil {
			return err
		}
	}

	return nil
}

func profileFieldDef(m MesgNum) [256]*field {
	return _fields[m]
}

func getFieldBySindex(index int, fields [256]*field) *field {
	for _, f := range fields {
		if f != nil && index == f.sindex {
			return f
		}
	}

	return fields[255]
}

// getEncodeMesgDef generates an appropriate encodeMesgDef to will encode all
// of the valid fields in mesg. Any fields which are set to their respective
// invalid value will be skipped (not present in the returned encodeMesgDef)
func getEncodeMesgDef(mesg reflect.Value, localMesgNum byte) *encodeMesgDef {
	mesgNum := getGlobalMesgNum(mesg.Type())
	allInvalid := getMesgAllInvalid(mesgNum)
	profileFields := profileFieldDef(mesgNum)

	if mesg.NumField() != allInvalid.NumField() {
		panic(fmt.Sprintf("mismatched number of fields in type %+v", mesg.Type()))
	}

	def := &encodeMesgDef{
		globalMesgNum: mesgNum,
		localMesgNum:  localMesgNum,
		fields:        make([]*field, 0, mesg.NumField()),
	}

	for i := 0; i < mesg.NumField(); i++ {
		fval := mesg.Field(i)
		field := getFieldBySindex(i, profileFields)

		// Don't encode invalid fields
		if fval.Kind() == reflect.Slice {
			if fval.IsNil() {
				continue
			}

			skip := true
			invalid := field.t.BaseType().Invalid()
			for i := 0; i < fval.Len(); i++ {
				if fval.Interface() != invalid {
					skip = false
					break
				}
			}
			if skip {
				continue
			}
		} else if fval.Interface() == allInvalid.Field(i).Interface() {
			continue
		}

		// FIXME: No message can exceed 255 bytes
		def.fields = append(def.fields, field)
	}

	return def
}

func (e *encoder) writeDefMesg(def *encodeMesgDef) error {
	hdr := mesgDefinitionMask | (def.localMesgNum & localMesgNumMask)
	err := binary.Write(e.w, e.arch, hdr)
	if err != nil {
		return err
	}

	err = binary.Write(e.w, e.arch, byte(0))
	if err != nil {
		return err
	}

	switch e.arch {
	case binary.LittleEndian:
		err = binary.Write(e.w, e.arch, byte(0))
	case binary.BigEndian:
		err = binary.Write(e.w, e.arch, byte(1))
	}
	if err != nil {
		return err
	}

	err = binary.Write(e.w, e.arch, def.globalMesgNum)
	if err != nil {
		return err
	}

	err = binary.Write(e.w, e.arch, byte(len(def.fields)))
	if err != nil {
		return err
	}

	for _, f := range def.fields {
		fdef := fieldDef{
			num:   f.num,
			size:  byte(f.t.BaseType().Size()),
			btype: f.t.BaseType(),
		}
		if fdef.btype == types.BaseString {
			fdef.size = f.length
		} else if f.t.Array() {
			fdef.size = fdef.size * f.length
		}

		err := binary.Write(e.w, e.arch, fdef)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *encoder) encodeDefAndDataMesg(mesg reflect.Value) error {
	mesg = reflect.Indirect(mesg)
	if !mesg.IsValid() {
		return nil
	}

	// We'll always just use local ID 0, for simplicity
	// We know the full file contents up-front, so no need to interleave
	def := getEncodeMesgDef(mesg, 0)

	err := e.writeDefMesg(def)
	if err != nil {
		return err
	}

	err = e.writeMesg(mesg, def)
	if err != nil {
		return err
	}

	return err
}

func (e *encoder) encodeFile(file reflect.Value) error {
	for i := 0; i < file.NumField(); i++ {
		v := file.Field(i)
		switch v.Kind() {
		case reflect.Struct, reflect.Ptr:
			err := e.encodeDefAndDataMesg(reflect.Indirect(v))
			if err != nil {
				return err
			}
		case reflect.Slice:
			var def *encodeMesgDef
			for j := 0; j < v.Len(); j++ {
				v2 := reflect.Indirect(v.Index(j))

				// Not necessary that the first message will have all defined fields that may appear in the following messages
				// So we have to build a model first by iterating though all the message and collecting valid definition fields
				if def == nil {
					// map to collect field definitions
					mfields := make(map[byte]*field)
					for k := 0; k < v.Len(); k++ {
						r := reflect.Indirect(v.Index(k))
						def = getEncodeMesgDef(r, 0)
						for _, f := range def.fields {
							mfields[f.num] = f
						}
					}
					// should not be nil at this point, but just in case
					if def != nil {
						def.fields = make([]*field, 0, len(mfields))
						for _, f := range mfields {
							def.fields = append(def.fields, f)
						}
						err := e.writeDefMesg(def)
						if err != nil {
							return err
						}
					} else {
						return fmt.Errorf("cannot create definition message for %+v", v.Interface())
					}
				}

				err := e.writeMesg(v2, def)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// Encode writes the given FIT file into the given Writer. file.CRC and
// file.Header.CRC will be updated to the correct values.
func Encode(w io.Writer, file *File, arch binary.ByteOrder) error {
	buf := &bytes.Buffer{}
	enc := &encoder{
		w:    buf,
		arch: arch,
	}

	// XXX: Is there a better way to do this with reflection?
	var data reflect.Value
	switch file.Type() {
	case FileTypeActivity:
		activity, err := file.Activity()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*activity)
	case FileTypeDevice:
		device, err := file.Device()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*device)
	case FileTypeSettings:
		settings, err := file.Settings()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*settings)
	case FileTypeSport:
		sport, err := file.Sport()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*sport)
	case FileTypeWorkout:
		workout, err := file.Workout()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*workout)
	case FileTypeCourse:
		course, err := file.Course()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*course)
	case FileTypeSchedules:
		schedules, err := file.Schedules()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*schedules)
	case FileTypeWeight:
		weight, err := file.Weight()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*weight)
	case FileTypeTotals:
		totals, err := file.Totals()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*totals)
	case FileTypeGoals:
		goals, err := file.Goals()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*goals)
	case FileTypeBloodPressure:
		bloodPressure, err := file.BloodPressure()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*bloodPressure)
	case FileTypeMonitoringA:
		monitoringA, err := file.MonitoringA()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*monitoringA)
	case FileTypeActivitySummary:
		activitySummary, err := file.ActivitySummary()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*activitySummary)
	case FileTypeMonitoringDaily:
		monitoringDaily, err := file.MonitoringDaily()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*monitoringDaily)
	case FileTypeMonitoringB:
		monitoringB, err := file.MonitoringB()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*monitoringB)
	case FileTypeSegment:
		segment, err := file.Segment()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*segment)
	case FileTypeSegmentList:
		segmentList, err := file.SegmentList()
		if err != nil {
			return fmt.Errorf("encode failed: %w", err)
		}
		data = reflect.ValueOf(*segmentList)
	default:
		return fmt.Errorf("encode failed: Unknown filetype '%v'", file.Type())
	}

	// Encode the data
	err := enc.encodeDefAndDataMesg(reflect.ValueOf(file.FileId))
	if err != nil {
		return fmt.Errorf("encode failed: FileId: %w", err)
	}

	err = enc.encodeDefAndDataMesg(reflect.ValueOf(file.FileCreator))
	if err != nil {
		return fmt.Errorf("encode failed: FileCreator: %w", err)
	}

	err = enc.encodeDefAndDataMesg(reflect.ValueOf(file.TimestampCorrelation))
	if err != nil {
		return fmt.Errorf("encode failed: TimestampCorrelation: %w", err)
	}

	err = enc.encodeFile(data)
	if err != nil {
		return fmt.Errorf("encode failed: %vFile: %w", file.Type(), err)
	}

	file.Header.DataSize = uint32(buf.Len())
	hdr, err := file.Header.MarshalBinary()
	if err != nil {
		return fmt.Errorf("encode failed: Header: %w", err)
	}

	// Calculate file CRC
	crc := dyncrc16.New()

	_, err = crc.Write(hdr)
	if err != nil {
		return fmt.Errorf("encode failed: header crc calc: %w", err)
	}

	_, err = crc.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("encode failed: data crc calc: %w", err)
	}

	file.CRC = crc.Sum16()

	// Write out the data
	_, err = w.Write(hdr)
	if err != nil {
		return fmt.Errorf("encode failed: writing header: %w", err)
	}

	_, err = w.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("encode failed: writing data: %w", err)
	}

	err = binary.Write(w, binary.LittleEndian, file.CRC)
	if err != nil {
		return fmt.Errorf("encode failed: writing crc: %w", err)
	}

	return nil
}
