package config

import (
	"fmt"
	"testing"
)

var config = Config{
	Name: "./app/config/config.toml",
}

func TestConfig(t *testing.T) {
	config.InitConfig()

	fmt.Println(Conf.GetString("cache.env"))
	fmt.Println(Conf.GetString("cache.redis.host"))
	fmt.Println(Conf.GetString("cache.redis.port"))

}
