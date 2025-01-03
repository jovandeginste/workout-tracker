package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolToHTML(t *testing.T) {
	assert.Equal(t, "<i class=\"text-green-500 icon-[fa-solid--check]\"></i>", BoolToHTML(true))
	assert.Equal(t, "<i class=\"text-rose-500 icon-[fa-solid--times]\"></i>", BoolToHTML(false))
}
