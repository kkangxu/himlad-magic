package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type zapLogger struct {
	Path        string
	Level       string
	MaxFileSize int
	MaxBackups  int
	MaxAge      int
	Compress    bool
	Caller      bool
	sugar       *zap.SugaredLogger
}

func newZapAdapter(path, level string, options ...LogOption) *zapLogger {
	zapLog := &zapLogger{
		Path:        path,
		Level:       level,
		MaxFileSize: 1024,
		MaxBackups:  3,
		MaxAge:      7,
		Compress:    true,
		Caller:      false,
	}

	for _, opt := range options {
		opt.apply(zapLog)
	}

	return zapLog.build()
}

func (z *zapLogger) build() *zapLogger {

	var level zapcore.Level
	switch z.Level {
	case zap.DebugLevel.String():
		level = zap.DebugLevel
	case zap.InfoLevel.String():
		level = zap.InfoLevel
	case zap.WarnLevel.String():
		level = zap.WarnLevel
	case zap.ErrorLevel.String():
		level = zap.ErrorLevel
	case zap.DPanicLevel.String():
		level = zap.DPanicLevel
	case zap.PanicLevel.String():
		level = zap.PanicLevel
	case zap.FatalLevel.String():
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), zapcore.AddSync(z.newLumberjack()), level)

	l := zap.New(core)
	if z.Caller {
		l = l.WithOptions(zap.AddCaller(), zap.AddCallerSkip(2), zap.AddStacktrace(zapcore.ErrorLevel))
	}

	z.sugar = l.Sugar()

	return z
}

func (z *zapLogger) newLumberjack() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   z.Path,
		MaxSize:    z.MaxFileSize,
		MaxBackups: z.MaxBackups,
		MaxAge:     z.MaxAge,
		Compress:   z.Compress,
	}
}

func (z *zapLogger) setMaxFileSize(size int) {
	z.MaxFileSize = size
}

func (z *zapLogger) setMaxBackups(n int) {
	z.MaxBackups = n
}

func (z *zapLogger) setMaxAge(age int) {
	z.MaxAge = age
}

func (z *zapLogger) setCompress(compress bool) {
	z.Compress = compress
}

func (z *zapLogger) setCaller(caller bool) {
	z.Caller = caller
}

func (z *zapLogger) Debug(args ...interface{}) {
	z.sugar.Debug(args...)
}

func (z *zapLogger) Info(args ...interface{}) {
	z.sugar.Info(args...)
}

func (z *zapLogger) Warn(args ...interface{}) {
	z.sugar.Warn(args...)
}

func (z *zapLogger) Error(args ...interface{}) {
	z.sugar.Error(args...)
}

func (z *zapLogger) DPanic(args ...interface{}) {
	z.sugar.DPanic(args...)
}

func (z *zapLogger) Panic(args ...interface{}) {
	z.sugar.Panic(args...)
}

func (z *zapLogger) Fatal(args ...interface{}) {
	z.sugar.Fatal(args...)
}

func (z *zapLogger) Debugf(template string, args ...interface{}) {
	z.sugar.Debugf(template, args...)
}

func (z *zapLogger) Infof(template string, args ...interface{}) {
	z.sugar.Infof(template, args...)
}

func (z *zapLogger) Warnf(template string, args ...interface{}) {
	z.sugar.Warnf(template, args...)
}

func (z *zapLogger) Errorf(template string, args ...interface{}) {
	z.sugar.Errorf(template, args...)
}

func (z *zapLogger) DPanicf(template string, args ...interface{}) {
	z.sugar.DPanicf(template, args...)
}

func (z *zapLogger) Panicf(template string, args ...interface{}) {
	z.sugar.Panicf(template, args...)
}

func (z *zapLogger) Fatalf(template string, args ...interface{}) {
	z.sugar.Fatalf(template, args...)
}

func (z *zapLogger) Debugw(msg string, kvs ...interface{}) {
	z.sugar.Debugw(msg, kvs...)
}

func (z *zapLogger) Infow(msg string, kvs ...interface{}) {
	z.sugar.Infow(msg, kvs...)
}

func (z *zapLogger) Warnw(msg string, kvs ...interface{}) {
	z.sugar.Warnw(msg, kvs...)
}

func (z *zapLogger) Errorw(msg string, kvs ...interface{}) {
	z.sugar.Errorw(msg, kvs...)
}

func (z *zapLogger) DPanicw(msg string, kvs ...interface{}) {
	z.sugar.DPanicw(msg, kvs...)
}

func (z *zapLogger) Panicw(msg string, kvs ...interface{}) {
	z.sugar.Panicw(msg, kvs...)
}

func (z *zapLogger) Fatalw(msg string, kvs ...interface{}) {
	z.sugar.Fatalw(msg, kvs...)
}
