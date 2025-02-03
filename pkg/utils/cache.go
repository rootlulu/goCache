package utils

import "sync"

type Cache[K comparable, V any] struct {
	m  map[K]V
	mu sync.Mutex
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		m: make(map[K]V),
	}
}

func (c *Cache[K, V]) Get(key K) V {
	return c.m[key]
}

func (c *Cache[K, V]) GetBool(key K) (V, bool) {
	val, ok := c.m[key]
	return val, ok
}

func (c *Cache[K, V]) Set(key K, val V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = val
}

func (c *Cache[K, V]) Del(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
}

func (c *Cache[K, V]) Length() int {
	return len(c.m)
}

func (c *Cache[K, V]) Range() <-chan struct {
	Key K
	Val V
} {
	ch := make(chan struct {
		Key K
		Val V
	})
	go func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		defer close(ch)

		for k, v := range c.m {
			ch <- struct {
				Key K
				Val V
			}{k, v}
		}
	}()
	return ch
}
