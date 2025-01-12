package app

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/invopop/ctxi18n"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"

	"github.com/labstack/echo/v4"
)

func (a *App) setContext(ctx echo.Context) {
	ctx.Set("version", &a.Version)
	ctx.Set("config", &a.Config)
	ctx.Set("echo", a.echo)

	lctx, _ := ctxi18n.WithLocale(ctx.Request().Context(), langFromContextString(ctx))
	if lctx == nil {
		lctx, _ = ctxi18n.WithLocale(ctx.Request().Context(), "en")
	}

	ctx.SetRequest(ctx.Request().WithContext(lctx))
}

func (a *App) setUserFromContext(ctx echo.Context) error {
	if err := a.setUser(ctx); err != nil {
		return fmt.Errorf("error validating user: %w", err)
	}

	u := a.getCurrentUser(ctx)
	if u.IsAnonymous() || !u.IsActive() {
		return errors.New("user not found or active")
	}

	return nil
}

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

	c.Set("user_language", dbUser.Profile.Language)
	c.Set("user_info", dbUser)

	return nil
}

func (a *App) getCurrentUser(c echo.Context) *database.User {
	d := c.Get("user_info")

	u, ok := d.(*database.User)
	if !ok {
		u = database.AnonymousUser()
	}

	u.SetContext(c.Request().Context())

	return u
}

func (a *App) getRouteSegment(c echo.Context) (*database.RouteSegment, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, err
	}

	rs, err := database.GetRouteSegment(a.db, id)
	if err != nil {
		return nil, err
	}

	return rs, nil
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

func (a *App) getEquipment(c echo.Context) (*database.Equipment, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, err
	}

	w, err := a.getCurrentUser(c).GetEquipment(a.db, id)
	if err != nil {
		return nil, err
	}

	return w, nil
}
