package _map

import (
	"fmt"
	"sync"
)

// Напишите функцию GetOrCreate, которая или создает новый элемент мапы,
// если его ещё не было и возвращает его значение, или просто возвращает
// Важно учесть, что код должен корректно работать в конкурентной среде.

type CurrentMap struct {
	data map[string]string
	m    sync.RWMutex
}

func NewConcurrentMap() *CurrentMap {

	return &CurrentMap{
		data: make(map[string]string),
		m:    sync.RWMutex{},
	}
}

func (c *CurrentMap) GerOrCreate(key, val string) string {
	c.m.RLock()
	v, ok := c.data[key]
	c.m.RUnlock()

	if ok {
		return val
	}

	c.m.Lock()
	defer c.m.Unlock()
	if v, ok = c.data[key]; ok {
		return v
	}

	c.data[key] = v

	return val
}

func MainGetOrCreate() {
	cm := NewConcurrentMap()

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		val := cm.GerOrCreate("key1", "value1")
		fmt.Println("Goroutine 1 got:", val)
	}()

	go func() {
		defer wg.Done()
		val := cm.GerOrCreate("key1", "value2")
		fmt.Println("Goroutine 2 got:", val)
	}()

	wg.Wait()

}
