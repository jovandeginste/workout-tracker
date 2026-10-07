package database

import (
	"container/list"
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/go-gorm/caches/v4"
)

var ErrInvalidDataInCache = errors.New("invalid cache data")

const (
	// maxCacheEntries bounds the number of distinct queries kept in memory.
	maxCacheEntries = 100
	// maxCacheValueBytes skips caching results above this size (e.g. heatmap
	// queries returning hundreds of thousands of GPS points) so one expensive
	// query can't dominate the cache's memory budget.
	maxCacheValueBytes = 512 * 1024
)

// memoryCacher is a bounded, LRU-evicted cache. The previous implementation
// used an unbounded sync.Map with no eviction, which let large query results
// (notably the heatmap endpoints) accumulate in memory indefinitely until a
// write cleared the whole cache, eventually exhausting server memory.
type memoryCacher struct {
	mu    sync.Mutex
	store map[string][]byte
	order *list.List
	elems map[string]*list.Element
}

func (c *memoryCacher) init() {
	if c.store == nil {
		c.store = make(map[string][]byte)
		c.order = list.New()
		c.elems = make(map[string]*list.Element)
	}
}

func (c *memoryCacher) Get(_ context.Context, key string, q *caches.Query[any]) (*caches.Query[any], error) {
	c.mu.Lock()
	b, ok := c.store[key]

	if ok {
		c.order.MoveToFront(c.elems[key])
	}
	c.mu.Unlock()

	if !ok {
		return nil, nil //nolint:nilnil
	}

	if err := q.Unmarshal(b); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidDataInCache, err)
	}

	return q, nil
}

func (c *memoryCacher) Store(_ context.Context, key string, val *caches.Query[any]) error {
	res, err := val.Marshal()
	if err != nil {
		return err
	}

	if len(res) > maxCacheValueBytes {
		return nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, exists := c.elems[key]; exists {
		c.order.MoveToFront(elem)
		c.store[key] = res

		return nil
	}

	c.store[key] = res
	c.elems[key] = c.order.PushFront(key)

	for c.order.Len() > maxCacheEntries {
		oldest := c.order.Back()
		if oldest == nil {
			break
		}

		oldestKey, _ := oldest.Value.(string)
		c.order.Remove(oldest)
		delete(c.elems, oldestKey)
		delete(c.store, oldestKey)
	}

	return nil
}

func (c *memoryCacher) Invalidate(_ context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.store = make(map[string][]byte)
	c.order = list.New()
	c.elems = make(map[string]*list.Element)

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
