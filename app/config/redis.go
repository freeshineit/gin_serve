package config

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

// setup redis
func SetupRedisConnection() *redis.Client {

	redisConfig, err := GetRedisConfig()

	if err != nil {
		log.Println(err)
		panic("redis config fail...")
	}

	Client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password, // no password set
		DB:       int(redisConfig.Db),  // use default DB
	})

	log.Println("Redis ping pong success...")

	return Client
}
