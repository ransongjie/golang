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