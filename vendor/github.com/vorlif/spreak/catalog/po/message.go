package po

// Message is a representation of a single message in a catalog.
type Message struct {
	Comment  *Comment
	Context  string
	ID       string
	IDPlural string
	Str      map[int]string
}

type Messages map[string]map[string]*Message

func NewMessage() *Message {
	return &Message{
		Comment: &Comment{
			References: make([]*Reference, 0),
			Flags:      make([]string, 0),
		},
		Context:  "",
		ID:       "",
		IDPlural: "",
		Str:      make(map[int]string),
	}
}

func (m *Message) AddReference(ref *Reference) {
	if m.Comment == nil {
		m.Comment = NewComment()
	}

	m.Comment.AddReference(ref)
}

func (m *Message) Less(q *Message) bool {
	if m.Comment != nil && q.Comment != nil {
		return m.Comment.Less(q.Comment)
	}

	if a, b := m.ID, q.ID; a != b {
		return a > b
	}

	return false
}

func (m *Message) Merge(other *Message) {
	if other == nil {
		return
	}

	if m.Comment == nil {
		m.Comment = NewComment()
	}
	m.Comment.Merge(other.Comment)

	if m.IDPlural == "" && other.IDPlural != "" {
		m.IDPlural = other.IDPlural
	}
}

func (m Messages) Add(msg *Message) {
	if _, ok := m[msg.Context]; !ok {
		m[msg.Context] = make(map[string]*Message)
	}

	if _, ok := m[msg.Context][msg.ID]; ok {
		m[msg.Context][msg.ID].Merge(msg)
	} else {
		m[msg.Context][msg.ID] = msg
	}
}
