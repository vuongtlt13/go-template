package logger

import (
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger represents the logger interface
type Logger interface {
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	With(args ...interface{}) *zap.Logger
	Sync() error
}

// AppLogger wraps zap.Logger
type AppLogger struct {
	zapLogger *zap.Logger
}

var (
	instance Logger
	once     sync.Once
)

// GetLogger returns the singleton logger instance
func GetLogger() Logger {
	once.Do(func() {
		instance = newLogger()
	})
	return instance
}

// newLogger creates a new logger instance
func newLogger() Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig

	logger, err := config.Build()
	if err != nil {
		// If we can't build the logger, fallback to a basic one
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			zapcore.InfoLevel,
		)
		logger = zap.New(core)
	}

	return &AppLogger{zapLogger: logger.WithOptions(zap.AddCallerSkip(1))}
}

// NewDevelopmentLogger creates a logger suitable for development (with more verbose output)
func NewDevelopmentLogger() Logger {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig = encoderConfig

	logger, err := config.Build()
	if err != nil {
		// If we can't build the logger, fallback to a basic one
		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			zapcore.DebugLevel,
		)
		logger = zap.New(core)
	}

	return &AppLogger{zapLogger: logger.WithOptions(zap.AddCallerSkip(1))}
}

// Info implements Logger
func (l *AppLogger) Info(msg string, args ...interface{}) {
	l.zapLogger.Info(msg, toZapFields(args...)...)
}

// Warn implements Logger
func (l *AppLogger) Warn(msg string, args ...interface{}) {
	l.zapLogger.Warn(msg, toZapFields(args...)...)
}

// Error implements Logger
func (l *AppLogger) Error(msg string, args ...interface{}) {
	l.zapLogger.Error(msg, toZapFields(args...)...)
}

// Fatal implements Logger
func (l *AppLogger) Fatal(msg string, args ...interface{}) {
	l.zapLogger.Fatal(msg, toZapFields(args...)...)
}

// Debug implements Logger
func (l *AppLogger) Debug(msg string, args ...interface{}) {
	l.zapLogger.Debug(msg, toZapFields(args...)...)
}

// With implements Logger
func (l *AppLogger) With(args ...interface{}) *zap.Logger {
	return l.zapLogger.With(toZapFields(args...)...)
}

// Sync implements Logger
func (l *AppLogger) Sync() error {
	return l.zapLogger.Sync()
}

// toZapFields converts a list of key-value pairs to zap.Fields
func toZapFields(args ...interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(args)/2)
	for i := 0; i < len(args)-1; i += 2 {
		key, ok := args[i].(string)
		if !ok {
			continue
		}
		fields = append(fields, zap.Any(key, args[i+1]))
	}
	return fields
}

// Field creation helpers
func Float64(key string, val float64) zapcore.Field {
	return zap.Float64(key, val)
}

func String(key string, val string) zapcore.Field {
	return zap.String(key, val)
}

func Int(key string, val int) zapcore.Field {
	return zap.Int(key, val)
}

func Error(err error) zapcore.Field {
	return zap.Error(err)
}

func Duration(key string, val time.Duration) zapcore.Field {
	return zap.Duration(key, val)
}
