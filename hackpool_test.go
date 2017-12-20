package hackpool

import (
	"testing"
	"fmt"
)

func TestHackPool(t *testing.T) {

	var hp *HackPool
	numGoroutine := 2
	taskCount := 10000

	hp = New(numGoroutine, func(i interface{}) {

		fmt.Println(i.(int))

		//time.Sleep(time.Second * 2)
	})

	// 由于同步往有缓冲通道塞数据,元素个数超过chan定义的长度,会造成死锁,所以必须异步塞数数据
	// 异步塞数据的好处是,如果chan通道已满,就等待chan被消费
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
