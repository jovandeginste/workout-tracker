package catalog

import (
	"fmt"
	"strings"

	"golang.org/x/text/language"
)

// Catalog represents a collection of messages (translations) for a language and a domain.
// Normally it is a PO or MO file.
type Catalog interface {
	// GetTranslation Returns a translation for an ID within a given context.
	GetTranslation(ctx, msgID string) (string, error)
	// GetPluralTranslation Returns a translation within a given context.
	// Here n is a number that should be used to determine the plural form.
	GetPluralTranslation(ctx, msgID string, n interface{}) (string, error)

	Language() language.Tag
}

// A Decoder reads and decodes catalogs for a language and a domain from a byte array.
type Decoder interface {
	Decode(lang language.Tag, domain string, data []byte) (Catalog, error)
}

func NewErrMissingContext(lang language.Tag, domain, context string) *ErrMissingContext {
	return &ErrMissingContext{
		Language: lang,
		Domain:   domain,
		Context:  context,
	}
}

// ErrMissingContext is the error returned when a matching context was not found for a language and domain.
type ErrMissingContext struct {
	Language language.Tag
	Domain   string
	Context  string
}

func (e *ErrMissingContext) Error() string { return e.String() }

func (e *ErrMissingContext) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("spreak: context not found: lang=%q ", e.Language))
	if e.Domain != "" {
		b.WriteString(fmt.Sprintf("domain=%q ", e.Domain))
	}
	if e.Context != "" {
		b.WriteString(fmt.Sprintf("ctx=%q ", e.Context))
	} else {
		b.WriteString("ctx='' (empty string)")
	}
	return b.String()
}

func NewErrMissingMessageID(lang language.Tag, domain, context, msgID string) *ErrMissingMessageID {
	return &ErrMissingMessageID{
		Language: lang,
		Domain:   domain,
		Context:  context,
		MsgID:    msgID,
	}
}

// ErrMissingMessageID is the error returned when a matching message was not found for a language and domain.
type ErrMissingMessageID struct {
	Language language.Tag
	Domain   string
	Context  string
	MsgID    string
}

func (e *ErrMissingMessageID) Error() string { return e.String() }

func (e *ErrMissingMessageID) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("spreak: msgID not found: lang=%q ", e.Language))
	if e.Domain != "" {
		b.WriteString(fmt.Sprintf("domain=%q ", e.Domain))
	}
	if e.Context != "" {
		b.WriteString(fmt.Sprintf("ctx=%q ", e.Context))
	}
	b.WriteString(fmt.Sprintf("msgID=%q", e.MsgID))
	return b.String()
}

func NewErrMissingTranslation(lang language.Tag, domain, context, msgID string, idx int) *ErrMissingTranslation {
	return &ErrMissingTranslation{
		Language: lang,
		Domain:   domain,
		Context:  context,
		MsgID:    msgID,
		Idx:      idx,
	}
}

// ErrMissingTranslation is the error returned when there is no translation for a domain of a language for a message.
type ErrMissingTranslation struct {
	Language language.Tag
	Domain   string
	Context  string
	MsgID    string
	Idx      int
}

func (e *ErrMissingTranslation) Error() string { return e.String() }

func (e *ErrMissingTranslation) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("spreak: translation not found: lang=%q ", e.Language))
	if e.Domain != "" {
		b.WriteString(fmt.Sprintf("domain=%q ", e.Domain))
	}
	if e.Context != "" {
		b.WriteString(fmt.Sprintf("ctx=%q ", e.Context))
	}
	b.WriteString(fmt.Sprintf("msgID=%q idx=%d", e.MsgID, e.Idx))
	return b.String()
}
