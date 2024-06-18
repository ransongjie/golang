package main

import "fmt"

type Shape interface {
	draw()
}

type Circle struct{}

type Square struct{}

func (c *Circle) draw() {
	fmt.Println("画出一个圆形")
}

func (c *Square) draw() {
	fmt.Println("画出一个矩形")
}

// 装饰器
type ColorDecorator struct {
	shape Shape
}

func (cd *ColorDecorator) draw() {
	cd.shape.draw()
}

func (cd *ColorDecorator) fill(color string) {
	fmt.Println("填充", color)
}

func main() {
	circle := &Circle{}
	colorDecorator := &ColorDecorator{circle}
	colorDecorator.draw()
	colorDecorator.fill("红色")
}
