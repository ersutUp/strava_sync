//gorm的日志对接到logurs
package log

import (
	"context"
	"github.com/sirupsen/logrus"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

type Gorm2logrus struct {
	SlowThreshold time.Duration
}

func (l *Gorm2logrus) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return l
}

func (l *Gorm2logrus) Info(ctx context.Context, msg string, data ...interface{}) {
	if logrus.InfoLevel <= logrus.GetLevel() {
		logrus.
			WithField(LogForm, Gorm).
			Infof(msg, append([]interface{}{}, data...)...)
	}
}
func (l *Gorm2logrus) Warn(ctx context.Context, msg string, data ...interface{}) {
	if logrus.WarnLevel <= logrus.GetLevel() {
		logrus.
			WithField(LogForm, Gorm).
			Warnf(msg, append([]interface{}{}, data...)...)
	}
}
func (l *Gorm2logrus) Error(ctx context.Context, msg string, data ...interface{}) {
	if logrus.ErrorLevel <= logrus.GetLevel() {
		logrus.
			WithField(LogForm, Gorm).
			Errorf(msg, append([]interface{}{}, data...)...)
	}
}
func (l *Gorm2logrus) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	sql, rows := fc()

	var log = logrus.
		WithField(LogForm, Gorm).
		WithField("sqlHandleMS", float64(elapsed.Nanoseconds())/1e6).
		WithField("sql", sql).
		WithField("rows", rows)

	switch {
	case err != nil:
		if logrus.ErrorLevel <= logrus.GetLevel() {
			log.Error(err)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0:
		if logrus.WarnLevel <= logrus.GetLevel() {
			log.WithField("slowSql", true).Warn("SLOW SQL >=", l.SlowThreshold)
		}
	default:
		if logrus.InfoLevel <= logrus.GetLevel() {
			log.Info()
		}
	}
}
