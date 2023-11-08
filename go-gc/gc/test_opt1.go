package gc

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"sync/atomic"
	"time"
)

var (
	stop  int32
	count int64
	sum   time.Duration
)

func GCOptimize1() {
	testOptimize1()
}

/*
concat 函数负责拼接一些长度不确定的字符串。
并且为了快速完成任务，出于某种原因，在两个嵌套的 for 循环中一口气创建了 800 个 goroutine。
在 main 函数中，启动了一个 goroutine 并在程序结束前不断的触发 GC，并尝试输出 GC 的平均执行时间
*/
func testOptimize1() {
	// go tool trace 相关 start
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	// go tool trace 相关 end

	go func() {
		var t time.Time
		// 原子性操作
		for atomic.LoadInt32(&stop) == 0 {
			t = time.Now()
			runtime.GC()         // 调用gc
			sum += time.Since(t) // gc总执行时间
			count++              // 调用gc次数
		}
		// 输出gc平均执行时间
		// GC spend avg: 1.10755ms
		fmt.Printf("GC spend avg: %v\n", time.Duration(int64(sum)/count))
	}()

	// 将执行800个过routine
	concat()
	// 原子性操作
	atomic.StoreInt32(&stop, 1)
}

// concat 函数负责拼接一些长度不确定的字符串
// 在两个嵌套的 for 循环中一口气创建了 800 个 goroutine
func concat() {
	for n := 0; n < 100; n++ {
		for i := 0; i < 8; i++ {
			go func() {
				s := "Go GC"
				s += " " + "Hello"
				s += " " + "World"
				_ = s
			}()
		}
	}
}

// 优化 testOptimize1
func testOptimize1More() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	go func() {
		var t time.Time
		for atomic.LoadInt32(&stop) == 0 {
			t = time.Now()
			runtime.GC()
			sum += time.Since(t)
			count++
		}
		// GC spend avg: 786.825µs. 提升30%左右
		fmt.Printf("GC spend avg: %v\n", time.Duration(int64(sum)/count))
	}()

	// 将执行800个过routine
	concatOpt()
	atomic.StoreInt32(&stop, 1)
}

// 同时创建大量goroutine对调度器产生的压力确实不小，我们不妨将这一产生速率减慢，一批一批地创建 goroutine
func concatOpt() {
	wg := sync.WaitGroup{}
	for n := 0; n < 100; n++ {
		wg.Add(8)
		for i := 0; i < 8; i++ {
			go func() {
				s := "Go GC"
				s += " " + "Hello"
				s += " " + "World"
				_ = s
				wg.Done()
			}()
		}
		wg.Wait() //8个goroutine执行完毕才创建下1批
	}
}

// 预先分配足够内存》减少内存分配，复用内存
// 原因在于 + 运算符会随着字符串长度的增加而申请更多的内存，并将内容从原来的内存位置拷贝到新的内存位置，造成大量不必要的内存分配，
// 先提前分配好足够的内存，再慢慢地填充，也是一种减少内存分配、复用内存形式的一种表现。
func concatOpt2() {
	wg := sync.WaitGroup{}
	for n := 0; n < 100; n++ {
		wg.Add(8)
		for i := 0; i < 8; i++ {
			go func() {
				s := make([]byte, 0, 20)
				s = append(s, "Go GC"...)
				s = append(s, ' ')
				s = append(s, "Hello"...)
				s = append(s, ' ')
				s = append(s, "World"...)
				_ = string(s)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
