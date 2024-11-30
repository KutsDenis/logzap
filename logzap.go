package logzap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// Init инициализирует zap-логгер с конфигурацией для нужной среды.
func Init(env string) {
	var config zap.Config
	if env == "prod" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	Logger, err = config.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
}

// Sync завершает все буферизованные записи логов. Должен быть вызван с defer в функции main.
func Sync() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}
