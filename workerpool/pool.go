package workerpool

import (
	"fmt"
	"runtime"
	"sync"
)

// Pool 协程池结构体
type Pool struct {
	maxWorkers int
	taskQueue  chan Task
	wg         sync.WaitGroup
	quit       chan bool
}

// NewPool 创建一个新的协程池
func NewPool(maxWorkers int, queueSize int) *Pool {
	if maxWorkers <= 0 {
		maxWorkers = runtime.NumCPU()
	}
	if queueSize <= 0 {
		queueSize = maxWorkers * 2
	}

	return &Pool{
		maxWorkers: maxWorkers,
		taskQueue:  make(chan Task, queueSize),
		quit:       make(chan bool),
	}
}

// Submit 提交一个任务到协程池
func (p *Pool) Submit(task Task) {
	p.taskQueue <- task
}

// Start 启动协程池，开始执行
func (p *Pool) Start() {
	for i := 0; i < p.maxWorkers; i++ {
		p.wg.Add(1)
		go func(id int) {
			defer p.wg.Done()
			for { // 每个worker 死循环，读去任务队列，进行处理
				select {
				case task, ok := <-p.taskQueue: // 即使 taskQueue 被关闭后，仍然能读取数据
					if ok {
						task.Do()
					} else { // p.taskQueue 被关闭后，且 管道中的数据被全部消费完成，才会走到else
						fmt.Println("worker ", id, " complete.")
						return
					}
				case <-p.quit:
					fmt.Println("worker ", id, " quit.")
					return
				}
			}
		}(i)
	}
	p.wg.Wait()
}

func (p *Pool) ProduceFinish() {
	fmt.Println("ProduceFinish stop submiting task")
	close(p.taskQueue)
}

// Stop 停止协程池
func (p *Pool) Stop() {
	fmt.Println(" pool stop, stop doing task.")
	close(p.quit)
}
