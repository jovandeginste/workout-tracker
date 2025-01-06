package partials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPageOptions(t *testing.T) {
	po := NewPageOptions()

	assert.Empty(t, po.Scripts)
	assert.Empty(t, po.Styles)
}

func TestPageOptions_WithScripts(t *testing.T) {
	po := NewPageOptions().WithScripts("/a.js", "/b.js")

	assert.Contains(t, po.Scripts, "/a.js")
	assert.Contains(t, po.Scripts, "/b.js")
}

func TestPageOptions_WithStyles(t *testing.T) {
	po := NewPageOptions().WithStyles("/a.css", "/b.css")

	assert.Contains(t, po.Styles, "/a.css")
	assert.Contains(t, po.Styles, "/b.css")
}

func TestPageOptions_WithScriptsAndStyles(t *testing.T) {
	po := NewPageOptions().
		WithScripts("/a.js", "/b.js").
		WithStyles("/a.css", "/b.css")

	assert.Contains(t, po.Scripts, "/a.js")
	assert.Contains(t, po.Scripts, "/b.js")

	assert.Contains(t, po.Styles, "/a.css")
	assert.Contains(t, po.Styles, "/b.css")
}
