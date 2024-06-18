package singleton

// 饿汉式

type MySingleton struct{}

// 包级别变量
var instance = &MySingleton{}

// 包级别方法
func GetInstance() *MySingleton {
	return instance
}
