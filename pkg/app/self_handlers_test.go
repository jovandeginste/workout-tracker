package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
