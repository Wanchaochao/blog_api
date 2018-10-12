package logger

import (
	"time"

	"github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"
)

type Config struct {
	LogName       string `toml:"log_name" json:"log_name"`
	LogPath       string `toml:"log_path" json:"log_path"`
	LogMode       string `toml:"log_mode" json:"log_mode"`
	LogLevel      string `toml:"log_level" json:"log_level"`
	LogDetail     bool   `toml:"log_detail" json:"log_detail"`
	LogMaxFiles   int    `toml:"log_max_files" json:"log_max_files"`
	LogSentryDSN  string `toml:"log_sentry_dsn" json:"log_sentry_dsn"`
	LogSentryType string `toml:"log_sentry_type" json:"log_sentry_type"`
}

var (
	// std is the name of the standard logger in stdlib `log`
	std           ILogger
	defaultConfig *Config
)

func init() {
	defaultConfig = &Config{
		LogName:  "app",
		LogMode:  "std",
		LogLevel: "debug",
	}
	std = newLogger(defaultConfig)
}

func Setting(option func(*Config)) {
	option(defaultConfig)
	std = newLogger(defaultConfig)
}

func NewLogger(options ...func(*Config)) ILogger {
	//copy default config
	conf := *defaultConfig
	for _, option := range options {
		option(&conf)
	}
	return newLogger(&conf)
}

func newLogger(conf *Config) ILogger {
	return NewLogrusLogger(func(l *LogrusLogger) {
		l.Level, _ = logrus.ParseLevel(conf.LogLevel)

		if conf.LogDetail {
			hook, err := NewLineHook(conf)
			if err == nil {
				l.Hooks.Add(hook)
			}
		}

		if conf.LogMode == "file" {
			hook, err := NewFileHook(conf)
			if err == nil {
				l.Hooks.Add(hook)
			}
		}

		if conf.LogSentryDSN != "" {
			l.Fingerprint = true
			tags := map[string]string{
				"type": conf.LogSentryType,
			}

			hook, err := logrus_sentry.NewWithTagsSentryHook(conf.LogSentryDSN, tags, []logrus.Level{
				logrus.PanicLevel,
				logrus.FatalLevel,
				logrus.ErrorLevel,
				logrus.WarnLevel,
				logrus.InfoLevel,
			})
			hook.Timeout = 1 * time.Second
			hook.StacktraceConfiguration.Enable = true

			if err == nil {
				l.Hooks.Add(hook)
			}
		}
	})
}

func Debugf(str string, args ...interface{}) {
	std.Debugf(str, args...)
}

func Infof(str string, args ...interface{}) {
	std.Infof(str, args...)
}

func Warnf(str string, args ...interface{}) {
	std.Warnf(str, args...)
}

func Errorf(str string, args ...interface{}) {
	std.Errorf(str, args...)
}

func Fatalf(str string, args ...interface{}) {
	std.Fatalf(str, args...)
}

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Warn(args ...interface{}) {
	std.Warn(args...)
}

func Error(args ...interface{}) {
	std.Error(args...)
}

func Fatal(args ...interface{}) {
	std.Fatal(args...)
}
