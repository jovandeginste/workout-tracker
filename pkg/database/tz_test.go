package database

import (
	"testing"

	"github.com/ringsaturn/tzf"
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

	// Verify that tzFinder is initialized (i.e., not the zero value tzf.F{}).
	// Since tzf.F is a struct containing data structures (R-tree), a successful
	// initialization should result in a non-zero value.
	var zeroF tzf.F
	require.NotEqual(t, zeroF, tzFinder, "tzFinder should be initialized after successful call")
}
