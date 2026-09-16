package app

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestRouteSegmentConfigMalformedForm(t *testing.T) {
	a := configuredApp(t)
	u := defaultUser(a.db)
	original := u.Profile.RouteSegmentConfig
	req := httptest.NewRequest(http.MethodPost, "/user/profile/route-segment-config",
		strings.NewReader("route_segment_trend_period=30&invalid=%zz"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c := a.echo.NewContext(req, rec)
	c.Set("user_info", u)

	require.NoError(t, a.userProfileRouteSegmentConfigUpdateHandler(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Empty(t, rec.Body.String())
	assert.Equal(t, original, u.Profile.RouteSegmentConfig)
}

func TestRouteSegmentConfigSaveFailure(t *testing.T) {
	a := configuredApp(t)
	u := defaultUser(a.db)
	require.NoError(t, u.Save(a.db))
	require.NoError(t, u.Profile.Save(a.db))
	require.NoError(t, a.db.Callback().Update().Before("gorm:update").Register("test:fail_profile_save", func(db *gorm.DB) {
		db.AddError(errors.New("internal persistence details"))
	}))
	req := httptest.NewRequest(http.MethodPost, "/user/profile/route-segment-config",
		strings.NewReader("route_segment_trend_period=30"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c := a.echo.NewContext(req, rec)
	c.Set("user_info", u)

	require.NoError(t, a.userProfileRouteSegmentConfigUpdateHandler(c))
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Empty(t, rec.Body.String())
}
