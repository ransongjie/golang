package main

import "fmt"

// 分离计算器和策略（加减乘除）
type Operation interface {
	do(a int, b int)
}

type Add struct{}
type Sub struct{}
type Mul struct{}
type Div struct{}
type CalCulator struct {
	operation Operation
}

func (this *Add) do(a int, b int) {
	fmt.Println(a + b)
}

func (this *Sub) do(a int, b int) {
	fmt.Println(a - b)
}

func (this *Mul) do(a int, b int) {
	fmt.Println(a * b)
}

func (this *Div) do(a int, b int) {
	fmt.Println(a / b)
}

func (this *CalCulator) calculate(a int, b int) {
	this.operation.do(a, b)
}

func main() {
	add := &Add{}
	calCulator := &CalCulator{add}
	calCulator.calculate(1, 2)
}
