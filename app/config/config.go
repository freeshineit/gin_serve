package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Secret jwt
var Secret = "qwertyuiopasdfghjklzxcvbnm,.1234"

// OneDayOfHours jwt
var OneDayOfHours = 60 * 60 * 24

type Config struct {
	Name string
}

// Global Config data
var Conf *viper.Viper

// 读取配置
func (c *Config) InitConfig() error {

	Conf = viper.GetViper()

	if c.Name != "" {
		Conf.SetConfigFile(c.Name)
	} else {
		Conf.AddConfigPath("app/config")
		Conf.SetConfigName("config")
	}
	Conf.SetConfigType("toml")

	// 从环境变量总读取
	// viper.AutomaticEnv()
	// viper.SetEnvPrefix("web")
	// viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	if err := Conf.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Config file not found; ignore error if desired")
			// panic("viper.ConfigFileNotFoundError,Config file not found; ignore error if desired")
			return err
		} else {
			// Config file was found but another error was produced
			fmt.Println("Config file was found but another error was produced")
			// panic("Config file was found but another error was produced")
			return err
		}
	}

	return nil
}
