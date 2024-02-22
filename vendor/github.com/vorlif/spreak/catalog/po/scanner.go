package po

import (
	"bytes"
	"io"
	"regexp"
	"strings"
)

var (
	reComment                = regexp.MustCompile(`^#`)              // #
	rePrevMsgContextComments = regexp.MustCompile(`^#\|\s+msgctxt`)  // #| msgctxt
	rePrevMsgIDComments      = regexp.MustCompile(`^#\|\s+msgid`)    // #| msgid
	rePrevStringLineComments = regexp.MustCompile(`^#\|\s+".*"\s*$`) // #| "message"

	reMsgContext   = regexp.MustCompile(`^msgctxt\s+".*"\s*$`)           // msgctxt
	reMsgID        = regexp.MustCompile(`^msgid\s+".*"\s*$`)             // msgid
	reMsgIDPlural  = regexp.MustCompile(`^msgid_plural\s+".*"\s*$`)      // msgid_plural
	reMsgStr       = regexp.MustCompile(`^msgstr\s*".*"\s*$`)            // msgstr
	reMsgStrPlural = regexp.MustCompile(`^msgstr\s*(\[\d+])\s*".*"\s*$`) // msgstr[0]
	reMsgLine      = regexp.MustCompile(`^\s*".*"\s*$`)                  // "message"
	reBlankLine    = regexp.MustCompile(`^\s*$`)                         //
)

type scanner struct {
	lines []string
	pos   int

	lastToken token
}

func newScanner(content string) *scanner {
	content = strings.Replace(content, "\r", "", -1)
	return &scanner{
		lines: strings.Split(content, "\n"),
		pos:   0,
	}
}

func (s *scanner) scan() (tok token, lit string) {
	line, err := s.read()
	defer func() {
		s.lastToken = tok
	}()

	if err != nil {
		return eof, ""
	}

	if tokk, l := s.scanRegex(line); tokk != none {
		return tokk, l
	}

	if !reComment.MatchString(line) {
		return failure, line
	}

	line = strings.TrimSpace(line)
	if len(line) == 1 {
		// comment without content
		return commentTranslator, line
	}

	if len(line) == 2 {
		// special comment without content
		switch line[1] {
		case '.', ':', ',', '|':
			return s.scan()
		}
	}

	switch line[1] {
	case '.':
		return commentExtracted, line
	case ':':
		return commentReference, line
	case ',':
		return commentFlags, line
	case '|':
		// #| "..."
		if rePrevMsgContextComments.MatchString(line) {
			return commentPrevContext, line
		} else if rePrevMsgIDComments.MatchString(line) {
			return commentPrevMsgID, line
		} else if rePrevStringLineComments.MatchString(line) {
			switch s.lastToken {
			case commentPrevContext, commentPrevContextLine:
				return commentPrevContextLine, line
			case commentPrevMsgID, commentPrevMsgIDLine:
				return commentPrevMsgIDLine, line
			default:
				return commentPrevUnknown, line
			}
		}

		return failure, line
	default:
		return commentTranslator, line
	}
}

func (s *scanner) scanRegex(line string) (token, string) {
	if reBlankLine.MatchString(line) {
		s.unread()
		return s.scanWhitespace()
	} else if reMsgID.MatchString(line) {
		return msgID, line
	} else if reMsgStr.MatchString(line) {
		return msgStr, line
	} else if reMsgIDPlural.MatchString(line) {
		return msgIDPlural, line
	} else if reMsgStrPlural.MatchString(line) {
		return msgStrPlural, line
	} else if reMsgContext.MatchString(line) {
		return msgContext, line
	} else if reMsgLine.MatchString(line) {
		// "..."
		switch s.lastToken {
		case msgID, msgIDLine:
			return msgIDLine, line
		case msgIDPlural, msgIDPluralLine:
			return msgIDPluralLine, line
		case msgStr, msgStrLine:
			return msgStrLine, line
		case msgStrPlural, msgStrPluralLine:
			return msgStrPluralLine, line
		case msgContext, msgContextLine:
			return msgContextLine, line
		}
	}

	return none, line
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *scanner) scanWhitespace() (tok token, lit string) {
	var buf bytes.Buffer
	currentLine, _ := s.read()
	buf.WriteString(currentLine)

	for {
		line, err := s.read()
		if err == io.EOF {
			break
		} else if !reBlankLine.MatchString(line) {
			s.unread()
			break
		}

		buf.WriteString(line)
	}

	return whitespace, buf.String()
}

func (s *scanner) read() (string, error) {
	if s.pos >= len(s.lines) {
		return "", io.EOF
	}

	line := s.lines[s.pos]
	s.pos++
	return line, nil
}

func (s *scanner) currentLine() string {
	if s.pos >= len(s.lines) {
		return ""
	}

	return s.lines[s.pos]
}

func (s *scanner) unread() {
	if s.pos > 0 {
		s.pos--
	}
}
