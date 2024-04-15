package app

import (
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jovandeginste/workout-tracker/pkg/database"

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
	c.Set("user_language", dbUser.Profile.Language)

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

	a.addError(data, c)
	a.addNotice(data, c)

	return data
}

func (a *App) addWorkouts(u *database.User, data map[string]interface{}) error {
	if u == nil {
		return nil
	}

	w, err := u.GetWorkouts(a.db)
	if err != nil {
		return err
	}

	data["workouts"] = w

	return nil
}

func (a *App) addRecentWorkouts(data map[string]interface{}) error {
	w, err := database.GetRecentWorkouts(a.db, 20)
	if err != nil {
		return err
	}

	data["recentWorkouts"] = w

	return nil
}

func (a *App) addUsers(data map[string]interface{}) error {
	users, err := database.GetUsers(a.db)
	if err != nil {
		return err
	}

	data["users"] = users

	return nil
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
