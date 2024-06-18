package main

import "fmt"

/*
品牌和型号作手机的成员属性
*/

type Brand struct {
	name string
}
type Model struct {
	name string
}
type Phone struct {
	brand *Brand
	model *Model
}

func (b *Brand) getName() string {
	return b.name
}

func (m *Model) getName() string {
	return m.name
}

func (p *Phone) show() {
	fmt.Println(p.brand.getName(), p.model.getName())
}

func main() {
	brand := &Brand{"Huawei"}
	model := &Model{"Pura70"}
	phone := &Phone{brand: brand, model: model}
	phone.show()
}
