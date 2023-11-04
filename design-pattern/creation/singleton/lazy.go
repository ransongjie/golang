package singleton

import "sync"

// 懒汉式
type LazySingleton struct{}

var lazyInstance *LazySingleton

// 保证线程安全
var once sync.Once

func GetLazyInstance() *LazySingleton {
	// 线程安全的执行1次
	once.Do(func() {
		lazyInstance = &LazySingleton{}
	})
	return lazyInstance
}
