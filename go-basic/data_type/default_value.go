package data_type

import "fmt"

func RunDefaultValue() {
	//整数
	var aint int
	fmt.Println(aint) //0

	//浮点
	var afloat32 float32
	fmt.Println(afloat32) //0.0

	//复数
	var acomplex64 complex64
	fmt.Println(acomplex64) //0+0i

	//布尔
	var abool bool
	fmt.Println(abool) //false

	//数组
	var aintArray [3]int = [3]int{}
	fmt.Println(aintArray[0]) //0

	//字符串
	var astring string
	if astring == "" {
		fmt.Println("空串")
	}
}
