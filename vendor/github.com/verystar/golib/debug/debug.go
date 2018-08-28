package debug

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/verystar/golib/color"
)

var (
	debugFlag = "off"
	printTag  = ""
	savePath  = "./debug/"
	debugLock sync.Mutex
)

func Open(flag, tag string) {
	debugFlag = flag
	printTag = tag
}

func SavePath(p string) {
	savePath = p
}

func Debugf(format string, args ...interface{}) {
	if debugFlag == "on" {
		pc, file, line, ok := runtime.Caller(1)
		if !ok {
			file = "<?>"
			line = 0
		} else {
			if i := strings.LastIndex(file, "/"); i >= 0 {
				file = file[i+1:]
			}
		}
		fn := runtime.FuncForPC(pc).Name()
		fnparts := strings.Split(fn, "/")
		t := time.Now()
		hour, min, sec := t.Clock()
		nanosec := t.Nanosecond() / 1e3

		debugLock.Lock()
		fmt.Printf("DEBUG: [%02d:%02d:%02d.%04d] %s:%d(%s): %s\n",
			hour, min, sec, nanosec,
			file, line, color.Red(fnparts[len(fnparts)-1]),
			fmt.Sprintf(format, args...))
		debugLock.Unlock()
	}
}
