package concurrent

import (
	"fmt"
	"time"
)

func RunChan() {
	testChan1()
	testChan2()
	testChan3()
	testChan4()
	testChan5()
	testChan6()
}

// chan创建
func testChan1() {
	var ch chan int
	fmt.Println(ch)

	var ch1 chan int = make(chan int, 10) //缓冲区
	fmt.Println(ch1)

	var ch2 map[string]chan bool // key=string value=chan bool
	fmt.Println(ch2)
}

// chan发送与接收
func testChan2() {
	var ch chan int = make(chan int, 1)
	ch <- 10  //发送
	v := <-ch //接收
	fmt.Println(v)
}

// todo 阻塞非阻塞
func testChan3() {

}

//无缓冲的chan，不能保存数据，接收和发送goroutine必须同时准备好
// func RunChan4(){
//见channel1.go
// }

//有缓冲的chan，能够保存数据，接收和发送goroutine可以不同时准备好
// func RunChan5(){
//见channel2.go
// }

// 关闭chan, 判断chan是否已经被关闭
// 不需要显式关闭，没有goroutine持有时会自动关闭
func testChan4() {
	//有缓冲的chan
	var ch chan int = make(chan int, 1)
	ch <- 10
	v := <-ch
	fmt.Println(v)
	close(ch) //将往通道发送 该类型通道的默认值，int chan，发送0
	v, ok := <-ch
	if !ok {
		fmt.Println(v, " 关闭")
	}

	//无缓冲的chan
	var ch1 chan int = make(chan int)
	close(ch1) //
	v, ok = <-ch1
	if !ok {
		fmt.Println(v, " 关闭")
	}
}

// select
func testChan5() {
	var ch chan int
	var ch1 chan int = make(chan int, 1)
	select {
	case <-ch:
		fmt.Println("<-ch")
	case ch1 <- 10:
		fmt.Println("ch1<-10")
	default:
		fmt.Println("default")
	}
}

// chan超时处理
// 永远没有往 ch 中写入数据，永远无法从 ch 中读取到数据，导致整个 goroutine 永远阻塞。
// 使用select有一个chan获得结果即可退出阻塞
func testChan6() {
	//例子1
	var ch chan int = make(chan int, 1)
	var ch1 chan int = make(chan int, 1)
	ch1 <- 10
	select {
	case <-ch:
		fmt.Println("从ch中读取到数据")
	case <-ch1:
		fmt.Println("从ch1中读取到数据")
	}

	//例子2
	ch = make(chan int)
	ch1 = make(chan int, 1)
	go func() {
		time.Sleep(1e9 * 5) // 睡眠 5 秒
		ch1 <- 10
	}()
	select {
	case <-ch: //主线程被阻塞到这里
		fmt.Println("从ch中读取到数据")
	case <-ch1: //5s钟之后从这里退出
		fmt.Println("从ch1中读取到数据")
	}
}
