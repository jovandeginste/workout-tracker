package i18n

import "strings"

// Code is used to represent a language code which follows the
// ISO 639-1 standard, with sub-tags aggregated with hyphens,
// as defined in [RFC5646](https://datatracker.ietf.org/doc/html/rfc5646).
// Examples include:
//
// fr, en-US, es-419, az-Arab, x-pig-latin, man-Nkoo-GN
type Code string

// String returns the string variant of the code.
func (c Code) String() string {
	return string(c)
}

// Base returns the base language code, without any subtags.
func (c Code) Base() Code {
	out := strings.SplitN(c.String(), "-", 2)
	return Code(out[0])
}

// ParseAcceptLanguage provides an ordered set of codes extracted
// from an HTTP "Accept-Language" header as defined in RFC9110.
// Current implementation will ignore quality values and instead
// just assume the order of the provided codes is valid.
func ParseAcceptLanguage(txt string) []Code {
	list := make([]Code, 0)
	for _, s := range strings.Split(txt, ",") {
		s = strings.TrimSpace(s)

		// Remove any quality values.
		if i := strings.Index(s, ";"); i > 0 {
			s = s[:i]
		}

		list = append(list, Code(s))
	}
	return list
}
