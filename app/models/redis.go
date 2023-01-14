package models

import (
	"fmt"
	"go_python_serve/app/config"
	"log"

	"github.com/go-redis/redis/v8"
)

type ConfigRedis struct {
	Host     string
	Port     int32
	Db       int32
	Password string
}

func GetRedisConfig() (ConfigRedis, error) {
	return ConfigRedis{
		Host:     config.Conf.GetString("cache.redis.host"),
		Port:     config.Conf.GetInt32("cache.redis.port"),
		Db:       config.Conf.GetInt32("cache.redis.db"),
		Password: config.Conf.GetString("cache.redis.password"),
	}, nil
}

// RedisClient redis client instance
var RedisClient *redis.Client

// var ClusterClient *redis.ClusterClient

// InitRedis init redis.
// https://pkg.go.dev/github.com/go-redis/redis/v7
func InitRedis() {

	redisConfig, err := GetRedisConfig()

	if err != nil {
		log.Println(err)
		panic("redis config fail...")
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password, // no password set
		DB:       int(redisConfig.Db),  // use default DB
	})

	log.Println("Redis ping pong success...")
}
