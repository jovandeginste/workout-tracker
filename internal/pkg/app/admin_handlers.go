package app

import (
	"net/http"

	"github.com/jovandeginste/workout-tracker/internal/database"
	"github.com/labstack/echo/v4"
)

func (a *App) adminRoutes(e *echo.Group) *echo.Group {
	adminGroup := e.Group("/admin")
	adminGroup.Use(a.ValidateAdminMiddleware)

	adminGroup.GET("", a.adminRootHandler).Name = "admin"
	adminGroup.POST("/config", a.adminConfigUpdateHandler).Name = "admin-config-update"

	adminUsersGroup := adminGroup.Group("/users")
	adminUsersGroup.GET("/:id/edit", a.adminUserEditHandler).Name = "admin-user-edit"
	adminUsersGroup.POST("/:id", a.adminUserUpdateHandler).Name = "admin-user-update"
	adminUsersGroup.POST("/:id/delete", a.adminUserDeleteHandler).Name = "admin-user-delete"
	adminUsersGroup.GET("/:id", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, a.echo.Reverse("admin-user-edit", c.Param("id")))
	}).Name = "admin-user-show"

	return adminGroup
}

func (a *App) adminRootHandler(c echo.Context) error {
	data := a.defaultData(c)

	if err := a.addUsers(data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), err)
	}

	return c.Render(http.StatusOK, "admin_root.html", data)
}

func (a *App) adminUserEditHandler(c echo.Context) error {
	data := a.defaultData(c)
	a.adminAddUser(data, c)

	return c.Render(http.StatusOK, "admin_user_edit.html", data)
}

func (a *App) adminUserUpdateHandler(c echo.Context) error {
	u, err := a.getUser(c)
	if err != nil {
		return a.redirectWithError(c, "/admin", err)
	}

	u.Name = c.FormValue("name")
	u.Username = c.FormValue("username")
	u.Admin = isChecked(c.FormValue("admin"))
	u.Active = isChecked(c.FormValue("active"))

	if c.FormValue("password") != "" {
		if err := u.SetPassword(c.FormValue("password")); err != nil {
			return a.redirectWithError(c, a.echo.Reverse("admin-user-show", c.Param("id")), err)
		}
	}

	if err := u.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("admin-user-show", c.Param("id")), err)
	}

	a.setNotice(c, "The user '%s' has been updated.", u.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("admin-user-show", c.Param("id")))
}

func (a *App) adminUserDeleteHandler(c echo.Context) error { //nolint:dupl
	u, err := a.getUser(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("admin"), err)
	}

	if err := u.Delete(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("admin-user-show", c.Param("id")), err)
	}

	a.setNotice(c, "The user '%s' has been deleted.", u.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("admin"))
}

func (a *App) adminConfigUpdateHandler(c echo.Context) error {
	var cnf database.Config

	if err := c.Bind(&cnf); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("admin"), err)
	}

	if err := cnf.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("admin"), err)
	}

	if err := a.ResetConfiguration(); err != nil {
		return err
	}

	a.setNotice(c, "Config updated")

	return c.Redirect(http.StatusFound, a.echo.Reverse("admin"))
}

func isChecked(value string) bool {
	return value == "on"
}
