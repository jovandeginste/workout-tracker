package catalog

import (
	"encoding/json"
	"fmt"

	"golang.org/x/text/language"

	"github.com/vorlif/spreak/catalog/cldrplural"
)

type jsonDecoder struct{}

// NewJSONDecoder returns a new Decoder for reading JSON files.
// The structure follows a key-value structure, where the key is either an ID or the singular text of the source language.
// For singular-only texts, the value is a string with a translation.
// For plural texts it is an object with the CLDR plural forms and the matching translations.
func NewJSONDecoder() Decoder { return jsonDecoder{} }

func (jsonDecoder) Decode(lang language.Tag, domain string, data []byte) (Catalog, error) {
	var messages jsonFile
	if err := json.Unmarshal(data, &messages); err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, fmt.Errorf("spreak: File contains no translations lang=%v domain=%q", lang, domain)
	}

	catl := &jsonCatalog{
		lookupMap: make(map[string]map[string]*jsonMessage),
		domain:    domain,
		language:  lang,
	}

	catl.pluralSet, _ = cldrplural.ForLanguage(lang)

	for key, msg := range messages {
		if key == "" || msg == nil || msg.Other == "" {
			continue
		}

		if _, ok := catl.lookupMap[msg.Context]; !ok {
			catl.lookupMap[msg.Context] = make(map[string]*jsonMessage)
		}

		catl.lookupMap[msg.Context][key] = msg
	}

	return catl, nil
}

type jsonCatalog struct {
	// Map for a quick lookup of messages.
	// First key is the context and second the msg key (e.g. lookup["context"]["app.name"]).
	lookupMap map[string]map[string]*jsonMessage
	domain    string
	language  language.Tag
	pluralSet *cldrplural.RuleSet
}

func (m *jsonCatalog) GetTranslation(ctx, msgID string) (string, error) {
	tr, err := m.getTranslation(ctx, msgID, cldrplural.Other)
	if err != nil {
		return msgID, err
	}

	return tr, nil
}

func (m *jsonCatalog) GetPluralTranslation(ctx, msgID string, n interface{}) (string, error) {
	cat := m.pluralSet.Evaluate(n)
	tr, err := m.getTranslation(ctx, msgID, cat)
	if err != nil {
		return msgID, err
	}

	return tr, nil
}

func (m jsonCatalog) Language() language.Tag { return m.language }

func (m *jsonCatalog) getTranslation(ctx, key string, cat cldrplural.Category) (string, error) {
	if ctx != "" {
		key += "_" + ctx
	}
	if _, hasCtx := m.lookupMap[ctx]; !hasCtx {
		return "", NewErrMissingContext(m.language, m.domain, ctx)
	}

	if _, hasMsg := m.lookupMap[ctx][key]; !hasMsg {
		return "", NewErrMissingMessageID(m.language, m.domain, ctx, key)
	}

	msg := m.lookupMap[ctx][key]
	tr := msg.getTranslation(cat)
	if tr == "" {
		return "", NewErrMissingTranslation(m.language, m.domain, ctx, key, int(cat))
	}

	return tr, nil
}

type jsonFile map[string]*jsonMessage

type jsonMessage struct {
	Comment string `json:"comment,omitempty"`
	Context string `json:"context,omitempty"`

	Zero  string `json:"zero,omitempty"`
	One   string `json:"one,omitempty"`
	Two   string `json:"two,omitempty"`
	Few   string `json:"few,omitempty"`
	Many  string `json:"many,omitempty"`
	Other string `json:"other"`
}

func (m *jsonMessage) getTranslation(cat cldrplural.Category) string {
	switch cat {
	case cldrplural.Zero:
		return m.Zero
	case cldrplural.One:
		return m.One
	case cldrplural.Two:
		return m.Two
	case cldrplural.Few:
		return m.Few
	case cldrplural.Many:
		return m.Many
	default:
		return m.Other
	}
}

type jsonMessageAlias jsonMessage

func (m *jsonMessage) UnmarshalJSON(data []byte) error {
	var other string
	if err := json.Unmarshal(data, &other); err == nil {
		m.Other = other
		return nil
	}

	aux := &struct{ *jsonMessageAlias }{jsonMessageAlias: (*jsonMessageAlias)(m)}
	return json.Unmarshal(data, aux)
}
