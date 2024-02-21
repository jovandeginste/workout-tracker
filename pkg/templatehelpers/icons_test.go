package templatehelpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIconForDefault(t *testing.T) {
	for _, dummy := range []string{
		"",
		"invalid-value-to-get-icon",
		";invalid-value-to-get-icon",
	} {
		d := IconFor(dummy)
		assert.Contains(t, d, "icon-solid")
		assert.Contains(t, d, "icon-baseline")
		assert.Contains(t, d, "icon-space-sm")
		assert.Contains(t, d, "icon-before")
		assert.Contains(t, d, "icon-question")
	}
}

func TestIconFor(t *testing.T) {
	assert.Contains(t, IconFor("distance"), "icon-road")
	assert.Contains(t, IconFor("close"), "icon-xmark")
	assert.Contains(t, IconFor("dashboard"), "icon-chart-line")
	assert.Contains(t, IconFor("running"), "icon-person-running")
}
