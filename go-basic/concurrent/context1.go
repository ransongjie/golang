package concurrent

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*
type Context interface {
	// 返回 Context 的截止时间
    Deadline() (deadline time.Time, ok bool)
    // 返回 一个通道。这个通道在 Context 被取消或超时时关闭
    Done() <-chan struct{}
    // 返回 Context 的错误信息
    Err() error
    // 返回 Context 中与 key 关联的值
    Value(key interface{}) interface{}
}
*/

func RunContext() {
	testContext9()
}

// 创建和传递context
func testContext1() {
	// 创建一个 根Context
	rootContext := context.Background()
	// 创建一个 带有超时时间的Context (ctx)，这里设置超时时间为2秒
	// return_1=context, return_2=cancelFunc
	ctx, cancel := context.WithTimeout(rootContext, 2*time.Second)

	// return 时执行
	defer cancel()

	// 在新的goroutine中执行任务
	go func(ctx context.Context) {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("任务完成")
		case <-ctx.Done():
			fmt.Println("任务取消或超时")
		}
	}(ctx) // go func(形参){}(实参)

	// 等待一段时间，模拟程序运行
	time.Sleep(5 * time.Second)
}

// 使用 WithValue 传递数据
func testContext2() {
	// 创建一个根Context
	rootContext := context.Background()

	// 使用WithValue传递数据
	ctx := context.WithValue(rootContext, "onekey", 123)

	testContext2a(ctx)
}

func testContext2a(ctx context.Context) {
	// 从Context中获取数据
	if userID, ok := ctx.Value("onekey").(int); ok {
		fmt.Println("value: ", userID)
	} else {
		fmt.Println("onekey 不存在")
	}
}

// 设置请求超时时间
func testContext3() {
	// 创建一个根Context
	rootContext := context.Background()

	// 创建一个超时时间为2秒的Context
	timeoutCtx, _ := context.WithTimeout(rootContext, 2*time.Second)

	// 创建一个手动取消的Context
	cancelCtx, cancel := context.WithCancel(rootContext)
	defer cancel()

	// 在新的goroutine中执行任务 入参 timeoutCtx
	go func(ctx context.Context) {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("任务完成")
		case <-ctx.Done():
			fmt.Println("任务取消或超时")
		}
	}(timeoutCtx)

	// 在另一个goroutine中执行任务 入参 cancelCtx
	go func(ctx context.Context) {
		select { // select 只会执行其中1个 case
		case <-time.After(1 * time.Second):
			fmt.Println("另一个任务完成")
		case <-ctx.Done():
			fmt.Println("另一个任务取消")
		}
	}(cancelCtx)

	// 等待一段时间，模拟程序运行
	time.Sleep(8 * time.Second)
}

// 处理请求取消
func testContext4() {
	// 创建一个根Context
	rootContext := context.Background()

	// 创建一个可以手动取消的Context
	ctx, cancel := context.WithCancel(rootContext)
	defer cancel()

	// 在新的goroutine中执行任务
	go func(ctx context.Context) {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("任务完成")
		case <-ctx.Done():
			fmt.Println("任务取消")
		}
	}(ctx)

	// 等待一段时间，手动取消任务
	time.Sleep(2 * time.Second)
	cancel()

	// 等待一段时间，模拟程序运行
	time.Sleep(1 * time.Second)
}

// 使用 Background WithTimeout WithCancel 创建各种 context
func testContext5() {
	// 创建一个根Context
	rootContext := context.Background()

	// 创建一个超时时间为2秒的Context
	timeoutCtx, _ := context.WithTimeout(rootContext, 2*time.Second)

	// 创建一个手动取消的Context
	cancelCtx, cancel := context.WithCancel(rootContext)
	defer cancel()

	// 在新的goroutine中执行任务
	go func(ctx context.Context) {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("任务完成")
		case <-ctx.Done():
			fmt.Println("任务取消或超时")
		}
	}(timeoutCtx)

	// 在另一个goroutine中执行任务
	go func(ctx context.Context) {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("另一个任务完成")
		case <-ctx.Done():
			fmt.Println("另一个任务取消")
		}
	}(cancelCtx)

	// 等待一段时间，模拟程序运行
	time.Sleep(5 * time.Second)
}

// /Context 在并发中的应用///
// 使用 Context 控制多个协程
func testContext6() {
	// 创建一个根Context
	rootContext := context.Background()

	// 创建一个可以手动取消的Context
	ctx, cancel := context.WithCancel(rootContext)
	defer cancel()

	// 使用WaitGroup等待所有任务完成
	// wg.Add(1) // 每次 +1
	// defer wg.Done() // 每次 -1
	// wg.Wait() // 为0退出 等待所有任务完成
	var wg sync.WaitGroup

	// 启动多个协程执行任务
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			select {
			case <-time.After(time.Duration(id) * time.Second):
				fmt.Println("任务", id, "完成")
			case <-ctx.Done():
				fmt.Println("任务", id, "取消")
			}
		}(i)
	}

	// 等待一段时间，然后手动取消任务
	time.Sleep(2 * time.Second)
	cancel()

	// 等待所有任务完成
	wg.Wait()
}

// HTTP 请求中的 Context 使用
// 需要从浏览器发送请求
func testContext7() {
	http.HandleFunc("/", testContext7a)
	http.ListenAndServe(":8080", nil)
}
func testContext7a(w http.ResponseWriter, r *http.Request) {
	// 获取到请求的 Context
	ctx := r.Context()

	select {
	case <-time.After(2 * time.Second):
		fmt.Fprintln(w, "Hello, World!")
	case <-ctx.Done(): // 在 Context 被取消或超时关闭 时返回一个 context
		err := ctx.Err()
		fmt.Println("Server:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// 数据库操作中的 Context 使用
func testContext8() {
	// 连接数据库
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database")
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer db.Close()

	// 创建一个Context，设置超时时间为5秒
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 在Context的超时时间内执行数据库查询
	// 查询时间超过 5 秒，Context 会接收到取消信号，可以在其中执行处理查询超时的逻辑。
	rows, err := db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		fmt.Println("数据库查询失败:", err)
		return
	}
	defer rows.Close()

	// 处理查询结果
	for rows.Next() {
		// 处理每一行数据
	}
}

// 使用 Context 在多个微服务之间进行数据传递和超时控制
// 创建了一个简单的微服务模拟，它接收来自 reqCh 通道的请求，并将处理结果发送到 resCh 通道。
// 用带有 5 秒超时时间的 Context 来确保请求不会无限期等待，同时也能够处理超时的情况
func testContext9() {
	// 创建根 Context
	rootContext := context.Background()

	// 创建用于请求和响应的通道
	reqCh := make(chan Request)
	resCh := make(chan Response)

	// 启动微服务
	go microservice(rootContext, reqCh, resCh)

	// 创建带有5秒超时时间的Context
	ctx, cancel := context.WithTimeout(rootContext, 5*time.Second)
	defer cancel()

	// 发送请求到微服务
	for i := 1; i <= 3; i++ {
		req := Request{ID: i}
		reqCh <- req

		select {
		case <-ctx.Done():
			fmt.Println("Request timed out!")
			return
		case res := <-resCh:
			fmt.Println(res.Message)
		}
	}
}

type Request struct {
	ID int
}

type Response struct {
	Message string
}

func microservice(ctx context.Context, reqCh chan Request, resCh chan Response) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Microservice shutting down...")
			return
		case req := <-reqCh:
			// 模拟处理请求的耗时操作
			time.Sleep(2 * time.Second)
			response := Response{Message: fmt.Sprintf("Processed request with ID %d", req.ID)}
			resCh <- response
		}
	}
}
