// Package ctxi18n is used to internationalize applications using the context
// for the locale.
package ctxi18n

import (
	"context"
	"errors"
	"io/fs"

	"github.com/invopop/ctxi18n/i18n"
)

var (
	// DefaultLocale defines the default or fallback locale code to use
	// if no other match inside the packages list was found.
	DefaultLocale i18n.Code = "en"
)

var (
	locales *i18n.Locales
)

var (
	// ErrMissingLocale implies that the requested locale was not found
	// in the current index.
	ErrMissingLocale = errors.New("locale not defined")
)

func init() {
	locales = new(i18n.Locales)
}

// Load walks through all the files in provided File System and prepares
// an internal global list of locales ready to use.
func Load(fs fs.FS) error {
	return locales.Load(fs)
}

// LoadWithDefault performs the regular load operation, but will merge
// the default locale with every other locale, ensuring that every text
// has at least the value from the default locale.
func LoadWithDefault(fs fs.FS, locale i18n.Code) error {
	return locales.LoadWithDefault(fs, locale)
}

// Get provides the Locale object for the matching code.
func Get(code i18n.Code) *i18n.Locale {
	return locales.Get(code)
}

// Match attempts to find the best possible matching locale based on the
// locale string provided. The locale string is parsed according to the
// "Accept-Language" header format defined in RFC9110.
func Match(locale string) *i18n.Locale {
	return locales.Match(locale)
}

// WithLocale tries to match the provided code with a locale and ensures
// it is available inside the context.
func WithLocale(ctx context.Context, locale string) (context.Context, error) {
	l := locales.Match(locale)
	if l == nil {
		l = locales.Get(DefaultLocale)
		if l == nil {
			return nil, ErrMissingLocale
		}
	}
	return l.WithContext(ctx), nil
}

// Locale provides the locale object currently stored in the context.
func Locale(ctx context.Context) *i18n.Locale {
	return i18n.GetLocale(ctx)
}
