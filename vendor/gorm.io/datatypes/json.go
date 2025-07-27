package datatypes

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// JSON defined JSON data type, need to implements driver.Valuer, sql.Scanner interface
type JSON json.RawMessage

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return string(j), nil
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = JSON("null")
		return nil
	}
	var bytes []byte
	if s, ok := value.(fmt.Stringer); ok {
		bytes = []byte(s.String())
	} else {
		switch v := value.(type) {
		case []byte:
			if len(v) > 0 {
				bytes = make([]byte, len(v))
				copy(bytes, v)
			}
		case string:
			bytes = []byte(v)
		default:
			return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
		}
	}

	result := json.RawMessage(bytes)
	*j = JSON(result)
	return nil
}

// MarshalJSON to output non base64 encoded []byte
func (j JSON) MarshalJSON() ([]byte, error) {
	return json.RawMessage(j).MarshalJSON()
}

// UnmarshalJSON to deserialize []byte
func (j *JSON) UnmarshalJSON(b []byte) error {
	result := json.RawMessage{}
	err := result.UnmarshalJSON(b)
	*j = JSON(result)
	return err
}

func (j JSON) String() string {
	return string(j)
}

// GormDataType gorm common data type
func (JSON) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
func (JSON) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}

func (js JSON) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	if len(js) == 0 {
		return gorm.Expr("NULL")
	}

	data, _ := js.MarshalJSON()

	switch db.Dialector.Name() {
	case "mysql":
		if v, ok := db.Dialector.(*mysql.Dialector); ok && !strings.Contains(v.ServerVersion, "MariaDB") {
			return gorm.Expr("CAST(? AS JSON)", string(data))
		}
	}

	return gorm.Expr("?", string(data))
}

// JSONQueryExpression json query expression, implements clause.Expression interface to use as querier
type JSONQueryExpression struct {
	column      string
	keys        []string
	hasKeys     bool
	equals      bool
	likes       bool
	equalsValue interface{}
	extract     bool
	path        string
}

// JSONQuery query column as json
func JSONQuery(column string) *JSONQueryExpression {
	return &JSONQueryExpression{column: column}
}

// Extract extract json with path
func (jsonQuery *JSONQueryExpression) Extract(path string) *JSONQueryExpression {
	jsonQuery.extract = true
	jsonQuery.path = path
	return jsonQuery
}

// HasKey returns clause.Expression
func (jsonQuery *JSONQueryExpression) HasKey(keys ...string) *JSONQueryExpression {
	jsonQuery.keys = keys
	jsonQuery.hasKeys = true
	return jsonQuery
}

// Keys returns clause.Expression
func (jsonQuery *JSONQueryExpression) Equals(value interface{}, keys ...string) *JSONQueryExpression {
	jsonQuery.keys = keys
	jsonQuery.equals = true
	jsonQuery.equalsValue = value
	return jsonQuery
}

// Likes return clause.Expression
func (jsonQuery *JSONQueryExpression) Likes(value interface{}, keys ...string) *JSONQueryExpression {
	jsonQuery.keys = keys
	jsonQuery.likes = true
	jsonQuery.equalsValue = value
	return jsonQuery
}

// Build implements clause.Expression
func (jsonQuery *JSONQueryExpression) Build(builder clause.Builder) {
	if stmt, ok := builder.(*gorm.Statement); ok {
		switch stmt.Dialector.Name() {
		case "mysql", "sqlite":
			switch {
			case jsonQuery.extract:
				builder.WriteString("JSON_EXTRACT(")
				builder.WriteQuoted(jsonQuery.column)
				builder.WriteByte(',')
				builder.AddVar(stmt, prefix+jsonQuery.path)
				builder.WriteString(")")
			case jsonQuery.hasKeys:
				if len(jsonQuery.keys) > 0 {
					builder.WriteString("JSON_EXTRACT(")
					builder.WriteQuoted(jsonQuery.column)
					builder.WriteByte(',')
					builder.AddVar(stmt, jsonQueryJoin(jsonQuery.keys))
					builder.WriteString(") IS NOT NULL")
				}
			case jsonQuery.equals:
				if len(jsonQuery.keys) > 0 {
					builder.WriteString("JSON_EXTRACT(")
					builder.WriteQuoted(jsonQuery.column)
					builder.WriteByte(',')
					builder.AddVar(stmt, jsonQueryJoin(jsonQuery.keys))
					builder.WriteString(") = ")
					if value, ok := jsonQuery.equalsValue.(bool); ok {
						builder.WriteString(strconv.FormatBool(value))
					} else {
						stmt.AddVar(builder, jsonQuery.equalsValue)
					}
				}
			case jsonQuery.likes:
				if len(jsonQuery.keys) > 0 {
					builder.WriteString("JSON_EXTRACT(")
					builder.WriteQuoted(jsonQuery.column)
					builder.WriteByte(',')
					builder.AddVar(stmt, jsonQueryJoin(jsonQuery.keys))
					builder.WriteString(") LIKE ")
					if value, ok := jsonQuery.equalsValue.(bool); ok {
						builder.WriteString(strconv.FormatBool(value))
					} else {
						stmt.AddVar(builder, jsonQuery.equalsValue)
					}
				}
			}
		case "postgres":
			switch {
			case jsonQuery.extract:
				builder.WriteString(fmt.Sprintf("json_extract_path_text(%v::json,", stmt.Quote(jsonQuery.column)))
				stmt.AddVar(builder, jsonQuery.path)
				builder.WriteByte(')')
			case jsonQuery.hasKeys:
				if len(jsonQuery.keys) > 0 {
					stmt.WriteQuoted(jsonQuery.column)
					stmt.WriteString("::jsonb")
					for _, key := range jsonQuery.keys[0 : len(jsonQuery.keys)-1] {
						stmt.WriteString(" -> ")
						stmt.AddVar(builder, key)
					}

					stmt.WriteString(" ? ")
					stmt.AddVar(builder, jsonQuery.keys[len(jsonQuery.keys)-1])
				}
			case jsonQuery.equals:
				if len(jsonQuery.keys) > 0 {
					builder.WriteString(fmt.Sprintf("json_extract_path_text(%v::json,", stmt.Quote(jsonQuery.column)))

					for idx, key := range jsonQuery.keys {
						if idx > 0 {
							builder.WriteByte(',')
						}
						stmt.AddVar(builder, key)
					}
					builder.WriteString(") = ")

					if _, ok := jsonQuery.equalsValue.(string); ok {
						stmt.AddVar(builder, jsonQuery.equalsValue)
					} else {
						stmt.AddVar(builder, fmt.Sprint(jsonQuery.equalsValue))
					}
				}
			case jsonQuery.likes:
				if len(jsonQuery.keys) > 0 {
					builder.WriteString(fmt.Sprintf("json_extract_path_text(%v::json,", stmt.Quote(jsonQuery.column)))

					for idx, key := range jsonQuery.keys {
						if idx > 0 {
							builder.WriteByte(',')
						}
						stmt.AddVar(builder, key)
					}
					builder.WriteString(") LIKE ")

					if _, ok := jsonQuery.equalsValue.(string); ok {
						stmt.AddVar(builder, jsonQuery.equalsValue)
					} else {
						stmt.AddVar(builder, fmt.Sprint(jsonQuery.equalsValue))
					}
				}
			}
		}
	}
}

// JSONOverlapsExpression JSON_OVERLAPS expression, implements clause.Expression interface to use as querier
type JSONOverlapsExpression struct {
	column clause.Expression
	val    string
}

// JSONOverlaps query column as json
func JSONOverlaps(column clause.Expression, value string) *JSONOverlapsExpression {
	return &JSONOverlapsExpression{
		column: column,
		val:    value,
	}
}

// Build implements clause.Expression
// only mysql support JSON_OVERLAPS
func (json *JSONOverlapsExpression) Build(builder clause.Builder) {
	if stmt, ok := builder.(*gorm.Statement); ok {
		switch stmt.Dialector.Name() {
		case "mysql":
			builder.WriteString("JSON_OVERLAPS(")
			json.column.Build(builder)
			builder.WriteString(",")
			builder.AddVar(stmt, json.val)
			builder.WriteString(")")
		}
	}
}

type columnExpression string

func Column(col string) columnExpression {
	return columnExpression(col)
}

func (col columnExpression) Build(builder clause.Builder) {
	if stmt, ok := builder.(*gorm.Statement); ok {
		switch stmt.Dialector.Name() {
		case "mysql", "sqlite", "postgres":
			builder.WriteString(stmt.Quote(string(col)))
		}
	}
}

const prefix = "$."

func jsonQueryJoin(keys []string) string {
	if len(keys) == 1 {
		return prefix + keys[0]
	}

	n := len(prefix)
	n += len(keys) - 1
	for i := 0; i < len(keys); i++ {
		n += len(keys[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(prefix)
	b.WriteString(keys[0])
	for _, key := range keys[1:] {
		b.WriteString(".")
		b.WriteString(key)
	}
	return b.String()
}

// JSONSetExpression json set expression, implements clause.Expression interface to use as updater
type JSONSetExpression struct {
	column     string
	path2value map[string]interface{}
	mutex      sync.RWMutex
}

// JSONSet update fields of json column
func JSONSet(column string) *JSONSetExpression {
	return &JSONSetExpression{column: column, path2value: make(map[string]interface{})}
}

// Set return clause.Expression.
//
//	{
//		"age": 20,
//		"name": "json-1",
//		"orgs": {"orga": "orgv"},
//		"tags": ["tag1", "tag2"]
//	}
//
//	// In MySQL/SQLite, path is `age`, `name`, `orgs.orga`, `tags[0]`, `tags[1]`.
//	DB.UpdateColumn("attr", JSONSet("attr").Set("orgs.orga", 42))
//
//	// In PostgreSQL, path is `{age}`, `{name}`, `{orgs,orga}`, `{tags, 0}`, `{tags, 1}`.
//	DB.UpdateColumn("attr", JSONSet("attr").Set("{orgs, orga}", "bar"))
func (jsonSet *JSONSetExpression) Set(path string, value interface{}) *JSONSetExpression {
	jsonSet.mutex.Lock()
	jsonSet.path2value[path] = value
	jsonSet.mutex.Unlock()
	return jsonSet
}

// Build implements clause.Expression
// support mysql, sqlite and postgres
func (jsonSet *JSONSetExpression) Build(builder clause.Builder) {
	if stmt, ok := builder.(*gorm.Statement); ok {
		switch stmt.Dialector.Name() {
		case "mysql":

			var isMariaDB bool
			if v, ok := stmt.Dialector.(*mysql.Dialector); ok {
				isMariaDB = strings.Contains(v.ServerVersion, "MariaDB")
			}

			builder.WriteString("JSON_SET(")
			builder.WriteQuoted(jsonSet.column)
			for path, value := range jsonSet.path2value {
				builder.WriteByte(',')
				builder.AddVar(stmt, prefix+path)
				builder.WriteByte(',')

				if _, ok := value.(clause.Expression); ok {
					stmt.AddVar(builder, value)
					continue
				}

				rv := reflect.ValueOf(value)
				if rv.Kind() == reflect.Ptr {
					rv = rv.Elem()
				}
				switch rv.Kind() {
				case reflect.Slice, reflect.Array, reflect.Struct, reflect.Map:
					b, _ := json.Marshal(value)
					if isMariaDB {
						stmt.AddVar(builder, string(b))
						break
					}
					stmt.AddVar(builder, gorm.Expr("CAST(? AS JSON)", string(b)))
				case reflect.Bool:
					builder.WriteString(strconv.FormatBool(rv.Bool()))
				default:
					stmt.AddVar(builder, value)
				}
			}
			builder.WriteString(")")

		case "sqlite":
			builder.WriteString("JSON_SET(")
			builder.WriteQuoted(jsonSet.column)
			for path, value := range jsonSet.path2value {
				builder.WriteByte(',')
				builder.AddVar(stmt, prefix+path)
				builder.WriteByte(',')

				if _, ok := value.(clause.Expression); ok {
					stmt.AddVar(builder, value)
					continue
				}

				rv := reflect.ValueOf(value)
				if rv.Kind() == reflect.Ptr {
					rv = rv.Elem()
				}
				switch rv.Kind() {
				case reflect.Slice, reflect.Array, reflect.Struct, reflect.Map:
					b, _ := json.Marshal(value)
					stmt.AddVar(builder, gorm.Expr("JSON(?)", string(b)))
				default:
					stmt.AddVar(builder, value)
				}
			}
			builder.WriteString(")")

		case "postgres":
			var expr clause.Expression = columnExpression(jsonSet.column)
			for path, value := range jsonSet.path2value {
				if _, ok = value.(clause.Expression); ok {
					expr = gorm.Expr("JSONB_SET(?,?,?)", expr, path, value)
					continue
				} else {
					b, _ := json.Marshal(value)
					expr = gorm.Expr("JSONB_SET(?,?,?)", expr, path, string(b))
				}
			}
			stmt.AddVar(builder, expr)
		}
	}
}

func JSONArrayQuery(column string) *JSONArrayExpression {
	return &JSONArrayExpression{
		column: column,
	}
}

type JSONArrayExpression struct {
	contains    bool
	in          bool
	column      string
	keys        []string
	equalsValue interface{}
}

// Contains checks if column[keys] contains the value given. The keys parameter is only supported for MySQL and SQLite.
func (json *JSONArrayExpression) Contains(value interface{}, keys ...string) *JSONArrayExpression {
	json.contains = true
	json.equalsValue = value
	json.keys = keys
	return json
}

// In checks if columns[keys] is in the array value given. This method is only supported for MySQL and SQLite.
func (json *JSONArrayExpression) In(value interface{}, keys ...string) *JSONArrayExpression {
	json.in = true
	json.keys = keys
	json.equalsValue = value
	return json
}

// Build implements clause.Expression
func (json *JSONArrayExpression) Build(builder clause.Builder) {
	if stmt, ok := builder.(*gorm.Statement); ok {
		switch stmt.Dialector.Name() {
		case "mysql":
			switch {
			case json.contains:
				builder.WriteString("JSON_CONTAINS(" + stmt.Quote(json.column) + ",JSON_ARRAY(")
				builder.AddVar(stmt, json.equalsValue)
				builder.WriteByte(')')
				if len(json.keys) > 0 {
					builder.WriteByte(',')
					builder.AddVar(stmt, jsonQueryJoin(json.keys))
				}
				builder.WriteByte(')')
			case json.in:
				builder.WriteString("JSON_CONTAINS(JSON_ARRAY")
				builder.AddVar(stmt, json.equalsValue)
				builder.WriteByte(',')
				if len(json.keys) > 0 {
					builder.WriteString("JSON_EXTRACT(")
				}
				builder.WriteQuoted(json.column)
				if len(json.keys) > 0 {
					builder.WriteByte(',')
					builder.AddVar(stmt, jsonQueryJoin(json.keys))
					builder.WriteByte(')')
				}
				builder.WriteByte(')')
			}
		case "sqlite":
			switch {
			case json.contains:
				builder.WriteString("EXISTS(SELECT 1 FROM json_each(")
				builder.WriteQuoted(json.column)
				if len(json.keys) > 0 {
					builder.WriteByte(',')
					builder.AddVar(stmt, jsonQueryJoin(json.keys))
				}
				builder.WriteString(") WHERE value = ")
				builder.AddVar(stmt, json.equalsValue)
				builder.WriteString(") AND json_array_length(")
				builder.WriteQuoted(json.column)
				if len(json.keys) > 0 {
					builder.WriteByte(',')
					builder.AddVar(stmt, jsonQueryJoin(json.keys))
				}
				builder.WriteString(") > 0")
			case json.in:
				builder.WriteString("CASE WHEN json_type(")
				builder.WriteQuoted(json.column)
				if len(json.keys) > 0 {
					builder.WriteByte(',')
					builder.AddVar(stmt, jsonQueryJoin(json.keys))
				}
				builder.WriteString(") = 'array' THEN NOT EXISTS(SELECT 1 FROM json_each(")
				builder.WriteQuoted(json.column)
				if len(json.keys) > 0 {
					builder.WriteByte(',')
					builder.AddVar(stmt, jsonQueryJoin(json.keys))
				}
				builder.WriteString(") WHERE value NOT IN ")
				builder.AddVar(stmt, json.equalsValue)
				builder.WriteString(") ELSE ")
				if len(json.keys) > 0 {
					builder.WriteString("json_extract(")
				}
				builder.WriteQuoted(json.column)
				if len(json.keys) > 0 {
					builder.WriteByte(',')
					builder.AddVar(stmt, jsonQueryJoin(json.keys))
					builder.WriteByte(')')
				}
				builder.WriteString(" IN ")
				builder.AddVar(stmt, json.equalsValue)
				builder.WriteString(" END")
			}
		case "postgres":
			switch {
			case json.contains:
				builder.WriteString(stmt.Quote(json.column))
				builder.WriteString(" ? ")
				builder.AddVar(stmt, json.equalsValue)
			}
		}
	}
}
