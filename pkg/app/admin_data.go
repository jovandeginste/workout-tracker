package app

import (
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

func (a *App) getUser(c echo.Context) (*database.User, error) {
	id, err := cast.ToUint64E(c.Param("id"))
	if err != nil {
		return nil, err
	}

	u, err := database.GetUserByID(a.db, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
