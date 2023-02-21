package config

type LoggerConfig struct {
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Level      string `mapstructure:"level"`
	Compress   bool   `mapstructure:"compress"`
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
	Db       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
}

type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	Issuer     string `mapstructure:"issuer"`
	JWTExpires int    `mapstructure:"jwt_expires"`
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

type EmailConfig struct {
	Password string `mapstructure:"password"`
}

type Config struct {
	Logger LoggerConfig `mapstructure:"logger"`
	Mysql  MysqlConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
	JWT    JWTConfig    `mapstructure:"jwt"`
	App    AppConfig    `mapstructure:"app"`
	Proxy  ProxyConfig  `mapstructure:"proxy"`
	Socket SocketConfig `mapstructure:"socket"`
	Email  EmailConfig  `mapstructure:"email"`
	Name   string
}
