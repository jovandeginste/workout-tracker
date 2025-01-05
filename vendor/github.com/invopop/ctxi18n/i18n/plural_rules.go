package i18n

// Standard pluralization rule keys.
const (
	DefaultRuleKey = "default"
)

// PluralRule defines a simple method that expects a dictionary and number and
// will find a matching dictionary entry.
type PluralRule func(d *Dict, num int) *Dict

const (
	zeroKey  = "zero"
	oneKey   = "one"
	otherKey = "other"
)

var rules = map[string]PluralRule{
	// Most languages can use this rule
	DefaultRuleKey: func(d *Dict, n int) *Dict {
		if n == 0 {
			v := d.Get(zeroKey)
			if v != nil {
				return v
			}
		}
		if n == 1 {
			return d.Get(oneKey)
		}
		return d.Get(otherKey)
	},
}

// GetRule provides the PluralRule for the given key.
func GetRule(key string) PluralRule {
	return rules[key]
}

// mapPluralRule is used to map a language code into a pluralization rule.
func mapPluralRule(_ Code) PluralRule {
	return rules[DefaultRuleKey]
}
