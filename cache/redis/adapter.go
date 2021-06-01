package cache

import (
	"github.com/go-redis/redis/v8"
)

type redisCacheAdapter struct {
}

func (a redisCacheAdapter) get(key string) {

}

func (a redisCacheAdapter) set(key string, value interface{}) {

}

func getInstance(config redisConfig) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
