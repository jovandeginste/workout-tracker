package localize

// MsgID is an alias type for string which is used by xspreak for extracting strings
// MsgID and Singular are synonymous and both represent the ID to identify the message.
type MsgID = string

// Singular is an alias type for string which is used by xspreak for extracting strings
// MsgID and Singular are synonymous and both represent the ID to identify the message.
type Singular = string

// Plural is an alias type for string which is used by xspreak for extracting strings.
type Plural = string

// Context is an alias type for string which is used by xspreak for extracting strings.
type Context = string

// Domain is an alias type for string which is used by xspreak for extracting strings.
type Domain = string

// Key is an alias type for string which is used by xspreak for extracting strings.
// The exported key is not stored as text, only the key is stored.
// Should only be used when messages are identified by keys and not by strings. For example when using JSON.
type Key = string

// PluralKey is an alias type for string which is used by xspreak for extracting strings.
// It represents a MsgId AND a PluralId which identify a message.
// The exported key is not stored as text, only the key is stored.
// Should only be used when messages are identified by keys and not by strings. For example when using JSON.
type PluralKey = string

// A Localizable allows access to a message that needs to be translated.
//
// An implementation can be passed directly to a spreak.Localizer or spreak.Locale via
// l.Localize(impl) to obtain a translation.
type Localizable interface {
	// GetMsgID specifies the message id (singular) for which the message should be translated.
	GetMsgID() string
	// GetPluralID specifies the plural for which the message should be translated.
	GetPluralID() string
	// GetContext specifies the context for which the message should be translated.
	GetContext() string
	// GetVars is optional, can be used to pass parameters.
	GetVars() []interface{}
	GetCount() int
	// HasDomain specifies whether the domain of Domain() is to be used.
	// If false the default domain is used.
	HasDomain() bool
	// GetDomain specifies the domain for which the message should be translated.
	GetDomain() string
}

// Message is a simple struct representing a message without a domain.
// Can be converted to a translated string by spreak.Localizer or spreak.Locale.
type Message struct {
	Singular Singular
	Plural   Plural
	Context  Context
	Vars     []interface{}
	Count    int
}

var _ Localizable = (*Message)(nil)

func (m *Message) GetMsgID() string { return m.Singular }

func (m *Message) GetPluralID() string { return m.Plural }

func (m *Message) GetContext() string { return m.Context }

func (m *Message) GetVars() []interface{} { return m.Vars }

func (m *Message) GetCount() int { return m.Count }

func (m *Message) HasDomain() bool { return false }

func (m *Message) GetDomain() string { return "" }
