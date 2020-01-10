package task

import (
	"errors"
	"log"
	"sync/atomic"
)

//任务定义
//handler 回调函数
//Params  参数
type Task struct {
	Handler func(v ...interface{})
	Params  []interface{}
}

//任务池
type Pool struct {
	capacity       uint64
	runningWorkers uint64
	state          int64
	taskC          chan *Task
	PanicHandler   func(interface{})
}

var ErrInvalidPoolCap = errors.New("invalid pool cap")

const (
	RUNNING = 1
	STOPED  = 0
)

//创建一个进程池
func NewPool(capacity uint64) (*Pool, error) {
	if capacity <= 0 {
		return nil, ErrInvalidPoolCap
	}

	return &Pool{
		capacity: capacity,
		state:    RUNNING,
		taskC:    make(chan *Task, capacity),
	}, nil
}

//创建协程池中的协程
func (p *Pool) build() {
	p.incr()

	go func() {
		defer func() {
			p.decr()

			if r := recover(); r != nil {
				if p.PanicHandler != nil {
					p.PanicHandler(r)
				} else {
					log.Printf("Worker panic: %s\n", r)
				}
			}
		}()

		for {
			select {
			case task, ok := <-p.taskC:
				if !ok {
					return
				}

				task.Handler(task.Params...)
			}
		}
	}()
}

func (p *Pool) incr() {
	atomic.AddUint64(&p.runningWorkers, 1)
}

func (p *Pool) decr() {
	atomic.AddUint64(&p.runningWorkers, ^uint64(0))
}

func (p *Pool) GetRunningWorkers() uint64 {
	return atomic.LoadUint64(&p.runningWorkers)
}

func (p *Pool) GetCap() uint64 {
	return atomic.LoadUint64(&p.capacity)
}

var ErrPoolAlreadyClosed = errors.New("pool already closed")

func (p *Pool) Put(task *Task) error {

	if p.state == STOPED {
		return ErrPoolAlreadyClosed
	}

	//创建进程池
	if p.capacity > p.GetRunningWorkers() {
		p.build()
	}

	//投递任务
	p.taskC <- task

	return nil
}

func (p *Pool) Close() {
	p.state = STOPED

	for len(p.taskC) > 0 {
	}

	close(p.taskC)
}
