package main

import "fmt"

type Shape interface {
	draw()
}

type Circle struct {
	name string
}

type Square struct {
	name string
}

func (p *Circle) draw() {
	fmt.Println(p.name)
}

func (p *Square) draw() {
	fmt.Println(p.name)
}

// 工厂
type Factory interface {
	create() Shape
}

type CircleFactory struct{}

type SquareFactory struct{}

func (p *CircleFactory) create() Shape {
	return &Circle{"圆形"}
}

func (p *SquareFactory) create() Shape {
	return &Square{"矩形"}
}

// main
func main() {
	circleFactory := &CircleFactory{}
	circle := circleFactory.create()
	circle.draw()
}
