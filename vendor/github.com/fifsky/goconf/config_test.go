package goconf

import (
	"testing"
	"path/filepath"

	"github.com/tidwall/gjson"
)

func TestNewConfig(t *testing.T) {
	_, err := NewConfig("./testdata/")
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewConfig("./testdata2/")
	if err == nil {
		t.Fatalf("testdata2 must return not found error")
	}
}

func TestConfig_Get(t *testing.T) {
	conf, err := NewConfig("./testdata/")
	if err != nil {
		t.Fatal(err)
	}

	if ret, err := conf.Get("dev.name.last"); err != nil || ret.String() != "Anderson" {
		t.Errorf("get key %s test error", "dev.name.last")
	}

	if ret, err := conf.Get("dev.age"); err != nil || ret.Int() != 37 {
		t.Errorf("get key %s test error", "dev.age")
	}

	if ret, err := conf.Get("dev.childen"); err != nil || ret.IsArray() {
		t.Errorf("get key %s test error", "dev.childen")
	}

	if ret, err := conf.Get("dev.friends.1.age"); err != nil || ret.Int() != 68 {
		t.Errorf("get key %s test error", "dev.friends.1.age")
	}

	if ret, err := conf.Get("prod.widget.window.width"); err != nil || ret.Int() != 500 {
		t.Errorf("get key %s test error", "prod.window.width")
	}

	if ret, err := conf.Get("prod.image2.alignment"); err != nil || ret.String() != "" {
		t.Errorf("get key %s value must empty", "prod.image2.alignment")
	}

	if _, err := conf.Get("dev2.notfound"); err == nil {
		t.Errorf("get key %s must return err", "dev2.notfound")
	}
}

func TestConfig_MustGet(t *testing.T) {
	conf, err := NewConfig("./testdata/")
	if err != nil {
		t.Fatal(err)
	}

	if conf.MustGet("dev.name.last").String() != "Anderson" {
		t.Errorf("must get key %s test error", "dev.name.last")
	}

	if conf.MustGet("dev.age").Int() != 37 {
		t.Errorf("must get key %s test error", "dev.age")
	}

	if conf.MustGet("dev.childen").IsArray() {
		t.Errorf("must get key %s test error", "dev.childen")
	}

	if conf.MustGet("dev.friends.1.age").Int() != 68 {
		t.Errorf("must get key %s test error", "dev.friends.1.age")
	}

	if conf.MustGet("prod.widget.window.width").Int() != 500 {
		t.Errorf("must get key %s test error", "prod.window.width")
	}

	if conf.MustGet("prod.image2.alignment").String() != "" {
		t.Errorf("must get key %s value must empty", "prod.image2.alignment")
	}

	if conf.MustGet("dev.name.notfound").String() != "" {
		t.Errorf("must get key %s must return empty", "dev.name.notfound")
	}

	if conf.MustGet("dev.name.notfound2").Int() != 0 {
		t.Errorf("must get key %s must return 0", "dev.name.notfound2")
	}

	if conf.MustGet("dev2.notfound").String() != "" {
		t.Errorf("must get key %s must return empty", "dev2.notfound")
	}

	if conf.MustGet("dev2.notfound2").Int() != 0 {
		t.Errorf("must get key %s must return 0", "dev2.notfound2")
	}
}

func TestConfig_Cache(t *testing.T) {
	conf, err := NewConfig("./testdata/")
	if err != nil {
		t.Fatal(err)
	}

	name := conf.MustGet("dev.name.first").String()

	file := filepath.Join(conf.Path, "dev"+conf.Ext)

	cache, ok := conf.cache.Load(file)
	if !ok || !cache.(gjson.Result).IsObject() {
		t.Fatalf("config cache miss")
	}

	cacheName := cache.(gjson.Result).Get("name.first").String()

	if name != "" && name != cacheName {
		t.Fatalf("cache value not match [%s:%s]", name, cacheName)
	}
}

func TestConfig_Unmarshal(t *testing.T) {
	type database struct {
		Driver string `json:"driver"`
		Host   string `json:"host"`
		Port   int    `json:"port"`
	}

	type configDemo struct {
		Database database `json:"database"`
	}

	conf, err := NewConfig("./testdata/")
	if err != nil {
		t.Fatal(err)
	}

	{
		app := &configDemo{}
		err = conf.Unmarshal("json5", app)

		if err != nil {
			t.Fatal(err)
		}

		if app.Database.Host != "localhost" {
			t.Fatalf("Unmarshal struct host must return %s", "localhost")
		}

		if app.Database.Port != 3306 {
			t.Fatalf("Unmarshal struct port must return %d", 3306)
		}
		//fmt.Println(app)
	}

	{
		db := &database{}
		err = conf.Unmarshal("json5.database", db)

		if err != nil {
			t.Fatal(err)
		}

		if db.Host != "localhost" {
			t.Fatalf("Unmarshal Xpath struct host must return %s", "localhost")
		}
	}
}

func TestConfig_UnmarshalSlice(t *testing.T) {
	type config struct {
		Children []string
	}

	app := config{}

	conf, err := NewConfig("./testdata/")
	if err != nil {
		t.Fatal(err)
	}
	err = conf.Unmarshal("dev.children", &app.Children)

	if err != nil {
		t.Fatalf("UnmarshalSlice error:%s", err)
	}

	if len(app.Children) == 0 {
		t.Fatalf("UnmarshalSlice error")
	}
}

func TestConfig_Load(t *testing.T) {
	type log struct {
		LogName string `json:"log_name"`
		LogPath string `json:"log_path"`
	}

	type db struct {
		Name string `json:"name"`
		Host string `json:"host"`
		Port string `json:"port"`
	}

	type config struct {
		Log log `conf:"log"`
		DB  db  `conf:"db"`
	}

	conf, err := NewConfig("./testdata/")
	if err != nil {
		t.Fatal(err)
	}

	app := &config{}
	err = conf.Load(app)

	if err != nil {
		t.Error(err)
	}

	if app.Log.LogName != "log" {
		t.Fatalf("config load log.name must get 'log' but get '%s'", app.Log.LogName)
	}

	if app.DB.Host != "localhost" {
		t.Fatalf("config load db.host must get 'localhost' but get '%s'", app.DB.Host)
	}
}
