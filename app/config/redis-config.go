package config

import (
	"errors"
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

func CloseRedisConnection(client *redis.Client) {
	client.Close()
}

type RedisConfig struct {
	Host     string
	Port     int32
	Db       int32
	Password string
}

func GetRedisConfig() (RedisConfig, error) {

	host := Conf.GetString("cache.redis.host")
	port := Conf.GetInt32("cache.redis.port")
	db := Conf.GetInt32("cache.redis.db")
	password := Conf.GetString("cache.redis.password")

	if host == "" || port == 0 || db == 0 || password == "" {
		return RedisConfig{
			Host:     host,
			Port:     port,
			Db:       db,
			Password: password,
		}, errors.New("redis config missing parameter")
	}

	return RedisConfig{
		Host:     host,
		Port:     port,
		Db:       db,
		Password: password,
	}, nil
}
