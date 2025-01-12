package database

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/go-gorm/caches/v4"
)

var ErrInvalidDataInCache = errors.New("invalid cache data")

type memoryCacher struct {
	store *sync.Map
}

func (c *memoryCacher) init() {
	if c.store == nil {
		c.store = &sync.Map{}
	}
}

func (c *memoryCacher) Get(ctx context.Context, key string, q *caches.Query[any]) (*caches.Query[any], error) {
	val, ok := c.store.Load(key)
	if !ok {
		// Not stored in cache
		return nil, nil //nolint:nilnil
	}

	b, ok := val.([]byte)
	if !ok {
		return nil, ErrInvalidDataInCache // invalid data stored in cache
	}

	if err := q.Unmarshal(b); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidDataInCache, err) //  invalid data stored in cache
	}

	return q, nil
}

func (c *memoryCacher) Store(ctx context.Context, key string, val *caches.Query[any]) error {
	res, err := val.Marshal()
	if err != nil {
		return err
	}

	c.store.Store(key, res)

	return nil
}

func (c *memoryCacher) Invalidate(ctx context.Context) error {
	c.store = &sync.Map{}
	return nil
}

func NewMemoryCache() *caches.Caches {
	c := &memoryCacher{}
	c.init()

	cachesPlugin := &caches.Caches{Conf: &caches.Config{
		Cacher: c,
		Easer:  true,
	}}

	return cachesPlugin
}
