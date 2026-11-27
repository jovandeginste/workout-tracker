package app

import (
	"errors"
	"net/http"
	"strings"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/views/user"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

var ErrLoginFailed = errors.New("username or password incorrect")

func (a *App) addRoutesUsers(g *echo.Group) {
	usersGroup := g.Group("/users")
	a.GET(usersGroup, "/suggest", a.userSuggestHandler, "user-suggest")
	a.GET(usersGroup, "/:id", a.userShowHandler, "user-show")
}

// userSuggestHandler returns the users matching the (partial) username, for the
// autocomplete of the username filters.
func (a *App) userSuggestHandler(c *echo.Context) error {
	users, err := database.GetUsers(a.db)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return Render(c, http.StatusOK, user.Suggestions(filterUsers(users, c.QueryParam("username"))))
}

const userSuggestionLimit = 20

func filterUsers(users []*database.User, query string) []*database.User {
	query = strings.ToLower(strings.TrimSpace(query))

	matches := make([]*database.User, 0, len(users))

	for _, u := range users {
		if query != "" &&
			!strings.Contains(strings.ToLower(u.Username), query) &&
			!strings.Contains(strings.ToLower(u.Name), query) {
			continue
		}

		matches = append(matches, u)

		if len(matches) == userSuggestionLimit {
			break
		}
	}

	return matches
}

// userSigninHandler will be executed after SignInForm submission.
func (a *App) userSigninHandler(c *echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	storedUser, err := database.GetUser(a.db, username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		a.logger.Error("login: lookup user", "error", err)
		return a.redirectWithError(c, a.Reverse("user-login"), ErrLoginFailed)
	}

	// LDAP usernames are case-insensitive. Only fall back to a folded lookup
	// for LDAP accounts; local usernames keep their exact-match behaviour.
	if storedUser == nil {
		storedUser, err = database.GetUserFold(a.db, username)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			a.logger.Error("login: folded lookup user", "error", err)
			return a.redirectWithError(c, a.Reverse("user-login"), ErrLoginFailed)
		}

		if storedUser != nil && storedUser.AuthSource != database.AuthSourceLDAP {
			storedUser = nil
		}
	}

	var u *database.User

	switch {
	case storedUser == nil:
		u, err = a.loginWithLDAP(c, username, password)
	case storedUser.AuthSource == database.AuthSourceLDAP:
		u, err = a.loginWithLDAP(c, storedUser.Username, password)
	default:
		if !storedUser.ValidLogin(password) {
			err = ErrLoginFailed
		} else {
			u = storedUser
		}
	}

	if err != nil {
		return a.redirectWithError(c, a.Reverse("user-login"), ErrLoginFailed)
	}

	return a.finishLogin(c, u)
}

// userSignoutHandler will log a user out
func (a *App) userSignoutHandler(c *echo.Context) error {
	a.clearTokenCookie(c)

	if err := a.sessionManager.Destroy(c.Request().Context()); err != nil {
		return a.redirectWithError(c, a.Reverse("user-login"), err)
	}

	return c.Redirect(http.StatusFound, a.Reverse("user-login"))
}

// userRegisterHandler will be executed after registration submission.
func (a *App) userRegisterHandler(c *echo.Context) error {
	if a.Config.RegistrationDisabled {
		return a.redirectWithError(c, a.Reverse("user-login"), errors.New("registration is disabled"))
	}

	// Initiate a new User struct.
	u := new(database.User)

	// Parse the submitted data and fill the User struct with the data from the registration form.
	if err := c.Bind(u); err != nil {
		return a.redirectWithError(c, a.Reverse("user-login"), err)
	}

	// The form may contain an AuthSource field; never let a registration
	// choose an external provider.
	u.AuthSource = database.AuthSourceLocal

	if err := u.SetPassword(c.FormValue("password")); err != nil {
		return a.redirectWithError(c, a.Reverse("user-login"), err)
	}

	applyUserDefaults(u)
	// ensure user is not admin and not active by default
	u.Admin = false
	u.Active = false

	if err := u.Create(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("user-login"), err)
	}

	a.addNoticeT(c, "user.auth.activation-needed")

	return c.Redirect(http.StatusFound, a.Reverse("user-login"))
}

func applyUserDefaults(u *database.User) {
	u.Profile.Theme = BrowserTheme
	u.Profile.TotalsShow = database.WorkoutTypeList{DefaultTotalsShow}
	u.Profile.Language = BrowserLanguage
	u.Profile.RouteSegmentConfig.TrendPeriod = database.DefaultRouteSegmentTrendPeriod
	u.Profile.RouteSegmentConfig.TrendBreakDetection = true
}

func (a *App) userShowHandler(c *echo.Context) error {
	u, err := a.getUser(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("dashboard"), err)
	}

	viewer := a.getCurrentUser(c)

	return Render(c, http.StatusOK, user.Show(u, nil, nil, viewer))
}
