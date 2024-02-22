package spreak

import (
	"errors"

	"golang.org/x/text/language"

	"github.com/vorlif/spreak/localize"
)

// A Localizer holds the catalogs of all domains for a language and provides an interface for their use.
// It has a default domain, which can differ from the bundle default domain and is used
// if no domain is specified for translations.
// A number of supported languages can be specified at creation time,
// where the language matcher of the bundle decides which language fits best.
// For this language the Localizer then offers the possibility to translate.
// If no language fits, the fallback language is used.
// If no fallback language is specified, the source language is used.
// For web applications, a Localizer can be created for each request, which can be disposed of at the end of the request.
type Localizer struct {
	bundle        *Bundle
	locale        *locale
	defaultDomain string
	localeFound   bool
}

// A KeyLocalizer is a wrapper for a Localizer, which can be used when looking up translations by key and not by source language.
// It provides the same methods as a Localizer, but does not require the specification of a plural text.
// The usage is useful for example when the translations are loaded from JSON files with a key-value structure.
type KeyLocalizer struct {
	*Localizer
}

// NewLocalizerForDomain creates a new Localizer for a language and a default domain,
// which is used if no domain is specified.
// Multiple languages can be passed and the best matching language is searched for.
// If no matching language is found, a Localizer is created which returns the original messages.
// Valid languages are strings or language.Tag. All other inputs are dropped.
func NewLocalizerForDomain(bundle *Bundle, domain string, lang ...interface{}) *Localizer {
	var tags []language.Tag

	for _, accept := range lang {
		switch val := accept.(type) {
		case language.Tag:
			tags = append(tags, val)
		case string:
			desired, _, err := language.ParseAcceptLanguage(val)
			if err != nil {
				continue
			}
			tags = append(tags, desired...)
		default:
			continue
		}
	}

	if _, index, conf := bundle.languageMatcher.Match(tags...); conf > language.No {
		tag := bundle.languages[index]
		return newLocalizerFromTag(bundle, domain, tag)
	}

	return newLocalizerFromTag(bundle, domain)
}

// NewLocalizer operates like NewLocalizerForDomain, with the default domain of the bundle as domain.
func NewLocalizer(bundle *Bundle, lang ...interface{}) *Localizer {
	return NewLocalizerForDomain(bundle, bundle.defaultDomain, lang...)
}

// NewKeyLocalizerForDomain creates a new Localizer with NewLocalizerForDomain which is then wrapped by a KeyLocalizer.
func NewKeyLocalizerForDomain(bundle *Bundle, domain string, lang ...interface{}) *KeyLocalizer {
	return &KeyLocalizer{Localizer: NewLocalizerForDomain(bundle, domain, lang...)}
}

// NewKeyLocalizer creates a new Localizer with NewLocalizer which is then wrapped by a KeyLocalizer.
func NewKeyLocalizer(bundle *Bundle, langs ...interface{}) *KeyLocalizer {
	return &KeyLocalizer{Localizer: NewLocalizer(bundle, langs...)}
}

func newLocalizerFromTag(bundle *Bundle, domain string, tag ...language.Tag) *Localizer {
	l := &Localizer{
		bundle:        bundle,
		defaultDomain: domain,
	}

	for _, accept := range tag {
		if acceptLocale, hasLocale := bundle.locales[accept]; hasLocale {
			l.locale = acceptLocale
			l.localeFound = true
			break
		}
	}

	if l.locale == nil {
		l.locale = bundle.fallbackLocale
	}

	return l
}

// HasDomain checks whether a catalog has been loaded for a specified domain.
func (l *Localizer) HasDomain(domain string) bool {
	_, hasDomain := l.locale.domainCatalogs[domain]
	return hasDomain
}

// Domains returns a list of all domains for which a catalog was found.
func (l *Localizer) Domains() []string {
	domains := make([]string, 0, len(l.locale.domainCatalogs))
	for domain := range l.locale.domainCatalogs {
		domains = append(domains, domain)
	}
	return domains
}

// HasLocale returns whether a matching locale has been found and message translation can take place.
func (l *Localizer) HasLocale() bool { return l.localeFound }

// DefaultDomain returns the default domain.
// The default domain is used if a domain is not explicitly specified for a requested translation.
// If no default domain is specified, the default domain of the bundle is used.
func (l *Localizer) DefaultDomain() string { return l.defaultDomain }

// Language returns the language into which the translation of messages is performed.
// If no language is present, language.Und is returned.
func (l *Localizer) Language() language.Tag {
	return l.locale.language
}

// The Get function return the localized translation of message, based on the used locale
// current default domain and language of the locale.
// The message argument identifies the message to be translated.
// If no suitable translation exists and a fallback language has been provided, the text of this language will be returned.
// If no fallback is provided or no translation exists for the fallback language, the source message is returned.
func (l *Localizer) Get(message localize.Singular) string {
	t, _ := l.lookupSingularTranslation(l.defaultDomain, NoCtx, message)
	return t
}

// Getf operates like Get, but formats the message according to a format identifier and returns the resulting string.
func (l *Localizer) Getf(message localize.Singular, vars ...interface{}) string {
	t, _ := l.lookupSingularTranslation(l.defaultDomain, NoCtx, message, vars...)
	return t
}

// DGet operates like Get, but look the message up in the specified domain.
func (l *Localizer) DGet(domain localize.Domain, message localize.Singular) string {
	t, _ := l.lookupSingularTranslation(domain, NoCtx, message)
	return t
}

// DGetf operates like Get, but look the message up in the specified domain and
// formats the message according to a format identifier and returns the resulting string.
func (l *Localizer) DGetf(domain localize.Domain, message localize.Singular, vars ...interface{}) string {
	t, _ := l.lookupSingularTranslation(domain, NoCtx, message, vars...)
	return t
}

// NGet acts like Get, but consider plural forms.
// The plural formula is applied to n and return the resulting message (some languages have more than two plurals).
// If n is a floating-point number and the CLDR rules are used, the floating-point number should be represented and passed as a string for the best result.
func (l *Localizer) NGet(singular localize.Singular, plural localize.Plural, n interface{}) string {
	t, _ := l.lookupPluralTranslation(l.defaultDomain, NoCtx, singular, plural, n)
	return t
}

// NGetf operates like NGet, but formats the message according to a format identifier and returns the resulting string.
func (l *Localizer) NGetf(singular localize.Singular, plural localize.Plural, n interface{}, vars ...interface{}) string {
	t, _ := l.lookupPluralTranslation(l.defaultDomain, NoCtx, singular, plural, n, vars...)
	return t
}

// DNGet operates like NGet, but look the message up in the specified domain.
func (l *Localizer) DNGet(domain localize.Domain, singular localize.Singular, plural localize.Plural, n interface{}) string {
	t, _ := l.lookupPluralTranslation(domain, NoCtx, singular, plural, n)
	return t
}

// DNGetf operates like DNGet, but formats the message according to a format identifier and returns the resulting string.
func (l *Localizer) DNGetf(domain localize.Domain, singular localize.Singular, plural localize.Plural, n interface{}, vars ...interface{}) string {
	t, _ := l.lookupPluralTranslation(domain, NoCtx, singular, plural, n, vars...)
	return t
}

// PGet operates like Get, but restricted to the specified context.
func (l *Localizer) PGet(context localize.Context, message localize.Singular) string {
	t, _ := l.lookupSingularTranslation(l.defaultDomain, context, message)
	return t
}

// PGetf operates like PGet, but formats the message according to a format identifier and returns the resulting string.
func (l *Localizer) PGetf(context localize.Context, message localize.Singular, vars ...interface{}) string {
	t, _ := l.lookupSingularTranslation(l.defaultDomain, context, message, vars...)
	return t
}

// DPGet operates like Get, but look the message up in the specified domain and with the specified context.
func (l *Localizer) DPGet(domain localize.Domain, context localize.Context, message localize.Singular) string {
	t, _ := l.lookupSingularTranslation(domain, context, message)
	return t
}

// DPGetf operates like DPGet, but formats the message according to a format identifier and returns the resulting string.
func (l *Localizer) DPGetf(domain localize.Domain, context localize.Context, message localize.Singular, vars ...interface{}) string {
	t, _ := l.lookupSingularTranslation(domain, context, message, vars...)
	return t
}

// NPGet operates like NGet, but restricted to the specified context.
func (l *Localizer) NPGet(context localize.Context, singular localize.Singular, plural localize.Plural, n interface{}) string {
	t, _ := l.lookupPluralTranslation(l.defaultDomain, context, singular, plural, n)
	return t
}

// NPGetf operates like NPGet, but formats the message according to a format identifier and returns the resulting string.
func (l *Localizer) NPGetf(context localize.Context, singular localize.Singular, plural localize.Plural, n interface{}, vars ...interface{}) string {
	t, _ := l.lookupPluralTranslation(l.defaultDomain, context, singular, plural, n, vars...)
	return t
}

// DNPGet operates like NGet, but look the message up in the specified domain and with the specified context.
func (l *Localizer) DNPGet(domain localize.Domain, context localize.Context, singular localize.Singular, plural localize.Plural, n interface{}) string {
	t, _ := l.lookupPluralTranslation(domain, context, singular, plural, n)
	return t
}

// DNPGetf operates like DNPGet, but formats the message according to a format identifier and returns the resulting string.
func (l *Localizer) DNPGetf(domain localize.Domain, context localize.Context, singular localize.Singular, plural localize.Plural, n interface{}, vars ...interface{}) string {
	t, _ := l.lookupPluralTranslation(domain, context, singular, plural, n, vars...)
	return t
}

// LocalizeWithError translates structs that implement the interface localize.Localizable.
// If a suitable translation is found, it will be returned.
// If no matching translation is found, the original string with the matching plural form and an error are returned.
func (l *Localizer) LocalizeWithError(t localize.Localizable) (string, error) {
	if t == nil {
		return "<nil>", errors.New("spreak: Localizable is nil")
	}

	var vars []interface{}
	if len(t.GetVars()) > 0 {
		vars = append(vars, t.GetVars()...)
	}

	domain := l.defaultDomain
	if t.HasDomain() {
		domain = t.GetDomain()
	}

	if t.GetPluralID() != "" {
		return l.lookupPluralTranslation(domain, t.GetContext(), t.GetMsgID(), t.GetPluralID(), t.GetCount(), vars...)
	}

	return l.lookupSingularTranslation(domain, t.GetContext(), t.GetMsgID(), vars...)
}

// Localize acts like LocalizeWithError, but does not return an error.
func (l *Localizer) Localize(t localize.Localizable) string {
	translated, _ := l.LocalizeWithError(t)
	return translated
}

// LocalizeError translates the passed error and returns a new error of type localize.Error
// which wraps the original error.
// If no suitable translation is found, the original error is returned.
// By default, localized messages with the context "errors" are searched for.
// The query is limited to the current domain and the error context specified in the corresponding bundle.
// By default, this is the context "errors".
// Using WithErrorContext("other") during bundle creation to change the error context for a bundle.
func (l *Localizer) LocalizeError(err error) error {
	if err == nil {
		return nil
	}

	switch v := err.(type) {
	case localize.Localizable:
		translation, errT := l.LocalizeWithError(v)
		if errT != nil {
			return err
		}

		return &localize.Error{Translation: translation, Wrapped: err}
	default:
		translation, errT := l.lookupSingularTranslation(l.defaultDomain, l.bundle.errContext, err.Error())
		if errT != nil {
			return err
		}

		return &localize.Error{Translation: translation, Wrapped: err}
	}
}

func (l *Localizer) Print(format string, vars ...interface{}) string {
	return l.locale.printFunc(format, vars...)
}

func (l *Localizer) lookupSingularTranslation(domain localize.Domain, ctx localize.Context, msgID localize.Singular, vars ...interface{}) (string, error) {
	t, err := l.locale.lookupSingularTranslation(domain, ctx, msgID, vars...)
	if err == nil {
		if !l.localeFound {
			return t, errMissingLocale
		}
		return t, nil
	}

	errA := err
	if l.locale.language != l.bundle.fallbackLocale.language {
		t, err = l.bundle.fallbackLocale.lookupSingularTranslation(domain, ctx, msgID, vars...)
		if err == nil {
			return t, errA
		}
	}

	// The source locale always returns a text and never an error
	t, _ = l.bundle.sourceLocale.lookupSingularTranslation(domain, ctx, msgID, vars...)
	return t, errA
}

func (l *Localizer) lookupPluralTranslation(domain string, ctx localize.Context, singular localize.Singular, plural localize.Plural, n interface{}, vars ...interface{}) (string, error) {
	t, err := l.locale.lookupPluralTranslation(domain, ctx, singular, plural, n, vars...)
	if err == nil {
		if !l.localeFound {
			return t, errMissingLocale
		}
		return t, nil
	}

	errA := err
	t, err = l.bundle.fallbackLocale.lookupPluralTranslation(domain, ctx, singular, plural, n, vars...)
	if err == nil {
		return t, errA
	}

	t, _ = l.bundle.sourceLocale.lookupPluralTranslation(domain, ctx, singular, plural, n, vars...)
	return t, errA
}

func (l *KeyLocalizer) Get(key localize.Key) string {
	t, _ := l.lookupSingularTranslation(l.defaultDomain, NoCtx, key)
	return t
}

// Getf operates like Get, but formats the message according to a format identifier and returns the resulting string.
func (l *KeyLocalizer) Getf(key localize.Key, vars ...interface{}) string {
	t, _ := l.lookupSingularTranslation(l.defaultDomain, NoCtx, key, vars...)
	return t
}

// DGet operates like Get, but look the message up in the specified domain.
func (l *KeyLocalizer) DGet(domain localize.Domain, key localize.Key) string {
	t, _ := l.lookupSingularTranslation(domain, NoCtx, key)
	return t
}

// DGetf operates like Get, but look the message up in the specified domain and
// formats the message according to a format identifier and returns the resulting string.
func (l *KeyLocalizer) DGetf(domain localize.Domain, key localize.Key, vars ...interface{}) string {
	t, _ := l.lookupSingularTranslation(domain, NoCtx, key, vars...)
	return t
}

// NGet acts like Get, but consider plural forms.
// The plural formula is applied to n and return the resulting message (some languages have more than two plurals).
func (l *KeyLocalizer) NGet(key localize.PluralKey, n interface{}) string {
	t, _ := l.lookupPluralTranslation(l.defaultDomain, NoCtx, key, key, n)
	return t
}

// NGetf operates like NGet, but formats the message according to a format identifier and returns the resulting string.
func (l *KeyLocalizer) NGetf(key localize.PluralKey, n interface{}, vars ...interface{}) string {
	t, _ := l.lookupPluralTranslation(l.defaultDomain, NoCtx, key, key, n, vars...)
	return t
}

// DNGet operates like NGet, but look the message up in the specified domain.
func (l *KeyLocalizer) DNGet(domain localize.Domain, key localize.PluralKey, n interface{}) string {
	t, _ := l.lookupPluralTranslation(domain, NoCtx, key, key, n)
	return t
}

// DNGetf operates like DNGet, but formats the message according to a format identifier and returns the resulting string.
func (l *KeyLocalizer) DNGetf(domain localize.Domain, key localize.PluralKey, n interface{}, vars ...interface{}) string {
	t, _ := l.lookupPluralTranslation(domain, NoCtx, key, key, n, vars...)
	return t
}

// PGet operates like Get, but restricted to the specified context.
func (l *KeyLocalizer) PGet(context localize.Context, key localize.Key) string {
	t, _ := l.lookupSingularTranslation(l.defaultDomain, context, key)
	return t
}

// PGetf operates like PGet, but formats the message according to a format identifier and returns the resulting string.
func (l *KeyLocalizer) PGetf(context localize.Context, key localize.Key, vars ...interface{}) string {
	t, _ := l.lookupSingularTranslation(l.defaultDomain, context, key, vars...)
	return t
}

// DPGet operates like Get, but look the message up in the specified domain and with the specified context.
func (l *KeyLocalizer) DPGet(domain localize.Domain, context localize.Context, key localize.Key) string {
	t, _ := l.lookupSingularTranslation(domain, context, key)
	return t
}

// DPGetf operates like DPGet, but formats the message according to a format identifier and returns the resulting string.
func (l *KeyLocalizer) DPGetf(domain localize.Domain, context localize.Context, key localize.Key, vars ...interface{}) string {
	t, _ := l.lookupSingularTranslation(domain, context, key, vars...)
	return t
}

// NPGet operates like NGet, but restricted to the specified context.
func (l *KeyLocalizer) NPGet(context localize.Context, key localize.PluralKey, n interface{}) string {
	t, _ := l.lookupPluralTranslation(l.defaultDomain, context, key, key, n)
	return t
}

// NPGetf operates like NPGet, but formats the message according to a format identifier and returns the resulting string.
func (l *KeyLocalizer) NPGetf(context localize.Context, key localize.PluralKey, n interface{}, vars ...interface{}) string {
	t, _ := l.lookupPluralTranslation(l.defaultDomain, context, key, key, n, vars...)
	return t
}

// DNPGet operates like NGet, but look the message up in the specified domain and with the specified context.
func (l *KeyLocalizer) DNPGet(domain localize.Domain, context localize.Context, key localize.PluralKey, n interface{}) string {
	t, _ := l.lookupPluralTranslation(domain, context, key, key, n)
	return t
}

// DNPGetf operates like DNPGet, but formats the message according to a format identifier and returns the resulting string.
func (l *KeyLocalizer) DNPGetf(domain localize.Domain, context localize.Context, key localize.PluralKey, n interface{}, vars ...interface{}) string {
	t, _ := l.lookupPluralTranslation(domain, context, key, key, n, vars...)
	return t
}
