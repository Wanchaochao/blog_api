package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fifsky/goconf"
	"github.com/verystar/golib/debug"
	"github.com/verystar/logger"
	"fmt"
	"blog/util"
)

type common struct {
	Env         string `json:"env"`
	ConfigPath  string `json:"config_path"`
	StoragePath string `json:"storage_path"`
	Debug       string `json:"debug"`
	StatDB      string `json:"stat_db"`
	Port        string `json:"port"`
	Token 		string `json:"token"`
}

//type slbConfig struct {
//	AccessKeyId     string `json:"access_key_id"`
//	AccessKeySecret string `json:"access_key_secret"`
//	RegionId        string `json:"region_id"`
//}

type app struct {
	Common    common        `conf:"common"`
	Log       logger.Config `conf:"log"`
	//Slb       slbConfig     `conf:"slb"`
	StartTime time.Time
}

var App = &app{
	StartTime: time.Now(),
}

func init() {
	argsInit()
	Load(ExtArgs)
}

func Load(args map[string]string) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "prod"
	}

	appPath := args["config"]
	if appPath == "" {
		//获得程序路径从里面获取到app的路径
		execpath, err := os.Getwd()
		if err == nil {
			src := "src/blog" // 项目目录
			appPath = execpath[0 : strings.Index(execpath, src)+ len(src)]
		}
	}

	App.Common.ConfigPath = filepath.Join(appPath, "config")

	conf, err := goconf.NewConfig(filepath.Join(App.Common.ConfigPath, env))
	if err != nil {
		logger.Fatalf("json config path error %s", err.Error())
	}

	//load config
	if err := conf.Load(App); err != nil {
		log.Fatal("Config Error:", err)
	}

	if !filepath.IsAbs(App.Common.StoragePath) {
		App.Common.StoragePath = filepath.Join(appPath, App.Common.StoragePath)
	}

	//debug model
	if args["debug"] != "" {
		App.Common.Debug = args["debug"]
	}

	//debug
	if App.Common.Debug == "on" {
		debug.Open("on", args["debug-tag"])
		debug.SavePath(filepath.Join(App.Common.StoragePath, "debug"))
		//log level
		App.Log.LogLevel = "debug"
		//log model
		App.Log.LogMode = "std"
	}


	logger.Setting(func(c *logger.Config) {
		c.LogMode = App.Log.LogMode
		c.LogLevel = App.Log.LogLevel
		c.LogMaxFiles = App.Log.LogMaxFiles
		c.LogPath = filepath.Join(App.Common.StoragePath, "logs")
		c.LogSentryDSN = App.Log.LogSentryDSN
		c.LogSentryType = App.Log.LogSentryType
	})
	fmt.Println(util.JsonEncode(App))
}
