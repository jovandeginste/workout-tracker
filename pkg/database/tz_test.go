package database

import (
	"testing"

	"github.com/ringsaturn/tzf/v2"
	"github.com/stretchr/testify/require"
)

func TestInitTZFinder_Success(t *testing.T) {
	// Save and restore the global variable to ensure test isolation
	var oldFinder tzf.F
	oldFinder, tzFinder = tzFinder, oldFinder
	defer func() { tzFinder = oldFinder }()

	// Call the function under test
	err := InitTZFinder()
	require.NoError(t, err, "InitTZFinder should succeed in standard environment")

	// Verify that tzFinder is initialized (i.e., not the nil finder).
	var zeroF tzf.F
	require.NotEqual(t, zeroF, tzFinder, "tzFinder should be initialized after successful call")
}
