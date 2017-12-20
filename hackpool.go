package hackpool

import (
	"sync"
)

type HackPool struct {
	state    chan struct{}
	messages chan interface{}
	numGo    int
	function func(interface{})
}

func New(numGoroutine int, function func(interface{})) *HackPool {
	t := &HackPool{
		state:    make(chan struct{}, numGoroutine),
		messages: make(chan interface{}, numGoroutine),
		numGo:    numGoroutine,
		function: function,
	}
	return t
}

func (c *HackPool) QueueCount() int {
	return len(c.messages)
}

func (c *HackPool) Push(data interface{}) {
	c.messages <- data
}

func (c *HackPool) Close() {
	close(c.messages)
}

func (c *HackPool) Run() {

	var wg sync.WaitGroup

	// 阻塞,等待message有数据或close
	for v := range c.messages {

		// 重点:当state被赋值,上一个协程才会结束
		c.state <- struct{}{}

		wg.Add(1)

		go func(value interface{}) {

			c.function(value)

			// 阻塞, 等待state被赋值.
			// 增加state的目的是为了保证每个函数能够执行完,也就是保证同时执行的函数数量
			<-c.state

			wg.Done()
		}(v)
	}

	wg.Wait()
}
