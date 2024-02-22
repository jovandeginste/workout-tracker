package util

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func ToTime(i interface{}) (t time.Time, err error) {
	i = Indirect(i)
	if i == nil {
		return time.Time{}, errors.New("time is nil")
	}

	switch v := i.(type) {
	case time.Time:
		return v, nil
	case time.Duration:
		return time.Now().Add(v), nil
	}

	if num, err := ToNumber(i); err == nil {
		return time.Unix(int64(num), 0), nil
	}
	return time.Time{}, fmt.Errorf("unable to cast %#v of type %T to Time", i, i)
}

func ToNumber(n interface{}) (float64, error) {
	n = Indirect(n)
	if n == nil {
		return 0, errors.New("number is nil")
	}

	switch nt := n.(type) {
	case uint:
		return float64(nt), nil
	case uint8:
		return float64(nt), nil
	case uint16:
		return float64(nt), nil
	case uint32:
		return float64(nt), nil
	case uint64:
		return float64(nt), nil
	case int:
		return float64(nt), nil
	case int8:
		return float64(nt), nil
	case int16:
		return float64(nt), nil
	case int32:
		return float64(nt), nil
	case int64:
		return float64(nt), nil
	case float32:
		return float64(nt), nil
	case float64:
		return nt, nil
	case complex64:
		return float64(real(nt)), nil
	case complex128:
		return real(nt), nil
	case string:
		res, err := strconv.ParseFloat(nt, 64)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to float64", n, n)
		}
		return res, nil
	}

	if num, errC := ToNumber(fmt.Sprintf("%d", n)); errC == nil {
		return num, nil
	}

	return ToNumber(fmt.Sprintf("%v", n))
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// Indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func Indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}
