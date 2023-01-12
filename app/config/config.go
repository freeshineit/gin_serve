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

type ConfigLogger struct {
	Dir       string
	HttpPath  string
	ErrorPath string
}

func GetLogger() (ConfigLogger, error) {
	fmt.Println(Conf.GetString("logger.dir"))

	return ConfigLogger{
		Dir:       Conf.GetString("logger.dir"),
		HttpPath:  Conf.GetString("logger.http_path"),
		ErrorPath: Conf.GetString("logger.error_path"),
	}, nil //
}

type ConfigMysql struct {
	Host          string
	Port          string
	Database      string
	User          string
	Password      string
	ConnectionMax int32
}

func GetMysql() (ConfigMysql, error) {

	return ConfigMysql{
		Host:          Conf.GetString("database.mysql.host"),
		Port:          Conf.GetString("database.mysql.port"),
		Database:      Conf.GetString("database.mysql.database"),
		User:          Conf.GetString("database.mysql.user"),
		Password:      Conf.GetString("database.mysql.password"),
		ConnectionMax: Conf.GetInt32("database.mysql.connection_max"),
	}, nil
}

type ConfigRedis struct {
	Host     string
	Port     int32
	Db       int32
	Password string
}

func GetRedis() (ConfigRedis, error) {
	return ConfigRedis{
		Host:     Conf.GetString("cache.redis.host"),
		Port:     Conf.GetInt32("cache.redis.port"),
		Db:       Conf.GetInt32("cache.redis.db"),
		Password: Conf.GetString("cache.redis.password"),
	}, nil
}
