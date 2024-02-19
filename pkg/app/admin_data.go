package app

import (
	"strconv"

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

func (a *App) adminAddUser(data map[string]interface{}, c echo.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		a.addError(data, c)
		return
	}

	user, err := database.GetUserByID(a.db, userID)
	if err != nil {
		a.addError(data, c)
		return
	}

	data["adminUser"] = user
}

func (a *App) getUser(c echo.Context) (*database.User, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, err
	}

	u, err := database.GetUserByID(a.db, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
