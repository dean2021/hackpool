# HackPool

非常优雅的协程库

## Example


```go
package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/dean2020/hackpool"
	"strconv"
)

// 回调函数
func call_function(v interface{}) {
	resp, err := http.Get(v.(string))
	if err == nil {
		fmt.Println(v, resp.StatusCode)
	} else {
		fmt.Println(err)
	}
}

func main() {

	thread_count := 10
	task_count := 100
	wp := hackpool.New(thread_count, call_function)
	go func() {
		for i := 0; i < task_count; i++ {
			wp.Push(i)
		}
		wp.Close() //关闭任务队列，跑完本次任务就退出。若不关闭，则可以一直往里写任务
	}()
  
	// 跑起来! 伙计
	wp.Run()
}
```

## Installation

    go get github.com/dean2020/hackpool

## License

This project is copyleft of [CSOIO](http://www.csoio.com/) and released under the GPL 3 license.

