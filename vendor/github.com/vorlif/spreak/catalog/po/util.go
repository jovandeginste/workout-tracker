package po

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/vorlif/spreak/internal/util"
)

const nbsp = 0xA0

func DecodePoString(text string) string {
	lines := strings.Split(text, "\n")
	for idx := 0; idx < len(lines); idx++ {
		left := strings.Index(lines[idx], `"`)
		right := strings.LastIndex(lines[idx], `"`)
		if left < 0 || right < 0 || left == right {
			lines[idx] = ""
			continue
		}
		line := lines[idx][left+1 : right]

		res := unescapePoBytes([]byte(line))
		lines[idx] = string(res)
	}
	return strings.Join(lines, "")
}

func EncodePoString(s string, lim int) string {

	var lines []string
	if lim <= 0 {
		lines = encodePoStringWithoutWrap(s)
	} else {
		lines = encodePoStringWithWrap(s, lim)
	}

	// Single line msgid / entry starts with msgid "text"
	if len(lines) == 1 {
		// Single line
		return lines[0]
	} else if len(lines) == 2 && (lines[1] == "" || lines[1] == `""`) {
		// Single line with newline at end
		return lines[0]
	}

	lines = append(lines, "")
	copy(lines[1:], lines)
	lines[0] = `""`

	lastIdx := len(lines) - 1
	if lines[lastIdx] == "" || lines[lastIdx] == `""` {
		lines = lines[:lastIdx]
	}

	return strings.Join(lines, "\n")
}

func encodePoStringWithWrap(s string, pageWidth int) []string {
	lines := make([]string, 0, 2) // The average message is two lines long

	lineBuf := &util.StringBuffer{}
	var lastWordBuf util.StringBuffer
	var currentSpaceBuf util.StringBuffer

	var currentLen int

	createNewLine := func() {
		poLine := fmt.Sprintf(`"%s"`, escapePoString(lineBuf.String()))
		lines = append(lines, poLine)
		lineBuf.Reset()
		currentLen = 0
	}

	for _, char := range s {
		if char == '\n' {
			lastWordBuf.WriteInto(lineBuf)
			currentSpaceBuf.WriteInto(lineBuf)
			lineBuf.WriteRune('\n')
			createNewLine()
			continue
		}

		currentLen = lineBuf.Len()
		if unicode.IsSpace(char) && char != nbsp {
			if currentLen+lastWordBuf.Len()+currentSpaceBuf.Len()-1 >= pageWidth && currentLen > 0 {
				createNewLine()
			}

			currentSpaceBuf.WriteRune(char)

			if currentLen+lastWordBuf.Len()+currentSpaceBuf.Len()-1 >= pageWidth && currentLen > 0 {
				createNewLine()
			}
		} else {
			if currentSpaceBuf.Len() > 0 {
				if currentLen+lastWordBuf.Len()+currentSpaceBuf.Len()+1 >= pageWidth && currentLen == 0 {
					lastWordBuf.WriteInto(lineBuf)
					currentSpaceBuf.WriteInto(lineBuf)
					createNewLine()
				}

				lastWordBuf.WriteInto(lineBuf)
				currentSpaceBuf.WriteInto(lineBuf)
			}

			lastWordBuf.WriteRune(char)
		}
	}

	lastWordBuf.WriteInto(lineBuf)
	currentSpaceBuf.WriteInto(lineBuf)
	if remain := lineBuf.String(); remain != "" || len(lines) == 0 {
		poLine := fmt.Sprintf(`"%s"`, escapePoString(lineBuf.String()))
		lines = append(lines, poLine)
	}

	return lines
}

func encodePoStringWithoutWrap(s string) []string {
	lines := make([]string, 0, 2) // The average message is two lines long

	init := make([]byte, 0, len(s))
	buf := bytes.NewBuffer(init)

	for _, char := range s {
		// A newline closes the line
		if char == '\n' {
			poLine := fmt.Sprintf(`"%s\n"`, escapePoString(buf.String()))
			lines = append(lines, poLine)
			buf.Reset()
			continue
		}

		buf.WriteRune(char)
	}

	if remain := buf.String(); remain != "" || len(lines) == 0 {
		poLine := fmt.Sprintf(`"%s"`, escapePoString(buf.String()))
		lines = append(lines, poLine)
	}

	return lines
}

// Adapted from https://cs.opensource.google/go/go/+/refs/tags/go1.18.3:src/encoding/json/tables.go;drc=fc66cae490a0cd8b8cefefbc0ace7c3fb030f779;l=15

// safeSet holds the value true if the ASCII character with the given array
// position can be represented inside a po string without any further
// escaping.
//
// All values are true except for the ASCII control characters (0-31), the
// double quote ("), and the backslash character ("\").
var safeSet = [utf8.RuneSelf]bool{
	' ':      true,
	'!':      true,
	'"':      false,
	'#':      true,
	'$':      true,
	'%':      true,
	'&':      true,
	'\'':     true,
	'(':      true,
	')':      true,
	'*':      true,
	'+':      true,
	',':      true,
	'-':      true,
	'.':      true,
	'/':      true,
	'0':      true,
	'1':      true,
	'2':      true,
	'3':      true,
	'4':      true,
	'5':      true,
	'6':      true,
	'7':      true,
	'8':      true,
	'9':      true,
	':':      true,
	';':      true,
	'<':      true,
	'=':      true,
	'>':      true,
	'?':      true,
	'@':      true,
	'A':      true,
	'B':      true,
	'C':      true,
	'D':      true,
	'E':      true,
	'F':      true,
	'G':      true,
	'H':      true,
	'I':      true,
	'J':      true,
	'K':      true,
	'L':      true,
	'M':      true,
	'N':      true,
	'O':      true,
	'P':      true,
	'Q':      true,
	'R':      true,
	'S':      true,
	'T':      true,
	'U':      true,
	'V':      true,
	'W':      true,
	'X':      true,
	'Y':      true,
	'Z':      true,
	'[':      true,
	'\\':     false,
	']':      true,
	'^':      true,
	'_':      true,
	'`':      true,
	'a':      true,
	'b':      true,
	'c':      true,
	'd':      true,
	'e':      true,
	'f':      true,
	'g':      true,
	'h':      true,
	'i':      true,
	'j':      true,
	'k':      true,
	'l':      true,
	'm':      true,
	'n':      true,
	'o':      true,
	'p':      true,
	'q':      true,
	'r':      true,
	's':      true,
	't':      true,
	'u':      true,
	'v':      true,
	'w':      true,
	'x':      true,
	'y':      true,
	'z':      true,
	'{':      true,
	'|':      true,
	'}':      true,
	'~':      true,
	'\u007f': true,
}

var hex = "0123456789abcdef"

func escapePoString(s string) string {
	data := make([]byte, 0, len(s))
	buf := bytes.NewBuffer(data)

	start := 0
	for i := 0; i < len(s); {
		if b := s[i]; b < utf8.RuneSelf {
			if safeSet[b] {
				i++
				continue
			}
			if start < i {
				buf.WriteString(s[start:i])
			}
			buf.WriteByte('\\')
			switch b {
			case '\\', '"':
				buf.WriteByte(b)
			case '\n':
				buf.WriteByte('n')
			case '\r':
				buf.WriteByte('r')
			case '\t':
				buf.WriteByte('t')
			case '\a':
				buf.WriteByte('a')
			case '\b':
				buf.WriteByte('b')
			case '\f':
				buf.WriteByte('f')
			case '\v':
				buf.WriteByte('v')
			default:
				buf.WriteByte('x')
				buf.WriteByte(hex[b>>4])
				buf.WriteByte(hex[hex[b&0xF]])
			}
			i++
			start = i
			continue
		}

		c, size := utf8.DecodeRuneInString(s[i:])
		if c == utf8.RuneError && size == 1 {
			if start < i {
				buf.WriteString(s[start:i])
			}
			buf.WriteString(`\ufffd`)
			i += size
			start = i
			continue
		}

		i += size
	}

	if start < len(s) {
		buf.WriteString(s[start:])
	}

	return buf.String()
}

func unescapePoBytes(s []byte) []byte {
	// Check for unusual characters. If there are none,
	// then no unquoting is needed, so return a slice of the
	// original bytes.
	r := 0
	for r < len(s) {
		c := s[r]
		if c == '\\' || c == '"' || c < ' ' {
			break
		}
		if c < utf8.RuneSelf {
			r++
			continue
		}
		rr, size := utf8.DecodeRune(s[r:])
		if rr == utf8.RuneError && size == 1 {
			break
		}
		r += size
	}
	if r == len(s) {
		return s
	}

	b := make([]byte, len(s)+2*utf8.UTFMax)
	w := copy(b, s[0:r])
	for r < len(s) {
		// Out of room? Can only happen if s is full of
		// malformed UTF-8 and we're replacing each
		// byte with RuneError.
		if w >= len(b)-2*utf8.UTFMax {
			nb := make([]byte, (len(b)+utf8.UTFMax)*2)
			copy(nb, b[0:w])
			b = nb
		}
		switch c := s[r]; {
		case c == '\\':
			if r+1 >= len(s) {
				break
			}
			r++
			switch s[r] {
			case 'a':
				b[w] = '\a'
				r++
				w++
			case '"', '\\', '/', '\'':
				b[w] = s[r]
				r++
				w++
			case 'b':
				b[w] = '\b'
				r++
				w++
			case 'f':
				b[w] = '\f'
				r++
				w++
			case 'n':
				b[w] = '\n'
				r++
				w++
			case 'r':
				b[w] = '\r'
				r++
				w++
			case 't':
				b[w] = '\t'
				r++
				w++
			case 'v':
				b[w] = '\v'
				r++
				w++
			default:
				b[w] = '\\'
				w++
			}

		// ASCII
		case c < utf8.RuneSelf:
			b[w] = c
			r++
			w++

		// Coerce to well-formed UTF-8.
		default:
			rr, size := utf8.DecodeRune(s[r:])
			r += size
			w += utf8.EncodeRune(b[w:], rr)
		}
	}
	return b[0:w]
}
