package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	"fmt"
)

var _ ILogger = (*LogrusLogger)(nil)

// LogrusLogger file logger
type LogrusLogger struct {
	*logrus.Logger
	Fingerprint bool
}

// NewFileLogger providers a file logger based on logrus
func NewLogrusLogger(option func(l *LogrusLogger)) ILogger {
	l := &LogrusLogger{
		Logger: &logrus.Logger{
			Out: os.Stderr,
			Formatter: &logrus.JSONFormatter{
				TimestampFormat: "2006-01-02 15:04:05",
			},
			Hooks: make(logrus.LevelHooks),
		},
	}
	option(l)
	return l
}

func (l *LogrusLogger) withFinger(format string) ILogger {
	if l.Fingerprint {
		return l.Logger.WithFields(logrus.Fields{
			"fingerprint": []string{format},
		})
	}

	return l.Logger
}

func (l *LogrusLogger) Debugf(format string, args ...interface{}) {
	l.withFinger(format).Debugf(format, args...)
}

func (l *LogrusLogger) Infof(format string, args ...interface{}) {
	l.withFinger(format).Infof(format, args...)
}

func (l *LogrusLogger) Warnf(format string, args ...interface{}) {
	l.withFinger(format).Warnf(format, args...)
}

func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	l.withFinger(format).Errorf(format, args...)
}

func (l *LogrusLogger) Fatalf(format string, args ...interface{}) {
	l.withFinger(format).Fatalf(format, args...)
}

func (l *LogrusLogger) Debug(args ...interface{}) {
	l.withFinger(argsFormat(args...)).Debug(args...)
}

func (l *LogrusLogger) Info(args ...interface{}) {
	l.withFinger(argsFormat(args...)).Info(args...)
}

func (l *LogrusLogger) Warn(args ...interface{}) {
	l.withFinger(argsFormat(args...)).Warn(args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	l.withFinger(argsFormat(args...)).Error(args...)
}

func (l *LogrusLogger) Fatal(args ...interface{}) {
	l.withFinger(argsFormat(args...)).Fatal(args...)
}

func argsFormat(args ...interface{}) string {
	format := ""
	if len(args) > 0 {
		format = fmt.Sprint(args[0])
	}

	return format
}
