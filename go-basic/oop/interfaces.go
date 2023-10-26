package oop

//方法的聚合体 》类型的聚合体
/*
只含有方法的接口
只含有类型的接口
含有方法和类型的接口
*/
import (
	"fmt"
)

// ////////////////////只含有方法的接口//////////////////////////
// 只含有方法的接口
type Phoe interface {
	call()
}
type HuaweiPhone struct {
}

func (huaweiPhone HuaweiPhone) call() {
	fmt.Println("I am Huawei, I can call you!")
}

// ////////////////////只含有类型的接口//////////////////////////

// 简化写法
type ASlice[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16] []T
type AIFace interface { // 接口成了类型集合体
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16
}
type ASlice2[T AIFace] []T

// 组合写法
type IntIFace interface {
	int | int8 | int16 | int32 | int64
}
type UintIFace interface {
	uint | uint8 | uint16 | uint32 | uint64
}
type FloatIFace interface {
	float32 | float64
}
type BSlice[T IntIFace | UintIFace | FloatIFace | string] []T

// 类型并集 |
type MyUnion interface {
	int | int8 | int16
}

// 类型交集 换行
type MyInter interface {
	int | int8 | int16
	int
}
type MySliceInter[T MyInter] []T

var mySliceInter MySliceInter[int]

// var mySliceInter MySliceInter[int8]//报错

// 所有类型的集合 空接口
type AllType[T interface{}] []T

// 等价值写法
type AllType2[T any] []T

/*
go是强类型语言，在基本类型上的自定义类型，如何允许
~底层类型
不允许的情况
1. ~非基本类型
2. ~接口
允许的情况
1. ~[]byte
*/

// 不允许底层类型
var bSlice BSlice[int]

type BInt int

// var bSlice2 BSlice[BInt]//将报错

// 允许底层类型
type IntIFace2 interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type UintIFace2 interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
type FloatIFace2 interface {
	~float32 | ~float64
}
type BSlice2[T IntIFace2 | UintIFace2 | FloatIFace2 | ~string] []T

var cSlice BSlice2[int]

type BInt2 int

var cSlice2 BSlice2[BInt2]

// /////////////////////含有类型和方法的接口/////////////////////////
// 只含有类型限制的接口
type PPP interface {
	int | int8 | int32
}

// 含有类型限制的接口
// 限制为自定义基本类型
type IFI interface {
	~int | ~float32
	RInt() int
	RFloat() float32
}
type Intme int

func (im Intme) RInt() int {
	return 100
}
func (im Intme) RFloat() float32 {
	return 100.2
}
func RunTF1() {
	var im Intme = 10
	var rint int = im.RInt()
	var rfloat float32 = im.RFloat()
	fmt.Println(rint, ", ", rfloat)
}

// 含有类型限制的接口
// 限制为结构体
type Phone interface {
	Iphone | Huawei
	Name() string
	Price() float32
}
type Iphone struct{}
type Huawei struct{}

func (iphone Iphone) Name() string {
	return "Iphone"
}
func (iphone Iphone) Price() float32 {
	return 200.2
}
func (huawei Huawei) Name() string {
	return "Huawei"
}
func (huawei Huawei) Price() float32 {
	return 200.1
}
func RunTF2() {
	var iphone Iphone = Iphone{}
	var name string = iphone.Name()
	var price float32 = iphone.Price()
	fmt.Println(name, ", ", price)

	var huawei Huawei = Huawei{}
	var name2 string = huawei.Name()
	var price2 float32 = huawei.Price()
	fmt.Println(name2, ", ", price2)
}

// 基本接口，只含有方法的接口
type Fruit interface {
	Name() string
}
type Cabbage struct{}

func (cabbage Cabbage) Name() string {
	return "cabbage"
}
func RunTF3() {
	var cabbage Fruit = Cabbage{}
	var name string = cabbage.Name()
	fmt.Println(name)
}

// /////////////////////例子/////////////////////////
// 可比较 comparable golang内置接口
// == !=
type MyMapC[K comparable, V any] map[K]V

// 可排序 ordered
// > < >= <=
// Ordered 代表所有可比大小排序的类型
type Ordered interface {
	Integer | Float | ~string
}
type Integer interface {
	Signed | Unsigned
}
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
type Float interface {
	~float32 | ~float64
}
