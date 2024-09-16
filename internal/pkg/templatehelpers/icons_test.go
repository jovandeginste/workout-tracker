package templatehelpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIconFor_Default(t *testing.T) {
	for _, dummy := range []string{
		"",
		"invalid-value-to-get-icon",
		";invalid-value-to-get-icon",
	} {
		d := IconFor(dummy)
		assert.Contains(t, d, "icon-decoration")
	}
}

func TestIconFor_Types(t *testing.T) {
	assert.Contains(t, IconFor("distance"), "fa6-solid--road")
	assert.Contains(t, IconFor("close"), "fa6-solid--xmark")
	assert.Contains(t, IconFor("dashboard"), "fa6-solid--chart-line")
	assert.Contains(t, IconFor("running"), "fa6-solid--person-running")
}
