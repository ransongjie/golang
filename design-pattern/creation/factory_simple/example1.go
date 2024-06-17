package main

import (
	"fmt"
)

// 接口
type IShape interface {
	draw()
}

// 嵌入结构体
type Shape struct {
	name string
}

// 圆形
type Circle struct {
	Shape
}

func (p *Circle) draw() {
	fmt.Println(p.name)
}

// 矩形
type Square struct {
	Shape
}

func (p *Square) draw() {
	fmt.Println(p.name)
}

// 包级别方法，类似其它语言的静态方法
func createShape(shapeType string) IShape {
	switch shapeType { //无case穿透
	case "circle":
		return &Circle{Shape{"圆形"}} //使用嵌入结构体包装公共属性
	case "square":
		return &Square{Shape{"矩形"}}
	}
	return nil
}

func main() {
	circle := createShape("circle")
	circle.draw()
}
