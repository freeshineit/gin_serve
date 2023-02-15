package config

type LoggerConfig struct {
	Dir       string `mapstructure:"dir"`
	Path      string `mapstructure:"path"`
	ErrorPath string `mapstructure:"errorPath"`
}

type MysqlConfig struct {
	Host          string `mapstructure:"host"`
	Port          string `mapstructure:"port"`
	Database      string `mapstructure:"database"`
	User          string `mapstructure:"user"`
	Password      string `mapstructure:"password"`
	ConnectionMax int32  `mapstructure:"connectionMax"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int32  `mapstructure:"port"`
	Db       int32  `mapstructure:"db"`
	Password string `mapstructure:"password"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Issuer string `mapstructure:"issuer"`
}

type AppConfig struct {
	Port string `mapstructure:"port"`
}

type ProxyConfig struct {
	Port string `mapstructure:"port"`
}

type SocketConfig struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	Logger LoggerConfig `mapstructure:"logger"`
	Mysql  MysqlConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
	JWT    JWTConfig    `mapstructure:"jwt"`
	App    AppConfig    `mapstructure:"app"`
	Proxy  ProxyConfig  `mapstructure:"proxy"`
	Socket SocketConfig `mapstructure:"socket"`
	Name   string
}
