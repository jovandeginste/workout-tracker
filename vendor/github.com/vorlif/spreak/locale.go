package spreak

import (
	"golang.org/x/text/language"

	"github.com/vorlif/spreak/catalog"
	"github.com/vorlif/spreak/catalog/poplural"
	"github.com/vorlif/spreak/localize"
)

// Deprecated: Will be removed with v1.0. A Localizer should be used, as it offers the same functionalities.
type Locale = Localizer

// NewLocale creates a new locale for a language and the default domain of the bundle.
// If a locale is not found, an error is returned.
// Deprecated: Will be removed with v1.0. Use NewLocalizer instead.
func NewLocale(bundle *Bundle, lang language.Tag) (*Locale, error) {
	return NewLocaleWithDomain(bundle, lang, bundle.defaultDomain)
}

// NewLocaleWithDomain creates a new locale for a language and a default domain.
// If no locale is found, an error is returned.
// Deprecated: Will be removed with v1.0. Use NewLocalizerForDomain instead.
func NewLocaleWithDomain(bundle *Bundle, lang language.Tag, domain string) (*Locale, error) {
	l := NewLocalizerForDomain(bundle, domain, lang)
	if !l.HasLocale() {
		return nil, newMissingLanguageError(lang)
	}
	return l, nil
}

type locale struct {
	bundle           *Bundle
	language         language.Tag
	domainCatalogs   map[string]catalog.Catalog
	printFunc        PrintFunc
	pluralFunc       poplural.PluralFunc
	isSourceLanguage bool
}

func buildLocale(bundle *Bundle, lang language.Tag, catalogs map[string]catalog.Catalog) *locale {
	l := &locale{
		bundle:         bundle,
		language:       lang,
		domainCatalogs: catalogs,
		printFunc:      bundle.printer.GetPrintFunc(lang),
	}

	l.pluralFunc, _ = poplural.ForLanguage(lang)

	return l
}

func buildSourceLocale(bundle *Bundle, sourceLang language.Tag) *locale {
	l := &locale{
		bundle:           bundle,
		language:         sourceLang,
		printFunc:        bundle.printer.GetPrintFunc(sourceLang),
		isSourceLanguage: true,
	}
	l.pluralFunc, _ = poplural.ForLanguage(sourceLang)
	return l
}

func (l *locale) lookupSingularTranslation(domain localize.Domain, ctx localize.Context, msgID localize.Singular, vars ...interface{}) (string, error) {
	if l.isSourceLanguage {
		return l.printFunc(msgID, vars...), nil
	}

	catl, err := l.getCatalog(domain)
	if err != nil {
		if l.bundle.missingCallback != nil {
			l.bundle.missingCallback(err)
		}

		return "", err
	}

	translation, errT := catl.GetTranslation(ctx, msgID)
	if errT != nil {
		if l.bundle.missingCallback != nil {
			l.bundle.missingCallback(errT)
		}

		return "", errT
	}

	return l.printFunc(translation, vars...), nil
}

func (l *locale) lookupPluralTranslation(domain string, ctx localize.Context, singular localize.Singular, plural localize.Plural, n interface{}, vars ...interface{}) (string, error) {
	if l.isSourceLanguage {
		return l.printSourceMessage(singular, plural, n, vars), nil
	}

	catl, err := l.getCatalog(domain)
	if err != nil {
		if l.bundle.missingCallback != nil {
			l.bundle.missingCallback(err)
		}

		return "", err
	}

	translation, errT := catl.GetPluralTranslation(ctx, singular, n)
	if errT != nil {
		if l.bundle.missingCallback != nil {
			l.bundle.missingCallback(errT)
		}

		return "", errT
	}

	return l.printFunc(translation, vars...), nil
}

func (l *locale) printSourceMessage(singular, plural string, n interface{}, vars []interface{}) string {
	idx := l.pluralFunc(n)
	if idx == 0 || plural == "" {
		return l.printFunc(singular, vars...)
	}

	return l.printFunc(plural, vars...)
}

func (l *locale) getCatalog(domain string) (catalog.Catalog, error) {
	if _, hasDomain := l.domainCatalogs[domain]; !hasDomain {
		err := &ErrMissingDomain{
			Language: l.language,
			Domain:   domain,
		}
		return nil, err
	}

	return l.domainCatalogs[domain], nil
}
