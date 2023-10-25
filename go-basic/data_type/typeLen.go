package data_type

import (
	"fmt"
	"unsafe"
)

func RunTypeLen() {
	//整数
	a:=int8(1)
	b:=uint8(1)
	c:=int16(1)
	d:=uint16(1)
	e:=int32(1)
	f:=uint32(1)
	g:=int64(1)
	h:=uint64(1)
	i:=int(1)
	j:=uint(1)
	abyte:=byte(1)
	arune:=rune('a')
	fmt.Println("int8: ", unsafe.Sizeof(a))// 1
	fmt.Println("uint8: ", unsafe.Sizeof(b))// 1
	fmt.Println("int16: ", unsafe.Sizeof(c))// 2
	fmt.Println("uint16: ", unsafe.Sizeof(d))// 2
	fmt.Println("int32: ", unsafe.Sizeof(e))// 4
	fmt.Println("uint32: ", unsafe.Sizeof(f))// 4
	fmt.Println("int64: ", unsafe.Sizeof(g))// 8
	fmt.Println("uint64: ", unsafe.Sizeof(h))// 8
	fmt.Println("int: ", unsafe.Sizeof(i))//64位机器是 8Byte
	fmt.Println("uint: ", unsafe.Sizeof(j))//64位机器是 8Byte
	fmt.Println("byte: ",unsafe.Sizeof(abyte))// 1
	fmt.Println("rune: ",unsafe.Sizeof(arune))// 4
	
	//浮点
	aa:=float32(1.0)
	bb:=float64(1.0)
	fmt.Println("float32: ", unsafe.Sizeof(aa))// 4
	fmt.Println("float64: ", unsafe.Sizeof(bb))// 8

	//bool
	ab:=true
	fmt.Println("bool: ", unsafe.Sizeof(ab))// 1Byte

	//复数
	k:=complex64(1+2i)
	fmt.Println("complex64: ", unsafe.Sizeof(k))// 8Byte
	m:=complex128(1+2i)
	fmt.Println("complex128: ", unsafe.Sizeof(m))// 16Byte

	//指针
	o:=uintptr(1)//存储指针值，变量地址，做地址运算
	fmt.Println("uintptr: ", unsafe.Sizeof(o))// 64位机器是 8Byte

	//字符串
	s:="abc"
	fmt.Println("string: ", unsafe.Sizeof(s))// 64位机器是 16Byte

	//interface{}
	var t interface{}
	fmt.Println("interface{}: ", unsafe.Sizeof(t))// 64位机器是 16Byte

	//数组
	as:=[10]int{}
	fmt.Println("array: ", unsafe.Sizeof(as))// 64位机器是 10*8Byte

	//slice
	ss:=[]int{}
	fmt.Println("slice: ", unsafe.Sizeof(ss))// 64位机器是 24Byte

	//map
	mm:=map[string]int{}
	fmt.Println("map: ", unsafe.Sizeof(mm))// 64位机器是 8Byte
}
