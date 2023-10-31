package gc

import (
	"os"
	"runtime/trace"
)

// 内存泄漏
func TestLeak() {
	traceLeak()
}

func traceLeak() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	// testLeak1()
	testLeak2()
	// testLeak3()
}

// 全局变量 导致的内存泄漏
var cache = map[interface{}]interface{}{}

func testLeak1() {
	for i := 0; i < 10000; i++ {
		m := make([]byte, 1<<10)
		cache[i] = m
	}
}

// goroutine 导致的内存泄漏
func testLeak2() {
	for i := 0; i < 100000; i++ {
		go func() {
			select {}
		}()
	}
}

// 由 chan 阻塞 goroutine 导致的内存泄漏
var ch = make(chan struct{})

func testLeak3() {
	for i := 0; i < 100000; i++ {
		// 没有接收方，goroutine 会一直阻塞
		go func() { ch <- struct{}{} }()
	}
}
