# sessions

From `gin-crontrib/sessions` only cookies are supported

https://github.com/gin-contrib/sessions/

## Usage

### Start using it

Download and install it:

```bash
$ go get github.com/ilibs/sessions
```

Import it in your code:

```go
import "github.com/ilibs/sessions"
```

## Examples

#### cookie-based

```go
package main

import (
	"github.com/ilibs/sessions"
	"github.com/ilibs/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}
```
