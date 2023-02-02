package config

type LoggerConfig struct {
	Dir       string
	HttpPath  string
	ErrorPath string
}

func GetLoggerConfig() LoggerConfig {

	dir := Conf.GetString("logger.dir")
	httpPath := Conf.GetString("logger.http_path")
	errorPath := Conf.GetString("logger.error_path")

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
