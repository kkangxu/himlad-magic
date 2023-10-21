package morm

import (
	"context"
	"time"

	"gorm.io/gorm/logger"
)

// OrmLogger 自定义LOG
type OrmLogger struct {
	loglevel logger.LogLevel
}

// LogMode 设置log等级
func (t *OrmLogger) LogMode(level logger.LogLevel) logger.Interface {
	t.loglevel = level
	return t
}

// Info 输出info等级的log
func (t *OrmLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if t.loglevel >= logger.Info {
		// log.WithContext(ctx).Info(fmt.Sprintf(msg, data...))
	}
}

// Warn 输出warn等级的log
func (t *OrmLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if t.loglevel >= logger.Warn {
		// log.WithContext(ctx).Warn(fmt.Sprintf(msg, data...))
	}
}

// Error 输出error等级的log
func (t *OrmLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if t.loglevel >= logger.Error {
		// log.WithContext(ctx).Error(fmt.Sprintf(msg, data...))
	}
}

// Trace 输出trace等级的log
func (t *OrmLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if t.loglevel >= logger.Info {
		// sql, rows := fc()
		// log.WithContext(ctx).Debug(
		// 	fmt.Sprintf("sql: %s, affected rows: %d", sql, rows),
		// 	log.Time("time", begin),
		// )
	}
}

