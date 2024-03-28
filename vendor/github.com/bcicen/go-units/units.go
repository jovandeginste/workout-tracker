// Package units is a library for manipulating and converting between various units of measurement
package units

import (
	"errors"
	"sort"
	"strings"
)

// Return all registered Units, sorted by name and quantity
func All() []Unit {
	units := make(UnitList, 0, len(unitMap))
	for _, u := range unitMap {
		units = append(units, u)
	}
	sort.Sort(units)
	return []Unit(units)
}

// MustConvertFloat converts a provided float from one Unit to another, panicking on error
func MustConvertFloat(x float64, from, to Unit) Value {
	val, err := ConvertFloat(x, from, to)
	if err != nil {
		panic(err)
	}
	return val
}

// ConvertFloat converts a provided float from one Unit to another
func ConvertFloat(x float64, from, to Unit) (Value, error) {
	path, err := ResolveConversion(from, to)
	if err != nil {
		return Value{}, err
	}

	for _, c := range path {
		x = c.Fn(x)
	}

	return Value{x, to}, nil
}

// Find Unit matching name or symbol provided
func Find(s string) (Unit, error) {

	// first try case-sensitive match
	for _, u := range unitMap {
		if matchUnit(s, u, true) {
			return u, nil
		}
	}

	// then case-insensitive
	for _, u := range unitMap {
		if matchUnit(s, u, false) {
			return u, nil
		}
	}

	// finally, try stripping plural suffix
	if strings.HasSuffix(s, "s") || strings.HasSuffix(s, "S") {
		s = s[:len(s)-1]
		for _, u := range unitMap {
			if matchUnit(s, u, false) {
				return u, nil
			}
		}
	}

	return Unit{}, errors.New("unit \"" + s + "\" not found")
}

func matchUnit(s string, u Unit, matchCase bool) bool {
	for _, name := range u.Names() {
		if matchCase {
			if name == s {
				return true
			}
		} else {
			if strings.EqualFold(s, name) {
				return true
			}
		}
	}

	return false
}
