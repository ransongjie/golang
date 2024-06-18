package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Shape interface {
	getName() string
	setName(name string)
	getColor() string
	setColor(color string)
	getX() float64
	setX(x float64)
	getY() float64
	setY(y float64)
}

type Circle struct {
	name   string
	color  string
	x      float64
	y      float64
	radius float64
}

func (c *Circle) getName() string {
	return c.name
}
func (c *Circle) setName(name string) {
	c.name = name
}
func (c *Circle) getColor() string {
	return c.color
}
func (c *Circle) setColor(color string) {
	c.color = color
}
func (c *Circle) getX() float64 {
	return c.x
}
func (c *Circle) setX(x float64) {
	c.x = x
}
func (c *Circle) getY() float64 {
	return c.y
}
func (c *Circle) setY(y float64) {
	c.y = y
}
func (c *Circle) setRadius(radius float64) {
	c.radius = radius
}
func (c *Circle) getRadius() float64 {
	return c.radius
}

func (c *Circle) show() {
	fmt.Println(c.name,
		", color=", c.color,
		", x=", c.x,
		", y=", c.y,
		", radius=", c.radius)
}

// 包级别变量，类似静态变量
var colorShape = map[string]Shape{}

// 包级别方法，类似静态方法
func getCircle(color string) Shape {
	if _, ok := colorShape[color]; !ok {
		circle := &Circle{}
		circle.setName("圆形")
		circle.setColor(color)
		colorShape[color] = circle
	}
	return colorShape[color]
}

const l int = 5

var colors [l]string = [l]string{"red", "blue", "yellow", "black", "green"}

func randomColor() string {
	rand.Seed(time.Now().UnixNano()) //一次调用，到处生效
	idx := rand.Intn(l)
	return colors[idx]
}

func randomXY() float64 {
	return rand.Float64() * 100
}

func main() {
	//创建大量对象
	for i := 0; i < 1000; i++ {
		circle := getCircle(randomColor()).(*Circle) //类型断言为Circle
		circle.setX(randomXY())
		circle.setY(randomXY())
		circle.setRadius(10.2)
	}

	fmt.Println(len(colorShape))

	circle := getCircle(randomColor()).(*Circle)
	circle.show()
}
