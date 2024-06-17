package main

import "fmt"

// 形状
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

// 颜色
type Color interface {
	fill()
}

type Red struct {
	name string
}

type Blue struct {
	name string
}

func (p *Red) fill() {
	fmt.Println(p.name)
}

func (p *Blue) fill() {
	fmt.Println(p.name)
}

// 工厂
type AbstractFactory interface {
	createShape(shapeType string) Shape
	createColor(colorType string) Color
}

type ShapeFactory struct{}
type ColorFactory struct{}

// golang无静态成员方法
func (p *ShapeFactory) createShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return &Circle{"圆形"}
	case "square":
		return &Square{"矩形"}
	}
	return nil
}

// 必须实现接口的所有方法
func (p *ShapeFactory) createColor(colorType string) Color {
	return nil
}

func (p *ColorFactory) createShape(shapeType string) Shape {
	return nil
}

func (p *ColorFactory) createColor(colorType string) Color {
	switch colorType {
	case "red":
		return &Red{"红色"}
	case "blue":
		return &Blue{"蓝色"}
	}
	return nil
}

// 超级工厂
type HyperFactory struct{}

func (P *HyperFactory) createFactory(typ string) AbstractFactory {
	switch typ {
	case "shapeFactory":
		return &ShapeFactory{}
	case "colorFactory":
		return &ColorFactory{}
	}
	return nil
}

func main() {
	hyperFactory := &HyperFactory{}
	shapeFactory := hyperFactory.createFactory("shapeFactory")
	circle := shapeFactory.createShape("circle")
	circle.draw()
}
