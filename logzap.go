package logzap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const prodENV = "prod"

// Logger - структура логгера.
type Logger struct {
	zapLogger *zap.Logger
}

// NewLogger создает и возвращает новый логгер.
func NewLogger(env string) *Logger {
	config := getConfig(env)

	var options []zap.Option
	if env == prodENV {
		options = append(options, zap.WithCaller(false))
	}

	zapLogger, err := config.Build(options...)
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	return &Logger{zapLogger: zapLogger}
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

// Debug логирует сообщение с уровнем Debug.
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.zapLogger.Debug(msg, fields...)
}

// Info логирует сообщение с уровнем Info.
func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.zapLogger.Info(msg, fields...)
}

// Warn логирует сообщение с уровнем Warn.
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.zapLogger.Warn(msg, fields...)
}

// Error логирует сообщение с уровнем Error.
func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.zapLogger.Error(msg, fields...)
}

// DPanic логирует сообщение с уровнем DPanic.
func (l *Logger) DPanic(msg string, fields ...zap.Field) {
	l.zapLogger.DPanic(msg, fields...)
}

// Panic логирует сообщение с уровнем Panic.
func (l *Logger) Panic(msg string, fields ...zap.Field) {
	l.zapLogger.Panic(msg, fields...)
}

// Fatal логирует сообщение с уровнем Fatal.
func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.zapLogger.Fatal(msg, fields...)
}

// Sync завершает все буферизованные записи логов.
func (l *Logger) Sync() {
	_ = l.zapLogger.Sync()
}
