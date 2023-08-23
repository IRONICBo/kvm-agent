package log

import (
	"fmt"
	"kvm-agent/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func InitLogger(config config.App) error {
	_logger, err := NewLoggerWithConfig(config)
	if err != nil {
		return err
	}

	logger = _logger

	return nil
}

func NewLoggerWithConfig(config config.App) (*zap.Logger, error) {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/agent.log", config.LogFile),
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   true,
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		TimeKey:      "time",
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(lumberjackLogger),
		zap.InfoLevel,
	)

	// Change to debug level
	if config.Debug {
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	}

	_logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return _logger, nil
}

func GetLogger() *zap.Logger {
	return logger
}

func CloseLogger() {
	if logger != nil {
		err := logger.Sync()
		if err != nil {
			fmt.Println("Logger", "close failed, err:", err.Error())
		}
	}
}

// Info log info.
func Info(Operation string, args ...interface{}) {
	logger.Sugar().Infof("[Operation:%s], %v", Operation, args)
}

// Error log error.
func Error(Operation string, args ...interface{}) {
	logger.Sugar().Errorf("[Operation:%s], %v", Operation, args)
}

// Debug log debug.
func Debug(Operation string, args ...interface{}) {
	logger.Sugar().Debugf("[Operation:%s], %v", Operation, args)
}

// Panic log panic.
func Panic(Operation string, args ...interface{}) {
	logger.Sugar().Panicf("[Operation:%s], %v", Operation, args)
}

// Infof log info with format.
func Infof(Operation string, format string, args ...interface{}) {
	_format := fmt.Sprintf("[Operation:%s], %s", Operation, format)
	logger.Sugar().Infof(_format, args)
}

// Errorf log error with format.
func Errorf(Operation string, format string, args ...interface{}) {
	_format := fmt.Sprintf("[Operation:%s], %s", Operation, format)
	logger.Sugar().Errorf(_format, args)
}

// Debugf log debug with format.
func Debugf(Operation string, format string, args ...interface{}) {
	_format := fmt.Sprintf("[Operation:%s], %s", Operation, format)
	logger.Sugar().Debugf(_format, args)
}

// Panicf log panic with format.
func Panicf(Operation string, format string, args ...interface{}) {
	_format := fmt.Sprintf("[Operation:%s], %s", Operation, format)
	logger.Sugar().Panicf(_format, args)
}
