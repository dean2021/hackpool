package hackpool

import (
	"testing"
	"fmt"
)

func TestHackPool(t *testing.T) {

	var hp *HackPool
	numGoroutine := 2
	taskCount := 100

	hp = New(numGoroutine, func(i interface{}) {
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
