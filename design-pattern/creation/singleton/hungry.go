package singleton

// 饿汉式

type MySingleton struct{}

// 全局变量
var instance = &MySingleton{}

// 全局方法
func GetInstance() *MySingleton {
	return instance
}
