package reflect

import (
	"fmt"
	"reflect"
)

func RunReflect() {
	testReflectType()
	testReflectValue()
}

type MyInt int

// reflect.Type 的方法
func testReflectType() {
	// reflect.TypeOf() 适用于任意类型
	var a int = 10
	aType := reflect.TypeOf(a) // 反射类型 源码interface
	fmt.Println(aType)         // int

	var b MyInt = 20
	bType := reflect.TypeOf(b) //反射类型
	fmt.Println(bType)         // reflect.MyInt

	bKind := reflect.ValueOf(b).Kind() //go基本类型
	fmt.Println(bKind)                 //int

	// reflect.TypeOf(as).Elem() 适用于 array, slice, map, chan, pointer
	var as []int = []int{}
	fmt.Println(reflect.TypeOf(as).Elem()) // int

	// reflect.TypeOf(amap).Key() 适用于 map的key
	var amap map[string]int = map[string]int{}
	fmt.Println(reflect.TypeOf(amap).Key()) //string

	// reflect.MapOf() 创建map反射类型。reflect.MakeMap(map反射类型) 创建map反射值
	bmap := reflect.MakeMap(reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0)))
	fmt.Println(bmap) //map[]
}

// reflect.Value 的方法
func testReflectValue() {
	var a string = "xcrj"
	aValue := reflect.ValueOf(a) // 反射值 源码struct
	fmt.Println(aValue)          //xcrj
}
