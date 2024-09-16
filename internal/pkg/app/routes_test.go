package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jovandeginste/workout-tracker/internal/database"
	session "github.com/spazzymoto/echo-scs-session"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func configuredApp(t *testing.T) *App {
	t.Helper()
	t.Setenv("WT_DATABASE_DRIVER", "memory")

	a := defaultApp(t)

	t.Run("should self-configure", func(t *testing.T) {
		require.NoError(t, a.Configure())
	})

	return a
}

func defaultUser(db *gorm.DB) *database.User {
	u := &database.User{
		Username: "my-username",
		Password: "my-password",
		Name:     "my-name",
	}

	u.SetDB(db)

	return u
}

func TestRoute_UserRender(t *testing.T) {
	t.Run("should render for the user", func(t *testing.T) {
		a := configuredApp(t)

		e := a.echo

		req := httptest.NewRequest(http.MethodGet, e.Reverse("dashboard"), nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user_info", defaultUser(a.db))

		s := session.LoadAndSave(a.sessionManager)
		h := s(a.dashboardHandler)

		require.NoError(t, h(c))
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Dashboard for my-name")
	})
}

func TestRoute_UserRenderLang(t *testing.T) {
	langTests := map[string]string{
		"en": "Dashboard for",
		"nl": "Dashboard voor",
	}

	for lang, expected := range langTests {
		t.Run("should render in "+lang+" for the user", func(t *testing.T) {
			a := configuredApp(t)

			e := a.echo

			req := httptest.NewRequest(http.MethodGet, e.Reverse("dashboard"), nil)
			rec := httptest.NewRecorder()

			req.Header.Set("Accept-Language", lang)

			c := e.NewContext(req, rec)
			c.Set("user_info", defaultUser(a.db))

			s := session.LoadAndSave(a.sessionManager)
			h := s(a.dashboardHandler)

			require.NoError(t, h(c))
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), expected+" my-name")
		})
	}
}

func TestRoute_NoUserRedirect(t *testing.T) {
	t.Run("should redirect", func(t *testing.T) {
		a := configuredApp(t)

		e := a.echo

		req := httptest.NewRequest(http.MethodGet, e.Reverse("dashboard"), nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		s := session.LoadAndSave(a.sessionManager)
		h := s(a.dashboardHandler)

		require.Error(t, h(c))
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestRoute_NoUserAccessLogin(t *testing.T) {
	t.Run("should render a login page", func(t *testing.T) {
		a := configuredApp(t)

		e := a.echo

		req := httptest.NewRequest(http.MethodGet, e.Reverse("user-login"), nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		s := session.LoadAndSave(a.sessionManager)
		h := s(a.userLoginHandler)

		require.NoError(t, h(c))
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `<button id="signin" type="submit">`)
	})
}

func TestRoute_NoUserAccessLoginLang(t *testing.T) {
	langTests := map[string]string{
		"en": "Sign in",
		"nl": "Aanmelden",
	}

	for lang, expected := range langTests {
		t.Run("should render login page in "+lang, func(t *testing.T) {
			a := configuredApp(t)

			e := a.echo

			req := httptest.NewRequest(http.MethodGet, e.Reverse("user-login"), nil)
			rec := httptest.NewRecorder()

			req.Header.Set("Accept-Language", lang)

			c := e.NewContext(req, rec)

			s := session.LoadAndSave(a.sessionManager)
			h := s(a.userLoginHandler)

			require.NoError(t, h(c))
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), expected)
		})
	}
}
