package generic

import (
	"fmt"
)

/**
类型形参，实参，约束
类型形参(Type parameter)
类型实参(Type argument)
类型形参列表(Type parameter list)
类型约束(Type constraint)
泛型类型(Generic type)
实例化(Instantiation)
*/

/**
  泛型 slice map chan
  泛型 func
  泛型 struct interface
  泛型 receiver（成员方法）
*/

func RunGeneric() {
	testGeneric1()
	testGeneric2()
	testGeneric3()
	testGeneric4()
	testGeneric5()
	testGeneric6()
	testGeneric7()
	testGeneric8()
}

// 泛型 slice
func testGeneric1() {
	// 下面的类型不能接受多种类型的参数
	type IntSlice []int         //新定义类型IntSlice int slice
	type Float32Slice []float32 //新定义类型Float32Slice float slice
	type StrSlice []string      //新定义类型StrSlice string slice

	/**
	T 类型形参
	[T int|float32|string] 类型形参列表
	int|float32|string 类型约束
	IFSSlice 泛型类型
	*/
	type IFSSlice[T int | float32 | string] []T

	/**
	IFSSlice[int] int类型实参
	T 类型形参
	ifss 实例
	*/
	var ifss IFSSlice[int] = []int{}
	ifss = append(ifss, 1)
	fmt.Println(ifss)
}

// 泛型 map
func testGeneric2() {
	type IntStrMap map[int]string

	type IFSMap[K int | float32 | string, V int | float32 | string] map[K]V

	var iFSMap IFSMap[string, int] = map[string]int{}
	iFSMap["xcrj"] = 0
	fmt.Println(iFSMap)
}

// 泛型 chan
func testGeneric3() {
	type IntChan chan int

	type IFSChan[T int | float32 | string] chan T

	//make(chan int,10) 最大capacity
	var ifsChan IFSChan[int] = make(chan int, 1)
	ifsChan <- 10
	val := <-ifsChan
	fmt.Println(val)
}

// 泛型 func，约束放到方法名后
func testGeneric4() {
	var a int = 10
	var b int = 20
	var r = add[int](a, b)
	fmt.Println(r)
}
func add[T int | int64 | float32 | float64](a T, b T) (r T) {
	r = a + b
	return r
}

// 泛型 struct
func testGeneric5() {
	var stu Student[int] = Student[int]{123, "xcrj"}
	fmt.Println(stu)
}

type Student[T int | string] struct {
	Code T
	Name string
}

// 泛型 interface, 由方法集合 向 类型集合转变
func testGeneric6() {
	var apple Apple[int] = FujiApple{"fuji apple"}
	apple.GetPrice(10)
}

type Apple[T int | string] interface {
	GetPrice(price T)
}
type FujiApple struct {
	Name string
}

func (fujiApple FujiApple) GetPrice(price int) {
	fmt.Println(fujiApple.Name, ": ", price)
}

// 泛型 receiver，成员方法名前
func testGeneric7() {
	var ifSlice IFSlice[int] = []int{}
	ifSlice = append(ifSlice, 1)
	ifSlice = append(ifSlice, 2)
	ifSlice = append(ifSlice, 3)
	var sum int = ifSlice.sum()
	fmt.Println(sum)
}

// 切片
type IFSlice[T int | float32] []T

// 给切片增加方法。入参，出参，函数体都可以使用泛型
func (ifSlice IFSlice[T]) sum() T {
	var sum T
	for _, num := range ifSlice {
		sum += num
	}
	return sum
}

// 泛型 receiver
// 泛型队列
// interface{} 代码任意类型 可以用 any 代替
func testGeneric8() {
	var iQue Queue[int]
	iQue.Offer(1)
	iQue.Offer(2)
	iQue.Offer(3)
	v, ok := iQue.Poll()
	fmt.Println(v, ", ", ok)
	fmt.Println(iQue.Size())
}

type Queue[T interface{}] struct {
	elements []T
}

// 入队列
func (queue *Queue[T]) Offer(element T) {
	queue.elements = append(queue.elements, element)
}

// 出队列
// 出参2：队列是否为空
func (queue *Queue[T]) Poll() (T, bool) {
	var r T
	if len(queue.elements) == 0 {
		return r, true
	}
	r = queue.elements[0]
	queue.elements = queue.elements[1:]
	return r, len(queue.elements) == 0
}

// 队列元素个数
func (queue *Queue[T]) Size() int {
	return len(queue.elements)
}
