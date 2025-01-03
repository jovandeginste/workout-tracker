package app

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func defaultAPIUser(db *gorm.DB) *database.User {
	u := defaultUser(db)
	u.APIKey = "my-api-key"
	u.Profile.APIActive = true
	u.Save(db)
	u.Profile.Save(db)

	return u
}

func TestAPI_WhoAmI(t *testing.T) { //nolint:funlen
	a := configuredApp(t)
	e := a.echo
	ts := httptest.NewServer(e)
	url := ts.URL + e.Reverse("api-whoami")
	u := defaultAPIUser(a.db)

	t.Run("with valid authorization header", func(t *testing.T) {
		client := &http.Client{}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)

		req.Header.Set("Authorization", "Bearer my-api-key")

		res, err := client.Do(req)
		require.NoError(t, err)

		if res != nil {
			defer res.Body.Close()
		}

		b, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Contains(t, string(b), u.Username)
	})

	t.Run("with invalid authorization header", func(t *testing.T) {
		client := &http.Client{}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)

		req.Header.Set("Authorization", "Bearer wrong-api-key")

		res, err := client.Do(req)
		require.NoError(t, err)

		if res != nil {
			defer res.Body.Close()
		}

		b, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
		assert.Contains(t, string(b), "Unauthorized")
	})

	t.Run("with valid query parameter", func(t *testing.T) {
		client := &http.Client{}

		req, err := http.NewRequest(http.MethodGet, url+"?api-key=my-api-key", nil)
		require.NoError(t, err)

		res, err := client.Do(req)
		require.NoError(t, err)

		if res != nil {
			defer res.Body.Close()
		}

		b, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Contains(t, string(b), u.Username)
	})

	t.Run("with invalid query parameter", func(t *testing.T) {
		client := &http.Client{}

		req, err := http.NewRequest(http.MethodGet, url+"?api-key=wrong-api-key", nil)
		require.NoError(t, err)

		res, err := client.Do(req)
		require.NoError(t, err)

		if res != nil {
			defer res.Body.Close()
		}

		b, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
		assert.Contains(t, string(b), "Unauthorized")
	})
}
