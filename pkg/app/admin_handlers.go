package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *App) adminRootHandler(c echo.Context) error {
	data := a.defaultData(c)
	a.adminAddUsers(data, c)

	return c.Render(http.StatusOK, "admin_root.html", data)
}
