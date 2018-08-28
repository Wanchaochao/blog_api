package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewLineHook(t *testing.T) {
	log := NewLogrusLogger(func(l *LogrusLogger) {
		l.Level = logrus.DebugLevel
		hook, err := NewLineHook(&Config{
			LogName:   "test",
			LogLevel:  "debug",
			LogDetail: true,
		})
		if err == nil {
			l.Hooks.Add(hook)
		}
	})

	log.Info("hahahahahah")
}
