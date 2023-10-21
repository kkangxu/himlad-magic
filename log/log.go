package log

import (
	"sync"
)

var logger *Log
var once sync.Once

type Log struct {
	zapLogger *zapLogger
}

// Init init logger
func Init(path, level string, options ...LogOption) {
	once.Do(func() {
		logger = &Log{zapLogger: newZapAdapter(path, level, options...)}
	})
}

// Sync flushes buffered logs (if any).
func Sync() {
	if logger == nil {
		return
	}
	logger.zapLogger.sugar.Sync()
}

type Logger interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	DPanicf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Debugw(msg string, kvs ...interface{})
	Infow(msg string, kvs ...interface{})
	Warnw(msg string, kvs ...interface{})
	Errorw(msg string, kvs ...interface{})
	DPanicw(msg string, kvs ...interface{})
	Panicw(msg string, kvs ...interface{})
	Fatalw(msg string, kvs ...interface{})
}

func Debug(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Debugf(template, args...)
}

func Debugw(msg string, kvs ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Debugw(msg, kvs...)
}

func Info(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Infof(template, args...)
}

func Infow(msg string, kvs ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Infow(msg, kvs...)
}

func Warn(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Warnf(template, args...)
}

func Warnw(msg string, kvs ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Warnw(msg, kvs...)
}

func Error(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Errorf(template, args...)
}

func Errorw(msg string, kvs ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Errorw(msg, kvs...)
}

func Panic(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Panicf(template, args...)
}

func Panicw(msg string, kvs ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Panicw(msg, kvs...)
}

func Fatal(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Fatalf(template, args...)
}

func Fatalw(msg string, kvs ...interface{}) {
	if logger == nil {
		return
	}

	logger.zapLogger.Fatalw(msg, kvs...)
}
