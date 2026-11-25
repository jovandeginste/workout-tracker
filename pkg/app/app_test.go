package app

import (
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/fsouza/slognil"
	appassets "github.com/jovandeginste/workout-tracker/v2/assets"
	"github.com/jovandeginste/workout-tracker/v2/pkg/version"
	apptranslations "github.com/jovandeginste/workout-tracker/v2/translations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func defaultApp(t *testing.T) *App {
	t.Helper()
	t.Setenv("WT_LOGGING", "false")

	a := NewApp(version.Version{RefName: "test"})

	a.Assets = appassets.FS()
	a.Translations = apptranslations.FS()

	return a
}

func TestApp_RandomJWTError(t *testing.T) {
	a1 := defaultApp(t)
	s1 := a1.jwtSecret()
	assert.NotEmpty(t, s1)

	a2 := defaultApp(t)
	s2 := a2.jwtSecret()
	assert.NotEqual(t, s1, s2)
}

func TestApp_NewApp(t *testing.T) {
	a := defaultApp(t)
	assert.NotNil(t, a.rawLogger)
	assert.NotNil(t, a.logger)
	assert.IsType(t, slognil.Handler{}, a.logger.Handler())
	assert.Equal(t, "test", a.Version.RefName)
}

func TestApp_Configure(t *testing.T) {
	a := defaultApp(t)
	assert.Nil(t, a.db)

	t.Setenv("WT_DATABASE_DRIVER", "memory")
	require.NoError(t, a.Configure())

	assert.Equal(t, "memory", a.Config.DatabaseDriver)
	assert.NotNil(t, a.db)
}

func TestApp_ConfigureLoginCustomTextFromEnv(t *testing.T) {
	t.Setenv("WT_DATABASE_DRIVER", "memory")
	t.Setenv("WT_LOGIN_CUSTOM_TEXT", "# Welcome\n\nPlease contact the administrator.")

	a := defaultApp(t)
	require.NoError(t, a.Configure())

	assert.Equal(t, "# Welcome\n\nPlease contact the administrator.", a.Config.LoginCustomText)
	assert.True(t, a.Config.Overrides.LoginCustomText)
	assert.False(t, a.Config.Overrides.RegistrationDisabled)
	assert.False(t, a.Config.Overrides.SocialsDisabled)
}

func TestApp_ConfigureLoginCustomTextFromConfigFile(t *testing.T) {
	t.Chdir(t.TempDir())
	require.NoError(t, os.WriteFile(
		filepath.Join(".", "workout-tracker.yaml"),
		[]byte("login_custom_text: \"# From the config file\"\n"),
		0o600,
	))

	t.Setenv("WT_DATABASE_DRIVER", "memory")

	a := defaultApp(t)
	require.NoError(t, a.Configure())

	assert.Equal(t, "# From the config file", a.Config.LoginCustomText)
	assert.True(t, a.Config.Overrides.LoginCustomText)
	assert.False(t, a.Config.Overrides.RegistrationDisabled)
	assert.False(t, a.Config.Overrides.SocialsDisabled)
}

func TestApp_LoginCustomTextDatabase(t *testing.T) {
	a := configuredApp(t)
	a.Config.LoginCustomText = "# Stored welcome"

	require.NoError(t, a.Config.Save(a.db))
	require.NoError(t, a.ResetConfiguration())

	assert.Equal(t, "# Stored welcome", a.Config.LoginCustomText)
}

func TestApp_NewLogger(t *testing.T) {
	l := newLogger(false)
	assert.IsType(t, slognil.Handler{}, l.Handler())

	l = newLogger(true)
	assert.IsType(t, &slog.JSONHandler{}, l.Handler())
}

func TestApp_RandomJWTErrorIdemPotent(t *testing.T) {
	a := defaultApp(t)
	s1 := a.jwtSecret()
	assert.NotEmpty(t, s1)

	s2 := a.jwtSecret()
	assert.Equal(t, s1, s2)
}
