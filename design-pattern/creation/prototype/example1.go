package main

import "fmt"

/*
Go语言本身并不支持传统意义上的“原型模式”（Prototype Pattern），
这是一种在面向对象编程中常见的设计模式，用于创建对象的原型实例，然后复制这些原型来创建新对象。
*/
type Apple struct {
	name  string
	price float64
}

func (a *Apple) show() {
	fmt.Println(a.name, ", ", a.price)
}

// clone func
func (a *Apple) clone() *Apple {
	return &Apple{name: a.name, price: a.price}
}

func main() {
	prototype := &Apple{name: "烟台苹果", price: 23.3}
	apple := prototype.clone()
	apple.show()
}
