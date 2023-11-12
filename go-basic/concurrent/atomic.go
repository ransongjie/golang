package concurrent

import (
	"fmt"
	"sync/atomic"
)

// 测试 sync/atomic 包

func RunAtomic() {
	testAtomic1()
}

func testAtomic1() {
	var num int32 = 0
	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt32(&num, 1)
		}()
	}
	for atomic.LoadInt32(&num) != 10 {
		// 等待所有 goroutine 完成
	}
	fmt.Println(num) // 10
}
