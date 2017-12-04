package hackpool

import (
	"sync"
)

type HackPool struct {
	c            chan struct{}
	queue        chan interface{}
	thread_count int
	callfunc     func(interface{})
}

func New(thread_count int, callfunc func(interface{})) *HackPool {
	t := &HackPool{
		c:            make(chan struct{}, thread_count),
		queue:        make(chan interface{}, thread_count*2),
		thread_count: thread_count,
		callfunc:     callfunc,
	}
	return t
}

func (this *HackPool) QueueCount() int {
	return len(this.queue)
}

func (this *HackPool) Push(data interface{}) {
	this.queue <- data
}

func (this *HackPool) Close() {
	close(this.queue)
}

func (this *HackPool) Run() {
	var wg sync.WaitGroup
	for {
		v, ok := <-this.queue
		if ok {
			this.c <- struct{}{}
			go func(x interface{}) {
				wg.Add(1)
				this.callfunc(x)
				wg.Done()
				<-this.c
			}(v)
		} else {
			break
		}
	}
	wg.Wait()
}
