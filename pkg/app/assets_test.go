package app

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssetHeadersWithAssetDir(t *testing.T) {
	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, "logo.png"), []byte("not a png"), 0o600))

	t.Setenv("WT_DATABASE_DRIVER", "memory")
	a := defaultApp(t)
	a.AssetDir = dir
	require.NoError(t, a.Configure())

	req := httptest.NewRequest(http.MethodGet, a.Reverse("assets")+"/logo.png", nil)
	rec := httptest.NewRecorder()
	a.echo.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "no-cache", rec.Header().Get(echo.HeaderCacheControl))
	assert.NotEmpty(t, rec.Header().Get(echo.HeaderLastModified))
}
