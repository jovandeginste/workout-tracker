package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func configuredApp(t *testing.T) *App {
	t.Helper()
	t.Setenv("WT_DATABASE_DRIVER", "memory")

	a := defaultApp(t)

	require.NoError(t, a.Configure())

	return a
}

func defaultUser(db *gorm.DB) *database.User {
	u := &database.User{
		UserData: database.UserData{
			Username: "my-username",
			Name:     "my-name",
			Active:   true,
		},
		UserSecrets: database.UserSecrets{
			Password: "my-password",
		},
	}

	u.SetDB(db)

	return u
}

func TestRoute_HealthCheck(t *testing.T) {
	t.Run("should pass health check", func(t *testing.T) {
		a := configuredApp(t)
		req := httptest.NewRequest(http.MethodGet, a.Reverse("health"), nil)
		rec := httptest.NewRecorder()

		a.echo.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "OK", rec.Body.String())
	})
}

func TestRoute_DashboardWebRoot(t *testing.T) {
	tests := []struct {
		name    string
		webRoot string
		path    string
	}{
		{name: "default", path: "/"},
		{name: "root slash", webRoot: "/", path: "/"},
		{name: "prefix", webRoot: "/wot", path: "/wot"},
		{name: "prefix trailing slash", webRoot: "/wot/", path: "/wot"},
		{name: "nested prefix", webRoot: "fitness/workouts", path: "/fitness/workouts"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("WT_WEB_ROOT", tt.webRoot)
			a := configuredApp(t)
			assert.Equal(t, tt.path, a.Reverse("dashboard"))

			u := defaultUser(a.db)
			require.NoError(t, u.Save(a.db))
			tokenReq := httptest.NewRequest(http.MethodGet, tt.path, nil)
			tokenRec := httptest.NewRecorder()
			require.NoError(t, a.createToken(u, a.echo.NewContext(tokenReq, tokenRec)))
			tokenResponse := tokenRec.Result()
			defer tokenResponse.Body.Close()
			cookies := tokenResponse.Cookies()
			require.Len(t, cookies, 1)

			paths := []string{tt.path}
			if tt.path != "/" {
				paths = append(paths, tt.path+"/")
			}

			for _, requestPath := range paths {
				t.Run(requestPath, func(t *testing.T) {
					req := httptest.NewRequest(http.MethodGet, requestPath, nil)
					rec := httptest.NewRecorder()
					a.echo.ServeHTTP(rec, req)
					assert.Equal(t, http.StatusFound, rec.Code)
					assert.Equal(t, a.Reverse("user-signout"), rec.Header().Get("Location"))

					req = httptest.NewRequest(http.MethodGet, requestPath, nil)
					req.AddCookie(cookies[0])
					rec = httptest.NewRecorder()
					a.echo.ServeHTTP(rec, req)
					assert.Equal(t, http.StatusOK, rec.Code)
					assert.Contains(t, rec.Body.String(), "Dashboard for my-name")
				})
			}
		})
	}
}

func TestRoute_UserRender(t *testing.T) {
	t.Run("should render for the user", func(t *testing.T) {
		a := configuredApp(t)

		req := httptest.NewRequest(http.MethodGet, a.Reverse("dashboard"), nil)
		rec := httptest.NewRecorder()

		c := a.echo.NewContext(req, rec)
		a.setContext(c)
		c.Set("user_info", defaultUser(a.db))

		s := sessionLoadAndSave(a.sessionManager)
		h := s(a.dashboardHandler)

		require.NoError(t, h(c))
		assert.Empty(t, c.Get("errors"))
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

			req := httptest.NewRequest(http.MethodGet, a.Reverse("dashboard"), nil)
			rec := httptest.NewRecorder()

			req.Header.Set("Accept-Language", lang)

			c := a.echo.NewContext(req, rec)
			a.setContext(c)
			c.Set("user_info", defaultUser(a.db))

			s := sessionLoadAndSave(a.sessionManager)
			h := s(a.dashboardHandler)

			require.NoError(t, h(c))
			assert.Empty(t, c.Get("errors"))
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), expected+" my-name")
		})
	}
}

func TestRoute_NoUserRedirect(t *testing.T) {
	t.Run("should redirect", func(t *testing.T) {
		a := configuredApp(t)

		req := httptest.NewRequest(http.MethodGet, a.Reverse("dashboard"), nil)
		rec := httptest.NewRecorder()

		c := a.echo.NewContext(req, rec)
		a.setContext(c)
		s := sessionLoadAndSave(a.sessionManager)
		h := s(a.dashboardHandler)

		require.NoError(t, h(c))
		assert.NotEmpty(t, a.sessionManager.Get(c.Request().Context(), "errors"))
		assert.Equal(t, http.StatusFound, rec.Code)
	})
}

func TestRoute_NoUserAccessLogin(t *testing.T) {
	t.Run("should render a login page", func(t *testing.T) {
		a := configuredApp(t)

		req := httptest.NewRequest(http.MethodGet, a.Reverse("user-login"), nil)
		rec := httptest.NewRecorder()

		c := a.echo.NewContext(req, rec)
		a.setContext(c)

		s := sessionLoadAndSave(a.sessionManager)
		h := s(a.userLoginHandler)

		require.NoError(t, h(c))
		assert.Empty(t, c.Get("errors"))
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

			req := httptest.NewRequest(http.MethodGet, a.Reverse("user-login"), nil)
			rec := httptest.NewRecorder()

			req.Header.Set("Accept-Language", lang)

			c := a.echo.NewContext(req, rec)
			a.setContext(c)

			s := sessionLoadAndSave(a.sessionManager)
			h := s(a.userLoginHandler)

			require.NoError(t, h(c))
			assert.Empty(t, c.Get("errors"))
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), expected)
		})
	}
}

func TestLangFromContextString_PrefixMatching(t *testing.T) {
	a := configuredApp(t)

	tests := []struct {
		acceptLanguage string
		expectedMatch  string
	}{
		// Direct matches
		{"nl", "nl"},
		{"en", "en"},
		{"pt-BR", "pt-BR"},
		{"zh-Hans", "zh-Hans"},
		{"nb-NO", "nb-NO"},

		// Prefix matches (e.g. requesting nl-BE when only nl is supported, or pt-PT when only pt is supported, or pt-BR, etc.)
		{"nl-BE", "nl"}, // nl-BE is not directly supported, matches prefix "nl"
		{"en-GB", "en"}, // en-GB matches prefix "en"
		{"es-MX", "es"}, // es-MX matches prefix "es"
		{"de-CH", "de"}, // de-CH matches prefix "de"
		{"pt-PT", "pt"}, // pt-PT matches prefix "pt" (since "pt" is in helpers.SupportedLanguages)

		// Ordered priority (first match wins or is preferred)
		{"nl-BE,en-US;q=0.8", "nl,en"},
		{"fr-CH,nl-NL;q=0.9", "fr,nl"},
	}

	for _, tt := range tests {
		t.Run(tt.acceptLanguage, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("Accept-Language", tt.acceptLanguage)
			rec := httptest.NewRecorder()
			c := a.echo.NewContext(req, rec)

			actual := a.langFromContextString(c)
			assert.Equal(t, tt.expectedMatch, actual)
		})
	}
}

func TestRoute_UserExportAll(t *testing.T) {
	t.Run("should export all files as a zip", func(t *testing.T) {
		a := configuredApp(t)

		u := defaultUser(a.db)
		require.NoError(t, u.Save(a.db))

		workout := &database.Workout{
			UserID: u.ID,
			Name:   "Morning Run",
			GPX: &database.GPXData{
				Filename: "run.gpx",
				Content:  []byte("<gpx>test</gpx>"),
				Checksum: []byte("checksum-123"),
			},
		}
		require.NoError(t, a.db.Save(workout).Error)

		req := httptest.NewRequest(http.MethodGet, a.Reverse("user-export-all"), nil)
		rec := httptest.NewRecorder()

		c := a.echo.NewContext(req, rec)
		a.setContext(c)
		c.Set("user_info", u)

		s := sessionLoadAndSave(a.sessionManager)
		h := s(a.userExportAllHandler)

		require.NoError(t, h(c))
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/zip", rec.Header().Get("Content-Type"))
		assert.Contains(t, rec.Header().Get("Content-Disposition"), "attachment; filename=\"workouts-export.zip\"")

		bodyBytes := rec.Body.Bytes()
		assert.NotEmpty(t, bodyBytes)
	})
}
