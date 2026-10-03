package database

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/go-gorm/caches/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemoryCacher_StoreAndGet(t *testing.T) {
	c := &memoryCacher{}
	c.init()

	ctx := context.Background()
	stored := &caches.Query[any]{Dest: "some-value", RowsAffected: 1}

	require.NoError(t, c.Store(ctx, "key-1", stored))

	got := &caches.Query[any]{}
	result, err := c.Get(ctx, "key-1", got)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "some-value", got.Dest)
	assert.Equal(t, int64(1), got.RowsAffected)
}

func TestMemoryCacher_GetMissingKey(t *testing.T) {
	c := &memoryCacher{}
	c.init()

	result, err := c.Get(context.Background(), "does-not-exist", &caches.Query[any]{})
	require.NoError(t, err)
	assert.Nil(t, result)
}

func TestMemoryCacher_SkipsOversizedValues(t *testing.T) {
	c := &memoryCacher{}
	c.init()

	ctx := context.Background()
	huge := &caches.Query[any]{Dest: strings.Repeat("x", maxCacheValueBytes+1)}

	require.NoError(t, c.Store(ctx, "huge", huge))

	result, err := c.Get(ctx, "huge", &caches.Query[any]{})
	require.NoError(t, err)
	assert.Nil(t, result, "oversized values must not be cached")
}

func TestMemoryCacher_EvictsOldestBeyondCap(t *testing.T) {
	c := &memoryCacher{}
	c.init()

	ctx := context.Background()
	total := maxCacheEntries + 5

	for i := range total {
		key := fmt.Sprintf("key-%d", i)
		require.NoError(t, c.Store(ctx, key, &caches.Query[any]{Dest: i}))
	}

	c.mu.Lock()
	entries := len(c.store)
	c.mu.Unlock()

	assert.LessOrEqual(t, entries, maxCacheEntries)

	// the oldest keys should have been evicted
	result, err := c.Get(ctx, "key-0", &caches.Query[any]{})
	require.NoError(t, err)
	assert.Nil(t, result)

	// the most recent key should still be present
	result, err = c.Get(ctx, fmt.Sprintf("key-%d", total-1), &caches.Query[any]{})
	require.NoError(t, err)
	require.NotNil(t, result)
}

func TestMemoryCacher_Invalidate(t *testing.T) {
	c := &memoryCacher{}
	c.init()

	ctx := context.Background()
	require.NoError(t, c.Store(ctx, "key-1", &caches.Query[any]{Dest: "value"}))

	require.NoError(t, c.Invalidate(ctx))

	result, err := c.Get(ctx, "key-1", &caches.Query[any]{})
	require.NoError(t, err)
	assert.Nil(t, result)
}
