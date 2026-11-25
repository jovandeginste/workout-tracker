package templatehelpers

import (
	"io/fs"
	"strings"
	"testing"

	apptranslations "github.com/jovandeginste/workout-tracker/v2/translations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToLanguageInformation(t *testing.T) {
	assert.Equal(t, "👽", LanguageToFlag("unknown-CODE"))
	assert.Equal(t, "🌐🇺🇸🇬🇧", LanguageToFlag("en"))
	assert.Equal(t, "🇳🇱🇧🇪", LanguageToFlag("nl"))
	assert.Equal(t, "🇨🇳", LanguageToFlag("zh-Hans"))
	assert.Equal(t, "🇧🇷", LanguageToFlag("pt-BR"))
	assert.Equal(t, "🇵🇹", LanguageToFlag("pt"))
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

func TestAllSupportedLanguagesHaveFlag(t *testing.T) {
	files, err := fs.ReadDir(apptranslations.FS(), ".")
	require.NoError(t, err)

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		if !strings.HasSuffix(name, ".yaml") {
			continue
		}
		langCode := strings.TrimSuffix(name, ".yaml")
		langCode = strings.ReplaceAll(langCode, "_", "-")

		flag := LanguageToFlag(langCode)
		assert.NotEqual(t, "👽", flag, "Translation %s should have a flag mapped", langCode)
	}
}
