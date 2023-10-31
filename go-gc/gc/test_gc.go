package gc

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/trace"
	"time"
)

// 观察GC
func TestGC() {
	// gcWay1()
	// gcWay2()
	// gcWay3()
	gcWay4()
}

func gcWay1() {
	for n := 1; n < 100000; n++ {
		allocate()
	}
}
func allocate() {
	_ = make([]byte, 1<<20)
}

func gcWay2() {
	//创建trace.out文件
	f, _ := os.Create("trace.out")
	defer f.Close()
	//启动trace
	trace.Start(f)
	defer trace.Stop()
}

func gcWay3() {
	ticker := time.NewTicker(time.Second)
	s := debug.GCStats{}
	for range ticker.C {
		allocate()
		debug.ReadGCStats(&s)
		fmt.Printf("gc %d last@%v, PauseTotal %v\n", s.NumGC, s.LastGC, s.PauseTotal)
	}
	ticker.Stop()
}

func gcWay4() {
	t := time.NewTicker(time.Second)
	s := runtime.MemStats{}
	for range t.C {
		allocate()
		runtime.ReadMemStats(&s)
		fmt.Printf("gc %d last@%v, next_heap_size@%vMB\n", s.NumGC, time.Unix(int64(time.Duration(s.LastGC).Seconds()), 0), s.NextGC/(1<<20))
	}
}
