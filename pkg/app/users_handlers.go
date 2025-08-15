package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/views/user"
	"github.com/labstack/echo/v4"
)

var ErrLoginFailed = errors.New("username or password incorrect")

func (a *App) addRoutesUsers(e *echo.Group) {
	usersGroup := e.Group("/users")
	usersGroup.GET("/:id", a.userShowHandler).Name = "user-show"
}

// userSigninHandler will be executed after SignInForm submission.
func (a *App) userSigninHandler(c echo.Context) error {
	// Initiate a new User struct.
	u := new(database.User)

	// Parse the submitted data and fill the User struct with the data from the SignIn form.
	if err := c.Bind(u); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), err)
	}

	storedUser, err := database.GetUser(a.db, u.Username)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), err)
	}

	if !storedUser.ValidLogin(c.FormValue("password")) {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), ErrLoginFailed)
	}

	// If password is correct, generate tokens and set cookies.
	a.sessionManager.Put(c.Request().Context(), "username", u.Username)

	if err := a.createToken(storedUser, c); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), ErrLoginFailed)
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("dashboard"))
}

// userSignoutHandler will log a user out
func (a *App) userSignoutHandler(c echo.Context) error {
	a.clearTokenCookie(c)

	if err := a.sessionManager.Destroy(c.Request().Context()); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), err)
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("user-login"))
}

// userRegisterHandler will be executed after registration submission.
func (a *App) userRegisterHandler(c echo.Context) error {
	if a.Config.RegistrationDisabled {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), errors.New("registration is disabled"))
	}

	// Initiate a new User struct.
	u := new(database.User)

	// Parse the submitted data and fill the User struct with the data from the registration form.
	if err := c.Bind(u); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), err)
	}

	if err := u.SetPassword(c.FormValue("password")); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), err)
	}

	u.Profile.Theme = BrowserTheme
	u.Profile.TotalsShow = DefaultTotalsShow
	u.Profile.Language = BrowserLanguage
	// ensure user is not admin and not active by default
	u.Admin = false
	u.Active = false

	if err := u.Create(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-login"), err)
	}

	a.addNoticeT(c, "translation.Your_account_has_been_created_but_needs_to_be_activated")

	return c.Redirect(http.StatusFound, a.echo.Reverse("user-login"))
}

func (a *App) userShowHandler(c echo.Context) error {
	u, err := a.getUser(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("dashboard"), err)
	}

	if u.IsAnonymous() {
		return a.redirectWithError(
			c,
			a.echo.Reverse("dashboard"),
			fmt.Errorf("user id '%s' not found", c.Param("id")),
		)
	}

	w, err := u.GetWorkouts(a.db)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), err)
	}

	return Render(c, http.StatusOK, user.Show(u, nil, w, nil))
}
