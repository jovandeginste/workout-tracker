package i18n

import (
	"context"
	"encoding/json"
	"fmt"
)

// Locale holds the internationalization entries for a specific locale.
type Locale struct {
	code Code
	dict *Dict
	rule PluralRule
}

const (
	missingDictOut      = "!(MISSING: %s)"
	localeKey      Code = "locale"
)

// DefaultText when detected as an argument to a translation
// function will be used if no language match is found.
type DefaultText string

// Default when used as an argument to a translation function
// ensure the provided txt is used as a default value if no
// language match is found.
func Default(txt string) DefaultText {
	return DefaultText(txt)
}

// NewLocale creates a new locale with the provided key and dictionary.
func NewLocale(code Code, dict *Dict) *Locale {
	l := &Locale{
		code: code,
		dict: dict,
	}
	l.rule = mapPluralRule(code)
	return l

}

// Code returns the language code of the locale.
func (l *Locale) Code() Code {
	return l.code
}

// T provides the value from the dictionary stored by the locale.
func (l *Locale) T(key string, args ...any) string {
	return interpolate(key, l.dict.Get(key), args...)
}

// N uses the locale pluralization rules to determine which
// string value to provide based on the provided number.
func (l *Locale) N(key string, n int, args ...any) string {
	d := l.dict.Get(key)
	return interpolate(key, l.rule(d, n), args...)
}

// Has performs a check to see if the key exists in the locale.
// This is useful for checking if a key exists before attempting
// to use it when the Default function cannot be used.
func (l *Locale) Has(key string) bool {
	return l.dict.Has(key)
}

// PluralRule provides the pluralization rule for the locale.
func (l *Locale) PluralRule() PluralRule {
	return l.rule
}

// UnmarshalJSON attempts to load the locale from a JSON byte slice.
func (l *Locale) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	l.dict = new(Dict)
	if err := json.Unmarshal(data, l.dict); err != nil {
		return err
	}
	return nil
}

// WithContext inserts the locale into the context so that it can be
// loaded later with `GetLocale`.
func (l *Locale) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, localeKey, l)
}

// GetLocale retrieves the locale from the context.
func GetLocale(ctx context.Context) *Locale {
	if l, ok := ctx.Value(localeKey).(*Locale); ok {
		return l
	}
	return nil
}

func interpolate(key string, d *Dict, args ...any) string {
	var s string
	s, args = extractDefault(args)
	if d != nil {
		s = d.value
	}
	if s == "" {
		return missing(key)
	}
	if len(args) > 0 {
		switch arg := args[0].(type) {
		case M:
			return arg.Replace(s)
		default:
			return fmt.Sprintf(s, args...)
		}
	}
	return s
}

func extractDefault(args []any) (string, []any) {
	for i, arg := range args {
		if dt, ok := arg.(DefaultText); ok {
			return string(dt), append(args[:i], args[i+1:]...)
		}
	}
	return "", args
}

func missing(key string) string {
	return fmt.Sprintf(missingDictOut, key)
}
