# HackPool

非常优雅的协程库

## Example


```go
package main

import (
	"fmt"
	"time"
	"github.com/Greyh4t/hackpool"
)

func call_function(v interface{}) {
	fmt.Println(v)
	time.Sleep(time.Second * 1)
}

func main() {
	thread_count := 10
	queue_size := 100
	wp := New(thread_count, call_function)
	go func() {
		for i := 0; i < queue_size; i++ {
			wp.Push(i)
		}
		wp.Close() //关闭任务队列，跑完本次任务就退出。若不关闭，则可以一直往里写任务
	}()
	wp.Run()
}
```

## Installation

    go get github.com/dean2020/hackpool

## License

This project is copyleft of [CSOIO](http://www.csoio.com/) and released under the GPL 3 license.

