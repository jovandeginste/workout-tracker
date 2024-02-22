package humanize

import (
	"io/fs"

	"golang.org/x/text/language"

	"github.com/vorlif/spreak"
	"github.com/vorlif/spreak/catalog"
)

type localeLoader struct {
	locales map[language.Tag]*LocaleData
}

func newLoader(locales []*LocaleData) *localeLoader {
	loader := &localeLoader{
		locales: make(map[language.Tag]*LocaleData, len(locales)),
	}
	for _, loc := range locales {
		loader.locales[loc.Lang] = loc
	}
	return loader
}

func (loader *localeLoader) Load(lang language.Tag, domain string) (catalog.Catalog, error) {
	data, hasData := loader.locales[lang]
	if !hasData {
		return nil, spreak.NewErrNotFound(lang, "humanize", "domain=%q", domain)
	}

	content, err := fs.ReadFile(data.Fs, domain+".po")
	if err != nil {
		return nil, spreak.NewErrNotFound(lang, "humanize", "domain=%q,err=%v", domain, err)
	}

	dec := catalog.NewPoDecoder()
	catl, err := dec.Decode(lang, domain, content)
	if err != nil {
		return nil, err
	}

	return catl, nil
}
