package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// FileHook to send logs via syslog.
type FileHook struct {
	conf  *Config
	mu    *sync.RWMutex
	cache *sync.Map
}

func NewFileHook(conf *Config) (*FileHook, error) {

	if _, err := os.Stat(conf.LogPath); err != nil {
		err = os.MkdirAll(conf.LogPath, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("can't mkdirall directory: path = %v, err = %v", conf.LogPath, err)
		}
	}

	hook := &FileHook{
		conf:  conf,
		mu:    &sync.RWMutex{},
		cache: &sync.Map{},
	}
	return hook, nil
}

func (h *FileHook) Fire(entry *logrus.Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.conf.LogMaxFiles > 0 {
		delDate := time.Now().AddDate(0, 0, -h.conf.LogMaxFiles).Format("2006-01-02")
		os.Remove(h.conf.LogPath + h.conf.LogName + "-" + delDate + ".log")
	}

	d := time.Now().Format("2006-01-02")
	logFile := filepath.Join(h.conf.LogPath, h.conf.LogName+"-"+d+".log")

	var logWriter *os.File
	f, ok := h.cache.Load(logFile)
	if !ok {
		var err error
		logWriter, err = os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("can't open file: path = %v, err = %v", logFile, err)
		}
		h.cache.Store(logFile, logWriter)
	} else {
		logWriter = f.(*os.File)
	}

	entry.Logger.Out = logWriter
	return nil
}

func (h *FileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
