package main

import (
	"fmt"
	"sync"
)

type maps_methods interface {
	Insert(int, int)
	Get(int) (int, error)
	Delete(int) error
}

type Cache struct {
	mu   sync.RWMutex
	mapa map[int]int
}

func NewCache() Cache {
	return Cache{
		mu:   sync.RWMutex{},
		mapa: make(map[int]int),
	}
}

func (c *Cache) Insert(key, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.mapa[key] = value
}

func (c *Cache) Get(key int) (value int, err error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.mapa[key]
	if !ok {
		return 0, fmt.Errorf("Value doesn't exist")
	}
	return value, nil
}
func (c *Cache) Delete(key int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.mapa, key)
}

func main() {
	var wg sync.WaitGroup
	task_map := NewCache()

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			task_map.Insert(i, i*i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(task_map.mapa[1000])
}
