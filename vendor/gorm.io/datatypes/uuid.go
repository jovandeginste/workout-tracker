package datatypes

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// This datatype stores the uuid in the database as a string. To store the uuid
// in the database as a binary (byte) array, please refer to datatypes.BinUUID.
type UUID uuid.UUID

// NewUUIDv1 generates a UUID version 1, panics on generation failure.
func NewUUIDv1() UUID {
	return UUID(uuid.Must(uuid.NewUUID()))
}

// NewUUIDv4 generates a UUID version 4, panics on generation failure.
func NewUUIDv4() UUID {
	return UUID(uuid.Must(uuid.NewRandom()))
}

// GormDataType gorm common data type.
func (UUID) GormDataType() string {
	return "string"
}

// GormDBDataType gorm db data type.
func (UUID) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "mysql":
		return "LONGTEXT"
	case "postgres":
		return "UUID"
	case "sqlserver":
		return "NVARCHAR(128)"
	case "sqlite":
		return "TEXT"
	default:
		return ""
	}
}

// Scan is the scanner function for this datatype.
func (u *UUID) Scan(value interface{}) error {
	var result uuid.UUID
	if err := result.Scan(value); err != nil {
		return err
	}
	*u = UUID(result)
	return nil
}

// Value is the valuer function for this datatype.
func (u UUID) Value() (driver.Value, error) {
	return uuid.UUID(u).Value()
}

// String returns the string form of the UUID.
func (u UUID) String() string {
	return uuid.UUID(u).String()
}

// Equals returns true if string form of UUID matches other, false otherwise.
func (u UUID) Equals(other UUID) bool {
	return u.String() == other.String()
}

// Length returns the number of characters in string form of UUID.
func (u UUID) Length() int {
	return len(u.String())
}

// IsNil returns true if the UUID is a nil UUID (all zeroes), false otherwise.
func (u UUID) IsNil() bool {
	return uuid.UUID(u) == uuid.Nil
}

// IsEmpty returns true if UUID is nil UUID or of zero length, false otherwise.
func (u UUID) IsEmpty() bool {
	return u.IsNil() || u.Length() == 0
}

// IsNilPtr returns true if caller UUID ptr is nil, false otherwise.
func (u *UUID) IsNilPtr() bool {
	return u == nil
}

// IsEmptyPtr returns true if caller UUID ptr is nil or it's value is empty.
func (u *UUID) IsEmptyPtr() bool {
	return u.IsNilPtr() || u.IsEmpty()
}
