package cldrplural

import "golang.org/x/text/language"

var builtInRuleSets = make(map[string]*RuleSet, 40)

// ForLanguage returns the set of rules for a language.
// If no matching language is found, the English rule set and false are returned.
func ForLanguage(lang language.Tag) (*RuleSet, bool) {
	n := lang
	for !n.IsRoot() {
		if form, hasForm := builtInRuleSets[n.String()]; hasForm {
			return form, true
		}

		base, confidence := n.Base()
		if confidence >= language.High {
			if form, hasForm := builtInRuleSets[base.String()]; hasForm {
				return form, true
			}
		}

		n = n.Parent()
	}

	return builtInRuleSets[language.English.String()], false
}

func addRuleSet(langs []string, set *RuleSet) {
	for _, lang := range langs {
		tag := language.MustParse(lang)
		builtInRuleSets[tag.String()] = set
	}
}

func newCategories(categories ...Category) []Category {
	return categories
}
