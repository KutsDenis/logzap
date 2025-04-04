package logzap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const prodENV = "prod"

// NewLogger создает и возвращает новый логгер.
func NewLogger(env string) *zap.Logger {
	config := getConfig(env)

	var options []zap.Option
	if env == prodENV {
		options = append(options, zap.WithCaller(false))
	}

	zapLogger, err := config.Build(options...)
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	return zapLogger
}

// getConfig возвращает конфигурацию логгера в зависимости от окружения.
func getConfig(env string) zap.Config {
	var config zap.Config
	if env == prodENV {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return config
}
