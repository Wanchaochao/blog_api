package task

import (
	"context"
	"sync/atomic"
	"time"
)

type ITask interface {
	Do()
}

type Dispatcher struct {
	Context       context.Context
	TaskChan      chan ITask
	CtrChan       chan struct{}
	counter       int32
	IsStopAddTask bool
	StopChan      chan struct{}
	IoWait        time.Duration
}

func NewDispatcher(ctx context.Context, maxTaskNum int) *Dispatcher {
	d := &Dispatcher{
		Context:  ctx,
		TaskChan: make(chan ITask, maxTaskNum),
		CtrChan:  make(chan struct{}, maxTaskNum),
		StopChan: make(chan struct{}),
		IoWait:   50 * time.Millisecond,
	}
	go d.run()
	return d
}

func (d *Dispatcher) SetIoWait(duration time.Duration) {
	d.IoWait = duration
}
func (d *Dispatcher) Wait() {
	// 等待结束
	<-d.StopChan
}

func (d *Dispatcher) run() {
	for {
		select {
		case <-d.Context.Done():
			d.StopChan <- struct{}{} // 结束
			return
		case task := <-d.TaskChan:
			go func() {
				defer func() {
					atomic.AddInt32(&d.counter, -1)
					<-d.CtrChan
				}()
				task.Do()
				//atomic.AddInt32(&d.counter, -1)
			}()
		default:

			//fmt.Println("IsStopAddTask", d.IsStopAddTask, "counter", d.counter)
			if d.IsStopAddTask && atomic.LoadInt32(&d.counter) < 1 {
				d.StopChan <- struct{}{} // 结束
				return
			}
			time.Sleep(d.IoWait)
		}
	}
}

func (d *Dispatcher) AddTask(task ITask) {
	d.CtrChan <- struct{}{}
	atomic.AddInt32(&d.counter, 1)
	d.TaskChan <- task
}
