package po

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/vorlif/spreak/internal/util"
)

// An Encoder writes a po file to an output stream.
type Encoder struct {
	w                *bufio.Writer
	wrapWidth        int
	writeHeader      bool
	writeEmptyHeader bool
	writeReferences  bool
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w:                bufio.NewWriter(w),
		wrapWidth:        -1,
		writeHeader:      true,
		writeReferences:  true,
		writeEmptyHeader: true,
	}
}

// SetWrapWidth defines at which length the texts should be wrapped.
// To disable wrapping, the value can be set to -1.
// Default is -1.
func (e *Encoder) SetWrapWidth(wrapWidth int) {
	e.wrapWidth = wrapWidth
}

// SetWriteHeader sets whether a header should be written or not.
// Default is true.
func (e *Encoder) SetWriteHeader(write bool) {
	e.writeHeader = write
}

// SetWriteEmptyHeader sets whether a header without values should also be written or not.
// Default is true.
func (e *Encoder) SetWriteEmptyHeader(write bool) {
	e.writeEmptyHeader = write
}

// SetWriteReferences sets whether references to the origin of the text should be stored or not.
// Default is true.
func (e *Encoder) SetWriteReferences(write bool) {
	e.writeReferences = write
}

// Deprecated: Obsolete, it is always sorted, the method is removed with version 1.0.
func (e *Encoder) SetSort(_ bool) {}

func (e *Encoder) Encode(f *File) error {
	if f.Header != nil && e.writeHeader {
		if err := e.encodeHeader(f.Header); err != nil {
			return err
		}

		if _, err := e.w.WriteString("\n"); err != nil {
			return err
		}
	}

	if f.Messages != nil {
		var messages []*Message
		for ctx := range f.Messages {
			for _, msg := range f.Messages[ctx] {
				messages = append(messages, msg)
			}
		}
		sort.Slice(messages, func(i, j int) bool {
			return messages[j].Less(messages[i])
		})
		for _, msg := range messages {
			if err := e.encodeMessage(msg); err != nil {
				return err
			}

			if _, err := e.w.WriteString("\n"); err != nil {
				return err
			}
		}
	}

	return e.w.Flush()
}

type headerEntry struct {
	key   string
	value string
}

func (e *Encoder) encodeHeader(h *Header) error {
	if err := e.encodeComment(h.Comment); err != nil {
		return err
	}

	headers := []headerEntry{
		{HeaderProjectIDVersion, h.ProjectIDVersion},
		{HeaderReportMsgIDBugsTo, h.ReportMsgidBugsTo},
		{HeaderPOTCreationDate, h.POTCreationDate},
		{HeaderPORevisionDate, h.PORevisionDate},
		{HeaderLastTranslator, h.LastTranslator},
		{HeaderLanguageTeam, h.LanguageTeam},
		{HeaderLanguage, h.Language},
		{HeaderMIMEVersion, h.MimeVersion},
		{HeaderContentType, h.ContentType},
		{HeaderContentTransferEncoding, h.ContentTransferEncoding},
		{HeaderPluralForms, h.PluralForms},
		{HeaderXGenerator, h.XGenerator},
	}

	for k, v := range h.UnknownFields {
		headers = append(headers, headerEntry{k, v})
	}

	if !e.writeEmptyHeader {
		var hasHeader bool
		for _, header := range headers {
			if header.value != "" {
				hasHeader = true
				break
			}
		}
		if !hasHeader {
			return nil
		}
	}

	lines := make([]string, 0, len(headers))
	lines = append(lines, `msgid ""`+"\n")
	lines = append(lines, `msgstr ""`+"\n")
	for _, header := range headers {
		isOptional := header.key == HeaderMIMEVersion || header.key == HeaderPluralForms || header.key == HeaderXGenerator
		if isOptional && header.value == "" {
			continue
		}
		lines = append(lines, fmt.Sprintf(`"%s: %s\n"`+"\n", header.key, header.value))
	}

	for _, line := range lines {
		if _, err := e.w.WriteString(line); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encodeComment(c *Comment) error {
	if c == nil {
		return nil
	}

	if c.Translator != "" {
		for _, comment := range util.WrapString(c.Translator, e.wrapWidth) {
			if _, err := e.w.WriteString(fmt.Sprintf("# %s\n", comment)); err != nil {
				return err
			}
		}
	}

	if c.Extracted != "" {
		for _, comment := range util.WrapString(c.Extracted, e.wrapWidth) {
			if _, err := e.w.WriteString(fmt.Sprintf("#. %s\n", comment)); err != nil {
				return err
			}
		}
	}

	if len(c.References) > 0 && e.writeReferences {
		var buff bytes.Buffer
		for i, ref := range c.References {
			prefix := ""
			if i > 0 {
				prefix = " "
			}
			if ref.Line > 0 {
				buff.WriteString(fmt.Sprintf("%s%s:%d", prefix, ref.Path, ref.Line))
			} else {
				buff.WriteString(fmt.Sprintf("%s%s", prefix, ref.Path))
			}
		}

		for _, comment := range util.WrapString(buff.String(), e.wrapWidth) {
			if _, err := e.w.WriteString(fmt.Sprintf("#: %s\n", comment)); err != nil {
				return err
			}
		}
	}

	if len(c.Flags) > 0 {
		if _, err := e.w.WriteString(fmt.Sprintf("#, %s\n", strings.Join(c.Flags, ", "))); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encodeMessage(m *Message) error {
	if err := e.encodeComment(m.Comment); err != nil {
		return err
	}

	if m.Context != "" {
		ctx := fmt.Sprintf("msgctxt %s\n", EncodePoString(m.Context, e.wrapWidth))
		if _, err := e.w.WriteString(ctx); err != nil {
			return err
		}
	}

	msgID := fmt.Sprintf("msgid %s\n", EncodePoString(m.ID, e.wrapWidth))
	if _, err := e.w.WriteString(msgID); err != nil {
		return err
	}

	hasPlural := m.IDPlural != "" || len(m.Str) > 1
	if hasPlural {
		pluralID := fmt.Sprintf("msgid_plural %s\n", EncodePoString(m.IDPlural, e.wrapWidth))
		if _, err := e.w.WriteString(pluralID); err != nil {
			return err
		}
	}

	if err := e.encodeTranslations(hasPlural, m.Str); err != nil {
		return err
	}

	return nil
}

func (e *Encoder) encodeTranslations(plural bool, orig map[int]string) error {
	m := make(map[int]string, len(orig))
	for k, v := range orig {
		m[k] = EncodePoString(v, e.wrapWidth)
	}

	// We need at least one entry
	if len(m) == 0 {
		m[0] = `""`
	}

	if plural {
		if len(m) == 1 {
			// Plural needs at least two entries
			m[1] = `""`
		}

		keys := make([]int, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		for _, k := range keys {
			if _, err := e.w.WriteString(fmt.Sprintf("msgstr[%d] %s\n", k, m[k])); err != nil {
				return err
			}
		}
	} else {
		if _, err := e.w.WriteString(fmt.Sprintf("msgstr %s\n", m[0])); err != nil {
			return err
		}
	}

	return nil
}
