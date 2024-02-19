package app

import (
	"github.com/jovandeginste/workouts/pkg/database"
	"github.com/labstack/echo/v4"
)

func (a *App) adminAddUsers(data map[string]interface{}, c echo.Context) {
	users, err := database.GetUsers(a.db)
	if err != nil {
		a.addError(data, c)
		return
	}

	data["adminUsers"] = users
}
