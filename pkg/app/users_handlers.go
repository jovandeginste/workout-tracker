package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/labstack/echo/v4"
)

var (
	ErrLoginFailed   = errors.New("username or password incorrect")
	ErrInternalError = errors.New("something went wrong")
)

// userSigninHandler will be executed after SignInForm submission.
func (a *App) userSigninHandler(c echo.Context) error {
	// Initiate a new User struct.
	u := new(database.User)

	// Parse the submitted data and fill the User struct with the data from the SignIn form.
	if err := c.Bind(u); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), fmt.Errorf("%w: %s", ErrInternalError, err))
	}

	storedUser, err := database.GetUser(a.db, u.Username)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), fmt.Errorf("%w: %s", ErrInternalError, err))
	}

	if !storedUser.ValidLogin(c.FormValue("password")) {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), ErrLoginFailed)
	}

	// If password is correct, generate tokens and set cookies.
	a.sessionManager.Put(c.Request().Context(), "username", u.Username)

	if err := a.createToken(storedUser, c); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("dashboard"))
}

// userSignoutHandler will log a user out
func (a *App) userSignoutHandler(c echo.Context) error {
	a.clearTokenCookie(c)

	if err := a.sessionManager.Destroy(c.Request().Context()); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), fmt.Errorf("%w: %s", ErrInternalError, err))
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("user-login"))
}

// userRegisterHandler will be executed after registration submission.
func (a *App) userRegisterHandler(c echo.Context) error {
	// Initiate a new User struct.
	u := new(database.User)

	// Parse the submitted data and fill the User struct with the data from the registration form.
	if err := c.Bind(u); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), fmt.Errorf("%w: %s", ErrInternalError, err))
	}

	if err := u.SetPassword(c.FormValue("password")); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), fmt.Errorf("%w: %s", ErrInternalError, err))
	}

	if err := u.Create(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), fmt.Errorf("%w: %s", ErrInternalError, err))
	}

	a.setNotice(c, "Your account has been created, but needs to be activated.")

	return c.Redirect(http.StatusFound, a.echo.Reverse("user-login"))
}

func (a *App) userProfileHandler(c echo.Context) error {
	data := a.defaultData(c)
	return c.Render(http.StatusOK, "user_profile.html", data)
}

func (a *App) userShowHandler(c echo.Context) error {
	data := a.defaultData(c)

	u, err := a.getUser(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("dashboard"), err)
	}

	if u == nil {
		return a.redirectWithError(
			c,
			a.echo.Reverse("dashboard"),
			fmt.Errorf("user id '%s' not found", c.Param("id")),
		)
	}

	data["user"] = u

	if err := a.addWorkouts(u, data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), err)
	}

	if err := a.addUserStatistics(u, data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), err)
	}

	return c.Render(http.StatusOK, "user_show.html", data)
}
