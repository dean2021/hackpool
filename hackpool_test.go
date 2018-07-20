package hackpool

import (
	"testing"
	"fmt"
	"time"
)

func TestHackPool(t *testing.T) {

	var hp *HackPool
	concurrency := 2
	taskCount := 10000

	hp = New(concurrency, func(i interface{}) {
		fmt.Println(i.(int))
		time.Sleep(time.Second * 2)
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
