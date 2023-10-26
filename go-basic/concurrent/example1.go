package concurrent

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//打羽毛球游戏 一局定胜负 无缓冲chan

/*
先于main函数执行，同包，同文件下可以有多个init函数
设置随机种子，后面用到了随机数
*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

// swg用来等待程序结束
var swg sync.WaitGroup //全局变量

func RunBadminton() {
	var ch chan int = make(chan int) //无缓冲通道
	swg.Add(2)
	//先准备好 发送和接收goroutine
	go player("LinDan", ch)
	go player("LiZongWei", ch)
	//再往chan中放东西
	ch <- 1
	swg.Wait()
}

// 任务循环执行
func player(name string, ch chan int) {
	defer swg.Done() //函数退出时执行 -1

	for {
		//接球
		ballCnt, ok := <-ch //接收数据 接球
		if !ok {
			//通道已关闭，对方以一定概率丢球
			fmt.Println(name, "赢了一球", "赢了") //一局定胜负
			return
		}

		//丢球
		n := rand.Intn(100)
		if n%5 == 0 {
			fmt.Println(name, "丢了一球", "输了") //一局定胜负
			close(ch)                       //关闭通道
			return
		}

		fmt.Println(name, "成功接球", ballCnt, "次")
		ballCnt++
		ch <- ballCnt //打球向对手
	}
}
