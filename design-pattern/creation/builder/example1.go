package main

import "fmt"

/**
 * 不同的套餐有不同的项目（汉堡、饮料）
 * 项目（打包方式、名字、价格）
 * 汉堡：鸡肉汉堡、蔬菜汉堡
 * 饮料：百事可乐、可口可乐
 * 打包方式：瓶子、包装纸
 */

// 打包
type Pack interface {
	pack() string
}

type Bottle struct{}

type Wrapper struct{}

func (p *Bottle) pack() string {
	return "瓶子"
}

func (p *Wrapper) pack() string {
	return "包装纸"
}

// 项目
type Packing interface {
	packing() Pack
}

type Item interface {
	name() string
	price() float32
	Packing // 嵌套接口，避免Burger需要实现所有Item的方法
}

// 汉堡
type Burger struct{}

type ChickenBurger struct{ Burger }

type VegBurger struct{ Burger }

func (b *Burger) packing() Pack {
	return &Wrapper{}
}

func (c *ChickenBurger) name() string {
	return "鸡肉汉堡"
}

func (c *ChickenBurger) price() float32 {
	return 10.3
}

func (v *VegBurger) name() string {
	return "蔬菜汉堡"
}

func (v *VegBurger) price() float32 {
	return 5.3
}

// 饮料
type Drink struct{}

type Coke struct{ Drink }

type Pepsi struct{ Drink }

func (d *Drink) packing() Pack {
	return &Bottle{}
}

func (c *Coke) name() string {
	return "可口可乐"
}

func (c *Coke) price() float32 {
	return 3.5
}

func (p *Pepsi) name() string {
	return "百事可乐"
}

func (p *Pepsi) price() float32 {
	return 3.0
}

// 一顿饭
type Meal struct {
	items []Item
}

func (m *Meal) addItem(item Item) {
	m.items = append(m.items, item)
}

func (m *Meal) showItem() {
	for i := 0; i < len(m.items); i++ {
		fmt.Println(
			m.items[i].name(), ", ",
			m.items[i].price(), ", ",
			m.items[i].packing().pack())
	}
}

func (m *Meal) showCost() {
	sum := 0.0
	for i := 0; i < len(m.items); i++ {
		sum += float64(m.items[i].price())
	}
	fmt.Println(sum)
}

// 套餐构建者
type MealBuilder struct{}

func (mb *MealBuilder) buildVegetariansMeal() *Meal {
	meal := &Meal{}
	// meal.addItem(VegBurger{})报错 因为 指针receiver
	meal.addItem(&VegBurger{})
	meal.addItem(&Pepsi{})
	return meal
}

func (mb *MealBuilder) buildCarnivorousMeal() *Meal {
	meal := &Meal{}
	meal.addItem(&ChickenBurger{})
	meal.addItem(&Coke{})
	return meal
}

// main
func main() {
	mealBuilder := MealBuilder{}
	meal := mealBuilder.buildCarnivorousMeal()
	meal.showItem()
	meal.showCost()
}
