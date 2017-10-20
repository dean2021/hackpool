package hackpool

import (
	"sync"
)

type HackPool struct {
	c            chan struct{}
	queue        []interface{}
	thread_count int
	callfunc     func(interface{})
}

func New(thread_count int, callfunc func(interface{})) *HackPool {
	t := &HackPool{
		c:            make(chan struct{}, thread_count),
		thread_count: thread_count,
		callfunc:     callfunc,
	}
	return t
}

func (this *HackPool) Push(data interface{}) {
	this.queue = append(this.queue, data)
}

func (this *HackPool) Run() {
	var wg sync.WaitGroup
	for _, v := range this.queue {
		this.c <- struct{}{}
		go func(x interface{}) {
			wg.Add(1)
			this.callfunc(x)
			wg.Done()
			<-this.c
		}(v)
	}
	wg.Wait()
}
