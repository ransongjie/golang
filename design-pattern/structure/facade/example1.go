package main

import "fmt"

type Shape interface {
	draw()
}

type Circle struct{}
type Square struct{}

func (c *Circle) draw() {
	fmt.Println("Circle")
}

func (c *Square) draw() {
	fmt.Println("Square")
}

type ShapeFacade struct {
	circle Shape
	square Shape
}

func (sf *ShapeFacade) drawCircle() {
	sf.circle.draw()
}

func (sf *ShapeFacade) drawSquare() {
	sf.square.draw()
}

func main() {
	shapeFacade := &ShapeFacade{circle: &Circle{}, square: &Square{}}
	shapeFacade.drawCircle()
	shapeFacade.drawSquare()
}
