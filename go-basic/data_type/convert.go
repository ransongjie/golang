package data_type

import (
	"fmt"
	"strconv"
)

func RunConvert() {
	//整数》整数
	var cint int = 3
	var cint32 int32 = int32(cint)
	fmt.Println(cint32)

	//整数》浮点
	var aint int = 10
	var afloat32 float32 = float32(aint)
	fmt.Println(afloat32)

	//浮点》整数
	var bfloat32 float32 = 3.2
	var bint int = int(bfloat32)
	fmt.Println(bint)

	//整数》string
	var dint int = 3
	var dstring string = strconv.Itoa(dint) //int to a string
	fmt.Println(dstring)

	//string》整数
	var estring string = "3"
	eint, err := strconv.Atoi(estring) //a string to int
	fmt.Println(eint)
	fmt.Println(err)

	//浮点数》string
	var ffloat32 float32 = 2.3
	var fstring string = strconv.FormatFloat(float64(ffloat32), 'f', 2, 32) //保留2位小数，float32
	fmt.Println(fstring)

	//string》浮点数
	var hstring string = "2.3"
	hfloat32, err := strconv.ParseFloat(hstring, 32) // float32
	fmt.Println(hfloat32)
	fmt.Println(err)

	//数组》切片
	var aarray [3]int = [3]int{1, 2, 3}
	var aslice []int = aarray[:]
	aslice = append(aslice, 4)
	fmt.Println(aslice)
}
