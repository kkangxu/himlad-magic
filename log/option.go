package log

type LogOption interface {
	apply(*zapLogger)
}

type logOptionFunc func(*zapLogger)

func (f logOptionFunc) apply(zapLog *zapLogger) {
	f(zapLog)
}

func SetMaxFileSize(size int) LogOption {
	return logOptionFunc(func(zapLog *zapLogger) {
		zapLog.setMaxFileSize(size)
	})
}

func SetMaxBackups(n int) LogOption {
	return logOptionFunc(func(zapLog *zapLogger) {
		zapLog.setMaxBackups(n)
	})
}

func SetMaxAge(age int) LogOption {
	return logOptionFunc(func(zapLog *zapLogger) {
		zapLog.setMaxAge(age)
	})
}

func SetCompress(compress bool) LogOption {
	return logOptionFunc(func(zapLog *zapLogger) {
		zapLog.setCompress(compress)
	})
}

func SetCaller(caller bool) LogOption {
	return logOptionFunc(func(zapLog *zapLogger) {
		zapLog.setCaller(caller)
	})
}
