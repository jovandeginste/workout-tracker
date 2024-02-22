package spreak

import (
	"errors"
	"fmt"
	"io/fs"

	"golang.org/x/text/language"

	"github.com/vorlif/spreak/catalog"
)

// NoDomain is the domain which is used if no default domain is stored.
const NoDomain = ""

// NoCtx is the context which is used if no context is stored.
const NoCtx = ""

// ErrorsCtx ist the context under which translations for extracted errors are searched.
// Can be changed when creating a bundle with WithErrorContext.
const ErrorsCtx = "errors"

// MissingTranslationCallback is a callback which can be stored with WithMissingTranslationCallback for a bundle.
// Called when translations, domains, or other are missing.
// The call is not goroutine safe.
type MissingTranslationCallback func(err error)

// LanguageMatcherBuilder is a builder which creates a language matcher.
// It is an abstraction of language.NewMatcher of the language package and should return the same values.
// Can be set when creating a bundle with WithLanguageMatcherBuilder for a bundle.
// The matcher is used, for example, when a new Localizer is created to determine the best matching language.
type LanguageMatcherBuilder func(t []language.Tag, options ...language.MatchOption) language.Matcher

// BundleOption is an option which can be passed when creating a bundle to customize its configuration.
type BundleOption func(opts *bundleBuilder) error

type setupAction func(options *bundleBuilder) error

type bundleBuilder struct {
	*Bundle

	sourceLanguage         language.Tag
	fallbackLanguage       language.Tag
	languageMatcherBuilder LanguageMatcherBuilder
	domainLoaders          map[string]Loader
	setupActions           []setupAction
}

// A Bundle is the central place to load and manage translations.
// It holds all catalogs for all domains and all languages.
// The bundle cannot be edited after creation and is goroutine safe.
// Typically, an application contains a bundle as a singleton.
// The catalog of the specified domains and languages will be loaded during the creation.
type Bundle struct {
	missingCallback MissingTranslationCallback
	printer         Printer
	defaultDomain   string
	errContext      string

	fallbackLocale  *locale
	sourceLocale    *locale
	languageMatcher language.Matcher

	languages []language.Tag
	locales   map[language.Tag]*locale
	domains   map[string]bool
}

// NewBundle creates a new bundle and returns it.
// An error is returned if something fails during creation.
// This is only the case if one of the options returns an error.
// A call without options will never return an error and can thus be used for testing or as a fallback.
// The catalog of the specified domains and languages will be loaded during the creation.
func NewBundle(opts ...BundleOption) (*Bundle, error) {
	builder := &bundleBuilder{
		Bundle: &Bundle{
			printer:       NewDefaultPrinter(),
			defaultDomain: NoDomain,
			errContext:    ErrorsCtx,

			languages: make([]language.Tag, 0),
			locales:   make(map[language.Tag]*locale),
			domains:   make(map[string]bool),
		},
		languageMatcherBuilder: language.NewMatcher,
		fallbackLanguage:       language.Und,
		sourceLanguage:         language.Und,

		domainLoaders: make(map[string]Loader),
		setupActions:  make([]setupAction, 0),
	}

	for _, opt := range opts {
		if opt == nil {
			return nil, errors.New("spreak.Bundle: try to create an bundle with a nil option")
		}
		if err := opt(builder); err != nil {
			return nil, err
		}
	}

	for domain := range builder.domainLoaders {
		builder.domains[domain] = false
	}

	for _, action := range builder.setupActions {
		if action == nil {
			return nil, errors.New("spreak.Bundle: try to create an bundle with a nil action")
		}
		if err := action(builder); err != nil {
			return nil, err
		}
	}

	builder.languageMatcher = builder.languageMatcherBuilder(builder.languages)
	builder.printer.Init(builder.languages)

	if sourceLocale, hasSource := builder.locales[builder.sourceLanguage]; hasSource {
		builder.Bundle.sourceLocale = sourceLocale
	} else {
		builder.Bundle.sourceLocale = buildSourceLocale(builder.Bundle, builder.sourceLanguage)
	}

	if fallbackLocale, hasFallback := builder.locales[builder.fallbackLanguage]; hasFallback {
		builder.Bundle.fallbackLocale = fallbackLocale
	} else {
		builder.Bundle.fallbackLocale = builder.Bundle.sourceLocale
	}

	return builder.Bundle, nil
}

// Domains returns a list of loaded domains.
// A domain is only loaded if at least one catalog is found in one language.
func (b *Bundle) Domains() []string {
	domains := make([]string, 0, len(b.domains))
	for domain, loaded := range b.domains {
		if loaded {
			domains = append(domains, domain)
		}
	}
	return domains
}

// CanLocalize indicates whether locales and domains have been loaded for translation.
func (b *Bundle) CanLocalize() bool {
	return len(b.locales) > 0 && len(b.Domains()) > 0
}

// SupportedLanguages returns all languages for which a catalog was found for at least one domain.
func (b *Bundle) SupportedLanguages() []language.Tag {
	languages := make([]language.Tag, 0, len(b.locales))
	for lang := range b.locales {
		languages = append(languages, lang)
	}
	return languages
}

// IsLanguageSupported indicates whether a language can be translated.
// The check is done by the bundle's matcher and therefore languages that are not returned by
// SupportedLanguages can be supported.
func (b *Bundle) IsLanguageSupported(lang language.Tag) bool {
	_, _, confidence := b.languageMatcher.Match(lang)
	return confidence > language.No
}

func (b *bundleBuilder) preloadLanguages(optional bool, languages ...interface{}) error {
	for _, accept := range languages {
		tag, errT := languageInterfaceToTag(accept)
		if errT != nil {
			return errT
		}

		_, err := b.createLocale(optional, tag)
		if err == nil {
			continue
		}

		if !optional {
			return err
		}

		var missErr *ErrMissingLanguage
		if errors.As(err, &missErr) {
			if b.missingCallback != nil {
				b.missingCallback(missErr)
			}

			continue
		}

		return err

	}

	return nil
}

func (b *bundleBuilder) createLocale(optional bool, lang language.Tag) (*locale, error) {
	if lang == language.Und {
		return nil, newMissingLanguageError(lang)
	}

	if lang == b.sourceLanguage {
		sourceLocale := buildSourceLocale(b.Bundle, b.sourceLanguage)
		b.locales[lang] = sourceLocale
		b.languages = append(b.languages, lang)
		return sourceLocale, nil
	}

	if cachedLocale, isCached := b.locales[lang]; isCached {
		return cachedLocale, nil
	}

	catalogs := make(map[string]catalog.Catalog, len(b.domainLoaders))

	for domain, domainLoader := range b.domainLoaders {
		catl, errD := domainLoader.Load(lang, domain)
		if errD != nil {
			var notFoundErr *ErrNotFound
			if errors.As(errD, &notFoundErr) {
				if b.missingCallback != nil {
					b.missingCallback(notFoundErr)
				}

				if optional {
					continue
				}
			}
			return nil, errD
		}

		catalogs[domain] = catl
		b.domains[domain] = true
	}

	if len(catalogs) == 0 {
		return nil, newMissingLanguageError(lang)
	}

	langLocale := buildLocale(b.Bundle, lang, catalogs)
	b.locales[lang] = langLocale
	b.languages = append(b.languages, lang)
	return langLocale, nil
}

// WithFallbackLanguage sets the fallback language to be used when creating Localizer if no suitable language is available.
// Should be used only if the fallback language is different from source language.
// Otherwise, it should not be set.
func WithFallbackLanguage(lang interface{}) BundleOption {
	return func(opts *bundleBuilder) error {
		tag, err := languageInterfaceToTag(lang)
		if err != nil {
			return err
		}

		opts.fallbackLanguage = tag
		opts.setupActions = append(opts.setupActions, func(builder *bundleBuilder) error {
			return builder.preloadLanguages(false, tag)
		})
		return nil
	}
}

// WithSourceLanguage sets the source language used for programming.
// If it is set, it will be considered as a matching language when creating a Localizer.
// Also, it will try to use the appropriate plural form and will not trigger any missing callbacks for the language.
// It is recommended to always set the source language.
func WithSourceLanguage(tag language.Tag) BundleOption {
	return func(opts *bundleBuilder) error {
		opts.sourceLanguage = tag
		opts.setupActions = append(opts.setupActions, func(builder *bundleBuilder) error {
			return builder.preloadLanguages(false, tag)
		})
		return nil
	}
}

// WithMissingTranslationCallback stores a MissingTranslationCallback which is called when a translation,
// domain or something else is missing.
// The call is not goroutine safe.
func WithMissingTranslationCallback(cb MissingTranslationCallback) BundleOption {
	return func(opts *bundleBuilder) error {
		opts.missingCallback = cb
		return nil
	}
}

// WithDefaultDomain sets the default domain which will be used if no domain is specified.
// By default, NoDomain (the empty string) is used.
func WithDefaultDomain(domain string) BundleOption {
	return func(opts *bundleBuilder) error {
		opts.defaultDomain = domain
		return nil
	}
}

// WithDomainLoader loads a domain via a specified loader.
func WithDomainLoader(domain string, l Loader) BundleOption {
	return func(opts *bundleBuilder) error {
		if _, found := opts.domainLoaders[domain]; found {
			return fmt.Errorf("spreak.Bundle: loader for domain %s already set", domain)
		}
		if l == nil {
			return errors.New("spreak.Bundle: loader of WithDomainLoader(..., loader) is nil")
		}
		opts.domainLoaders[domain] = l
		return nil
	}
}

// WithFilesystemLoader Loads a domain via a FilesystemLoader.
// The loader can be customized with options.
func WithFilesystemLoader(domain string, fsOpts ...FsOption) BundleOption {
	return func(opts *bundleBuilder) error {
		l, err := NewFilesystemLoader(fsOpts...)
		if err != nil {
			return err
		}

		if _, found := opts.domainLoaders[domain]; found {
			return fmt.Errorf("spreak: loader for domain %s already set", domain)
		}

		opts.domainLoaders[domain] = l
		return nil
	}
}

// WithDomainPath loads a domain from a specified path.
//
// This is a shorthand for WithFilesystemLoader(domain, WithPath(path)).
func WithDomainPath(domain string, path string) BundleOption {
	return WithFilesystemLoader(domain, WithPath(path))
}

// WithDomainFs loads a domain from a fs.FS.
//
// This is a shorthand for WithFilesystemLoader(domain, WithFs(fsys)).
func WithDomainFs(domain string, fsys fs.FS) BundleOption {
	if fsys == nil {
		return func(opts *bundleBuilder) error {
			return errors.New("spreak.Bundle: fsys of WithDomainFs(..., fsys) is nil")
		}
	}

	return WithFilesystemLoader(domain, WithFs(fsys))
}

// WithLanguage loads the catalogs of the domains for one or more languages.
// The passed languages must be of type string or language.Tag,
// all other values will abort the initialization of the bundle with an error.
// If a catalog file for a domain is not found for a language, it will be ignored.
// If a catlaog file for a domain is found but cannot be loaded, the bundle creation will fail and return errors.
//
// If you want to use a Localizer, you should pay attention to the order in which the languages are specified,
// otherwise unexpected behavior may occur.
// This is because the matching algorithm of the language.matcher can give unexpected results.
// See https://github.com/golang/go/issues/49176
func WithLanguage(languages ...interface{}) BundleOption {
	loadFunc := func(builder *bundleBuilder) error {
		return builder.preloadLanguages(true, languages...)
	}

	return func(opts *bundleBuilder) error {
		opts.setupActions = append(opts.setupActions, loadFunc)
		return nil
	}
}

// WithRequiredLanguage works like WithLanguage except that the creation of the bundle fails
// if a catalog for a language could not be found.
func WithRequiredLanguage(languages ...interface{}) BundleOption {
	loadFunc := func(builder *bundleBuilder) error {
		return builder.preloadLanguages(false, languages...)
	}

	return func(opts *bundleBuilder) error {
		opts.setupActions = append(opts.setupActions, loadFunc)
		return nil
	}
}

// WithPrinter sets a printer which creates a function for a language which converts a formatted string
// and variables into a string. (Like fmt.Sprintf).
func WithPrinter(p Printer) BundleOption {
	return func(opts *bundleBuilder) error {
		if p == nil {
			return errors.New("spreak.Bundle: printer of WithPrinter(...) is nil")
		}
		opts.printer = p
		return nil
	}
}

// WithPrintFunction sets a PrintFunc which converts a formatted string and variables to a string. (Like fmt.Sprintf).
func WithPrintFunction(printFunc PrintFunc) BundleOption {
	if printFunc != nil {
		printer := &printFunctionWrapper{f: printFunc}
		return WithPrinter(printer)
	}

	return func(opts *bundleBuilder) error {
		return errors.New("spreak.Bundle: parameter of WithPrintFunction(...) is nil")
	}
}

// WithLanguageMatcherBuilder sets a LanguageMatcherBuilder.
func WithLanguageMatcherBuilder(mc LanguageMatcherBuilder) BundleOption {
	return func(opts *bundleBuilder) error {
		if mc == nil {
			return errors.New("spreak.Bundle: MatchCreator of WithMatchCreator(...) is nil")
		}
		opts.languageMatcherBuilder = mc
		return nil
	}
}

// WithErrorContext set a context, which is used for the translation of errors.
// If no context is set, ErrorsCtx is used.
func WithErrorContext(ctx string) BundleOption {
	return func(opts *bundleBuilder) error {
		opts.errContext = ctx
		return nil
	}
}
