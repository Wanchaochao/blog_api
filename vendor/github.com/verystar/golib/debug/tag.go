package debug

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/verystar/golib/color"
)


var (
	debugFlag = "off"
	printTag  = ""
	savePath  = "./debug/"
)

func Open(flag, tag string) {
	debugFlag = flag
	printTag = tag
}

func SavePath(p string) {
	savePath = p
}

type DebugTagData struct {
	Key     string
	Data    interface{}
	Stack   CallStack
	Current string
}

type DebugTag struct {
	t    time.Time
	data []DebugTagData
	mu   sync.RWMutex
}

func NewDebugTag(options ...func(*DebugTag)) *DebugTag {
	debug := &DebugTag{}

	for _, option := range options {
		option(debug)
	}

	debug.Start()
	return debug
}

func (d *DebugTag) Start() {
	if debugFlag == "off" {
		return
	}
	d.t = time.Now()
}

func (d *DebugTag) Tag(key string, data ...interface{}) {
	if debugFlag == "off" {
		return
	}

	st := Callstack(2)
	t := time.Now().Sub(d.t).String()

	if printTag == "" || strings.Contains(key, printTag) {
		fmt.Println(color.Blue("[Debug Tag]("+t+")") + " -------------------------> " + key + " <-------------------------")
		fmt.Println(color.Green("File:" + st.File + ", Func:" + st.Func + ", Line:" + strconv.Itoa(st.LineNo)))
		if len(data) > 0 {
			format := strings.Repeat("===> %v\n", len(data))
			fmt.Println(color.Yellow(format, data...))
		}
	}

	d.mu.Lock()
	defer d.mu.Unlock()
	d.data = append(d.data, DebugTagData{
		Key:     key,
		Data:    data,
		Stack:   st,
		Current: t,
	})
}

func (d *DebugTag) Printer() {
	fmt.Println(color.Blue("[Debug]") + " -------------------------> " + time.Now().Format("2006-01-02 15:04:05") + " <-------------------------")
	buf, _ := json.MarshalIndent(d.data, "", "  ")
	fmt.Println(color.Yellow(string(buf)))
}

func (d *DebugTag) GetTagData() []DebugTagData {
	return d.data
}

func (d *DebugTag) Save(dir string, format string, prefix ...string) error {
	pre := ""
	if len(prefix) > 0 {
		pre = prefix[0] + "_"
	}
	if debugFlag == "off" {
		return nil
	}

	now := time.Now()
	s := now.Format(format)
	filename := filepath.Join(savePath, dir, pre+s+".log")
	//buf , err := json.Marshal(d.data)
	buf, err := json.MarshalIndent(d.data, "", "    ")
	if err != nil {
		return err
	}
	buffer := bytes.NewBufferString(fmt.Sprintf("\n[%v]\n", now.String()))
	buffer.Write(buf)
	buffer.WriteString("\n\n")
	return writeToFile(filename, buffer.Bytes())
}

func writeToFile(filename string, text []byte) error {
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.Write(text); err != nil {
		return err
	}
	return nil
}

func (d *DebugTag) SaveToSecond(dir string, prefix ...string) error {
	return d.Save(dir, "2006-01-02-15-04-05", prefix...)
}

func (d *DebugTag) SaveToMinute(dir string, prefix ...string) error {
	return d.Save(dir, "2006-01-02-15-04", prefix...)
}

func (d *DebugTag) SaveToHour(dir string, prefix ...string) error {
	return d.Save(dir, "2006-01-02-15", prefix...)
}

func (d *DebugTag) SaveToDay(dir string, prefix ...string) error {
	return d.Save(dir, "2006-01-02", prefix...)
}

var tagChan chan *tagData

func init() {
	tagChan = make(chan *tagData, 100)
	go func() {
		for t := range tagChan {
			tag(t.Key, t.Data...)
		}
	}()
}

type tagData struct {
	Key  string
	Data []interface{}
}

func Tag(key string, data ...interface{}) {
	tagChan <- &tagData{Key: key, Data: data}
}

func tag(key string, data ...interface{}) {
	st := Callstack(2)
	fmt.Println(color.Green("[Tag](%v) -------------------------> %s <-------------------------", time.Now().Format("2006-01-02 15:04:05"), key))
	fmt.Println(color.Green("File:%s, Func:%s, Line:%v", st.File, st.Func, st.LineNo))
	if len(data) > 0 {
		format := strings.Repeat("===> %v\n", len(data))
		fmt.Println(color.Yellow(format, data...))
	}
}
