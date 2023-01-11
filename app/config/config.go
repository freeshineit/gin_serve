package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

// 读取配置
func (c *Config) InitConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("app/config")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")

	// 从环境变量总读取
	// viper.AutomaticEnv()
	// viper.SetEnvPrefix("web")
	// viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	return viper.ReadInConfig()
}
