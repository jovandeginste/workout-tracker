package appviews

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmbedded(t *testing.T) {
	c, err := FS().Open("admin/admin_root.html")
	require.NoError(t, err)

	s, err := c.Stat()
	require.NoError(t, err)

	require.NotZero(t, s.Size())
}
