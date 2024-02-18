package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jovandeginste/workouts/pkg/database"
	"github.com/labstack/echo/v4"
)

var (
	ErrLoginFailed   = errors.New("username or password incorrect")
	ErrInternalError = errors.New("something went wrong")
)

func (a *App) loginError(c echo.Context, err error) error {
	a.setError(c, err.Error())

	return c.Redirect(http.StatusMovedPermanently, "/user/signin")
}

// SignIn will be executed after SignInForm submission.
func (a *App) SignIn(c echo.Context) error {
	// Initiate a new User struct.
	u := new(database.User)

	// Parse the submitted data and fill the User struct with the data from the SignIn form.
	if err := c.Bind(u); err != nil {
		return a.loginError(c, fmt.Errorf("%w: %s", ErrInternalError, err))
	}

	storedUser, err := database.GetUser(a.db, u.Username)
	if err != nil {
		return a.loginError(c, fmt.Errorf("%w: %s", ErrInternalError, err))
	}

	if !storedUser.ValidLogin(u.Password) {
		return a.loginError(c, ErrLoginFailed)
	}

	// If password is correct, generate tokens and set cookies.
	a.sessionManager.Put(c.Request().Context(), "username", u.Username)

	if err := a.createToken(storedUser, c); err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

// SignOut will log a user out
func (a *App) SignOut(c echo.Context) error {
	a.clearTokenCookie(c)

	if err := a.sessionManager.Destroy(c.Request().Context()); err != nil {
		return a.loginError(c, fmt.Errorf("%w: %s", ErrInternalError, err))
	}

	return c.Redirect(http.StatusMovedPermanently, "/user/signin")
}

// Register will be executed after registration submission.
func (a *App) Register(c echo.Context) error {
	// Initiate a new User struct.
	u := new(database.User)

	// Parse the submitted data and fill the User struct with the data from the registration form.
	if err := c.Bind(u); err != nil {
		return a.loginError(c, fmt.Errorf("%w: %s", ErrInternalError, err))
	}

	if err := u.IsValid(); err != nil {
		return a.loginError(c, err)
	}

	if err := u.CryptPassword(); err != nil {
		return a.loginError(c, fmt.Errorf("%w: %s", ErrInternalError, err))
	}

	if err := u.Create(a.db); err != nil {
		return a.loginError(c, fmt.Errorf("%w: %s", ErrInternalError, err))
	}

	a.setNotice(c, "Your account has been created, but needs to be activated.")

	return c.Redirect(http.StatusMovedPermanently, "/")
}
