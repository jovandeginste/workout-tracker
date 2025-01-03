package helpers

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
		d := iconFor(dummy)
		assert.Contains(t, d, "icon-decoration")
	}
}

func TestIconFor_Types(t *testing.T) {
	assert.Contains(t, iconFor("distance"), "fa6-solid--road")
	assert.Contains(t, iconFor("close"), "fa6-solid--xmark")
	assert.Contains(t, iconFor("dashboard"), "fa6-solid--chart-line")
	assert.Contains(t, iconFor("running"), "fa6-solid--person-running")
}
