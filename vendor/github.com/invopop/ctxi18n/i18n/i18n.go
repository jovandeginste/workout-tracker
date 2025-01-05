// Package i18n is responsible for keeping the key internationalization in one
// place.
package i18n

import (
	"context"
	"fmt"
	"strings"
)

const (
	missingLocaleOut = "!(MISSING LOCALE)"
)

type scopeType string

const (
	scopeKey scopeType = "scope"
)

// M stands for map and is a simple helper to make it easier to work with
// internationalization maps.
type M map[string]any

// T is responsible for translating a key into a string by extracting
// the local from the context.
func T(ctx context.Context, key string, args ...any) string {
	l := GetLocale(ctx)
	if l == nil {
		return missingLocaleOut
	}
	key = ExpandKey(ctx, key)
	return l.T(key, args...)
}

// N returns the pluralized translation of the provided key using n
// as the count.
func N(ctx context.Context, key string, n int, args ...any) string {
	l := GetLocale(ctx)
	if l == nil {
		return missingLocaleOut
	}
	key = ExpandKey(ctx, key)
	return l.N(key, n, args...)
}

// Has performs a check to see if the key exists in the locale.
func Has(ctx context.Context, key string) bool {
	l := GetLocale(ctx)
	if l == nil {
		return false
	}
	key = ExpandKey(ctx, key)
	return l.Has(key)
}

// WithScope is used to add a new scope to the context. To use this,
// use a `.` at the beginning of keys.
func WithScope(ctx context.Context, key string) context.Context {
	key = ExpandKey(ctx, key)
	return context.WithValue(ctx, scopeKey, key)
}

// ExpandKey extracts the current scope from the context and appends it
// to the start of the provided key.
func ExpandKey(ctx context.Context, key string) string {
	if !strings.HasPrefix(key, ".") {
		return key
	}
	scope, ok := ctx.Value(scopeKey).(string)
	if !ok {
		return key
	}
	return fmt.Sprintf("%s%s", scope, key)
}

// Replace is used to interpolate the matched keys in the provided
// string with their values in the map.
//
// Interpolation is performed using the `%{key}` pattern.
func (m M) Replace(in string) string {
	for k, v := range m {
		in = strings.Replace(in, fmt.Sprintf("%%{%s}", k), fmt.Sprint(v), -1)
	}
	return in
}
