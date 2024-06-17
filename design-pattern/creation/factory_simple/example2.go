package main

import "fmt"

type IShape interface {
	draw()
}

type Shape struct {
	name string
}

type Circle struct {
	Shape
}

func (p *Circle) draw() {
	fmt.Println(p.name)
}

type Square struct {
	Shape
}

func (p *Square) draw() {
	fmt.Println(p.name)
}

func getCircle() IShape {
	return &Circle{Shape{"圆形"}}
}

func getSquare() IShape {
	return &Circle{Shape{"矩形"}}
}

func main() {
	circle := getCircle()
	circle.draw()

	square := getSquare()
	square.draw()
}
