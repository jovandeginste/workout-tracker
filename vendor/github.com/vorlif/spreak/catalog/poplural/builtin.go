package poplural

import (
	"golang.org/x/text/language"
)

const fallbackRule = "nplurals=2; plural=n != 1;"

var (
	// "nplurals=2; plural=(n != 1);" -> Form.
	rawToBuiltIn = make(map[string]*Form)

	// language.Tag.String() -> Form.
	langToBuiltIn = make(map[string]*Form)
)

type PluralFunc = func(n interface{}) int

func ForLanguage(lang language.Tag) (PluralFunc, bool) {
	form, found := pluralRuleForLanguage(lang)
	return form.Evaluate, found
}

func pluralRuleForLanguage(lang language.Tag) (*Form, bool) {
	n := lang
	for !n.IsRoot() {
		if form, hasForm := langToBuiltIn[n.String()]; hasForm {
			return form, true
		}

		base, confidence := n.Base()
		if confidence >= language.High {
			if form, hasForm := langToBuiltIn[base.String()]; hasForm {
				return form, true
			}
		}

		n = n.Parent()
	}

	return rawToBuiltIn[fallbackRule], false
}

func registerLanguageForm(langs []string, rawRule string, form *Form) {
	rawToBuiltIn[rawRule] = form

	for _, lang := range langs {
		tag := language.MustParse(lang)
		langToBuiltIn[tag.String()] = form
	}
}

func registerRawForm(rawRule string, form *Form) {
	rawToBuiltIn[rawRule] = form
}
