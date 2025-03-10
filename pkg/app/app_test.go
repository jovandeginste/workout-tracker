package app

import (
	"log/slog"
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
