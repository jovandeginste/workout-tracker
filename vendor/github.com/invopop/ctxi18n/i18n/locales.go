package i18n

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/invopop/yaml"
)

// Locales is a map of language keys to their respective locale.
type Locales struct {
	list []*Locale
}

// Load walks through all the files in the provided File System
// and merges every one with the current list of locales.
func (ls *Locales) Load(src fs.FS) error {
	return fs.WalkDir(src, ".", func(path string, _ fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walking directory: %w", err)
		}

		switch filepath.Ext(path) {
		case ".yaml", ".yml", ".json":
			// good
		default:
			return nil
		}

		data, err := fs.ReadFile(src, path)
		if err != nil {
			return fmt.Errorf("reading file '%s': %w", path, err)
		}

		if err := yaml.Unmarshal(data, ls); err != nil {
			return fmt.Errorf("unmarshalling file '%s': %w", path, err)
		}

		return nil
	})
}

// LoadWithDefault performs the regular load operation, but follows up with
// a second operation that will ensure that default dictionary is merged with
// every other locale, thus ensuring that every text will have a fallback.
func (ls *Locales) LoadWithDefault(src fs.FS, locale Code) error {
	if err := ls.Load(src); err != nil {
		return err
	}

	l := ls.Get(locale)
	if l == nil {
		return fmt.Errorf("undefined default locale: %s", locale)
	}
	for _, loc := range ls.list {
		if loc == l {
			continue
		}
		loc.dict.Merge(l.dict)
	}

	return nil
}

// Get provides the define Locale object for the matching key.
func (ls *Locales) Get(code Code) *Locale {
	for _, loc := range ls.list {
		if loc.Code() == code {
			return loc
		}
	}
	return nil
}

// Match attempts to find the best possible matching locale based on the
// locale string provided. The locale string is parsed according to the
// "Accept-Language" header format defined in RFC9110.
func (ls *Locales) Match(locale string) *Locale {
	codes := ParseAcceptLanguage(locale)
	for _, code := range codes {
		for _, loc := range ls.list {
			if loc.Code() == code {
				return loc
			}
		}
	}
	return nil
}

// Codes provides a list of locale codes defined in the
// list.
func (ls *Locales) Codes() []Code {
	codes := make([]Code, len(ls.list))
	for i, l := range ls.list {
		codes[i] = l.Code()
	}
	return codes
}

// UnmarshalJSON attempts to load the locales from a JSON byte slice
// and merge them into any existing locales.
func (ls *Locales) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	aux := make(map[Code]*Dict)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	for c, v := range aux {
		if l := ls.Get(c); l != nil {
			l.dict.Merge(v)
		} else {
			ls.list = append(ls.list, NewLocale(c, v))
		}
	}
	return nil
}
