package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/invopop/ctxi18n"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/spf13/cast"

	"github.com/labstack/echo/v5"
)

func (a *App) setContext(ctx *echo.Context) {
	values := map[string]any{
		"version":        &a.Version,
		"config":         &a.Config,
		"echo":           a.echo,
		"sessionManager": a.sessionManager,
	}

	reqCtx := ctx.Request().Context()
	for key, value := range values {
		ctx.Set(key, value)
		reqCtx = context.WithValue(reqCtx, key, value)
	}

	lctx, _ := ctxi18n.WithLocale(reqCtx, a.langFromContextString(ctx))
	if lctx == nil {
		lctx, _ = ctxi18n.WithLocale(reqCtx, "en")
	}

	ctx.SetRequest(ctx.Request().WithContext(lctx))
}

func (a *App) setUserFromContext(ctx *echo.Context) error {
	if err := a.setUser(ctx); err != nil {
		return fmt.Errorf("error validating user: %w", err)
	}

	u := a.getCurrentUser(ctx)
	if u.IsAnonymous() || !u.IsActive() {
		return errors.New("user not found or active")
	}

	return nil
}

func (a *App) setUser(c *echo.Context) error {
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

	a.setCurrentUser(c, dbUser)

	return nil
}

func (a *App) setCurrentUser(c *echo.Context, u *database.User) {
	c.Set("user_language", u.Profile.Language)
	c.Set("user_info", u)
	c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), "user_info", u)))
}

// userForUsername returns the user whose items a list should show. Without a
// username, or with your own username, this is the current user.
func (a *App) userForUsername(c *echo.Context, username string) (*database.User, error) {
	current := a.getCurrentUser(c)

	if username == "" || username == current.Username {
		return current, nil
	}

	u, err := database.GetUser(a.db, username)
	if err != nil || u == nil {
		return nil, ErrUserNotFound
	}

	return u, nil
}

func (a *App) getCurrentUser(c *echo.Context) *database.User {
	d := c.Get("user_info")

	u, ok := d.(*database.User)
	if !ok {
		u = database.AnonymousUser()
	}

	u.SetContext(c.Request().Context())

	return u
}

func (a *App) getRouteSegment(c *echo.Context) (*database.RouteSegment, error) {
	id, err := cast.ToUint64E(c.Param("id"))
	if err != nil {
		return nil, err
	}

	rs, err := database.GetRouteSegment(a.db, id)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (a *App) getWorkout(c *echo.Context) (*database.Workout, error) {
	id, err := cast.ToUint64E(c.Param("id"))
	if err != nil {
		return nil, err
	}

	w, err := a.getCurrentUser(c).GetWorkout(a.db, id)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (a *App) getEquipment(c *echo.Context) (*database.Equipment, error) {
	id, err := cast.ToUint64E(c.Param("id"))
	if err != nil {
		return nil, err
	}

	w, err := a.getCurrentUser(c).GetEquipment(a.db, id)
	if err != nil {
		return nil, err
	}

	return w, nil
}
