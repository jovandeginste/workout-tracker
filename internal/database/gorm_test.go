package database

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGorm_Connect(t *testing.T) {
	db, err := Connect("memory", "", false, slog.Default())
	require.NoError(t, err)
	assert.NotNil(t, db)

	db, err = Connect("invalid-driver", "some-dsn", false, slog.Default())
	require.Error(t, err)
	assert.Nil(t, db)
}
