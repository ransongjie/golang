package main

import "fmt"

// 导出的结构体，字段和方法都可以在包外访问
type PublicStruct struct {
	PublicField  string // 导出的字段
	privateField string // 未导出的字段，只能在包内访问
}

func (p *PublicStruct) PublicMethod() string {
	return "Public Method"
}

func (p *PublicStruct) privateMethod() string {
	return "Private Method"
}

func main() {
	ps := PublicStruct{
		PublicField:  "Hello",
		privateField: "World",
	}

	fmt.Println(ps.PublicField)  // 可以访问
	fmt.Println(ps.privateField) //

	fmt.Println(ps.PublicMethod())  // 可以访问
	fmt.Println(ps.privateMethod()) //
}
