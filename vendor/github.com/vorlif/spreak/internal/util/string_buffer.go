package util

import "bytes"

type StringBuffer struct {
	buff bytes.Buffer
	len  int
}

func (b *StringBuffer) Len() int {
	return b.len
}

func (b *StringBuffer) Reset() {
	b.buff.Reset()
	b.len = 0
}

func (b *StringBuffer) WriteInto(w *StringBuffer) {
	_, _ = b.buff.WriteTo(&w.buff)
	w.len += b.len
	b.len = 0
}

func (b *StringBuffer) String() string {
	return b.buff.String()
}

func (b *StringBuffer) WriteRune(r rune) {
	b.buff.WriteRune(r)
	b.len++
}

func (b *StringBuffer) WriteString(s string) {
	_, _ = b.buff.WriteString(s)
	b.len += len(s)
}
