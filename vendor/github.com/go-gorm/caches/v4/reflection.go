package caches

import (
	"errors"
	"fmt"
	"reflect"

	"gorm.io/gorm/schema"
)

func SetPointedValue(dest interface{}, src interface{}) {
	reflect.ValueOf(dest).Elem().Set(reflect.ValueOf(src).Elem())
}

func deepCopy(src, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if srcVal.Type() != dstVal.Elem().Type() {
		return errors.New("src and dst must be of the same type")
	}

	return copyValue(srcVal, dstVal.Elem())
}

func copyValue(src, dst reflect.Value) error {
	switch src.Kind() {
	case reflect.Ptr:
		src = src.Elem()
		dst.Set(reflect.New(src.Type()))
		err := copyValue(src, dst.Elem())
		if err != nil {
			return err
		}

	case reflect.Struct:
		for i := 0; i < src.NumField(); i++ {
			if src.Type().Field(i).PkgPath != "" {
				return fmt.Errorf("%w: %+v", schema.ErrUnsupportedDataType, src.Type().Field(i).Name)
			}
			err := copyValue(src.Field(i), dst.Field(i))
			if err != nil {
				return err
			}
		}

	case reflect.Slice:
		newSlice := reflect.MakeSlice(src.Type(), src.Len(), src.Cap())
		for i := 0; i < src.Len(); i++ {
			err := copyValue(src.Index(i), newSlice.Index(i))
			if err != nil {
				return err
			}
		}
		dst.Set(newSlice)

	case reflect.Map:
		newMap := reflect.MakeMapWithSize(src.Type(), src.Len())
		for _, key := range src.MapKeys() {
			value := src.MapIndex(key)
			newValue := reflect.New(value.Type()).Elem()
			err := copyValue(value, newValue)
			if err != nil {
				return err
			}
			newMap.SetMapIndex(key, newValue)
		}
		dst.Set(newMap)

	default:
		dst.Set(src)
	}

	return nil
}
