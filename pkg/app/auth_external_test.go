package app

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jovandeginste/workout-tracker/v2/pkg/auth"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type fakeLDAP struct {
	identity *auth.Identity
	err      error
}

func (f *fakeLDAP) Authenticate(context.Context, string, string) (*auth.Identity, error) {
	return f.identity, f.err
}

func loginRequest(t *testing.T, a *App, username, password string) *httptest.ResponseRecorder {
	t.Helper()

	form := url.Values{"username": {username}, "password": {password}}
	req := httptest.NewRequest(http.MethodPost, a.Reverse("user-signin"), strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	a.echo.ServeHTTP(rec, req)

	return rec
}

func adminTokenCookie(t *testing.T, a *App) *http.Cookie {
	t.Helper()

	admin, err := database.GetUser(a.db, "admin")
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := a.echo.NewContext(req, rec)
	require.NoError(t, a.createToken(admin, c))

	cookies := rec.Result().Cookies()
	require.Len(t, cookies, 1)

	return cookies[0]
}

func adminFormRequest(t *testing.T, a *App, target string, form url.Values) *httptest.ResponseRecorder {
	t.Helper()

	req := httptest.NewRequest(http.MethodPost, target, strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	req.AddCookie(adminTokenCookie(t, a))
	rec := httptest.NewRecorder()

	a.echo.ServeHTTP(rec, req)

	return rec
}

func adminGetRequest(t *testing.T, a *App, target string) *httptest.ResponseRecorder {
	t.Helper()

	req := httptest.NewRequest(http.MethodGet, target, nil)
	req.AddCookie(adminTokenCookie(t, a))
	rec := httptest.NewRecorder()

	a.echo.ServeHTTP(rec, req)

	return rec
}

const testLocalPassword = "local-password"

func createLocalUser(t *testing.T, a *App, username string) *database.User {
	t.Helper()

	u := defaultUser(a.db)
	u.Username = username
	u.Active = true
	require.NoError(t, u.SetPassword(testLocalPassword))
	require.NoError(t, u.Create(a.db))

	return u
}

func createLDAPUser(t *testing.T, a *App, username string, active bool) *database.User {
	t.Helper()

	u := &database.User{
		UserData: database.UserData{
			Username:   username,
			Name:       username,
			AuthSource: database.AuthSourceLDAP,
			Active:     active,
		},
	}
	require.NoError(t, u.Create(a.db))

	return u
}

func TestConfigureAuthenticationLDAP(t *testing.T) {
	t.Setenv("WT_LDAP_ENABLED", "true")
	t.Setenv("WT_LDAP_URL", "ldaps://ldap.example.com:636")
	t.Setenv("WT_LDAP_USER_DN", "uid={username},ou=people,dc=example,dc=com")

	a := configuredApp(t)
	assert.NotNil(t, a.ldapAuth)
}

func TestConfigureAuthenticationLDAPPasswordFile(t *testing.T) {
	passwordFile := filepath.Join(t.TempDir(), "ldap-password")
	require.NoError(t, os.WriteFile(passwordFile, []byte("service-password\n"), 0o600))

	t.Setenv("WT_LDAP_ENABLED", "true")
	t.Setenv("WT_LDAP_URL", "ldaps://ldap.example.com:636")
	t.Setenv("WT_LDAP_BIND_DN", "cn=workout-tracker,ou=services,dc=example,dc=com")
	t.Setenv("WT_LDAP_BIND_PASSWORD_FILE", passwordFile)
	t.Setenv("WT_LDAP_BASE_DN", "ou=people,dc=example,dc=com")
	t.Setenv("WT_LDAP_USER_FILTER", "(uid={username})")

	a := configuredApp(t)
	assert.NotNil(t, a.ldapAuth)
}

func TestConfigureAuthenticationLDAPInvalid(t *testing.T) {
	t.Setenv("WT_LDAP_ENABLED", "true")

	a := defaultApp(t)
	t.Setenv("WT_DATABASE_DRIVER", "memory")
	require.Error(t, a.Configure())
}

func TestConfigureAuthenticationLDAPPasswordFileMissing(t *testing.T) {
	t.Setenv("WT_LDAP_ENABLED", "true")
	t.Setenv("WT_LDAP_URL", "ldaps://ldap.example.com:636")
	t.Setenv("WT_LDAP_BIND_DN", "cn=workout-tracker,ou=services,dc=example,dc=com")
	t.Setenv("WT_LDAP_BIND_PASSWORD_FILE", "/does/not/exist")
	t.Setenv("WT_LDAP_BASE_DN", "ou=people,dc=example,dc=com")
	t.Setenv("WT_LDAP_USER_FILTER", "(uid={username})")

	a := defaultApp(t)
	t.Setenv("WT_DATABASE_DRIVER", "memory")
	require.Error(t, a.Configure())
}

func TestLDAPLoginCreatesUserWhenAutoCreateEnabled(t *testing.T) {
	a := configuredApp(t)
	a.Config.LDAPAutoCreate = true
	a.Config.LDAPDefaultActive = true
	a.ldapAuth = &fakeLDAP{identity: &auth.Identity{Username: "ldap-user", Name: "LDAP User"}}

	rec := loginRequest(t, a, "ldap-user", "secret")
	require.Equal(t, http.StatusFound, rec.Code)
	assert.Equal(t, a.Reverse("dashboard"), rec.Header().Get("Location"))
	require.NotEmpty(t, rec.Result().Cookies())

	u, err := database.GetUser(a.db, "ldap-user")
	require.NoError(t, err)
	assert.Equal(t, database.AuthSourceLDAP, u.AuthSource)
	assert.True(t, u.Active)
	assert.Equal(t, "LDAP User", u.Name)
	assert.Equal(t, BrowserLanguage, u.Profile.Language)
	assert.Empty(t, u.Password)
}

func TestLDAPLoginDoesNotTakeOverLocalUser(t *testing.T) {
	a := configuredApp(t)

	u := createLocalUser(t, a, "jo")
	a.ldapAuth = &fakeLDAP{identity: &auth.Identity{Username: "jo", Name: "Jo from LDAP"}}

	rec := loginRequest(t, a, "jo", "ldap-password")
	require.Equal(t, http.StatusFound, rec.Code)
	assert.Equal(t, a.Reverse("user-login"), rec.Header().Get("Location"))

	stored, err := database.GetUser(a.db, "jo")
	require.NoError(t, err)
	assert.Equal(t, database.AuthSourceLocal, stored.AuthSource)
	assert.True(t, stored.ValidLogin("local-password"))
	assert.NotEqual(t, "Jo from LDAP", stored.Name)
	assert.NotEmpty(t, u.Password)
}

func TestLDAPLoginUnknownUserWhenAutoCreateDisabled(t *testing.T) {
	a := configuredApp(t)
	a.Config.LDAPAutoCreate = false
	a.ldapAuth = &fakeLDAP{identity: &auth.Identity{Username: "ldap-user", Name: "LDAP User"}}

	rec := loginRequest(t, a, "ldap-user", "secret")
	require.Equal(t, http.StatusFound, rec.Code)
	assert.Equal(t, a.Reverse("user-login"), rec.Header().Get("Location"))

	_, err := database.GetUser(a.db, "ldap-user")
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
}

func TestLDAPLoginAutoCreateInactiveUser(t *testing.T) {
	a := configuredApp(t)
	a.Config.LDAPAutoCreate = true
	a.Config.LDAPDefaultActive = false
	a.ldapAuth = &fakeLDAP{identity: &auth.Identity{Username: "ldap-user", Name: "LDAP User"}}

	rec := loginRequest(t, a, "ldap-user", "secret")
	require.Equal(t, http.StatusFound, rec.Code)
	assert.Equal(t, a.Reverse("user-login"), rec.Header().Get("Location"))

	u, err := database.GetUser(a.db, "ldap-user")
	require.NoError(t, err)
	assert.False(t, u.Active)
}

func TestLDAPLoginInactiveUser(t *testing.T) {
	a := configuredApp(t)

	createLDAPUser(t, a, "ldap-user", false)
	a.ldapAuth = &fakeLDAP{identity: &auth.Identity{Username: "ldap-user", Name: "LDAP User"}}

	rec := loginRequest(t, a, "ldap-user", "secret")
	require.Equal(t, http.StatusFound, rec.Code)
	assert.Equal(t, a.Reverse("user-login"), rec.Header().Get("Location"))
}

func TestLDAPLoginSyncsAdminGroup(t *testing.T) {
	a := configuredApp(t)
	a.Config.LDAPAdminGroup = "cn=workout-tracker-admins,ou=groups,dc=example,dc=com"

	u := createLDAPUser(t, a, "ldap-user", true)
	a.ldapAuth = &fakeLDAP{identity: &auth.Identity{Username: "ldap-user", Name: "LDAP User", Admin: true}}

	rec := loginRequest(t, a, "ldap-user", "secret")
	require.Equal(t, http.StatusFound, rec.Code)
	assert.Equal(t, a.Reverse("dashboard"), rec.Header().Get("Location"))

	stored, err := database.GetUser(a.db, "ldap-user")
	require.NoError(t, err)
	assert.True(t, stored.Admin)
	assert.Equal(t, "LDAP User", stored.Name)
	assert.Equal(t, u.ID, stored.ID)
}

func TestLocalLoginStillWorksWithLDAPEnabled(t *testing.T) {
	a := configuredApp(t)
	a.ldapAuth = &fakeLDAP{identity: &auth.Identity{Username: "someone", Name: "Someone"}}

	createLocalUser(t, a, "local-user")

	rec := loginRequest(t, a, "local-user", "local-password")
	require.Equal(t, http.StatusFound, rec.Code)
	assert.Equal(t, a.Reverse("dashboard"), rec.Header().Get("Location"))
}

func TestLocalLoginIsCaseSensitive(t *testing.T) {
	a := configuredApp(t)
	createLocalUser(t, a, "CaseUser")

	rec := loginRequest(t, a, "caseuser", "local-password")
	require.Equal(t, http.StatusFound, rec.Code)
	assert.Equal(t, a.Reverse("user-login"), rec.Header().Get("Location"))

	rec = loginRequest(t, a, "CaseUser", "local-password")
	require.Equal(t, http.StatusFound, rec.Code)
	assert.Equal(t, a.Reverse("dashboard"), rec.Header().Get("Location"))
}

func TestAdminNewUserPageShowsAuthSourcePills(t *testing.T) {
	a := configuredApp(t)

	rec := adminGetRequest(t, a, a.Reverse("admin-user-new"))
	require.Equal(t, http.StatusOK, rec.Code)

	body := rec.Body.String()
	assert.Contains(t, body, `name="auth_source"`)
	assert.Contains(t, body, `value="local"`)
	assert.Contains(t, body, `value="ldap"`)
}

func TestAdminCreateLDAPUser(t *testing.T) {
	a := configuredApp(t)

	rec := adminFormRequest(t, a, a.Reverse("admin-user-create"), url.Values{
		"username":    {"ldap-new"},
		"name":        {"LDAP New"},
		"auth_source": {database.AuthSourceLDAP},
		"active":      {"on"},
	})
	require.Equal(t, http.StatusFound, rec.Code)

	u, err := database.GetUser(a.db, "ldap-new")
	require.NoError(t, err)
	assert.Equal(t, database.AuthSourceLDAP, u.AuthSource)
	assert.True(t, u.Active)
	assert.False(t, u.Admin)
	assert.Empty(t, u.Password)
}

func TestAdminCreateLocalUserRequiresPassword(t *testing.T) {
	a := configuredApp(t)

	rec := adminFormRequest(t, a, a.Reverse("admin-user-create"), url.Values{
		"username":    {"local-new"},
		"name":        {"Local New"},
		"auth_source": {database.AuthSourceLocal},
		"active":      {"on"},
	})
	require.Equal(t, http.StatusFound, rec.Code)
	assert.Equal(t, a.Reverse("admin-user-new"), rec.Header().Get("Location"))

	_, err := database.GetUser(a.db, "local-new")
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
}

func TestAdminConvertLocalUserToLDAP(t *testing.T) {
	a := configuredApp(t)
	u := createLocalUser(t, a, "convert-me")

	rec := adminFormRequest(t, a, a.Reverse("admin-user-update", u.ID), url.Values{
		"username":    {u.Username},
		"name":        {u.Name},
		"auth_source": {database.AuthSourceLDAP},
		"active":      {"on"},
	})
	require.Equal(t, http.StatusFound, rec.Code)

	stored, err := database.GetUser(a.db, u.Username)
	require.NoError(t, err)
	assert.Equal(t, database.AuthSourceLDAP, stored.AuthSource)
	assert.Empty(t, stored.Password)
	assert.False(t, stored.ValidLogin("local-password"))
}

func TestAdminConvertLDAPUserToLocalRequiresPassword(t *testing.T) {
	a := configuredApp(t)
	u := createLDAPUser(t, a, "convert-me", true)

	rec := adminFormRequest(t, a, a.Reverse("admin-user-update", u.ID), url.Values{
		"username":    {u.Username},
		"name":        {u.Name},
		"auth_source": {database.AuthSourceLocal},
		"active":      {"on"},
	})
	require.Equal(t, http.StatusFound, rec.Code)

	stored, err := database.GetUser(a.db, u.Username)
	require.NoError(t, err)
	assert.Equal(t, database.AuthSourceLDAP, stored.AuthSource)
}

func TestAdminConvertLDAPUserToLocal(t *testing.T) {
	a := configuredApp(t)
	u := createLDAPUser(t, a, "convert-me", true)

	rec := adminFormRequest(t, a, a.Reverse("admin-user-update", u.ID), url.Values{
		"username":    {u.Username},
		"name":        {u.Name},
		"auth_source": {database.AuthSourceLocal},
		"password":    {"new-password"},
		"active":      {"on"},
	})
	require.Equal(t, http.StatusFound, rec.Code)

	stored, err := database.GetUser(a.db, u.Username)
	require.NoError(t, err)
	assert.Equal(t, database.AuthSourceLocal, stored.AuthSource)
	assert.True(t, stored.ValidLogin("new-password"))
}
