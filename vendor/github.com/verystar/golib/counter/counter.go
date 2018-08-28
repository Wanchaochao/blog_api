package counter

import (
	"sync/atomic"
	"time"
)

type Handler func(c *Counter, t time.Time, sum int32)

type Counter struct {
	Duration       int64 // 统计时长 单位秒
	Quite          int64 // 静默期 单位秒
	Max            int32 // 触发notify的最大值
	CurrentNum     int32 // 秒级别的统计数据量
	LastNotifyUnix int64 // 最后的notify时间
	Handler        Handler
	Mata           map[string]interface{}
}

func NewCounter(duration int64, quite int64, max int32) *Counter {
	c := &Counter{
		Duration: duration,
		Quite:    quite,
		Max:      max,
		Mata:     make(map[string]interface{}),
	}
	return c
}

func (c *Counter) Add(i int32) {
	atomic.AddInt32(&c.CurrentNum, i)
}

func (c *Counter) AddHandle(fn Handler) {
	c.Handler = fn
}

func (c *Counter) Notify(t time.Time, sum int32) {
	if c.Handler != nil {
		c.Handler(c, t, sum)
	}
}
