package ast

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

const (
	prefixNplurals = "nplurals"
	prefixPlural   = "plural"
	variableN      = "n"
)

type scanner struct {
	r *bufio.Reader
}

func newScanner(r io.Reader) *scanner {
	return &scanner{r: bufio.NewReader(r)}
}

func (s *scanner) scan() (tok Token, lit string) {
	ch := s.read()
	if ch == scannerEOF {
		return eof, ""
	}

	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanText()
	} else if isDigit(ch) {
		s.unread()
		return s.scanNumber()
	}

	switch ch {
	case scannerEOF:
		return eof, ""
	case '%':
		return Reminder, string(ch)
	case '?':
		return Question, string(ch)
	case ':':
		return colon, string(ch)
	case ';':
		return semicolon, string(ch)
	case '(':
		return leftBracket, string(ch)
	case ')':
		return rightBracket, string(ch)
	case '&', '!', '|', '<', '>', '=':
		nextCh := s.read()
		switch string([]rune{ch, nextCh}) {
		case "!=":
			return NotEqual, "!="
		case "&&":
			return LogicalAnd, "&&"
		case "==":
			return Equal, "=="
		case "||":
			return LogicalOr, "||"
		case ">=":
			return GreaterOrEqual, ">="
		case "<=":
			return LessOrEqual, "<="
		}

		s.unread()
		switch ch {
		case '<':
			return Less, string(ch)
		case '>':
			return Greater, string(ch)
		case '=':
			return assign, string(ch)
		}
	}

	return failure, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *scanner) scanWhitespace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		ch := s.read()
		if ch == scannerEOF {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		}

		buf.WriteRune(ch)
	}

	return whitespace, buf.String()
}

func (s *scanner) scanNumber() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		ch := s.read()
		if ch == scannerEOF {
			break
		} else if !isDigit(ch) && ch != '_' {
			s.unread()
			break
		}
		_, _ = buf.WriteRune(ch)
	}

	return Value, buf.String()
}

func (s *scanner) scanText() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		ch := s.read()
		if ch == scannerEOF {
			break
		} else if !isLetter(ch) && ch != '_' {
			s.unread()
			break
		}
		_, _ = buf.WriteRune(ch)
	}

	switch strings.ToLower(buf.String()) {
	case prefixNplurals:
		return nPlurals, buf.String()
	case prefixPlural:
		return plural, buf.String()
	case variableN:
		return Operand, buf.String()
	default:
		return failure, buf.String()
	}
}

func (s *scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return scannerEOF
	}
	return ch
}

func (s *scanner) unread() { _ = s.r.UnreadRune() }

func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' }

func isLetter(ch rune) bool { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') }

func isDigit(ch rune) bool { return ch >= '0' && ch <= '9' }

// eof represents a marker rune for the end of the reader.
var scannerEOF = rune(0)
