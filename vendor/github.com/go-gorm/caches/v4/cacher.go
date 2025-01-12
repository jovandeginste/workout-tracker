package caches

import (
	"context"
)

type Cacher interface {
	// Get impl should check if a specific key exists in the cache and return its value
	// look at Query.Marshal
	Get(ctx context.Context, key string, q *Query[any]) (*Query[any], error)
	// Store impl should store a cached representation of the val param
	// look at Query.Unmarshal
	Store(ctx context.Context, key string, val *Query[any]) error
	// Invalidate impl should invalidate all cached values
	// It will be called when INSERT / UPDATE / DELETE queries are sent to the DB
	Invalidate(ctx context.Context) error
}
