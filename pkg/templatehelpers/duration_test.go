package templatehelpers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHumanDuration(t *testing.T) {
	tests := map[time.Duration]string{
		0:                           "0s",
		1 * time.Second:             "1s",
		1 * time.Minute:             "1m",
		1 * time.Hour:               "1h",
		1*time.Hour + 1*time.Minute: "1h 1m",

		1*time.Hour + 1*time.Minute + 1*time.Second:   "1h 1m 1s",
		72*time.Hour + 30*time.Minute + 5*time.Second: "3d 30m 5s",
		100*time.Hour + 100*time.Second:               "4d 4h 1m 40s",
	}

	for d, expected := range tests {
		assert.Equal(t, expected, HumanDuration(d))
	}
}
