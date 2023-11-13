package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// global logger
var l Logger

type ZapLogger struct {
	logger *zap.Logger
}

func Initialize(isDebug bool) Logger {
	var logger *zap.Logger
	if isDebug {
		logger = newDevelopmentLogger()
	} else {
		logger = newProductionLogger()
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	l = &ZapLogger{
		logger: logger,
	}
	return l
}

func Debug(msg string, kvs map[string]interface{}) {
	l.Debug(msg, kvs)
}

func Debugf(msg string, args ...any) {
	l.Debugf(msg, args)
}

func Info(msg string, kvs map[string]any) {
	l.Info(msg, kvs)
}

func Infof(msg string, args ...any) {
	l.Infof(msg, args)
}

func Error(msg string, err error) {
	l.Error(msg, err)
}

func Errorf(msg string, args ...any) {
	l.Errorf(msg, args)
}

func Panic(msg string, err error) {
	l.Panic(msg, err)
}

func Panicf(msg string, args ...any) {
	l.Panicf(msg, args)
}

func (l *ZapLogger) Debug(msg string, kvs map[string]any) {
	fileds := make([]zap.Field, 0)
	for k, v := range kvs {
		fileds = append(fileds, zap.Any(k, v))
	}
	l.logger.Debug(msg, fileds...)
}

func (l *ZapLogger) Debugf(msg string, args ...any) {
	l.logger.Debug(fmt.Sprintf(msg, args))
}

func (l *ZapLogger) Info(msg string, kvs map[string]any) {
	fileds := make([]zap.Field, 0)
	for k, v := range kvs {
		fileds = append(fileds, zap.Any(k, v))
	}
	l.logger.Info(msg, fileds...)
}

func (l *ZapLogger) Infof(msg string, args ...any) {
	l.logger.Info(fmt.Sprintf(msg, args))
}

func (l *ZapLogger) Error(msg string, err error) {
	l.logger.Error(msg, zap.Error(err))
}

func (l *ZapLogger) Errorf(msg string, args ...any) {
	l.logger.Error(fmt.Sprintf(msg, args))
}

func (l *ZapLogger) Panic(msg string, err error) {
	l.logger.Panic(msg, zap.Error(err))
}

func (l *ZapLogger) Panicf(msg string, args ...any) {
	l.logger.Panic(fmt.Sprintf(msg, args))
}

func (l *ZapLogger) Shutdown() error {
	return l.logger.Sync()
}

func newProductionLogger() *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		// Filename:   "/var/log/mxshop/userweb/info.log",
		Filename:   "./log/info.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			MessageKey:    "msg",
			LevelKey:      "level",
			TimeKey:       "ts",
			CallerKey:     "caller",
			StacktraceKey: "trace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05"))
			},
			EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendInt64(int64(d) / 1000000)
			},
		}),
		w,
		zap.InfoLevel,
	)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
}

func newDevelopmentLogger() *zap.Logger {
	w := zapcore.AddSync(os.Stdout)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			MessageKey:    "msg",
			LevelKey:      "level",
			TimeKey:       "ts",
			CallerKey:     "caller",
			StacktraceKey: "trace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05"))
			},
			EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendInt64(int64(d) / 1000000)
			},
		}),
		w,
		zap.DebugLevel,
	)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
}
