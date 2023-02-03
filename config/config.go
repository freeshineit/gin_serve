package config

type LoggerConfig struct {
	Dir       string
	HttpPath  string
	ErrorPath string
}

type DatabaseConfig struct {
	Host          string
	Port          string
	Database      string
	User          string
	Password      string
	ConnectionMax int32
}
type RedisConfig struct {
	Host     string
	Port     int32
	Db       int32
	Password string
}

type JWTConfig struct {
	Secret string
	Issuer string
}

type AppConfig struct {
	Port string
}

type ProxyConfig struct {
	Port string
}

type SocketConfig struct {
	Port string
}

type Config struct {
	Logger   LoggerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
	App      AppConfig
	Proxy    ProxyConfig
	Socket   SocketConfig
	Name     string
}
