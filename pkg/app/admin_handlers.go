package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *App) adminRootHandler(c echo.Context) error {
	data := a.defaultData(c)

	if err := a.addUsers(data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), fmt.Errorf("%w: %s", ErrInternalError, err))
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
			return a.redirectWithError(c, a.echo.Reverse("admin-user-show", c.Param("id")), fmt.Errorf("%w: %s", ErrInternalError, err))
		}
	}

	if err := u.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("admin-user-show", c.Param("id")), fmt.Errorf("%w: %s", ErrInternalError, err))
	}

	a.setNotice(c, "The user '%s' has been updated.", u.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("admin-user-show", c.Param("id")))
}

func (a *App) adminUserDeleteHandler(c echo.Context) error {
	u, err := a.getUser(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("admin"), fmt.Errorf("%w: %s", ErrInternalError, err))
	}

	if err := u.Delete(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("admin-user-show", c.Param("id")), fmt.Errorf("%w: %s", ErrInternalError, err))
	}

	a.setNotice(c, "The user '%s' has been deleted.", u.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("admin"))
}

func isChecked(value string) bool {
	return value == "on"
}
