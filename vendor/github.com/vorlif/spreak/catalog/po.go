package catalog

import (
	"fmt"
	"strings"

	"golang.org/x/text/language"

	"github.com/vorlif/spreak/catalog/cldrplural"
	"github.com/vorlif/spreak/catalog/po"
	"github.com/vorlif/spreak/catalog/poplural"
	"github.com/vorlif/spreak/internal/mo"
)

const poCLDRHeader = "X-spreak-use-CLDR"

type gettextPluralFunction func(n interface{}) int

type poDecoder struct {
	useCLDRPlural bool
}

type moDecoder struct {
	useCLDRPlural bool
}

// NewPoDecoder returns a new Decoder for reading po files.
// If a plural forms header is set, it will be used.
// Otherwise, the CLDR plural rules are used to set the plural form.
// If there is no CLDR plural rule, the English plural rules will be used.
func NewPoDecoder() Decoder { return &poDecoder{} }

// NewPoCLDRDecoder creates a decoder for reading po files,
// which always uses the CLDR plural rules for determining the plural form.
// If no matching CLDR rule exists, the Po header rule is used. If no header exists,
// the english plural rules (1 is singular, otherwise plural) are used.
// Attention: The "Plural-Forms" header inside the Po file is ignored when using the CLDR rules.
// To ensure optimal compatibility with other applications, care should be taken to ensure that the Po header is compatible with the CLDR rules.
func NewPoCLDRDecoder() Decoder { return &poDecoder{useCLDRPlural: true} }

// NewMoDecoder returns a new Decoder for reading mo files.
// If a plural forms header is set, it will be used.
// Otherwise, the CLDR plural rules are used to set the plural form.
// If there is no CLDR plural rule, the English plural rules will be used.
func NewMoDecoder() Decoder { return &moDecoder{useCLDRPlural: false} }

// NewMoCLDRDecoder creates a decoder for reading mo files,
// which always uses the CLDR plural rules for determining the plural form.
// If no matching CLDR rule exists, the Mo header rule is used. If no header exists,
// the english plural rules (1 is singular, otherwise plural) are used.
// Attention: The "Plural-Forms" header inside the Mo file is ignored when using the CLDR rules.
// To ensure optimal compatibility with other applications, care should be taken to ensure that the Mo header is compatible with the CLDR rules.
func NewMoCLDRDecoder() Decoder { return &moDecoder{useCLDRPlural: true} }

func (d poDecoder) Decode(lang language.Tag, domain string, data []byte) (Catalog, error) {
	poFile, errParse := po.Parse(data)
	if errParse != nil {
		return nil, errParse
	}

	// We could check here if the language of the file matches the target language,
	// but leave it off to make loading more flexible.

	return buildGettextCatalog(poFile, lang, domain, d.useCLDRPlural)
}

func (d moDecoder) Decode(lang language.Tag, domain string, data []byte) (Catalog, error) {
	moFile, errParse := mo.ParseBytes(data)
	if errParse != nil {
		return nil, errParse
	}

	// We could check here if the language of the file matches the target language,
	// but leave it off to make loading more flexible.

	return buildGettextCatalog(moFile, lang, domain, d.useCLDRPlural)
}

func buildGettextCatalog(file *po.File, lang language.Tag, domain string, useCLDRPlural bool) (Catalog, error) {
	messages := make(messageLookupMap, len(file.Messages))

	for ctx := range file.Messages {
		if len(file.Messages[ctx]) == 0 {
			continue
		}

		if _, hasContext := messages[ctx]; !hasContext {
			messages[ctx] = make(map[string]*gettextMessage)
		}

		for msgID, poMsg := range file.Messages[ctx] {
			if msgID == "" {
				continue
			}

			if poMsg.Comment != nil && poMsg.Comment.HasFlag("fuzzy") {
				continue
			}

			d := &gettextMessage{
				Context:      poMsg.Context,
				ID:           poMsg.ID,
				IDPlural:     poMsg.IDPlural,
				Translations: poMsg.Str,
			}

			messages[poMsg.Context][poMsg.ID] = d
		}
	}

	catl := &gettextCatalog{
		language:     lang,
		translations: messages,
	}

	if useCLDRPlural {
		catl.pluralFunc = getCLDRPluralFunction(lang)
		return catl, nil
	}

	if file.Header != nil {
		if val := file.Header.Get(poCLDRHeader); strings.ToLower(val) == "true" {
			catl.pluralFunc = getCLDRPluralFunction(lang)
			return catl, nil
		}

		if file.Header.PluralForms != "" {
			forms, err := poplural.Parse(file.Header.PluralForms)
			if err != nil {
				return nil, fmt.Errorf("spreak.Decoder: plural forms for po file %v#%v could not be parsed: %w", lang, domain, err)
			}
			catl.pluralFunc = forms.Evaluate
			return catl, nil
		}
	}

	catl.pluralFunc = getCLDRPluralFunction(lang)
	return catl, nil
}

func getCLDRPluralFunction(lang language.Tag) func(a any) int {
	ruleSet, _ := cldrplural.ForLanguage(lang)

	catToForm := make(map[cldrplural.Category]int, len(ruleSet.Categories))
	for idx, cat := range ruleSet.Categories {
		catToForm[cat] = idx
	}

	return func(a any) int {
		cat := ruleSet.Evaluate(a)
		if form, ok := catToForm[cat]; ok {
			return form
		}

		return 0
	}
}

type gettextCatalog struct {
	language language.Tag

	translations messageLookupMap
	pluralFunc   gettextPluralFunction
	domain       string
}

type gettextMessage struct {
	Context      string
	ID           string
	IDPlural     string
	Translations map[int]string
}

// Map for a quick lookup of messages.
// First key is the context and second the MsgID (e.g. lookup["context"]["hello"]).
type messageLookupMap map[string]map[string]*gettextMessage

var _ Catalog = (*gettextCatalog)(nil)

func (c *gettextCatalog) GetTranslation(ctx, msgID string) (string, error) {
	msg, err := c.getMessage(ctx, msgID, 0)
	if err != nil {
		return msgID, err
	}

	return msg.Translations[0], nil
}

func (c *gettextCatalog) GetPluralTranslation(ctx, msgID string, n interface{}) (string, error) {
	idx := c.pluralFunc(n)
	msg, err := c.getMessage(ctx, msgID, idx)
	if err != nil {
		return msgID, err
	}

	return msg.Translations[idx], nil
}

func (c *gettextCatalog) Language() language.Tag { return c.language }

func (c *gettextCatalog) getMessage(ctx, msgID string, idx int) (*gettextMessage, error) {
	if _, hasCtx := c.translations[ctx]; !hasCtx {
		return nil, NewErrMissingContext(c.language, c.domain, ctx)
	}

	if _, hasMsg := c.translations[ctx][msgID]; !hasMsg {
		return nil, NewErrMissingMessageID(c.language, c.domain, ctx, msgID)
	}

	msg := c.translations[ctx][msgID]
	if tr, hasTranslation := msg.Translations[idx]; !hasTranslation || tr == "" {
		return nil, NewErrMissingTranslation(c.language, c.domain, ctx, msgID, idx)
	}

	return msg, nil
}
