package app

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/auth"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func (a *App) ConfigureAuthentication() error {
	if !a.Config.LDAPEnabled {
		return nil
	}

	password := a.Config.LDAPBindPassword
	if password == "" && a.Config.LDAPBindPasswordFile != "" {
		content, err := os.ReadFile(a.Config.LDAPBindPasswordFile)
		if err != nil {
			return fmt.Errorf("read ldap_bind_password_file: %w", err)
		}

		password = strings.TrimRight(string(content), "\r\n")
	}

	authenticator, err := auth.NewLDAP(auth.LDAPConfig{
		URL:                a.Config.LDAPURL,
		StartTLS:           a.Config.LDAPStartTLS,
		InsecureSkipVerify: a.Config.LDAPInsecureSkipVerify,
		CAFile:             a.Config.LDAPCAFile,
		BindDN:             a.Config.LDAPBindDN,
		BindPassword:       password,
		UserDN:             a.Config.LDAPUserDN,
		BaseDN:             a.Config.LDAPBaseDN,
		UserFilter:         a.Config.LDAPUserFilter,
		UsernameAttribute:  a.Config.LDAPUsernameAttribute,
		NameAttribute:      a.Config.LDAPNameAttribute,
		RequiredGroup:      a.Config.LDAPRequiredGroup,
		AdminGroup:         a.Config.LDAPAdminGroup,
		Timeout:            time.Duration(a.Config.LDAPTimeoutSeconds) * time.Second,
	})
	if err != nil {
		return err
	}

	a.ldapAuth = authenticator

	a.logger.Info("LDAP authentication enabled", "url", a.Config.LDAPURL)

	if a.Config.LDAPInsecureSkipVerify {
		a.logger.Warn("LDAP TLS certificate verification is disabled")
	}

	if !a.Config.RegistrationDisabled {
		a.logger.Warn("LDAP is enabled but registration is not disabled; local registrations can shadow LDAP usernames")
	}

	return nil
}

func (a *App) loginWithLDAP(c *echo.Context, username, password string) (*database.User, error) {
	if a.ldapAuth == nil {
		return nil, ErrLoginFailed
	}

	identity, err := a.ldapAuth.Authenticate(c.Request().Context(), username, password)
	if err != nil {
		a.logger.Warn("LDAP authentication failed", "username", username, "error", err)
		return nil, ErrLoginFailed
	}

	u, err := database.GetUser(a.db, identity.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u, err = database.GetUserFold(a.db, identity.Username)
	}

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			a.logger.Error("LDAP login: lookup user", "error", err)
			return nil, ErrLoginFailed
		}

		return a.createLDAPUser(identity)
	}

	return a.updateLDAPUser(u, identity)
}

func (a *App) createLDAPUser(identity *auth.Identity) (*database.User, error) {
	if !a.Config.LDAPAutoCreate {
		a.logger.Warn("LDAP login for unknown user; auto-create is disabled", "username", identity.Username)
		return nil, ErrLoginFailed
	}

	u := &database.User{
		UserData: database.UserData{
			Username:   identity.Username,
			Name:       identity.Name,
			AuthSource: database.AuthSourceLDAP,
			Active:     a.Config.LDAPDefaultActive,
			Admin:      identity.Admin,
		},
	}

	applyUserDefaults(u)
	u.Profile.User = u

	if err := u.Create(a.db); err != nil {
		if !errors.Is(err, gorm.ErrDuplicatedKey) {
			a.logger.Error("LDAP login: create user", "error", err)
			return nil, ErrLoginFailed
		}

		u, err = database.GetUser(a.db, identity.Username)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u, err = database.GetUserFold(a.db, identity.Username)
		}

		if err != nil {
			a.logger.Error("LDAP login: reload user after duplicate", "error", err)
			return nil, ErrLoginFailed
		}

		return a.updateLDAPUser(u, identity)
	}

	if !u.IsActive() {
		return nil, ErrLoginFailed
	}

	return u, nil
}

func (a *App) updateLDAPUser(u *database.User, identity *auth.Identity) (*database.User, error) {
	if u.AuthSource != database.AuthSourceLDAP {
		a.logger.Warn("LDAP login refused for local account", "username", u.Username)
		return nil, ErrLoginFailed
	}

	if !u.IsActive() {
		return nil, ErrLoginFailed
	}

	changed := false

	if u.Name != identity.Name {
		u.Name = identity.Name
		changed = true
	}

	if a.Config.LDAPAdminGroup != "" && u.Admin != identity.Admin {
		u.Admin = identity.Admin
		changed = true
	}

	if changed {
		if err := u.Save(a.db); err != nil {
			a.logger.Error("LDAP login: update user", "error", err)
			return nil, ErrLoginFailed
		}
	}

	return u, nil
}

func (a *App) finishLogin(c *echo.Context, u *database.User) error {
	a.sessionManager.Put(c.Request().Context(), "username", u.Username)

	if err := a.createToken(u, c); err != nil {
		a.logger.Error("login: create token", "error", err)
		return a.redirectWithError(c, a.Reverse("user-login"), ErrLoginFailed)
	}

	return c.Redirect(http.StatusFound, a.Reverse("dashboard"))
}
