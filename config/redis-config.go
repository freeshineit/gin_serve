package config

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

// setup redis
func SetupRedisConnection() *redis.Client {

	redisConfig := Conf.Redis

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password, // no password set
		DB:       int(redisConfig.Db),  // use default DB
	})

	log.Println("Redis ping pong success...")

	return RedisClient
}

func CloseRedisConnection(client *redis.Client) {
	client.Close()
}
