package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	data     map[int]int
	inflight map[int]chan struct{}
}

func NewCache() *Cache {
	return &Cache{
		data:     make(map[int]int),
		inflight: make(map[int]chan struct{}),
	}
}

func (c *Cache) Get(key int) int {
	c.mu.Lock()

	if val, ok := c.data[key]; ok {
		c.mu.Unlock()
		return val
	}

	if ch, ok := c.inflight[key]; ok {
		c.mu.Unlock()
		<-ch
		c.mu.Lock()
		val := c.data[key]
		c.mu.Unlock()
		return val
	}

	ch := make(chan struct{})
	c.inflight[key] = ch
	c.mu.Unlock()

	time.Sleep(1 * time.Second)
	val := fib(key)

	c.mu.Lock()
	c.data[key] = val
	delete(c.inflight, key)
	close(ch)
	c.mu.Unlock()

	return val
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	cache := NewCache()

	go func() { fmt.Println(cache.Get(10)) }()
	go func() { fmt.Println(cache.Get(10)) }()

	time.Sleep(2 * time.Second)
}