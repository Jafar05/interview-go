package Задачи

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu                sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	items             map[string]Item
}

type Item struct {
	Value      interface{}
	Created    time.Time
	Expiration int64
}

func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	cache := &Cache{
		items:             make(map[string]Item),
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
	}

	return cache
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	if duration == 0 {
		duration = c.defaultExpiration
	}

	var expiration int64

	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = Item{
		Value:      value,
		Created:    time.Now(),
		Expiration: expiration,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found || (item.Expiration > 0 && time.Now().UnixNano() > item.Expiration) {
		return nil, false
	}
	return item.Value, true
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, found := c.items[key]; !found {
		return errors.New("key not found")
	}
	delete(c.items, key)
	return nil
}

func (c *Cache) startGC() {
	ticker := time.NewTicker(c.cleanupInterval)
	defer ticker.Stop()

	for range ticker.C {
		c.cleanup()
	}
}

func (c *Cache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now().UnixNano()
	for k, item := range c.items {
		if item.Expiration > 0 && now > item.Expiration {
			delete(c.items, k)
		}
	}
}

func importCache() {
	// Создаем новый кэш с:
	// - Временем жизни элементов по умолчанию: 5 секунд
	// - Интервалом очистки: 10 секунд
	cache := New(5*time.Second, 10*time.Second)

	// Устанавливаем ключ "username" с значением "JohnDoe" на 3 секунды
	cache.Set("username", "JohnDoe", 3*time.Second)

	// Получаем значение по ключу "username"
	if value, found := cache.Get("username"); found {
		fmt.Println("Found:", value) // Ожидаем: Found: JohnDoe
	} else {
		fmt.Println("Key not found")
	}

	// Ждем 4 секунды, чтобы элемент устарел
	time.Sleep(4 * time.Second)

	// Проверяем снова
	if value, found := cache.Get("username"); found {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Key expired or not found") // Ожидаем: Key expired or not found
	}

	// Устанавливаем другой ключ без указания времени жизни
	cache.Set("session_id", "abc123", 0)

	// Получаем его значение
	if value, found := cache.Get("session_id"); found {
		fmt.Println("Session ID:", value) // Ожидаем: Session ID: abc123
	} else {
		fmt.Println("Key not found")
	}

	// Удаляем ключ "session_id"
	if err := cache.Delete("session_id"); err != nil {
		fmt.Println("Error deleting key:", err)
	} else {
		fmt.Println("Key 'session_id' deleted")
	}

	// Пробуем получить удаленный ключ
	if value, found := cache.Get("session_id"); found {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Key not found") // Ожидаем: Key not found
	}
}
