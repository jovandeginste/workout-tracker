package humanize

import (
	"io/fs"

	"golang.org/x/text/language"

	"github.com/vorlif/spreak"
)

const (
	djangoDomain = "django"
)

// Collection data structure which collects the translations and can create humanizers.
type Collection struct {
	bundle  *spreak.Bundle
	locales map[language.Tag]*FormatData
}

type options struct {
	bundleOptions []spreak.BundleOption
	locales       []*LocaleData
}

// A Humanizer represents a collection of functions for humanizing data structures for a chosen language.
type Humanizer struct {
	loc    *spreak.Localizer
	format *FormatData
}

// LocaleData represents a collection of information and data about a language.
type LocaleData struct {
	Lang   language.Tag
	Fs     fs.FS
	Format *FormatData
}

// FormatData represents a collection of formatting rules that belongs to a language.
// For example, the formatting of a date for a language.
//
// It is automatically provided when using a language from the locale package.
type FormatData struct {
	// DateFormat is the formatting to use for displaying dates
	// Fallback: 'N j, Y' (e.g. Feb. 4, 2003)
	DateFormat string
	// TimeFormat is the formatting to use for displaying time
	// Fallback: 'P' (e.g. 4 p.m.)
	TimeFormat string
	// DateTimeFormat is the formatting to use for displaying datetime
	// Fallback: 'N j, Y, P' (e.g. Feb. 4, 2003, 4 p.m.)
	DateTimeFormat string
	// YearMonthFormat is suitable for cases when only the year and month should be displayed.
	// Fallback: 'F Y'
	YearMonthFormat string
	// MonthDayFormat is suitable for cases when only the month and day should be displayed.
	// Fallback 'F j'
	MonthDayFormat string
	// Fallback: 'm/d/Y' (e.g. 12/31/2003)
	ShortDateFormat string
	// Fallback: 'm/d/Y P' (e.g. 12/31/2003 4 p.m.)
	ShortDatetimeFormat string
	// FirstDayOfWeek defines the day of the week on which the week starts.
	// e.d 0 = Sunday, 1 = Monday, etc...
	FirstDayOfWeek int
}

var fallbackFormat = &FormatData{
	DateFormat:          "N j, Y",
	TimeFormat:          "P",
	DateTimeFormat:      "N j, Y, P",
	YearMonthFormat:     "F Y",
	MonthDayFormat:      "F j",
	ShortDateFormat:     "m/d/Y",
	ShortDatetimeFormat: "m/d/Y P",
	FirstDayOfWeek:      0,
}

// An Option that can be used to customize the configuration when creating a Collection.
type Option func(opts *options) error

// WithLocale specifies which languages to support for humanization.
func WithLocale(data ...*LocaleData) Option {
	return func(opts *options) error {
		opts.locales = append(opts.locales, data...)
		return nil
	}
}

// WithBundleOption Allows to store custom options for the internally created spreak.Bundle.
func WithBundleOption(opt spreak.BundleOption) Option {
	return func(opts *options) error {
		opts.bundleOptions = append(opts.bundleOptions, opt)
		return nil
	}
}

// New creates a new Collection which holds the humanizers for the selected locales.
func New(opts ...Option) (*Collection, error) {
	o := &options{
		bundleOptions: nil,
		locales:       nil,
	}

	for _, opt := range opts {
		if err := opt(o); err != nil {
			return nil, err
		}
	}

	for _, d := range o.locales {
		if d.Format == nil {
			d.Format = fallbackFormat
		}
		d.Format.setDefaults()
	}

	loader := newLoader(o.locales)

	coll := &Collection{
		locales: make(map[language.Tag]*FormatData, len(o.locales)),
	}
	languages := make([]interface{}, 0, len(loader.locales))
	for tag, data := range loader.locales {
		languages = append(languages, tag)
		if data.Format != nil {
			coll.locales[tag] = data.Format
		}
	}

	o.bundleOptions = append(o.bundleOptions,
		spreak.WithSourceLanguage(language.English),
		spreak.WithDefaultDomain(djangoDomain),
		spreak.WithDomainLoader(djangoDomain, loader),
		spreak.WithLanguage(languages...),
	)

	bundle, err := spreak.NewBundle(o.bundleOptions...)
	if err != nil {
		return nil, err
	}

	coll.bundle = bundle
	return coll, nil
}

// MustNew is similar to New except it panics if an error happens.
func MustNew(opts ...Option) *Collection {
	collection, err := New(opts...)
	if err != nil {
		panic(err)
	}
	return collection
}

// CreateHumanizer creates a new humanizer.
// Multiple languages can be passed and a spreak.Localizer is created which decides which language is used.
func (p *Collection) CreateHumanizer(lang ...interface{}) *Humanizer {
	loc := spreak.NewLocalizer(p.bundle, lang...)

	if data, ok := p.locales[loc.Language()]; ok {
		return &Humanizer{loc: loc, format: data}
	}

	return &Humanizer{loc: loc, format: fallbackFormat}
}

func (f *FormatData) setDefaults() {
	if f.DateFormat == "" {
		f.DateFormat = fallbackFormat.DateFormat
	}
	if f.TimeFormat == "" {
		f.TimeFormat = fallbackFormat.TimeFormat
	}
	if f.DateTimeFormat == "" {
		f.DateTimeFormat = fallbackFormat.DateTimeFormat
	}
	if f.YearMonthFormat == "" {
		f.YearMonthFormat = fallbackFormat.YearMonthFormat
	}
	if f.MonthDayFormat == "" {
		f.MonthDayFormat = fallbackFormat.MonthDayFormat
	}
	if f.ShortDateFormat == "" {
		f.ShortDateFormat = fallbackFormat.ShortDateFormat
	}
	if f.ShortDatetimeFormat == "" {
		f.ShortDatetimeFormat = fallbackFormat.ShortDatetimeFormat
	}
}
