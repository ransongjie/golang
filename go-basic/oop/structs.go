package oop

import "fmt"

func RunStruct(){
	testStruct1()
	testStruct2()
	testStruct3()
}

func testStruct1(){
	var teacher Teacher=Teacher{"xcrj",30,20.2}
	fmt.Println(teacher)
	teacher.getName()
}
//结构体
type Teacher struct{
	//属性
	Name string
	Age int
	Salary float32
}
//成员方法
func (teacher Teacher) getName(){
	fmt.Println(teacher.Name)
}


func testStruct2(){
	//new 分配空间，默认初始化，返回指针类型
	var pteacher *Teacher=new(Teacher)
	pteacher.Name="xcrj"
	pteacher.Age=30
	pteacher.Salary=20.2
	fmt.Println(*pteacher)

	//&取址
	pteacher=&Teacher{"xcrj",30,20.2}
	fmt.Println(*pteacher)
}

//组合代替继承
//组合中部分与整体同生共死
//聚合中部分与整体不同生共死
func testStruct3(){
	car:=Car{Motor{"motor 520"},"BYD"}
	fmt.Println(car)
}
type Car struct{
	Motor
	Name string
}
type Motor struct{
	Name string
}