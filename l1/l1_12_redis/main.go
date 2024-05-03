package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Cache struct {
	client *redis.Client
}

func NewCache(addr string) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr: addr, // адрес и порт сервера Redis
	})
	return &Cache{client: client}
}

// Set сохраняет значение по ключу с опциональным временем жизни (TTL)
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

// Get извлекает значение по ключу
func (c *Cache) Get(key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}
func main() {
	cache := NewCache("localhost:6379")
	// Сохраняем значение в кеше с TTL 1 час
	err := cache.Set("user:1", "John Doe", time.Hour)
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}
	// Извлекаем и выводим значение из кеша
	value, err := cache.Get("user:1")
	if err != nil {
		fmt.Println("Error getting value:", err)
		return
	}
	fmt.Println("Got value:", value)
}
