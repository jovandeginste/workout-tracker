package po

import (
	"fmt"
	"sort"
	"strings"
)

type Reference struct {
	Path   string
	Line   int
	Column int
}

func (r Reference) String() string {
	if r.Line > 0 {
		return fmt.Sprintf("%s:%d:%d", r.Path, r.Line, r.Column)
	}

	return r.Path
}

func (r Reference) Equal(o *Reference) bool {
	return r.Path == o.Path && r.Line == o.Line && r.Column == o.Column
}

type Comment struct {
	Translator     string       // #  translator-comments
	Extracted      string       // #. extracted-comments
	References     []*Reference // #: src/file.go:210
	Flags          []string     // #, fuzzy,go-format,range:0..10
	PrevMsgContext string       // #| msgctxt previous-context
	PrevMsgID      string       // #| msgid previous-untranslated-string
}

func NewComment() *Comment {
	return &Comment{
		References: []*Reference{},
		Flags:      []string{},
	}
}

func (c *Comment) AddReference(ref *Reference) {
	if c.References == nil {
		c.References = make([]*Reference, 0)
	}
	c.References = append(c.References, ref)
	c.sort()
}

func (c *Comment) Less(q *Comment) bool {
	c.sort()
	for i := 0; i < len(c.References); i++ {
		if i >= len(q.References) {
			break
		}
		if c := strings.Compare(c.References[i].Path, q.References[i].Path); c != 0 {
			return c == 1
		}
		if a, b := c.References[i].Line, q.References[i].Line; a != b {
			return a > b
		}
		if a, b := c.References[i].Column, q.References[i].Column; a != b {
			return a > b
		}
	}
	return false
}

func (c *Comment) HasFlag(flag string) bool {
	for _, s := range c.Flags {
		if s == flag {
			return true
		}
	}
	return false
}

func (c *Comment) AddFlag(flag string) {
	if c.HasFlag(flag) {
		return
	}
	c.Flags = append(c.Flags, flag)
}

func (c *Comment) Merge(other *Comment) {
	if other == nil {
		return
	}

	if other.Translator != "" {
		left := strings.Split(c.Translator, "\n")
		right := strings.Split(other.Translator, "\n")
		res := mergeStringArrays(left, right)
		c.Translator = strings.TrimSpace(strings.Join(res, "\n"))
	}

	if other.Extracted != "" {
		left := strings.Split(c.Extracted, "\n")
		right := strings.Split(other.Extracted, "\n")
		res := mergeStringArrays(left, right)
		c.Extracted = strings.TrimSpace(strings.Join(res, "\n"))
	}

	c.Flags = mergeStringArrays(c.Flags, other.Flags)
	c.mergeReferences(other)

	c.sort()
}

func (c *Comment) mergeReferences(other *Comment) {
	newReferences := make([]*Reference, 0)

	for _, otherRef := range other.References {
		hasRef := false
		for _, ref := range c.References {
			if ref.Equal(otherRef) {
				hasRef = true
				break
			}
		}

		if !hasRef {
			newReferences = append(newReferences, otherRef)
		}
	}

	c.References = append(c.References, newReferences...)
}

func (c *Comment) sort() {
	sort.Strings(c.Flags)

	sort.Slice(c.References, func(i, j int) bool {
		if c.References[i].Path != c.References[j].Path {
			return c.References[i].Path < c.References[j].Path
		}

		if c.References[i].Line != c.References[j].Line {
			return c.References[i].Line < c.References[j].Line
		}

		return c.References[i].Column < c.References[j].Column
	})
}

func mergeStringArrays(left, right []string) []string {
	dst := make([]string, len(left), len(left)+len(right))
	copy(dst, left)

	for _, a := range right {
		hasA := false
		for _, b := range left {
			if b == a {
				hasA = true
				break
			}
		}

		if !hasA {
			dst = append(dst, a)
		}
	}

	return dst
}
