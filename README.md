# HackPool

北半球最优雅的协程池

## Example


```go
package main

import (
	"fmt"
	"github.com/poc-db/hackpool"
)

func main() {

	concurrency := 2   // 并发数
	taskCount := 100   

	hp := hackpool.New(concurrency, func(i interface{}) {
		fmt.Println(i.(int))
	})

	go func() {

		for i := 0; i < taskCount; i++ {
			hp.Push(i)
		}

		// 必须关闭,不然阻塞死锁
		hp.Close()
	}()

	// 跑起来! 伙计
	hp.Run()
}
```

## Installation

    go get github.com/poc-db/hackpool

## License

This project is copyleft of [CSOIO](http://www.csoio.com/) and released under the GPL 3 license.

