package redis

import (
	"github.com/go-redis/redis/v8"
)

func ConnRedis() *redis.Client {
	rd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return rd
}
