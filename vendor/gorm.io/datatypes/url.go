package datatypes

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type URL url.URL

func (u URL) Value() (driver.Value, error) {
	return u.String(), nil
}

func (u *URL) Scan(value interface{}) error {
	var us string
	switch v := value.(type) {
	case []byte:
		us = string(v)
	case string:
		us = v
	default:
		return errors.New(fmt.Sprint("Failed to parse URL:", value))
	}
	uu, err := url.Parse(us)
	if err != nil {
		return err
	}
	*u = URL(*uu)
	return nil
}

func (URL) GormDataType() string {
	return "url"
}

func (URL) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "TEXT"
}

func (u *URL) String() string {
	return (*url.URL)(u).String()
}

func (u URL) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

func (u *URL) UnmarshalJSON(data []byte) error {
	// ignore null
	if string(data) == "null" {
		return nil
	}
	uu, err := url.Parse(strings.Trim(string(data), `"'`))
	if err != nil {
		return err
	}
	*u = URL(*uu)
	return nil
}
