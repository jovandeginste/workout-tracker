package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *App) userProfileHandler(c echo.Context) error {
	data := a.defaultData(c)
	return c.Render(http.StatusOK, "user_profile.html", data)
}
