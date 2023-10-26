package concurrent

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

//学生做作业 有缓冲chan

const (
	stuNum  int = 10 //小班教学10个学生
	taskNum int = 6  //语数外政化生
)

var swg2 sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RunHomework() {
	swg2.Add(stuNum)

	tasks := make(chan string, taskNum) //有缓冲chan

	for i := 0; i < stuNum; i++ {
		go student("学生"+strconv.Itoa(i), tasks)
	}

	var homeworks [6]string = [6]string{"语文", "数学", "外语", "政治", "化学", "生物"}
	for i := 0; i < taskNum; i++ {
		tasks <- homeworks[i]
	}

	close(tasks)

	swg2.Wait()
}

func student(name string, tasks chan string) {
	defer swg2.Done()

	for { //循环
		task, ok := <-tasks

		if !ok { //管道已关闭
			fmt.Println(name, " 完成了所有作业")
			return
		}

		fmt.Println(name, " 开始做", task)
		//完成作业的时间 0~100 毫秒
		rnd := rand.Int63n(100)
		time.Sleep(time.Duration(rnd) * time.Millisecond)
		fmt.Println(name, " 完成作业", task)
	}
}
