package util

import (
	"strings"
	"unicode"
)

const nbsp = 0xA0

func WrapString(s string, pageWidth int) []string {
	if pageWidth <= 0 {
		return strings.Split(s, "\n")
	}

	lines := make([]string, 0, 10)

	currentLine := &StringBuffer{}
	var currentWordBuf StringBuffer
	var lastSpaceBuf StringBuffer

	for _, char := range s {
		if char == '\n' {
			if currentWordBuf.Len() == 0 {
				lastSpaceBuf.Reset()
			}
			lastSpaceBuf.WriteInto(currentLine)
			currentWordBuf.WriteInto(currentLine)
			lines = append(lines, currentLine.String())
			currentLine.Reset()
		} else if unicode.IsSpace(char) && char != nbsp {
			if currentWordBuf.Len() > 0 { // We had a word before and now a space
				lastSpaceBuf.WriteInto(currentLine)
				currentWordBuf.WriteInto(currentLine)
			}

			lastSpaceBuf.WriteRune(char)
		} else {
			if currentLine.Len()+lastSpaceBuf.Len()+currentWordBuf.Len() >= pageWidth {
				if currentLine.Len() > 0 {
					lines = append(lines, currentLine.String())
					currentLine.Reset()
					lastSpaceBuf.Reset()
				}

				if lastSpaceBuf.Len() > 0 {
					lastSpaceBuf.WriteInto(currentLine)
					currentWordBuf.WriteInto(currentLine)
				}
			}

			currentWordBuf.WriteRune(char)
		}
	}

	if currentWordBuf.Len() > 0 {
		lastSpaceBuf.WriteInto(currentLine)
		currentWordBuf.WriteInto(currentLine)
	}

	lines = append(lines, currentLine.String())
	return lines
}
