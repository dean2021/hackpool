package hackpool

import (
	"sync"
)

type HackPool struct {
	concurrency int
	queues      chan interface{}
	function    func(interface{})
}

func New(concurrency int, function func(interface{})) *HackPool {
	return &HackPool{
		concurrency: concurrency,
		queues:      make(chan interface{}),
		function:    function,
	}
}

func (c *HackPool) Push(data interface{}) {
	c.queues <- data
}

func (c *HackPool) Close() {
	close(c.queues)
}

func (c *HackPool) Run() {
	var wg sync.WaitGroup

	wg.Add(c.concurrency)

	for i := 0; i < c.concurrency; i++ {
		go func() {
			for v := range c.queues {
				c.function(v)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
