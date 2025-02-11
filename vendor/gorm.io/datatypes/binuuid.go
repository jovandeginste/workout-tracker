package datatypes

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// This datatype is similar to datatypes.UUID, major difference being that
// this datatype stores the uuid in the database as a binary (byte) array
// instead of a string. Developers may use either as per their preference.
type BinUUID uuid.UUID

// NewBinUUIDv1 generates a uuid version 1, panics on generation failure.
func NewBinUUIDv1() BinUUID {
	return BinUUID(uuid.Must(uuid.NewUUID()))
}

// NewBinUUIDv4 generates a uuid version 4, panics on generation failure.
func NewBinUUIDv4() BinUUID {
	return BinUUID(uuid.Must(uuid.NewRandom()))
}

// NewNilBinUUID generates a nil uuid.
func NewNilBinUUID() BinUUID {
	return BinUUID(uuid.Nil)
}

// BinUUIDFromString returns the BinUUID representation of the specified uuidStr.
func BinUUIDFromString(uuidStr string) BinUUID {
	return BinUUID(uuid.MustParse(uuidStr))
}

// GormDataType gorm common data type.
func (BinUUID) GormDataType() string {
	return "BINARY(16)"
}

// GormDBDataType gorm db data type.
func (BinUUID) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "mysql":
		return "BINARY(16)"
	case "postgres":
		return "BYTEA"
	case "sqlserver":
		return "BINARY(16)"
	case "sqlite":
		return "BLOB"
	default:
		return ""
	}
}

// Scan is the scanner function for this datatype.
func (u *BinUUID) Scan(value interface{}) error {
	valueBytes, ok := value.([]byte)
	if !ok {
		return errors.New("unable to convert value to bytes")
	}
	valueUUID, err := uuid.FromBytes(valueBytes)
	if err != nil {
		return err
	}
	*u = BinUUID(valueUUID)
	return nil
}

// Value is the valuer function for this datatype.
func (u BinUUID) Value() (driver.Value, error) {
	return uuid.UUID(u).MarshalBinary()
}

// String returns the string form of the UUID.
func (u BinUUID) Bytes() []byte {
	bytes, err := uuid.UUID(u).MarshalBinary()
	if err != nil {
		return nil
	}
	return bytes
}

// String returns the string form of the UUID.
func (u BinUUID) String() string {
	return uuid.UUID(u).String()
}

// Equals returns true if bytes form of BinUUID matches other, false otherwise.
func (u BinUUID) Equals(other BinUUID) bool {
	return bytes.Equal(u.Bytes(), other.Bytes())
}

// Length returns the number of characters in string form of UUID.
func (u BinUUID) LengthBytes() int {
	return len(u.Bytes())
}

// Length returns the number of characters in string form of UUID.
func (u BinUUID) Length() int {
	return len(u.String())
}

// IsNil returns true if the BinUUID is nil uuid (all zeroes), false otherwise.
func (u BinUUID) IsNil() bool {
	return uuid.UUID(u) == uuid.Nil
}

// IsEmpty returns true if BinUUID is nil uuid or of zero length, false otherwise.
func (u BinUUID) IsEmpty() bool {
	return u.IsNil() || u.Length() == 0
}

// IsNilPtr returns true if caller BinUUID ptr is nil, false otherwise.
func (u *BinUUID) IsNilPtr() bool {
	return u == nil
}

// IsEmptyPtr returns true if caller BinUUID ptr is nil or it's value is empty.
func (u *BinUUID) IsEmptyPtr() bool {
	return u.IsNilPtr() || u.IsEmpty()
}
