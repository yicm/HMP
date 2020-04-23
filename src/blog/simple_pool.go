package blog

import (
	"fmt"
	"sync"
)

type SimplePool struct {
	wg   sync.WaitGroup
	work chan func()
}

func NewSimplePoll(workers int) *SimplePool {
	p := &SimplePool{
		wg: sync.WaitGroup{},
		// 无缓冲chan,发送端会阻塞直到数据被接收；接收端会阻塞直到读取到数据
		// 有缓冲chan,当缓冲满时发送端阻塞，当缓冲空时接收端阻塞
		work: make(chan func()),
	}
	p.wg.Add(workers)
	//根据指定的并发量去读取管道并执行
	for i := 0; i < workers; i++ {
		go func() {
			defer func() {
				// 捕获异常 防止waitGroup阻塞
				if err := recover(); err != nil {
					fmt.Println(err)
					p.wg.Done()
				}
			}()
			// 从workChannel中取出任务执行
			// range遍历通道，如不关闭p.work，range遍历会阻塞
			for fn := range p.work {
				fn()
			}
			p.wg.Done()
		}()
	}
	return p
}

// 添加任务
func (p *SimplePool) Add(fn func()) {
	p.work <- fn
}

// 所有任务添加后结束
func (p *SimplePool) Run() {
	// 主动关闭chan
	close(p.work)
	p.wg.Wait()
}

// func taskSample(i int) func() {
// 	return func() {
// 		// todo something
// 	}
// }

// func TestSimplePool(t *testing.T) {
// 	p := NewSimplePoll(20)
// 	for i := 0; i < 100; i++ {
// 		p.Add(taskSample(i))
// 	}
// 	p.Run()
// }
