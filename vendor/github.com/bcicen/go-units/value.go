package units

import (
	"math"
	"strconv"
	"strings"
)

var DefaultFmtOptions = FmtOptions{true, false, 6}

type FmtOptions struct {
	Label     bool // if false, unit label/symbol will be omitted
	Short     bool // if true, use unit shortname or symbol
	Precision int  // maximum meaningful precision to truncate value
}

type Value struct {
	val  float64
	unit Unit
}

// NewValue creates a new Value instance
func NewValue(v float64, u Unit) Value { return Value{v, u} }

func (v Value) Unit() Unit     { return v.unit }
func (v Value) Float() float64 { return v.val }
func (v Value) String() string { return v.Fmt(DefaultFmtOptions) }

func (v Value) Fmt(opts FmtOptions) string {
	var label string

	if opts.Short {
		label = v.unit.Symbol
	} else {
		label = v.unit.Name
		// make label plural if needed
		if v.val > 1.0 {
			label = v.unit.PluralName()
		}
	}

	prec := opts.Precision
	// expand precision if needed to present meaningful value
	if v.val < 1 && v.val > 0 {
		prec = int((math.Log10(v.val)-0.5)*-1) + prec
	}

	vstr := strconv.FormatFloat(v.val, 'f', prec, 64)
	vstr = trimTrailing(vstr)

	if !opts.Label {
		return vstr
	}
	return vstr + " " + label
}

// MustConvert converts this Value to another Unit, panicking on error
func (v Value) MustConvert(to Unit) Value {
	newV, err := v.Convert(to)
	if err != nil {
		panic(err)
	}
	return newV
}

// Convert converts this Value to another Unit
func (v Value) Convert(to Unit) (Value, error) {
	// allow converting to same unit
	if v.unit.Name == to.Name {
		return v, nil
	}

	return ConvertFloat(v.val, v.unit, to)
}

// Trim trailing zeros from formatted float string
func trimTrailing(s string) string {
	if !strings.ContainsRune(s, '.') {
		return s
	}
	s = strings.TrimRight(s, "0")
	if s == "" {
		return "0"
	}
	if s[len(s)-1] == '.' {
		s = strings.TrimSuffix(s, ".")
	}
	return s
}
