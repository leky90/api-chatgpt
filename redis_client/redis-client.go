package redis_client

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitRedisClient() *redis.Client {
	// initialize Redis client
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// PING to test the connection
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis server: ", pong)

	return rdb
}

func GetRedisClient() *redis.Client {
	return rdb
}
