package app

import (
	"errors"
	"net/http"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/views/admin"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

var ErrPasswordRequired = errors.New("a password is required for local accounts")

func (a *App) adminRoutes(g *echo.Group) *echo.Group {
	adminGroup := g.Group("/admin")
	adminGroup.Use(a.ValidateAdminMiddleware)

	a.GET(adminGroup, "", a.adminRootHandler, "admin")
	a.POST(adminGroup, "/config", a.adminConfigUpdateHandler, "admin-config-update")

	adminUsersGroup := adminGroup.Group("/users")
	a.GET(adminUsersGroup, "/new", a.adminUserNewHandler, "admin-user-new")
	a.POST(adminUsersGroup, "", a.adminUserCreateHandler, "admin-user-create")
	a.GET(adminUsersGroup, "/:id/edit", a.adminUserEditHandler, "admin-user-edit")
	a.POST(adminUsersGroup, "/:id", a.adminUserUpdateHandler, "admin-user-update")
	a.GET(adminUsersGroup, "/:id/delete", a.adminUserDeleteConfirmHandler, "admin-user-delete-confirm")
	a.POST(adminUsersGroup, "/:id/delete", a.adminUserDeleteHandler, "admin-user-delete")
	a.GET(adminUsersGroup, "/:id", func(c *echo.Context) error {
		return c.Redirect(http.StatusFound, a.Reverse("admin-user-edit", c.Param("id")))
	}, "admin-user-show")

	return adminGroup
}

func (a *App) adminRootHandler(c *echo.Context) error {
	users, err := database.GetUsers(a.db)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("dashboard"), err)
	}

	return Render(c, http.StatusOK, admin.Root(users))
}

func (a *App) adminUserNewHandler(c *echo.Context) error {
	u := &database.User{
		UserData: database.UserData{
			AuthSource: database.AuthSourceLocal,
			Active:     true,
		},
	}

	return Render(c, http.StatusOK, admin.NewUser(u))
}

func (a *App) adminUserCreateHandler(c *echo.Context) error {
	u := &database.User{
		UserData: database.UserData{
			Username:   c.FormValue("username"),
			Name:       c.FormValue("name"),
			AuthSource: c.FormValue("auth_source"),
			Admin:      isChecked(c.FormValue("admin")),
			Active:     isChecked(c.FormValue("active")),
		},
	}

	if u.AuthSource == "" {
		u.AuthSource = database.AuthSourceLocal
	}

	switch u.AuthSource {
	case database.AuthSourceLocal:
		password := c.FormValue("password")
		if password == "" {
			return a.redirectWithError(c, a.Reverse("admin-user-new"), ErrPasswordRequired)
		}

		if err := u.SetPassword(password); err != nil {
			return a.redirectWithError(c, a.Reverse("admin-user-new"), err)
		}
	case database.AuthSourceLDAP:
		u.Password = ""
	default:
		return a.redirectWithError(c, a.Reverse("admin-user-new"), database.ErrInvalidAuthSource)
	}

	applyUserDefaults(u)
	u.Profile.User = u

	if err := u.Create(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("admin-user-new"), err)
	}

	a.addNoticeT(c, "user.notice.created", u.Name)

	return c.Redirect(http.StatusFound, a.Reverse("admin-user-edit", u.ID))
}

func (a *App) adminUserEditHandler(c *echo.Context) error {
	user, err := a.getUser(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("admin"), err)
	}

	if user == nil {
		return a.redirectWithError(c, a.Reverse("admin"), ErrUserNotFound)
	}

	return Render(c, http.StatusOK, admin.EditUser(user))
}

func (a *App) adminUserUpdateHandler(c *echo.Context) error {
	u, err := a.getUser(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("admin"), err)
	}

	u.Name = c.FormValue("name")
	u.Username = c.FormValue("username")
	u.Admin = isChecked(c.FormValue("admin"))
	u.Active = isChecked(c.FormValue("active"))

	if err := updateUserAuthSource(u, c.FormValue("auth_source"), c.FormValue("password")); err != nil {
		return a.redirectWithError(c, a.Reverse("admin-user-show", c.Param("id")), err)
	}

	if err := u.Save(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("admin-user-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "user.notice.updated", u.Name)

	return c.Redirect(http.StatusFound, a.Reverse("admin-user-show", c.Param("id")))
}

func updateUserAuthSource(u *database.User, source, password string) error {
	current := u.AuthSource
	if current == "" {
		current = database.AuthSourceLocal
	}

	if source == "" {
		source = current
	}

	switch source {
	case database.AuthSourceLocal:
		if current == database.AuthSourceLDAP && password == "" {
			return ErrPasswordRequired
		}

		u.AuthSource = database.AuthSourceLocal

		if password != "" {
			return u.SetPassword(password)
		}
	case database.AuthSourceLDAP:
		u.AuthSource = database.AuthSourceLDAP
		u.Password = ""
		u.Salt = ""
	default:
		return database.ErrInvalidAuthSource
	}

	return nil
}

func (a *App) adminUserDeleteHandler(c *echo.Context) error { //nolint:dupl
	u, err := a.getUser(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("admin"), err)
	}

	if err := u.Delete(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("admin-user-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "user.notice.deleted", u.Name)

	return a.redirectHTMXOrFound(c, a.Reverse("admin"))
}

func (a *App) adminConfigUpdateHandler(c *echo.Context) error {
	cnf, err := database.GetConfig(a.db)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return a.redirectWithError(c, a.Reverse("admin"), err)
		}

		cnf = &database.Config{}
	}

	// Unchecked checkboxes are not submitted, so reset them explicitly. Fields
	// that are overridden keep their stored value.
	if !a.Config.Overrides.RegistrationDisabled {
		cnf.RegistrationDisabled = false
	}

	if !a.Config.Overrides.SocialsDisabled {
		cnf.SocialsDisabled = false
	}

	if err := c.Bind(cnf); err != nil {
		return a.redirectWithError(c, a.Reverse("admin"), err)
	}

	if err := cnf.Save(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("admin"), err)
	}

	if err := a.ResetConfiguration(); err != nil {
		return a.redirectWithError(c, a.Reverse("admin"), err)
	}

	a.addNoticeT(c, "settings.updated")

	return c.Redirect(http.StatusFound, a.Reverse("admin"))
}

func isChecked(value string) bool {
	return value == "on"
}

func (a *App) adminUserDeleteConfirmHandler(c *echo.Context) error {
	u, err := a.getUser(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("admin"), err)
	}

	return Render(c, http.StatusOK, admin.DeleteModal(u))
}
