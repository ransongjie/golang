package singleton

import "sync"

// 懒汉式
type LazySingleton struct{}

var lazyInstance *LazySingleton

// 保证线程安全
var once sync.Once

func GetLazyInstance() *LazySingleton {
	// 线程安全的执行1次
	// lazyInstance为nil时，执行闭包中的代码。否则，直接返回
	once.Do(func() {
		lazyInstance = &LazySingleton{}
	})
	return lazyInstance
}
