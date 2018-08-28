package logger

import (
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"sync"
)

const skipFrames = 8

var isTest = false

type LineHook struct {
	conf *Config
	mu   *sync.RWMutex
}

func NewLineHook(conf *Config) (*LineHook, error) {
	hook := &LineHook{
		conf: conf,
		mu:   &sync.RWMutex{},
	}
	return hook, nil
}

func (h *LineHook) Fire(entry *logrus.Entry) error {
	if h.conf.LogDetail {
		h.mu.Lock()
		defer h.mu.Unlock()
		var skip = skipFrames
		//if logger test, skip must minus 1
		if isTest {
			skip -= 1
		}
		pc, file, lineno, _ := runtime.Caller(skip)
		fu := runtime.FuncForPC(pc)
		entry.Data["file"] = file
		entry.Data["func"] = path.Base(fu.Name())
		entry.Data["line"] = lineno
	}
	return nil
}

func (h *LineHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
