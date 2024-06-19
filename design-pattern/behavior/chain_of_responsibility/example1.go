package main

import "fmt"

// java static final
const INFO int = 1
const DEBUG int = 2
const ERROR int = 3

// 接口存放 公共多态方法
type Logger interface {
	log(level int, msg string)
	write(msg string)
}

// 嵌入式结构体存放 公共属性
type BaseLogger struct {
	level      int
	nextLogger Logger
}

// 嵌入式结构体存放 公共相同方法
func (bl *BaseLogger) setNextLogger(nextLogger Logger) {
	bl.nextLogger = nextLogger
}

type InfoLogger struct {
	BaseLogger
}
type DebugLogger struct {
	BaseLogger
}
type ErrorLogger struct {
	BaseLogger
}

func (il *InfoLogger) log(level int, msg string) {
	if level >= il.level {
		il.write(msg)
	}

	if il.nextLogger != nil {
		il.nextLogger.log(level, msg)
	}
}

func (dl *DebugLogger) log(level int, msg string) {
	if level >= dl.level {
		dl.write(msg)
	}

	if dl.nextLogger != nil {
		dl.nextLogger.log(level, msg)
	}
}

func (el *ErrorLogger) log(level int, msg string) {
	if level >= el.level {
		el.write(msg)
	}

	if el.nextLogger != nil {
		el.nextLogger.log(level, msg)
	}
}

func (il *InfoLogger) write(msg string) {
	fmt.Println("Info Logger: ", msg)
}

func (dl *DebugLogger) write(msg string) {
	fmt.Println("Debug Logger: ", msg)
}

func (el *ErrorLogger) write(msg string) {
	fmt.Println("Error Logger: ", msg)
}

// 工厂方法
func getChainOfLogger() Logger {
	infoLogger := &InfoLogger{BaseLogger{level: INFO}}
	debugLogger := &DebugLogger{BaseLogger{level: DEBUG}}
	errorLogger := &ErrorLogger{BaseLogger{level: ERROR}}

	errorLogger.setNextLogger(debugLogger)
	debugLogger.setNextLogger(infoLogger)
	return errorLogger
}

func main() {
	chainOfLogger := getChainOfLogger()
	chainOfLogger.log(INFO, "This is info level information")
	chainOfLogger.log(ERROR, "This is error level information")
}
