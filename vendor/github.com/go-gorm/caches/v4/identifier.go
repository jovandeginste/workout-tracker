package caches

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm/callbacks"

	"gorm.io/gorm"
)

const IdentifierPrefix = "gorm-caches::"

func buildIdentifier(db *gorm.DB) string {
	// Build query identifier,
	//	for that reason we need to compile all arguments into a string
	//	and concat them with the SQL query itself
	callbacks.BuildQuerySQL(db)
	query := db.Statement.SQL.String()
	queryArgs := valueToString(db.Statement.Vars)
	identifier := fmt.Sprintf("%s%s-%s", IdentifierPrefix, query, queryArgs)
	return identifier
}

func valueToString(value interface{}) string {
	valueOf := reflect.ValueOf(value)
	switch valueOf.Kind() {
	case reflect.Ptr:
		if valueOf.IsNil() {
			return "<nil>"
		}
		return valueToString(valueOf.Elem().Interface())
	case reflect.Map:
		var sb strings.Builder
		sb.WriteString("{")
		for i, key := range valueOf.MapKeys() {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf("%s: %s", valueToString(key.Interface()), valueToString(valueOf.MapIndex(key).Interface())))
		}
		sb.WriteString("}")
		return sb.String()
	case reflect.Slice:
		valueSlice := make([]interface{}, valueOf.Len())
		for i := range valueSlice {
			valueSlice[i] = valueToString(valueOf.Index(i).Interface())
		}
		return fmt.Sprintf("%v", valueSlice)
	default:
		return fmt.Sprintf("%v", value)
	}
}
