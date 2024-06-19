package main

import "fmt"

type Element interface {
	operate()
	getName() string
}

type Leaf struct {
	name string
}
type Composite struct {
	name string
	// elements [](*Element) 这里接口不定义为指针
	elements []Element
}

func (l *Leaf) getName() string {
	return l.name
}

func (l *Leaf) operate() {
	fmt.Println("叶子: ", l.name)
}

func (c *Composite) getName() string {
	return c.name
}

func (c *Composite) operate() {
	fmt.Println("组合: ", c.name)
	for i := 0; i < len(c.elements); i++ {
		c.elements[i].operate()
	}
}

func (c *Composite) add(e Element) {
	c.elements = append(c.elements, e)
}

func (c *Composite) del(e Element) {
	idx := -1
	for i := 0; i < len(c.elements); i++ {
		//利用接口方法获取结构体相同属性
		if c.elements[i].getName() == e.getName() {
			idx = i
			break
		}
	}
	if idx != -1 {
		c.elements = append(c.elements[:idx], c.elements[idx+1:]...)
	}
}

func (c *Composite) get(name string) Element {
	for i := 0; i < len(c.elements); i++ {
		//利用接口方法获取结构体相同属性
		if c.elements[i].getName() == name {
			return c.elements[i]
		}
	}
	return nil
}

func main() {
	leaf1 := &Leaf{"leaf1"}
	leaf2 := &Leaf{"leaf2"}
	leaf3 := &Leaf{"leaf3"}
	leaf4 := &Leaf{"leaf4"}
	leaf5 := &Leaf{"leaf5"}
	composite1 := &Composite{name: "composite1", elements: []Element{}}
	composite2 := &Composite{name: "composite2", elements: []Element{}}
	composite1.add(leaf1)
	composite1.add(leaf2)
	composite1.add(leaf3)
	composite2.add(leaf4)
	composite2.add(leaf5)
	composite1.add(composite2)
	composite1.operate()
}
