package app

import (
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jovandeginste/workouts/pkg/database"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
)

func (a *App) setUser(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return ErrInvalidJWTToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ErrInvalidJWTToken
	}

	dbUser, err := database.GetUser(a.db, claims["name"].(string))
	if err != nil {
		return err
	}

	if !dbUser.IsActive() {
		return ErrInvalidJWTToken
	}

	c.Set("user_info", dbUser)

	return nil
}

func (a *App) getCurrentUser(c echo.Context) *database.User {
	d := c.Get("user_info")
	if d == nil {
		return nil
	}

	u, ok := d.(*database.User)
	if !ok {
		return nil
	}

	return u
}

func (a *App) defaultData(c echo.Context) map[string]interface{} {
	data := map[string]interface{}{}

	data["version"] = a.Version
	// data["routes"] = a.Routes()
	spew.Dump(a.echo.Routes())

	a.addUserInfo(data, c)
	a.addError(data, c)
	a.addNotice(data, c)

	return data
}

func (a *App) addUserInfo(data map[string]interface{}, c echo.Context) {
	u := a.getCurrentUser(c)
	if u == nil {
		return
	}

	data["user"] = u
}

func (a *App) addWorkouts(data map[string]interface{}, c echo.Context) {
	w, err := a.getCurrentUser(c).GetWorkouts(a.db)
	if err != nil {
		a.addError(data, c)
	}

	data["workouts"] = w
}

func (a *App) addUserStatistics(data map[string]interface{}, c echo.Context) {
	data["UserStatistics"] = a.getCurrentUser(c).Statistics(a.db)
}

func (a *App) getWorkout(c echo.Context) (*database.Workout, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, err
	}

	w, err := a.getCurrentUser(c).GetWorkout(a.db, id)
	if err != nil {
		return nil, err
	}

	return w, nil
}
