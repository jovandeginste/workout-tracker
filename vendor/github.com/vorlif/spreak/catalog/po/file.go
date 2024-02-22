package po

import (
	"fmt"
	"strings"
)

const (
	HeaderProjectIDVersion        = "Project-Id-Version"
	HeaderReportMsgIDBugsTo       = "Report-Msgid-Bugs-To"
	HeaderPOTCreationDate         = "POT-Creation-Date"
	HeaderPORevisionDate          = "PO-Revision-Date"
	HeaderLastTranslator          = "Last-Translator"
	HeaderLanguageTeam            = "Language-Team"
	HeaderLanguage                = "Language"
	HeaderMIMEVersion             = "MIME-Version"
	HeaderContentType             = "Content-Type"
	HeaderContentTransferEncoding = "Content-Transfer-Encoding"
	HeaderPluralForms             = "Plural-Forms"
	HeaderXGenerator              = "X-Generator"
)

type File struct {
	Header   *Header
	Messages Messages
}

func NewFile() *File {
	return &File{
		Header:   &Header{},
		Messages: make(Messages),
	}
}

func (f *File) AddMessage(msg *Message) {
	if f.Messages == nil {
		f.Messages = make(Messages)
	}

	if msg.ID == "" {
		return
	}

	f.Messages.Add(msg)
}

func (f *File) GetMessage(ctx string, id string) *Message {
	if _, hasCtx := f.Messages[ctx]; !hasCtx {
		return nil
	}

	msg, ok := f.Messages[ctx][id]
	if !ok {
		return nil
	}

	return msg
}

func (f *File) String() string {
	if f.Header != nil {
		return fmt.Sprintf("po file %s %s", f.Header.ProjectIDVersion, f.Header.Language)
	}
	return "po file"
}

type Header struct {
	Comment                 *Comment // Header Comments
	ProjectIDVersion        string   // Project-Id-Version: PACKAGE VERSION
	ReportMsgidBugsTo       string   // Report-Msgid-Bugs-To: FIRST AUTHOR <EMAIL@ADDRESS>
	POTCreationDate         string   // POT-Creation-Date: YEAR-MO-DA HO:MI+ZONE
	PORevisionDate          string   // PO-Revision-Date: YEAR-MO-DA HO:MI+ZONE
	LastTranslator          string   // Last-Translator: FIRST AUTHOR <EMAIL@ADDRESS>
	LanguageTeam            string   // Language-Team:
	Language                string   // Language: de
	MimeVersion             string   // MIME-Version: 1.0
	ContentType             string   // Content-Type: text/plain; charset=UTF-8
	ContentTransferEncoding string   // Content-Transfer-Encoding: 8bit
	PluralForms             string   // Plural-Forms: nplurals=2; plural=(n != 1);
	XGenerator              string   // X-Generator: Poedit 3.0.1
	UnknownFields           map[string]string
}

func (h *Header) SetField(key, val string) {
	switch strings.ToUpper(key) {
	case strings.ToUpper(HeaderProjectIDVersion):
		h.ProjectIDVersion = val
	case strings.ToUpper(HeaderReportMsgIDBugsTo):
		h.ReportMsgidBugsTo = val
	case strings.ToUpper(HeaderPOTCreationDate):
		h.POTCreationDate = val
	case strings.ToUpper(HeaderPORevisionDate):
		h.PORevisionDate = val
	case strings.ToUpper(HeaderLastTranslator):
		h.LastTranslator = val
	case strings.ToUpper(HeaderLanguageTeam):
		h.LanguageTeam = val
	case strings.ToUpper(HeaderLanguage):
		h.Language = val
	case strings.ToUpper(HeaderMIMEVersion):
		h.MimeVersion = val
	case strings.ToUpper(HeaderContentType):
		h.ContentType = val
	case strings.ToUpper(HeaderContentTransferEncoding):
		h.ContentTransferEncoding = val
	case strings.ToUpper(HeaderPluralForms):
		h.PluralForms = val
	case strings.ToUpper(HeaderXGenerator):
		h.XGenerator = val
	default:
		if h.UnknownFields == nil {
			h.UnknownFields = make(map[string]string)
		}
		h.UnknownFields[key] = val
	}
}

func (h *Header) Get(key string) string {
	if h.UnknownFields == nil {
		return ""
	}

	key = strings.ToUpper(key)
	for unknownHeader, val := range h.UnknownFields {
		if strings.ToUpper(unknownHeader) == key {
			return val
		}
	}

	return ""
}
