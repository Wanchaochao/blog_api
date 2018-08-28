package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewFileHook(t *testing.T) {
	log := NewLogrusLogger(func(l *LogrusLogger) {
		l.Level = logrus.DebugLevel
		hook, err := NewFileHook(&Config{
			LogName:     "test",
			LogPath:     "/tmp/",
			LogLevel:    "debug",
			LogMaxFiles: 15,
			LogDetail:   true,
		})
		if err == nil {
			l.Hooks.Add(hook)
		}
	})

	log.Info("hahahahahah")
}

func BenchmarkNewFileHook(b *testing.B) {
	log := NewLogrusLogger(func(l *LogrusLogger) {
		l.Level = logrus.DebugLevel
		hook, err := NewFileHook(&Config{
			LogName:     "bench_test",
			LogPath:     "/tmp/",
			LogLevel:    "debug",
			LogMaxFiles: 15,
		})
		if err == nil {
			l.Hooks.Add(hook)
		}
	})

	for i := 0; i < b.N; i++ {
		log.Info("this is benchmark log content")
	}
}
