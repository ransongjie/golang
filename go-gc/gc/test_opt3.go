package gc

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func GCOptimize3() {
	testOptimize3()
}

// 简单的 Web 程序来说明复用内存的重要性。在这个程序中，每当产生一个 /example2的请求时，都会创建一段内存，并用于进行一些后续的工作
func testOptimize3() {
	//为了进行性能分析，我们还额外创建了一个监听 6060 端口的 goroutine，用于使用 pprof 进行分析
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	http.HandleFunc("/example2", func(w http.ResponseWriter, r *http.Request) {
		b := newBuf3()

		// 模拟执行一些工作
		for idx := range b {
			b[idx] = 1
		}

		fmt.Fprintf(w, "done, %v", r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)
}

func newBuf3() []byte {
	return make([]byte, 10<<20)
}
