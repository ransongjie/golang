package gc

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func GCOptimize2() {
	testOptimize2()
}

// 简单的 Web 程序来说明复用内存的重要性。在这个程序中，每当产生一个 /example2的请求时，都会创建一段内存，并用于进行一些后续的工作
func testOptimize2() {
	//为了进行性能分析，我们还额外创建了一个监听 6060 端口的 goroutine，用于使用 pprof 进行分析
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	http.HandleFunc("/example2", func(w http.ResponseWriter, r *http.Request) {
		b := newBuf()

		// 模拟执行一些工作
		for idx := range b {
			b[idx] = 1
		}

		fmt.Fprintf(w, "done, %v", r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)
}

func newBuf() []byte {
	return make([]byte, 10<<20)
}

// 可见 newBuf 产生的申请的内存过多，现在我们使用 sync.Pool 来复用 newBuf 所产生的对象
func testOptimize2More() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	http.HandleFunc("/example2", func(w http.ResponseWriter, r *http.Request) {
		b := bufPool.Get().([]byte) // 最后1个括号是类型断言
		for idx := range b {
			b[idx] = 0
		}
		fmt.Fprintf(w, "done, %v", r.URL.Path[1:])
		bufPool.Put(b) //
	})
	http.ListenAndServe(":8080", nil)
}

// 使用 sync.Pool 复用需要的 buf
var bufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 10<<20)
	},
}
