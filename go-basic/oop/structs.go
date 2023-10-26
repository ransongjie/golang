package oop

import "fmt"

func RunStruct() {
	testStruct1()
	testStruct2()
	testStruct3()
	testStruct4()
	testStruct5()
}

func testStruct1() {
	var teacher Teacher = Teacher{"xcrj", 30, 20.2}
	fmt.Println(teacher)
	teacher.getName()
}

//结构体
type Teacher struct {
	//属性
	Name   string
	Age    int
	Salary float32
}

//成员方法
func (teacher Teacher) getName() {
	fmt.Println(teacher.Name)
}

func testStruct2() {
	//new 分配空间，默认初始化，返回指针类型
	var pteacher *Teacher = new(Teacher)
	pteacher.Name = "xcrj"
	pteacher.Age = 30
	pteacher.Salary = 20.2
	fmt.Println(*pteacher)

	//&取址
	pteacher = &Teacher{"xcrj", 30, 20.2}
	fmt.Println(*pteacher)
}

//组合代替继承
//组合中部分与整体同生共死
//聚合中部分与整体不同生共死
func testStruct3() {
	car := Car{Motor{"motor 520"}, "BYD"}
	fmt.Println(car)
}

type Car struct {
	Motor
	Name string
}
type Motor struct {
	Name string
}

// 指针receiver。指针.成员方法
func testStruct4() {
	var weight *Set[string] = &Set[string]{} // 是每个单词
	weight.Add("xcrj")
	fmt.Println(&weight)
}

// Type receiver。Type.成员方法
func testStruct5() {
	var weight Set[string] = Set[string]{} // 是每个单词
	weight.Add2("xcrj")
	fmt.Println(weight)
}

// 空结构体不占用空间
type Void struct{}

// 变量
var void Void

// 泛型map
type Set[K string | int] map[K]Void

// 泛型成员方法
func (set *Set[K]) Add(key K) {
	(*set)[key] = void
}

// 泛型成员方法
func (set Set[K]) Add2(key K) {
	set[key] = void
}
