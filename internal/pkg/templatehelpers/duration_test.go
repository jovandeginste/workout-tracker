package templatehelpers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHumanDuration(t *testing.T) {
	tests := map[time.Duration]string{
		0:               "0s",
		1 * time.Second: "1s",
		1 * time.Minute: "60s",
		1 * time.Hour:   "60m",
		24 * time.Hour:  "24h",

		2 * time.Second: "2s",
		2 * time.Minute: "2m",
		2 * time.Hour:   "2h",
		48 * time.Hour:  "2d",

		1*time.Hour + 1*time.Minute:                   "61m",
		1*time.Hour + 1*time.Minute + 1*time.Second:   "61m 1s",
		25*time.Hour + 1*time.Minute + 1*time.Second:  "25h 1m",
		72*time.Hour + 30*time.Minute + 5*time.Second: "3d",
		100*time.Hour + 100*time.Second:               "4d 4h",
	}

	for d, expected := range tests {
		assert.Equal(t, expected, HumanDuration(d))
		assert.Equal(t, expected, HumanDuration(-d))
	}
}
