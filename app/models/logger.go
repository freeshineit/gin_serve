package models

import "go_python_serve/app/config"

type ConfigLogger struct {
	Dir       string
	HttpPath  string
	ErrorPath string
}

func GetLogger() (ConfigLogger, error) {
	// fmt.Println(Conf.GetString("logger.dir"))

	return ConfigLogger{
		Dir:       config.Conf.GetString("logger.dir"),
		HttpPath:  config.Conf.GetString("logger.http_path"),
		ErrorPath: config.Conf.GetString("logger.error_path"),
	}, nil //
}
