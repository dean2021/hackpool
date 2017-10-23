# HackPool

非常优雅的协程库

## Example


```go
package main

import (
	"fmt"

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

	// 协程(线程?^^)数量
	thread_count := 100
	wp := hackpool.New(thread_count, call_function)

	// 添加任务
	for i := 0; i < 100000; i++ {
		wp.Push("https://item.jd.com/" + strconv.Itoa(i) + ".html")
	}

	// 跑起来! 伙计
	wp.Run()
}
```

## Installation

    go get github.com/dean2020/hackpool

## License

This project is copyleft of [CSOIO](http://www.csoio.com/) and released under the GPL 3 license.

