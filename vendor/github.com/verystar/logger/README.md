# Logger
<a href="https://travis-ci.org/verystar/logger"><img src="https://travis-ci.org/verystar/logger.svg" alt="Build Status"></a>
<a href="https://codecov.io/gh/verystar/logger"><img src="https://codecov.io/gh/verystar/logger/branch/master/graph/badge.svg" alt="codecov"></a>
<a href="https://goreportcard.com/report/github.com/verystar/logger"><img src="https://goreportcard.com/badge/github.com/verystar/logger" alt="Go Report Card
"></a>
<a href="https://godoc.org/github.com/verystar/logger"><img src="https://godoc.org/github.com/verystar/logger?status.svg" alt="GoDoc"></a>
<a href="https://opensource.org/licenses/mit-license.php" rel="nofollow"><img src="https://badges.frapsoft.com/os/mit/mit.svg?v=103"></a>
</p>


Golang logger,Integrate logrus and sentry,support logroate

## Default logger

default log out to os.Stderr

```go
import "github.com/verystar/logger"

func main(){
    logger.Debug("debug printer %s","hehe")
    logger.Fatal("exit")
}
```

## Setting default logger
```
logger.Setting(func(c *logger.Config) {
    c.LogMode = "file"
    c.LogLevel = "info"
    c.LogMaxFiles = 15  //store for up to 15 days
    c.LogPath = "/tmp/logs/"
    c.LogSentryDSN = ""
    c.LogSentryType = ""
    c.LogDetail = true
})
```

## New logger

```
log := logger.NewLogger(func(c *logger.Config) {
    c.LogMode = "file"
    c.LogLevel = "info"
    c.LogMaxFiles = 15  //store for up to 15 days
    c.LogPath = "/tmp/logs/"
    c.LogSentryDSN = ""
    c.LogSentryType = ""
})

log.Debug("this is new log")
```

## Print filename and line no

if `LogDetail` is true,the log data add filename and line no

```
{"file":"/Users/fifsky/wwwroot/go/library/src/github.com/fifsky/goblog/handler/index.go","func":"handler.IndexGet","level":"debug","line":16,"msg":"[test]","time":"2018-08-02 22:37:02"}
```