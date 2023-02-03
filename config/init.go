package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

// Global Config data
var Conf *Config

func getConfig(v *viper.Viper) *Config {

	databaseConfig, err := getDatabaseConfig(v)
	if err != nil {
		panic(err.Error())
	}

	redisConfig, err := getRedisConfig(v)
	if err != nil {
		panic(err.Error())
	}

	Conf = &Config{
		Logger:   getLoggerConfig(v),
		Database: databaseConfig,
		Redis:    redisConfig,
		JWT:      getJWTConfig(v),
		App:      getAppConfig(v),
		Proxy:    getProxyConfig(v),
	}

	return Conf
}

// load config
func InitConfig(Name string) error {

	c := viper.GetViper()

	if Name != "" {
		c.SetConfigFile(Name)
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

	getConfig(c)

	return nil
}

func getRedisConfig(v *viper.Viper) (RedisConfig, error) {

	host := v.GetString("cache.redis.host")
	port := v.GetInt32("cache.redis.port")
	db := v.GetInt32("cache.redis.db")
	password := v.GetString("cache.redis.password")

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

func getJWTConfig(v *viper.Viper) JWTConfig {
	secret := v.GetString("jwt.secret")
	issuer := v.GetString("jwt.issuer")

	if secret == "" {
		secret = "xiaoshaoqq@gmail.com,.<>?"
	}

	if issuer == "" {
		issuer = "xiaoshaoqq@gmail.com"
	}

	return JWTConfig{
		Secret: secret,
		Issuer: issuer,
	}
}

func getDatabaseConfig(v *viper.Viper) (DatabaseConfig, error) {

	host := v.GetString("database.mysql.host")
	port := v.GetString("database.mysql.port")
	database := v.GetString("database.mysql.database")
	user := v.GetString("database.mysql.user")
	password := v.GetString("database.mysql.password")
	connection_max := v.GetInt32("database.mysql.connection_max")

	if host == "" || port == "" || database == "" || user == "" || password == "" {
		return DatabaseConfig{
			Host:     host,
			Port:     port,
			Database: database,
			User:     user,
			Password: password,
		}, errors.New("database config missing parameter")
	}

	return DatabaseConfig{
		Host:          host,
		Port:          port,
		Database:      database,
		User:          user,
		Password:      password,
		ConnectionMax: connection_max,
	}, nil
}

// get logger config
func getLoggerConfig(v *viper.Viper) LoggerConfig {

	dir := v.GetString("logger.dir")
	httpPath := v.GetString("logger.http_path")
	errorPath := v.GetString("logger.error_path")

	if dir == "" {
		dir = "logs/"
	}

	if httpPath == "" {
		httpPath = "http/"
	}

	if errorPath == "" {
		errorPath = "error/"
	}

	return LoggerConfig{
		Dir:       dir,
		HttpPath:  httpPath,
		ErrorPath: errorPath,
	}
}

func getAppConfig(v *viper.Viper) AppConfig {
	port := v.GetString("app.port")

	if port == "" {
		port = "8080"
	}

	return AppConfig{
		Port: port,
	}
}

func getProxyConfig(v *viper.Viper) ProxyConfig {
	port := v.GetString("proxy.port")

	if port == "" {
		port = "8081"
	}

	return ProxyConfig{
		Port: port,
	}
}
