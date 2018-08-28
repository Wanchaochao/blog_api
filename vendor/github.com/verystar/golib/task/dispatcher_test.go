package task

import (
	"testing"
	"context"
	"fmt"
	"time"
)

type demoTask struct {
	Index int
}

func (t *demoTask) Do() {
	fmt.Println("====>" , t.Index)
	time.Sleep(1 * time.Second)
}

func TestNewDispatcher(t *testing.T) {
	ctx := context.Background()
	d := NewDispatcher(ctx , 4)
	d.SetIoWait(2 * time.Second)
	for i := 0; i < 10 ; i++ {
		d.AddTask(&demoTask{
			Index:i,
		})
	}

	d.IsStopAddTask = true
	d.Wait()
}