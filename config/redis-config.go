package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var RedisClient *redis.Client

// setup redis
func SetupRedisConnection() *redis.Client {

	redisConfig := Conf.Redis

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.Db,       // use default DB
		// Username: "default",
		// DialTimeout:  10 * time.Second,
		// ReadTimeout:  30 * time.Second,
		// WriteTimeout: 30 * time.Second,
		// PoolSize:     10,
		// PoolTimeout:  30 * time.Second,
	})

	_, err := RedisClient.Ping(context.Background()).Result()

	if err != nil {
		fmt.Printf("Redis connect ping failed, err: %s \n", err.Error())
		zap.S().Errorf("Redis connect ping failed, err: %s", err.Error())
		return nil
	}

	fmt.Println("Redis connect ping success")
	return RedisClient
}

func CloseRedisConnection(client *redis.Client) {
	client.Close()
}
