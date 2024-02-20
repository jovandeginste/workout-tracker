package templatehelpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIconFor(t *testing.T) {
	assert.Equal(t, "icon-solid icon-baseline icon-space-sm icon-before icon-question", IconFor(""))
	assert.Contains(t, IconFor("distance"), "icon-road")
	assert.Contains(t, IconFor("close"), "icon-xmark")
}
