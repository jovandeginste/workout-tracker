package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	session "github.com/spazzymoto/echo-scs-session"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func configuredApp(t *testing.T) *App {
	t.Setenv("WT_DATABASE_DRIVER", "memory")

	a := defaultApp(t)
	require.NoError(t, a.Configure())

	return a
}

func TestRoute_NoAccess(t *testing.T) {
	a := configuredApp(t)

	e := a.echo

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	s := session.LoadAndSave(a.sessionManager)
	h := s(a.dashboardHandler)

	// Assertions
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusFound, rec.Code)
	}
}
