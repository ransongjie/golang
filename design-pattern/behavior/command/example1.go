package main

import "fmt"

// 士兵
type Solider struct {
	name string
}

func (s *Solider) shootPlane() {
	fmt.Println("Shoot plane")
}

func (s *Solider) shootTank() {
	fmt.Println("Shoot tank")
}

// 命令
type Order interface {
	do()
}

type ShootPlane struct {
	solider *Solider
}
type ShootTank struct {
	solider *Solider
}

func (this *ShootPlane) do() {
	this.solider.shootPlane()
}

func (this *ShootTank) do() {
	this.solider.shootTank()
}

// 将军
type Officer struct {
	orders []Order
}

func (this *Officer) addOrder(order Order) {
	this.orders = append(this.orders, order)
}

func (this *Officer) makeOrders() {
	for i := 0; i < len(this.orders); i++ {
		this.orders[i].do()
	}
}

// main
func main() {
	solider1 := &Solider{"solider1"}
	order1 := &ShootPlane{solider1}
	order2 := &ShootPlane{solider1}
	order3 := &ShootTank{solider1}
	officer := &Officer{}
	officer.addOrder(order1)
	officer.addOrder(order2)
	officer.addOrder(order3)
	officer.makeOrders()
}
