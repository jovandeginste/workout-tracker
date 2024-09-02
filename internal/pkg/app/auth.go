package app

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jovandeginste/workout-tracker/internal/database"
	"github.com/labstack/echo/v4"
)

var ErrInvalidJWTToken = errors.New("invalid JWT token")

func (a *App) createToken(u *database.User, c echo.Context) error {
	token := jwt.New(jwt.SigningMethodHS256)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ErrInvalidJWTToken
	}

	exp := time.Now().Add(time.Hour * 24 * 10)

	claims["name"] = u.Username
	claims["exp"] = exp.Unix()

	t, err := token.SignedString(a.jwtSecret())
	if err != nil {
		return err
	}

	a.setTokenCookie(t, exp, c)

	return nil
}

func (a *App) setTokenCookie(t string, exp time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = exp

	c.SetCookie(cookie)
}

func (a *App) clearTokenCookie(c echo.Context) {
	exp := time.Now()
	a.setTokenCookie("", exp, c)
}
