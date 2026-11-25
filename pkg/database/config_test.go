package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigCopy(t *testing.T) {
	t.Run("copies per user setting", func(t *testing.T) {
		c := Config{
			UserConfig: UserConfig{
				LoginCustomText: "# From the config file",
			},
		}
		from := Config{
			UserConfig: UserConfig{
				RegistrationDisabled: true,
				SocialsDisabled:      true,
				LoginCustomText:      "# From the database",
			},
		}

		c.copy(&from)

		assert.True(t, c.RegistrationDisabled)
		assert.True(t, c.SocialsDisabled)
		assert.Equal(t, "# From the config file", c.LoginCustomText)
	})

	t.Run("keeps non-zero destination values", func(t *testing.T) {
		c := Config{
			UserConfig: UserConfig{
				RegistrationDisabled: true,
				SocialsDisabled:      true,
				LoginCustomText:      "# From the config file",
			},
		}
		from := Config{
			UserConfig: UserConfig{
				LoginCustomText: "# From the database",
			},
		}

		c.copy(&from)

		assert.True(t, c.RegistrationDisabled)
		assert.True(t, c.SocialsDisabled)
		assert.Equal(t, "# From the config file", c.LoginCustomText)
	})
}
