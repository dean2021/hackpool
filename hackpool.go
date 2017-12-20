package hackpool

import (
	"sync"
)

type HackPool struct {
	state    chan struct{}
	message  chan interface{}
	numGo    int
	function func(interface{})
}

func New(numGoroutine int, function func(interface{})) *HackPool {
	t := &HackPool{
		state:    make(chan struct{}, numGoroutine),
		message:  make(chan interface{}, numGoroutine),
		numGo:    numGoroutine,
		function: function,
	}
	return t
}

func (c *HackPool) QueueCount() int {
	return len(c.message)
}

func (c *HackPool) Push(data interface{}) {
	c.message <- data
}

func (c *HackPool) Close() {
	close(c.message)
}

func (c *HackPool) Run() {

	var wg sync.WaitGroup

	for v := range c.message {

		c.state <- struct{}{}

		wg.Add(1)

		go func(value interface{}) {

			defer wg.Done()

			c.function(value)

			<-c.state

		}(v)
	}

	wg.Wait()
}
