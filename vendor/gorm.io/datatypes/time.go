package datatypes

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Time is time data type.
type Time time.Duration

// NewTime is a constructor for Time and returns new Time.
func NewTime(hour, min, sec, nsec int) Time {
	return newTime(hour, min, sec, nsec)
}

func newTime(hour, min, sec, nsec int) Time {
	return Time(
		time.Duration(hour)*time.Hour +
			time.Duration(min)*time.Minute +
			time.Duration(sec)*time.Second +
			time.Duration(nsec)*time.Nanosecond,
	)
}

// GormDataType returns gorm common data type. This type is used for the field's column type.
func (Time) GormDataType() string {
	return "time"
}

// GormDBDataType returns gorm DB data type based on the current using database.
func (Time) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "mysql":
		return "TIME"
	case "postgres":
		return "TIME"
	case "sqlserver":
		return "TIME"
	case "sqlite":
		return "TEXT"
	default:
		return ""
	}
}

// Scan implements sql.Scanner interface and scans value into Time,
func (t *Time) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		t.setFromString(string(v))
	case string:
		t.setFromString(v)
	case time.Time:
		t.setFromTime(v)
	default:
		return errors.New(fmt.Sprintf("failed to scan value: %v", v))
	}

	return nil
}

func (t *Time) setFromString(str string) {
	var h, m, s, n int
	fmt.Sscanf(str, "%02d:%02d:%02d.%09d", &h, &m, &s, &n)
	*t = newTime(h, m, s, n)
}

func (t *Time) setFromTime(src time.Time) {
	*t = newTime(src.Hour(), src.Minute(), src.Second(), src.Nanosecond())
}

// Value implements driver.Valuer interface and returns string format of Time.
func (t Time) Value() (driver.Value, error) {
	return t.String(), nil
}

// String implements fmt.Stringer interface.
func (t Time) String() string {
	if nsec := t.nanoseconds(); nsec > 0 {
		return fmt.Sprintf("%02d:%02d:%02d.%09d", t.hours(), t.minutes(), t.seconds(), nsec)
	} else {
		// omit nanoseconds unless any value is specified
		return fmt.Sprintf("%02d:%02d:%02d", t.hours(), t.minutes(), t.seconds())
	}
}

func (t Time) hours() int {
	return int(time.Duration(t).Truncate(time.Hour).Hours())
}

func (t Time) minutes() int {
	return int((time.Duration(t) % time.Hour).Truncate(time.Minute).Minutes())
}

func (t Time) seconds() int {
	return int((time.Duration(t) % time.Minute).Truncate(time.Second).Seconds())
}

func (t Time) nanoseconds() int {
	return int((time.Duration(t) % time.Second).Nanoseconds())
}

// MarshalJSON implements json.Marshaler to convert Time to json serialization.
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// UnmarshalJSON implements json.Unmarshaler to deserialize json data.
func (t *Time) UnmarshalJSON(data []byte) error {
	// ignore null
	if string(data) == "null" {
		return nil
	}
	t.setFromString(strings.Trim(string(data), `"`))
	return nil
}
