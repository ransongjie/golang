package data_type

import "fmt"

func RunString(){
	//字符串
	i:="xcrj"
	j:="xcrj"
	fmt.Println(i==j)// true
	ua:="xc"+"rj"
	fmt.Println(i==ua)// true
	u1:="xc"
	u2:="rj"
	ub:=u1+u2
	fmt.Println("字符串：",i==ub)// true
	uc:=u1+"rj"
	fmt.Println("字符串：",i==uc)// true
}