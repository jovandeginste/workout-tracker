package templatehelpers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNumericDuration(t *testing.T) {
	assert.InDelta(t, 1, NumericDuration(time.Second), 0)
	assert.InDelta(t, 3600, NumericDuration(time.Hour), 0)
}

func TestToLanguageInformation(t *testing.T) {
	assert.Equal(t, "🇺🇸", LanguageToFlag("en-GB"))
	assert.Equal(t, "🇨🇳", LanguageToFlag("zh-Hans"))
}

func TestCountryCodeToFlag(t *testing.T) {
	assert.Equal(t, "🇺🇦", CountryToFlag("UA"))
	assert.Equal(t, "🇧🇪", CountryToFlag("BE"))
}

func TestHumanDistanceKM(t *testing.T) {
	assert.Equal(t, "0", HumanDistanceKM(1.23))
	assert.Equal(t, "1.23", HumanDistanceKM(1234))
	assert.Equal(t, "1234.57", HumanDistanceKM(1234567))
}

func TestHumanDistance(t *testing.T) {
	assert.Equal(t, "0", HumanDistanceKM(1.23))
	assert.Equal(t, "1.23", HumanDistanceKM(1234))
	assert.Equal(t, "1234.57", HumanDistanceKM(1234567))
}

func TestHumanSpeedKPH(t *testing.T) {
	assert.Equal(t, "4.43", HumanSpeedKPH(1.23))
	assert.Equal(t, "10.01", HumanSpeedKPH(2.78))
	assert.Equal(t, "17.96", HumanSpeedKPH(4.99))
}

func TestHumanTempoKM(t *testing.T) {
	assert.Equal(t, "13:33", HumanTempoKM(1.23))
	assert.Equal(t, "5:59", HumanTempoKM(2.78))
	assert.Equal(t, "3:20", HumanTempoKM(4.99))
	assert.Equal(t, "5:01", HumanTempoKM(3.32))
}
