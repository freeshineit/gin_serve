package config

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// SetUpZapLogger 根据prod初始化 logger
func SetUpZapLogger(prod bool) *zap.Logger {

	// development environment
	if !prod {
		logger, _ := zap.NewDevelopment()
		return logger
	}

	writeSyncer := getLogWriter(Conf.Logger.Filename, Conf.Logger.MaxSize, Conf.Logger.MaxBackups, Conf.Logger.MaxAge, Conf.Logger.Compress)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(Conf.Logger.Level))
	if err != nil {
		return nil
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}
