# 观察Go GC的方式
## GODEBUG=gctrace=1
1. 调用 test_gc.go/gcWay1()
2. 运行命令行

Windows Command Prompt (CMD)
```shell
set GODEBUG=gctrace=1
go run main.go
```
Linux
```shell
go build -o main
GODEBUG=gctrace=1 ./main
```
## go tool trace
1. 运行 test_gc.go/gcWay2()
2. 运行命令行，自动跳转到web页面
```shell
go tool trace trace.out
```

## debug.ReadGCStats
通过代码的方式实现对感兴趣指标的监控
1. 运行 test_gc.go/gcWay3()

## runtime.ReadMemStats
通过代码的方式实现对感兴趣指标的监控
1. 运行 test_gc.go/gcWay4()

# 内存泄漏
- 全局变量 导致的内存泄漏
- goroutine 导致的内存泄漏
- 由 chan 阻塞 goroutine 导致的内存泄漏
1. 运行 test_leak.go
2. 运行命令行，自动跳转到web页面，等待一段时间
```shell
go tool trace trace.out
```
# 调优
## test_opt1
```text
检验：合理化内存分配的速度、提高赋值器的 CPU 利用率
实验：test_opt1
过程：
1. 运行 test_opt1
2. trace 文件分析 `go tool trace trace.out` 跳转到web页面
分析：web页面中可看的东西
- gc：
- mutator(赋值器 )的CPU利用率：越高越好，越低代表在gc collector上花费的时间越多
- goroutine分析：goroutine 的执行时间（红色部分）占其生命周期总时间非常短的一部分，但大部分时间都花费在调度器的等待上了（蓝色的部分）
结论：
- 同时创建大量goroutine对调度器产生的压力确实不小，我们不妨将这一产生速率减慢，一批一批地创建 goroutine。（协程也需要空间）当需要创建大量协程执行短期任务时，当一批协程执行任务完毕，再创建新的一批协程
- goroutine池
```
## test_opt2
```text
检验：降低并复用已经申请的内存
实验：test_opt2
过程：
1. 运行 test_opt1。其中 创建了一个监听 6060 端口的 goroutine，用于使用 pprof 进行分析
2. 命令行执行 `wget http://127.0.0.1:6060/debug/pprof/trace\?seconds\=20 -O trace.out`
- 使用 pprof 的 trace 来查看 GC 在此服务器中面对大量请求时候的状态，要使用 trace 可以通过访问 /debug/pprof/trace 路由来进行，其中 seconds 参数设置为 20s，并将 trace 的结果保存为 trace.out
3. 命令行执行 `ab -n 500 -c 100 http://127.0.0.1:8080/example2`
- 压测工具 ab，来同时产生 500 个请求（ -n 一共 500 个请求， -c一个时刻执行请求的数量，每次 100 个并发请求）
4. 命令行执行 `go tool pprof http://127.0.0.1:6060/debug/pprof/heap`
- 我们可以通过 go tool pprof 来查看究竟是谁分配了大量内存（使用 web 指令来使用浏览器打开统计信息的可视化图形）
结论：可见 newBuf 产生的申请的内存过多，现在我们使用 sync.Pool 来复用 newBuf 所产生的对象
```
## test_opt3
```text
检验：调整 GOGC 参数
实验：test_opt3
过程：
1. 运行 test_opt3, `GOGC=1000 ./main` 1000表示内存增长10倍
2. 命令行执行 `ab -n 500 -c 100 http://127.0.0.1:8080/example2`
结论：临时调整GOGC参数：需要时，增大 GOGC 的值，降低 GC 的运行频率
```