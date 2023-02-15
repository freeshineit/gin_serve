package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Global Config data
var Conf *Config

// load config
func InitConfig(nameFile string) error {
	c := viper.GetViper()

	if nameFile != "" {
		c.SetConfigFile(nameFile)
	} else {
		c.AddConfigPath("config")
		c.SetConfigName("config")
	}

	c.SetConfigType("toml")

	if err := c.ReadInConfig(); err != nil {
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

	// Unmarshal toml file content
	if err := viper.Unmarshal(&Conf); err != nil {
		fmt.Println("read config error:", err)
		return err
	}
	return nil
}
