package spreak

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/vorlif/spreak/catalog"
)

var (
	ErrRequireStringTag = errors.New("spreak: unsupported type, expecting string or language.Tag")

	errMissingLocale = errors.New("spreak: locale missing")
)

// Deprecated: Moved to catalog.NewPoDecoder and will be removed with v1.0.
func NewPoDecoder() catalog.Decoder { return catalog.NewPoDecoder() }

// Deprecated: Moved to catalog.NewMoDecoder and will be removed with v1.0.
func NewMoDecoder() catalog.Decoder { return catalog.NewMoDecoder() }

// Deprecated: Moved to catalog.ErrMissingContext and will be removed with v1.0.
type ErrMissingContext = catalog.ErrMissingContext

// Deprecated: Moved to catalog.ErrMissingMessageID and will be removed with v1.0.
type ErrMissingMessageID = catalog.ErrMissingMessageID

// Deprecated: Moved to catalog.ErrMissingTranslation and will be removed with v1.0.
type ErrMissingTranslation = catalog.ErrMissingTranslation

// PrintFunc formats according to a format specifier and returns the resulting string.
// Like fmt.Sprintf(...)
type PrintFunc func(str string, vars ...interface{}) string

// A Printer creates a PrintFunc for a language.
// Can be stored with WithPrinter when creating a bundle.
type Printer interface {
	Init(languages []language.Tag)
	GetPrintFunc(lang language.Tag) PrintFunc
}

type defaultPrinter struct {
	printers map[language.Tag]PrintFunc
}

// NewDefaultPrinter creates a printer which will be used if no printer was defined
// with WithPrinter when creating a bundle.
func NewDefaultPrinter() Printer {
	return &defaultPrinter{}
}

func (d *defaultPrinter) Init(languages []language.Tag) {
	d.printers = make(map[language.Tag]PrintFunc, len(languages))
	for _, lang := range languages {
		d.printers[lang] = defaultPrintFunc(lang)
	}
}

func (d *defaultPrinter) GetPrintFunc(lang language.Tag) PrintFunc {
	if printFunc, ok := d.printers[lang]; ok {
		return printFunc
	}

	return defaultPrintFunc(lang)
}

func defaultPrintFunc(lang language.Tag) PrintFunc {
	printer := message.NewPrinter(lang)
	return func(str string, vars ...interface{}) string {
		if len(vars) > 0 {
			return printer.Sprintf(str, vars...)
		}

		return str
	}
}

// Simple wrapper to use a PrinterFunction as a printer.
type printFunctionWrapper struct {
	f PrintFunc
}

var _ Printer = (*printFunctionWrapper)(nil)

func (p *printFunctionWrapper) Init(_ []language.Tag)                 {}
func (p *printFunctionWrapper) GetPrintFunc(_ language.Tag) PrintFunc { return p.f }

// ExpandLanguage returns possible filenames for a language tag without extension.
func ExpandLanguage(lang language.Tag) []string {
	expansions := make(map[string]bool, 4)
	expansions[lang.String()] = true

	base, baseConf := lang.Base()
	if baseConf > language.No {
		expansions[base.ISO3()] = true
		expansions[base.String()] = true
	}

	region, regionConf := lang.Region()
	if regionConf > language.No && baseConf > language.No {
		key := fmt.Sprintf("%s_%s", base.String(), region.String())
		expansions[key] = true

		key = fmt.Sprintf("%s-%s", base.String(), region.String())
		expansions[key] = true
	}

	script, scriptConf := lang.Script()
	if scriptConf > language.No && baseConf > language.No {
		key := fmt.Sprintf("%s_%s", base.String(), script.String())
		expansions[key] = true

		key = fmt.Sprintf("%s-%s", base.String(), script.String())
		expansions[key] = true
	}

	return stringMapKeys(expansions)
}

func parseLanguageName(lang string) (language.Tag, error) {
	if idx := strings.Index(lang, ":"); idx != -1 {
		lang = lang[:idx]
	}
	if idx := strings.Index(lang, "@"); idx != -1 {
		lang = lang[:idx]
	}
	if idx := strings.Index(lang, "."); idx != -1 {
		lang = lang[:idx]
	}
	lang = strings.TrimSpace(lang)
	return language.Parse(lang)
}

func languageInterfaceToTag(i interface{}) (language.Tag, error) {
	switch v := i.(type) {
	case string:
		tag, err := parseLanguageName(v)
		if err != nil {
			return language.Und, err
		}
		return tag, nil
	case language.Tag:
		return v, nil
	default:
		return language.Und, ErrRequireStringTag
	}
}

func stringMapKeys(m map[string]bool) []string {
	keys := make([]string, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	// Longest first
	sort.SliceStable(keys, func(i, j int) bool {
		if x, y := len(keys[i]), len(keys[j]); x != y {
			return x > y
		}

		return keys[i] < keys[j]
	})
	return keys
}

// ErrNotFound is the error returned by a loader if no matching context was found.
// If a loader returns any other error, the bundle creation will abort.
type ErrNotFound struct {
	Language   language.Tag
	Type       string
	Identifier string
}

func NewErrNotFound(lang language.Tag, source string, format string, vars ...interface{}) *ErrNotFound {
	return &ErrNotFound{
		Language:   lang,
		Type:       source,
		Identifier: fmt.Sprintf(format, vars...),
	}
}

func (e *ErrNotFound) Error() string { return e.String() }

func (e *ErrNotFound) String() string {
	return fmt.Sprintf("spreak: item of type %q for lang=%q could not be found: %s ", e.Type, e.Language, e.Identifier)
}

// ErrMissingLanguage is the error returned when a locale should be created and the matching language is not
// loaded or has no Catalog.
type ErrMissingLanguage struct {
	Language language.Tag
}

func newMissingLanguageError(lang language.Tag) *ErrMissingLanguage {
	return &ErrMissingLanguage{Language: lang}
}

func (e *ErrMissingLanguage) Error() string { return e.String() }

func (e *ErrMissingLanguage) String() string {
	return fmt.Sprintf("spreak: language not found: lang=%q ", e.Language)
}

// ErrMissingDomain is the error returned when a domain does not exist for a language.
type ErrMissingDomain struct {
	Language language.Tag
	Domain   string
}

func (e *ErrMissingDomain) Error() string { return e.String() }

func (e *ErrMissingDomain) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("spreak: domain not found: lang=%q ", e.Language))
	if e.Domain != "" {
		b.WriteString(fmt.Sprintf("domain=%q ", e.Domain))
	} else {
		b.WriteString("domain='' (empty string) ")
	}
	return b.String()
}
