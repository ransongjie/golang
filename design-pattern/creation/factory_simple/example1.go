package main

import (
	"fmt"
	"reflect"
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

// 形状工厂
type ShapeFactory struct{}

func (p *ShapeFactory) createShape(shapeType string) IShape {
	switch shapeType { //无case穿透
	case "circle":
		return &Circle{Shape{"圆形"}} //使用嵌入结构体包装公共属性
	case "square":
		return &Square{Shape{"矩形"}}
	}
	return nil
}

func main() {
	shapeFactory := ShapeFactory{}
	circle := shapeFactory.createShape("circle")
	fmt.Println(reflect.TypeOf(circle))
	circle.draw()
}
