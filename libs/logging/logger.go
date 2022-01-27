package logging

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var (
	loggerOnce sync.Once
	logger     *Logger
)

type Logger struct {
	logger       *zap.Logger
	loggerConfig LoggerConfig
}

type LoggerConfig struct {
	Config           zap.Config
	ContextFieldFunc func(ctx context.Context) []zap.Field
}

func InfoCtx(ctx context.Context, msg string, fields ...zap.Field) {
	contextFields := GetContextFields(ctx)
	allFields := append(contextFields, fields...)
	logger.logger.Info(msg, allFields...)
}

func ErrorCtx(ctx context.Context, msg string, fields ...zap.Field) {
	contextFields := GetContextFields(ctx)
	allFields := append(contextFields, fields...)
	logger.logger.Error(msg, allFields...)
}

func WarnCtx(ctx context.Context, msg string, fields ...zap.Field) {
	contextFields := GetContextFields(ctx)
	allFields := append(contextFields, fields...)
	logger.logger.Warn(msg, allFields...)
}

func DebugCtx(ctx context.Context, msg string, fields ...zap.Field) {
	contextFields := GetContextFields(ctx)
	allFields := append(contextFields, fields...)
	logger.logger.Debug(msg, allFields...)
}

func FatalCtx(ctx context.Context, msg string, fields ...zap.Field) {
	contextFields := GetContextFields(ctx)
	allFields := append(contextFields, fields...)
	logger.logger.Fatal(msg, allFields...)
}

func PanicCtx(ctx context.Context, msg string, fields ...zap.Field) {
	contextFields := GetContextFields(ctx)
	allFields := append(contextFields, fields...)
	logger.logger.Panic(msg, allFields...)
}

func DPanicCtx(ctx context.Context, msg string, fields ...zap.Field) {
	contextFields := GetContextFields(ctx)
	allFields := append(contextFields, fields...)
	logger.logger.DPanic(msg, allFields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.logger.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.logger.Warn(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.logger.Debug(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.logger.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.logger.Panic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	logger.logger.DPanic(msg, fields...)
}

func NewLogger(config LoggerConfig) *Logger {
	config.Config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	zapLogger, _ := config.Config.Build()

	return &Logger{
		logger:       zapLogger,
		loggerConfig: config,
	}
}

func GetLogger(config LoggerConfig) *Logger {
	loggerOnce.Do(func() {
		logger = NewLogger(config)
	})

	return logger
}

func GetContextFields(ctx context.Context) []zap.Field {
	if ctx != nil && logger.loggerConfig.ContextFieldFunc != nil {
		return logger.loggerConfig.ContextFieldFunc(ctx)
	}
	return nil
}
