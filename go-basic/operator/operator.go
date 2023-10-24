package operator

import "fmt"

//赋值运算符，算术运算符，比较运算符，逻辑运算符，位运算符，指针相关运算符
func RunOperator() {
	//赋值, =, +=, -= ...
	var aint int = 10
	var bint int = 20
	bint += 1
	bint -= 1
	fmt.Println(aint)
	fmt.Println(bint)

	//算数, +, -, *, /, %, ++, --
	fmt.Println(aint + bint)
	fmt.Println(aint - bint)
	fmt.Println(aint * bint)
	fmt.Println(aint / bint)
	fmt.Println(aint % bint)
	aint++
	aint--
	fmt.Println(aint)
	fmt.Println(aint)
	// ++aint//无先加，先减
	// --aint
	// fmt.Println(aint)
	// fmt.Println(aint)

	//关系, ==, !=, >, <, >=, <=
	fmt.Println(aint == bint)
	fmt.Println(aint != bint)
	fmt.Println(aint > bint)
	fmt.Println(aint < bint)
	fmt.Println(aint >= bint)
	fmt.Println(aint <= bint)

	//逻辑, &&, ||, !
	var abool bool = true
	var bbool bool = false
	fmt.Println(abool && bbool)
	fmt.Println(abool || bbool)
	fmt.Println(!abool)

	//位, &, |, ^ (异或), <<, >>
	var cint int = 10
	var dint int = 20
	fmt.Println(cint & dint)
	fmt.Println(cint | dint)
	fmt.Println(cint ^ dint)
	fmt.Println(cint >> 1)//不改变符号位
	fmt.Println(cint)//不改变原数值
	fmt.Println(cint << 1)//不改变符号位
	fmt.Println(cint)//不改变原数值

	//指针, &, *
	var fint int = 10
	var fintPointer *int = &fint
	fmt.Println(fintPointer)//指针址
	fmt.Println(*fintPointer)//指针值
}