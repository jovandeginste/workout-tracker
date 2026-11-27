package helpers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoolToHTML(t *testing.T) {
	assert.Equal(t, `<i class="text-green-500 icon-[fa-solid--check]"></i>`, BoolToHTML(true))
	assert.Equal(t, `<i class="text-rose-500 icon-[fa-solid--times]"></i>`, BoolToHTML(false))
}

func userLinkContext() context.Context {
	e := echo.New()

	_, _ = e.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/users/:id",
		Name:    "user-show",
		Handler: func(*echo.Context) error { return nil },
	})

	return context.WithValue(context.Background(), "echo", e)
}

func TestUserLink(t *testing.T) {
	ctx := userLinkContext()

	render := func(u *database.User) string {
		t.Helper()

		var b strings.Builder

		require.NoError(t, UserLink(u).Render(ctx, &b))

		return b.String()
	}

	t.Run("it renders nothing without a user", func(t *testing.T) {
		assert.Empty(t, render(nil))
	})

	t.Run("it renders nothing for a user without an ID", func(t *testing.T) {
		assert.Empty(t, render(&database.User{UserData: database.UserData{Username: "no-id"}}))
	})

	t.Run("it links the username to the user", func(t *testing.T) {
		u := &database.User{UserData: database.UserData{Username: "jo", Name: "Jo Vandeginste"}}
		u.ID = 42

		body := render(u)

		assert.Contains(t, body, `href="/users/42"`)
		assert.Contains(t, body, "jo")
		assert.NotContains(t, body, "return")
	})
}

func TestUserLinkWithName(t *testing.T) {
	ctx := userLinkContext()

	u := &database.User{UserData: database.UserData{Username: "jo", Name: "Jo Vandeginste"}}
	u.ID = 42

	var b strings.Builder

	require.NoError(t, UserLinkWithName(u).Render(ctx, &b))

	assert.Contains(t, b.String(), "jo")
	assert.Contains(t, b.String(), "Jo Vandeginste")
	assert.NotContains(t, b.String(), "return")

	b.Reset()
	require.NoError(t, UserLinkWithName(nil).Render(ctx, &b))
	assert.Empty(t, b.String())

	b.Reset()
	require.NoError(t, UserLinkWithName(&database.User{}).Render(ctx, &b))
	assert.Empty(t, b.String())
}

func TestIsIOS(t *testing.T) {
	ctxWithUA := func(ua string) context.Context {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("User-Agent", ua)

		return context.WithValue(context.Background(), "request", req)
	}

	t.Run("desktop browsers are not iOS", func(t *testing.T) {
		assert.False(t, IsIOS(ctxWithUA("Mozilla/5.0 (X11; Linux x86_64) Gecko/20100101 Firefox/130.0")))
	})

	t.Run("iOS browsers are iOS", func(t *testing.T) {
		for _, ua := range []string{
			"Mozilla/5.0 (iPhone; CPU iPhone OS 18_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.4 Mobile/15E148 Safari/604.1",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 18_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/126.0.6478.153 Mobile/15E148 Safari/604.1",
			"Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Mobile/15E148 Safari/604.1",
		} {
			assert.True(t, IsIOS(ctxWithUA(ua)))
		}
	})

	t.Run("without a request it is not iOS", func(t *testing.T) {
		assert.False(t, IsIOS(context.Background()))
	})
}
