package main

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
