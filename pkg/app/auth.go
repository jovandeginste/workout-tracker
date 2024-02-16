package app

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jovandeginste/workouts/pkg/user"
	"github.com/labstack/echo/v4"
)

var ErrInvalidJWTToken = errors.New("invalid JWT token")

func (a *App) createToken(u *user.User, c echo.Context) error {
	token := jwt.New(jwt.SigningMethodHS256)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ErrInvalidJWTToken
	}

	exp := time.Now().Add(time.Hour * 24 * 10)

	claims["name"] = u.Username
	claims["exp"] = exp.Unix()

	t, err := token.SignedString(a.jwtSecret)
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = exp

	c.SetCookie(cookie)

	return nil
}
