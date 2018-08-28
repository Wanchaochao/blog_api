package counter

import (
	"context"
	"errors"
	"sync/atomic"
	"time"
)

type Ceil struct {
	Num int32
}

type CeilList struct {
	Ceils map[*Counter]*Ceil
	Unix  int64
}

type Timeline struct {
	Container  [] *CeilList `json:"container"`    // 数据容器
	MaxMemTime int64        `json:"max_mem_time"` // 最大记忆秒数
	Counters   []*Counter   `json:"counters"`
	Context    context.Context
}

func NewTimeline(max_mem_time int64, ctx context.Context) *Timeline {
	l := &Timeline{
		Container:  make([] *CeilList, 0),
		MaxMemTime: max_mem_time,
		Counters:   make([]*Counter, 0),
		Context:    ctx,
	}
	return l
}

func (l *Timeline) AddCounter(c *Counter) error {
	if c.Duration > l.MaxMemTime {
		return errors.New("counter duration should not gt max_mem_time")
	}
	l.Counters = append(l.Counters, c)
	return nil
}

func (l *Timeline) Save() {

}

func (l *Timeline) Start() {
	go l.Loop()
}

func (l *Timeline) StartWithContext(ctx context.Context) {
	l.Context = ctx
	go l.Loop()
}

func (l *Timeline) Loop() {
	timer := time.NewTicker(time.Second)
	defer timer.Stop()

	for {
		select {
		case t := <-timer.C:
			unix := t.Unix()
			// 收集 Couter信息 ， 添加头部
			l.Push(unix)
			// 剔除尾部
			l.Pop(unix)
			// 通知所有的 符合条件的counter
			go l.Notify(t)
		case <-l.Context.Done():
			return
		}
	}
}

// 收集数据
func (l *Timeline) Push(unix int64) {
	cc := &CeilList{
		Ceils: make(map[*Counter]*Ceil),
		Unix:  unix,
	}
	for _, c := range l.Counters {
		cc.Ceils[c] = &Ceil{
			Num: c.CurrentNum,
		}
		// 重置
		atomic.StoreInt32(&c.CurrentNum, 0)
	}
	l.Container = append(l.Container, cc)
}

// pop过期数据
func (l *Timeline) Pop(unix int64) {

	ll := len(l.Container)
	if ll == 0 {
		return
	}

	// 当前时间 - 最大记忆时间 > 最后一条数据的 unix
	if unix-l.MaxMemTime > l.Container[0].Unix {
		// 删除最后一个
		l.Container = l.Container[1:]
	}
}
func (l *Timeline) Notify(t time.Time) {
	unix := t.Unix()
	for _, c := range l.Counters {

		// 如果是在静默期
		//fmt.Println("c.LastNotifyUnix+c.Quite > unix ", c.LastNotifyUnix+c.Quite > unix)
		if c.LastNotifyUnix+c.Quite > unix {
			continue
		}

		var sum int32
		for i := len(l.Container) - 1; i >= 0; i-- {
			con := l.Container[i]
			sum += con.Ceils[c].Num
			// 如果超出了 counter 的收集时间
			//fmt.Println("unix-c.Duration > con.Unix ", unix-c.Duration > con.Unix)
			if unix-c.Duration > con.Unix {
				break
			}
		}

		//fmt.Println("sum ", sum)
		if sum > c.Max {
			c.Notify(t, sum)
			c.LastNotifyUnix = unix
		}
	}
}

// 对counter 单独计数
func (l *Timeline) CounterSum(c *Counter, t time.Time) int32 {
	var sum int32
	unix := t.Unix()
	for i := len(l.Container) - 1; i >= 0; i-- {
		con := l.Container[i]
		sum += con.Ceils[c].Num
		// 如果超出了 counter 的收集时间
		if unix-c.Duration > con.Unix {
			break
		}
	}
	return sum
}
