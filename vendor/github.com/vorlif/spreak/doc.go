// Package spreak provides a simple translation facility based on the concepts of gettext.
//
// # Fundamentals
//
// Domain: A message domain is a set of translatable messages.
// Usually, every software package has its own message domain.
// The domain name is used to determine the message catalog where the translation is looked up.
//
// Default domain: The default domain is used if a domain is not explicitly specified for a requested translation.
// If no default domain is specified, the default domain of the bundle is used.
// If this was not specified either, the domain is NoDomain (an empty string).
//
// Context: Context can be added to strings to be translated.
// A context dependent translation lookup is when a translation for a given string is searched,
// that is limited to a given context. The translation for the same string in a different context can be different.
// The different translations of the same string in different contexts can be stored in the same MO file,
// and can be edited by the translator in the same PO file.
// The Context string is visible in the PO file to the translator.
// You should try to make it somehow canonical and never changing.
// Because every time you change an Context, the translator will have to review the translation of msgid.
//
// # Plurals
//
// For JSON files only the CLDR plural rules are supported.
// For po and mo files both gettext plural forms and CLDR plural rules are supported.
// The CLDR rules provides better support when floating point numbers are used.
// When using the CLDR plural rules with po files, a notation increasing from "Zero" to "Other" should be used.
// For example, if the used language supports "Zero", "Few" and "Other", Zero should be notated as entry 0, Few as entry 1 and Other as entry 2.
// It is also recommended to define a gettext compatible plural rule.
// On the website https://php-gettext.github.io/Languages/ you can find a list of gettext plural rules which are compatible to the CLDR plural rules.
//
// To use the CLDR rules in po/mo files you can either add a header "X-spreak-use-CLDR: true" or create a decoder with
// catalog.NewPoCLDRDecoder() / catalog.NewMoCLDRDecoder().
//
// For Polish with One, Few and Other, the structure of a Po file according to this convention could look like this:
//
//	msgid ""
//	msgstr ""
//	"Plural-Forms: n == 1 ? 0 : n % 10 >= 2 && n % 10 <= 4 && (n % 100 < 12 || n % 100 > 14) ? 1 : 2;\n"
//	"X-spreak-use-CLDR: true\n"
//
//	msgid "id"
//	msgid_plural "plural id"
//	msgstr[0] "Translation with the plural form One"
//	msgstr[1] "Translation with the plural form Few"
//	msgstr[2] "Translation with the plural form Other"
//
// If floating point numbers are used, it is recommended to pass them formatted as strings as they will be displayed later.
// For example, if the number n is to be displayed with two numbers after the decimal point, it should be formatted with fmt.Sprintf("%.2f", n).
package spreak
