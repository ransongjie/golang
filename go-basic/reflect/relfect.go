package reflect

import (
	"fmt"
	"reflect"
)

func RunReflect() {
	// testReflectType()
	// testReflectTypeFunc()
	// testReflectTypeFunc()
	// testReflectTypePtr()
	// testReflectTypeChan()
	testReflectTypeSlice()
	// testReflectValue()
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
}

// map
func testReflectTypeMap() {
	// reflect.TypeOf(amap).Key() 适用于 map的key
	var amap map[string]int = map[string]int{}
	fmt.Println(reflect.TypeOf(amap).Key()) //string

	// reflect.MapOf() 创建map反射类型。reflect.MakeMap(map反射类型) 创建map反射值
	bmap := reflect.MakeMap(reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0)))
	fmt.Println(bmap) //map[]
}

// func
func testReflectTypeFunc() {
	f := func(int, float64) string { return "" }
	// 入参类型
	fType := reflect.TypeOf(f)
	for i := 0; i < fType.NumIn(); i++ { // 入参个数
		inType := fType.In(i) // 第i个参数的Type
		fmt.Println(inType)   // int float64
	}

	// 出参类型
	for i := 0; i < fType.NumOut(); i++ {
		outType := fType.Out(i)
		fmt.Println(outType)
	}

	// 创建map
	// 定义函数类型
	funcType := reflect.FuncOf(
		[]reflect.Type{reflect.TypeOf(int(0)), reflect.TypeOf(int(0))}, // 入参类型
		[]reflect.Type{reflect.TypeOf(int(0))},                         // 出参类型
		false)                                                          // 非可变参数
	// 创建函数值
	funcValue := reflect.ValueOf(func(a, b int) int { return a + b }).Convert(funcType)
	// 调用函数
	result := funcValue.Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)})[0].Int()
	fmt.Println(result) // 输出: 3
}

// ptr
func testReflectTypePtr() {
	structType := reflect.TypeOf(TypePtr{})
	// 创建1个 指针类型的反射类型 对象
	ptrType := reflect.PtrTo(structType)
	// 同上
	ptrType2 := reflect.PointerTo(structType)
	// Struct type: reflect.TypePtr
	fmt.Println("Struct type:", structType)
	// Pointer to struct type: *reflect.TypePtr
	fmt.Println("Pointer to struct type:", ptrType)
	// Pointer to struct type: *reflect.TypePtr
	fmt.Println("Pointer to struct type:", ptrType2)
}

type TypePtr struct {
	Name int
}

// chan
func testReflectTypeChan() {
	var k = reflect.TypeOf(0)
	// 创建 chan reflect.Type
	fmt.Println(reflect.ChanOf(reflect.SendDir, k)) // chan<- int
}

// slice
func testReflectTypeSlice() {
	// 定义一个int类型的变量
	x := 1
	// 获取x的类型
	t := reflect.TypeOf(x)
	// 创建一个新的切片类型
	sliceType := reflect.SliceOf(t)
	fmt.Println(sliceType) // []int
	// 创建一个新的切片
	newSlice := reflect.MakeSlice(sliceType, 0, 0)
	// 打印新切片的类型和值
	fmt.Println(newSlice.Type(), newSlice.Interface()) //[]int []
}

// struct
func testReflectTypeStruct() {
	// struct 字段
	fields := []reflect.StructField{
		{
			Name: "Name",             //字段名称
			Type: reflect.TypeOf(""), // 字段类型
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(0),
		},
	}
	myStructType := reflect.StructOf(fields)
	fmt.Println(myStructType)
}

// array
func testReflectTypeArray() {
	// 定义元素类型和数组长度
	elemType := reflect.TypeOf(int(0))
	count := 3
	// 创建一个新的数组类型
	arrayType := reflect.ArrayOf(count, elemType)
	// 打印新的数组类型
	fmt.Println(arrayType)
}

// reflect.Value 的方法
func testReflectValue() {
	var a string = "xcrj"
	aValue := reflect.ValueOf(a) // 反射值 源码struct
	fmt.Println(aValue)          //xcrj
}
