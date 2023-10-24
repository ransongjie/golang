package data_type

import (
	"fmt"
	"unsafe"
)

func RunTypeLen() {
	//整数
	var aint int = 10
	len := unsafe.Sizeof(aint) //8
	fmt.Println(len)
	var auint uint = 20
	len = unsafe.Sizeof(auint) //8
	fmt.Println(len)

	//布尔
	var abool bool = true
	len = unsafe.Sizeof(abool) //1
	fmt.Println(len)
}
