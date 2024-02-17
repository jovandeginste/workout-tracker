package app

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jovandeginste/workouts/pkg/user"

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

	dbUser, err := user.GetUser(a.db, claims["name"].(string))
	if err != nil {
		return err
	}

	if !dbUser.IsActive() {
		return ErrInvalidJWTToken
	}

	c.Set("user_info", dbUser)

	return nil
}

func (a *App) getUser(c echo.Context) *user.User {
	d := c.Get("user_info")
	if d == nil {
		return nil
	}

	u, ok := d.(*user.User)
	if !ok {
		return nil
	}

	return u
}

func (a *App) defaultData(c echo.Context) map[string]interface{} {
	data := map[string]interface{}{}

	data["version"] = a.Version

	a.addUserInfo(data, c)
	a.addError(data, c)
	a.addNotice(data, c)

	return data
}

func (a *App) addUserInfo(data map[string]interface{}, c echo.Context) {
	u := a.getUser(c)
	if u == nil {
		return
	}

	data["user"] = u
}

func (a *App) addWorkouts(data map[string]interface{}, c echo.Context) {
	w, err := a.getUser(c).GetWorkouts(a.db)
	if err != nil {
		a.addError(data, c)
	}

	data["workouts"] = w
}
