package converters

import (
	"os"
	"testing"

	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tkrajina/gpxgo/gpx"
)

func rootExtension(t *testing.T, g *gpx.GPX, key string) (float64, bool) {
	t.Helper()

	node, ok := g.Extensions.GetNode(gpx.AnyNamespace, key)
	if !ok {
		return 0, false
	}

	value, err := cast.ToFloat64E(node.Data)
	require.NoError(t, err)

	return value, true
}

func TestParseFitKeepsSessionSummary(t *testing.T) {
	content, err := os.ReadFile("../database/testdata/issues/issue-0658-15718428901-activity.fit")
	require.NoError(t, err)

	g, err := ParseFit(content)
	require.NoError(t, err)

	distance, ok := rootExtension(t, g, ExtensionTotalDistance)
	require.True(t, ok)
	assert.InDelta(t, 52185.44, distance, 0.01)

	ascent, ok := rootExtension(t, g, ExtensionTotalAscent)
	require.True(t, ok)
	assert.InDelta(t, 262, ascent, 0.001)

	descent, ok := rootExtension(t, g, ExtensionTotalDescent)
	require.True(t, ok)
	assert.InDelta(t, 242, descent, 0.001)

	timer, ok := rootExtension(t, g, ExtensionTotalTimerTime)
	require.True(t, ok)
	assert.InDelta(t, 15771.708, timer, 0.01)

	// This file does not contain a maximum speed in its session.
	_, ok = rootExtension(t, g, ExtensionMaxSpeed)
	assert.False(t, ok)
}

func TestParseTCXKeepsLapSummary(t *testing.T) {
	content, err := os.ReadFile("../database/testdata/issues/issue-0860-18-98km-inline-skating-mapmyfitness.tcx")
	require.NoError(t, err)

	g, err := ParseTCX(content)
	require.NoError(t, err)

	distance, ok := rootExtension(t, g, ExtensionTotalDistance)
	require.True(t, ok)
	assert.InDelta(t, 18976.0969728, distance, 0.001)

	calories, ok := rootExtension(t, g, "total-calories")
	require.True(t, ok)
	assert.InDelta(t, 1327, calories, 0.001)

	maxSpeed, ok := rootExtension(t, g, ExtensionMaxSpeed)
	require.True(t, ok)
	assert.InDelta(t, 10.046285216, maxSpeed, 0.001)
}
