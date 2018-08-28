package task

import "sync"

func NewTaskPool(max int , handler func(*TaskPool))  *TaskPool {
	return &TaskPool{
		pool:make(chan  bool , max),
		wg:new(sync.WaitGroup),
		handler:handler,
	}
}
type TaskPool struct {
	Max int
	pool chan bool
	wg *sync.WaitGroup
	handler func(*TaskPool)
}

func (t *TaskPool)Add(i int)  {
	t.pool <- true
	t.wg.Add(i)
}

func (t *TaskPool)Done()  {
	<-t.pool
	t.wg.Done()
}


func (t *TaskPool)Run()  {
	t.handler(t)
	t.wg.Wait()
	close(t.pool)
}
