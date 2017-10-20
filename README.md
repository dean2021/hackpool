# HackPool

非常优雅的协程库

## Example


```go
package main

import (
	"fmt"
	"time"
	"github.com/dean2020/hackpool"
)
func call_function(v interface{}) {
	fmt.Println(v)
	time.Sleep(time.Second * 1)
}

func main() {
	thread_count := 10
	queue_size := 100
	wp := HackPool.New(thread_count, call_function)
	for i := 0; i < queue_size; i++ {
		wp.Push(i)
	}
	wp.Run()
}

```

## Installation

    go get github.com/dean2020/hackpool

## License

This project is copyleft of [CSOIO](http://www.csoio.com/) and released under the GPL 3 license.

