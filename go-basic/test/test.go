package test

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"unicode"
	"unsafe"
)

func RunTest() {
	test61()

	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// test7()
	// test8()
	// test9()
	// test10()
	// test11()
	// test12()
	// test13()
	// test14()
	// test15()
	// test16()
	// test17()
	// test18()
	// test19()
	// test20()
	// test21()
	// test22()
	// test23()
	// test24()
	// test25()
	// test26()
	// test27()
	// // test28()
	// test29()
	// test30()
	// test41()
	// test42()
	// test43()
	// test44()
}

// 测试从集合内外slice是否同1个
// 结论：不是同一个，集合中放入slice的指针
type Stack[T int | []int] []T

func test61() {
	stack := Stack[[]int]{}
	as := []int{}
	as = append(as, 10)
	stack = append(stack, as)
	ls := stack[0]
	ls = append(ls, 1, 2, 3, 4, 5, 6)
	fmt.Println(stack[0][2]) // runtime error: index out of range [2] with length 1
}

// 引用类型必须赋值
func test60() (rs []int) {
	var bs []int
	// runtime error: index out of range [0] with length 0
	bs[0] = 10
	fmt.Println(bs)

	// append会创建新slice
	rs = append(rs, 1)
	fmt.Println(rs)

	var as []int
	as = append(as, 1)
	fmt.Println(as)
	return
}

// nil slice 的len
func test59() {
	var as []int = nil
	fmt.Println(len(as)) // 0
}

// 二维切片make
func test58() {
	// 创建一个 3x3 的二维切片
	rows, cols := 3, 3
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}

	// 给二维切片赋值
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = i*cols + j
		}
	}

	// 打印二维切片
	for i := 0; i < rows; i++ {
		fmt.Println(matrix[i])
	}
}

// 数组长度必须是常量
func test57() {
	// 变量编译不通过
	// var a int = 10
	// var as [a]int = [a]int{}
	// fmt.Println(a, as)

	const b int = 10
	var bs [b]int = [b]int{}
	fmt.Println(b, bs)
}

func test56() {
	am := map[string]int{}
	am["xcrj"] = 18
	test56a(am)
	fmt.Println(am) //map[xcrj:20]
}

// 传递引用
func test56a(am map[string]int) {
	am["xcrj"] = 20
}

// struct 值传递，引用传递
func test55() {
	aa := AA{18, "xcrj"}
	test55a(aa)
	fmt.Println(aa) //{18 xcrj}
	test55b(&aa)
	fmt.Println(aa) //{200 xcrj}
}

// 将完整拷贝一份，函数内外非同一个对象
func test55a(aa AA) {
	aa.Age = 20
	fmt.Println(aa) //{20 xcrj}
}

// 传递指针，函数内外是同一个对象
func test55b(ap *AA) {
	ap.Age = 200
	fmt.Println(*ap) //{200 xcrj}
}

type AA struct {
	Age  int
	Name string
}

func test54() {
	a := 10.0
	b := 10.00000000009
	fmt.Println(a == b) // false
}

// 含有空结构体成员的 结构体 内存对齐
func test53() {
	fmt.Println(unsafe.Sizeof(C{})) // 16
	fmt.Println(unsafe.Sizeof(D{})) // 16
	fmt.Println(unsafe.Sizeof(E{})) // 24
	fmt.Println(unsafe.Sizeof(F{})) // 12
}

type C struct {
	a struct{}
	b int64
	c int64
}
type D struct {
	a int64
	b struct{}
	c int64
}
type E struct {
	a int64
	b int64
	c struct{}
}
type F struct {
	a int32
	b int32
	c struct{}
}

// 结构体 内存对齐
func test52() {
	fmt.Println(unsafe.Sizeof(A{})) // 24
	fmt.Println(unsafe.Sizeof(B{})) // 16
}

type A struct {
	a int32 // 4
	b int64 // 8
	c int32 // 4
}
type B struct {
	a int32 // 4
	b int32 // 4
	c int64 // 8
}

// unsafe.Alignof()
func test51() {
	// 基础类型，unsafe.Alignof(x) = min(64位机器/8，unsafe.Sizeof(x))
	fmt.Println(unsafe.Alignof(int(1)))        // 8 = min(8,8)
	fmt.Println(unsafe.Alignof(int32(1)))      // 4 = min(8,4)
	fmt.Println(unsafe.Alignof(int64(1)))      // 8 = min(8,8)
	fmt.Println(unsafe.Alignof(complex128(1))) // 8 = min(8,16)

	// struct类型，unsafe.Alignof(结构体)=max{unsafe.Alignof(每个字段)}
	fmt.Println(unsafe.Alignof(A{})) // 8
	fmt.Println(unsafe.Alignof(B{})) // 8
}

// unsafe.Pointer和uintptr的区别
func test50() {
	var w *W = new(W)
	//{0，0}
	fmt.Println(w.b, w.c)
	//通过指针运算给b变量赋值为10
	//uintptr(unsafe.Pointer(w)) unsafe.Pointer()类型不可以进行+-运算，uintptr可以
	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	*((*int)(b)) = 10
	//{10，0}
	fmt.Println(w.b, w.c)
}

type W struct {
	b int32
	c int64
}

// unsafe.Sizeof(x)
func test49() {
	var z bool
	fmt.Println(unsafe.Sizeof(z))                   //1B)
	fmt.Println(unsafe.Sizeof(int(1)))              // 8B
	fmt.Println(unsafe.Sizeof(uintptr(1)))          // 8
	fmt.Println(unsafe.Sizeof(map[string]string{})) // 8
	fmt.Println(unsafe.Sizeof(string("")))          // 16
	fmt.Println(unsafe.Sizeof([]string{}))          // 24
	var a interface{}
	fmt.Println(unsafe.Sizeof(a)) // 16
}

// slice是引用数据类型，作函数入参时，没有扩容的前提下（append造成扩容），函数内外的slice是同一个
func test48() {
	a := []int{1, 2, 3, 4, 5}
	//[1 2 3 4 5]
	fmt.Println(a)
	test48a(a)
	//[2 4 6 8 10]
	fmt.Println(a)
}
func test48a(a []int) {
	for k, v := range a {
		a[k] = v * 2
	}
	//[2 4 6 8 10]
	fmt.Println(a)
}

// 数组是基本数据类型，作函数入参时，函数内外的array不是同一个
func test47() {
	arr := [8]int{}
	for i := 0; i < 8; i++ {
		arr[i] = i
	}
	//[0 1 2 3 4 5 6 7]
	fmt.Println(arr)
	test47a(arr)
	//[0 1 2 3 4 5 6 7]
	fmt.Println(arr)
}
func test47a(arr [8]int) {
	for k, v := range arr {
		arr[k] = v * 2
	}
	//[0  2 4 6 8 10 12 14]
	fmt.Println(arr)
}

// 两个nil可能不相等。先比较类型再比较值
func test46() {
	var p *int = nil
	//interface{}变量被赋值为*int类型nil
	var q interface{} = p
	fmt.Println(p == q)   //true
	fmt.Println(p == nil) //true
	fmt.Println(q == nil) //false interface{}运行时类型是*int, nil是Type类型，类型不一致

	//interface{}变量被赋值为nil
	var o interface{} = nil
	fmt.Println(o == nil) //true
}

// byte 和 rune的类型
func test45() {
	var a byte
	typeOfA := reflect.TypeOf(a)
	//uint8 uint8
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	var b rune
	typeOfB := reflect.TypeOf(b)
	//int32 int32
	fmt.Println(typeOfB.Name(), typeOfB.Kind())
}

// func()(sum int) 返回值变量被默认初始化
func test1() {
	s := []int{1, 2, 3}
	sum := sum1(s)
	fmt.Println(sum)
}

// 返回值变量默认已经初始化, int返回值变量默认初始化为0
func sum1(nums []int) (sum int) {
	fmt.Println(sum)
	return
}

// 全局变量
var (
	ans []string = []string{}
)

func test2() {
	ans = append(ans, "xcrj")
	//[xcrj]
	fmt.Println(ans)
}

// string 获取 字符
func test3() {
	//string 获取 字符
	var str string = "xcrj"
	fmt.Println(str[1])
	//print format
	fmt.Printf("%c", str[1])
}

// slice string 拼接 join
func test4() {
	// string join 字符串join
	var ls []string = []string{}
	ls = append(ls, "128")
	ls = append(ls, "0")
	ls = append(ls, "0")
	ls = append(ls, "1")
	var str string = strings.Join(ls, ".")
	fmt.Println(str)
}

// golang 变量 常量
func test5() {
	var a int = 10
	const B int = 10
	fmt.Println(a)
}

// slice append 链式操作
func test6() {
	var lss [][]int = [][]int{}
	var ls []int = []int{}
	ls = append(ls, 1, 2, 3)
	//append(x,append)
	lss = append(lss, append([]int{}, ls...))
	fmt.Println(lss)
}

// 移除slice指定位置元素，移除第i个元素
func test7() {
	var s []int = []int{}
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	//删除2
	i := 1
	s = append(s[:i], s[i+1:]...)
	fmt.Println(s)
}

// slice copy
func test8() {
	// 在切片指定位置插入值
	var s []int = []int{}
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	i := 1
	// xcrj 注意 要添加1个元素 需要多append 2个元素，元素值随便
	// 要添加2个元素 需要多append 2个元素，
	s = append(s, 0)
	//[3,0]
	fmt.Println(s[i+1:])
	fmt.Println(s[i:])
	//[3,0], [2,3,0]
	//copy(param1,param2) 将param2 copy into param1中。拷贝元素个数=min(param1, param2)
	copy(s[i+1:], s[i:])
	//[3 0 2 3]
	fmt.Println(s)
	s[i] = 10
	//[1 10 2 3]
	fmt.Println(s)
}

// slice 中元素交换
func test9() {
	var as []int = []int{}
	as = append(as, 10)
	as = append(as, 20)
	// 类似python
	as[0], as[1] = as[1], as[0]
	fmt.Println(as)
}

// map v,ok:=amap["key"]
func test10() {
	var am map[int]int = map[int]int{}
	am[1] = 1
	// 值在前，TRUE FALSE在后
	a, ok := am[1]
	fmt.Println(a, " ", ok)
}

// slice 值传递
func test11() {
	// 若函数外部需要和函数内部同时看到变化，则不进行引用传递
	var as []int = []int{}
	as = append(as, 1)
	// 函数调用，值传递，函数中的as和函数外的as不是同一个
	test11a(as)
	// [1]
	fmt.Println(as)
}
func test11a(as []int) {
	// [1]
	fmt.Println(as)
	// 根本原因：append时，slice超出cap的情况下会分配一个新的底层array
	// 造成函数内外的as不是同一个
	as = append(as, 2)
	// [1 2]
	fmt.Println(as)
}

// slice 引用传递
func test12() {
	var as []int = []int{}
	as = append(as, 1)
	//函数调用，引用传递，函数中的as和函数外的as是同一个
	test12a(&as)
	// [1]
	fmt.Println(as)
}
func test12a(as *[]int) {
	// [1]
	fmt.Println(*as)
	*as = append(*as, 2)
	// [1 2]
	fmt.Println(*as)
}

// 深拷贝copy(slice,slice) 》 append(slice,slice) append引用
func test13() {
	var lss [][]int = [][]int{}
	var ls []int = []int{}
	ls = append(ls, 1)
	ls = append(ls, 2)
	ls = append(ls, 3)
	// 是append 引用，所以进行深拷贝, 注意仍然不能修改ls2
	var ls2 []int = make([]int, len(ls))
	copy(ls2, ls)
	lss = append(lss, ls2)
	fmt.Println(lss)
	ls[0] = 10
	fmt.Println(lss)
}

// append(slice,slice) append引用
func test14() {
	//slice append slice
	var lss [][]int = [][]int{}
	var ls []int = []int{}
	ls = append(ls, 1)
	ls = append(ls, 2)
	ls = append(ls, 3)
	// 是append 引用
	lss = append(lss, ls)
	fmt.Println(lss)
	ls[0] = 10
	fmt.Println(lss)
}

// 切片元素赋值, 必须使用 索引，不能使用range
func test15() {
	var fs []int = make([]int, 3)
	// for _,f:=range fs{
	// 	f=10;
	// }
	for i := 0; i < 3; i++ {
		fs[i] = 10
	}
	fmt.Println(fs)
}

// 二维slice创建，取值
func test16() {
	// xcrj 注意 二维slice的创建不需要 两个括号 {{}} 如果写两个括号就是初始化了二维slice的第一行数据
	var ss [][]int = [][]int{}
	// slice没有给len不能直接赋值，必须append
	// ss[0]=[]int{}
	ss = append(ss, []int{})
	// slice没有给len不能直接赋值，必须append
	// ss[0][1]=1
	// ss[0][2]=2
	ss[0] = append(ss[0], 1)
	ss[0] = append(ss[0], 2)
	fmt.Println(ss)
	//
	var as []int = ss[0]
	fmt.Println(as[0])
}

// slice range是顺序遍历的
func test17() {
	var as []int = []int{}
	as = append(as, 1)
	as = append(as, 2)
	as = append(as, 3)
	for _, a := range as {
		fmt.Println(a)
	}
}

// 格式化输出
func test18() {
	var c int
	var d float32
	//0
	fmt.Printf("%d", c)
	//0.0
	fmt.Printf("%0.1f", d)
}

// 全局变量默认值
var a int
var b float32

func test19() {
	//0
	fmt.Println(a)
	//0.0
	fmt.Println(b)
}

// 字符相减》类型转换为int
func test20() {
	var a rune = '2' //rune 等同于uint32
	var b rune = '0'
	var i int = int(a - b) //类型转换
	fmt.Println(i)
}

// slice 的索引值不能为 -数
func test21() {
	//slice -1
	var s []int = []int{}
	s = append(s, 1)
	s = append(s, 2)
	// slice 的索引值不能为 -数
	// fmt.Println(s[-1])
}

// 无返回值函数 java返回void
func test22() {
	//无返回值函数 java返回void
}

// slice 实现 队列
func test23() {
	//slice 实现 队列
	var queue []int = []int{}
	//入队
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	//出队
	var peak int = queue[0]
	fmt.Println(peak)
	queue = queue[1:]
}

// slice 实现 栈
func test24() {
	//slice 实现 栈
	var limit int = 10
	var stack []int = make([]int, 0)
	//push 压栈
	for i := 0; i < 3; i++ {
		stack = append(stack, i)
	}
	//pop 弹出
	var val int = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println(val)
	//peak 查看栈顶元素
	var pk int = stack[len(stack)-1]
	fmt.Println(pk)
	//栈空
	if len(stack) == 0 {
		fmt.Println("栈空")
	}
	//栈满
	if len(stack) > limit {
		fmt.Println("栈满")
	} else {
		fmt.Println("栈未满")
	}
}

// 判断是字符还是数字
func test25() {
	//判断字母 数字
	var a rune = 'a'
	var ab bool = unicode.IsLetter(a)
	fmt.Println(ab)

	var a1 rune = '2'
	var ab1 bool = unicode.IsDigit(a1)
	fmt.Println(ab1)
}

// 字母 大小写转换
func test26() {
	//unicode
	//字母 大写转小写
	var a rune = 'A'
	var al rune = unicode.ToLower(a)
	fmt.Printf("%c", al)
	//字母 大写转小写 反之
	var a1 rune = 'a'
	var al1 rune = unicode.ToUpper(a1)
	fmt.Printf("%c", al1)
}

// 字符串 大小写转换
func test27() {
	// strings
	//字符串 大写转小写
	var a string = "ABCD"
	var al string = strings.ToLower(a)
	fmt.Println(al)
	//字符串 大写转小写 反之
	var a1 string = "abcd"
	var al1 string = strings.ToUpper(a1)
	fmt.Println(al1)
}

// string 《》字符串数组[]rune
func test29() {
	// []rune切片转string
	var str1 string = "xcrj"
	// 写法 强制类型转换
	var cs []rune = []rune(str1)
	// 写法 强制类型转换
	var str2 string = string(cs)
	fmt.Println(str2)
}

// map key存在判断
func test30() {
	// 是否存在
	var am map[string]int = map[string]int{"xcrj": 1}
	v := am["xcrj"]
	fmt.Println(v)
	_, ok := am["xcrj"]
	fmt.Println(ok)
}

// 空结构体不占用空间
type void struct{}

var vv void

// 创建set
func test31() {
	// map[string]类型
	var set map[string]void = make(map[string]void)
	// set["xcrj"]=值
	set["xcrj"] = vv
	set["xcrj1"] = vv
	fmt.Println(set)
}

// 数组转slice slice转数组
func test32() {
	// array to slice
	var as [3]int = [3]int{1, 2, 3}
	var s1 []int = as[:]
	fmt.Println(s1)
	// slice to array
	var bs [3]int = [3]int{}
	copy(bs[:], s1) // copy(目的,原)
	fmt.Println(bs)
}

// 数组/切片判断相等 判等
func test33() {
	// 数组判断相等
	var a1 [3]int = [3]int{1, 2, 3}
	var a2 [3]int = [3]int{1, 2, 3}
	fmt.Println(a1 == a2)

	// byte切片判断相等
	var bs1 []byte = []byte{1, 2, 3}
	var bs2 []byte = []byte{1, 2, 3}
	fmt.Println(bytes.Equal(bs1, bs2))

	// 普通切片判断相等
	var s1 []int = []int{1, 2, 3}
	var s2 []int = []int{1, 2, 3}
	// fmt.Println(s1==s2)
	// 反射性能消耗大
	fmt.Println(reflect.DeepEqual(s1, s2))
	// 自己写判断
	var len int = len(s1)
	var eql bool = true
	for i := 0; i < len; i++ {
		if s1[i] != s2[i] {
			eql = false
			break
		}
	}
	fmt.Println(eql)
}

// string长度
func test34() {
	var astr string = "xcrj"
	fmt.Println(len(astr))
}

// 工厂模式
func test35() {
	var test14 Test14 = Constructor("xcrj")
	test14.ShowName()
}

type Test14 struct {
	// public属性
	Name string
}

// 工厂模式创建类，静态工厂
// public 方法
func Constructor(name string) Test14 {
	return Test14{name}
}
func (this *Test14) ShowName() {
	fmt.Println(this.Name)
}

// slice有序列表 set没有 自己根据map创建
func test36() {
	var as []int = []int{}
	// xcrj go 只有有序列表 没有 无序列表 set
	// 重复值 后者 不会 覆盖前者 有序列表
	as = append(as, 1)
	as = append(as, 1)
	// [1,1]
	fmt.Println(as)
}

// int slice排序
func test37() {
	var as []int = []int{}
	as = append(as, 1)
	as = append(as, 3)
	as = append(as, 2)
	// 不需要 as=sort.Ints(as)
	sort.Ints(as)
	fmt.Println(as)
}

// 遍历string的每个字符
func test38() {
	var str string = "abc"
	for _, c := range str {
		fmt.Printf("%c", c)
	}
}

// 创建空slice
func test39() {
	// 创建空slice 必须写0
	var as []int = make([]int, 0)
	as = append(as, 4)
	as = append(as, 5)
	as = append(as, 6)
}

// slice错误遍历方式
func test40() {
	var as []int = make([]int, 0)
	as = append(as, 4)
	// 必须 for _,a:=range as
	for a := range as {
		// 0 1 2, 默认输出索引值
		fmt.Println(a)
	}
}

// map getordefault [string]int, 不存在的key默认为0
func test41() {
	var amap map[string]int = make(map[string]int)
	amap["xcrj"] = 1
	v, ok := amap["xcrj1"]
	fmt.Println(v, " ", ok) //0, false
	// get or default
	var r int = amap["xcrj1"]
	fmt.Println(r)
	amap["xcrj1"] = r
	fmt.Println(amap)
}

// 字符串拼接
func test42() {
	var a string = "a"
	var b string = "b"
	fmt.Println(a + b) //ab
}

// a slice 插入 b slice 末尾
func test43() {
	var as []rune = make([]rune, 0)
	as = append(as, '1')
	var bs []rune = make([]rune, 0)
	bs = append(bs, '1')
	bs = append(bs, '2')

	var ras []rune = make([]rune, len(bs)-len(as))
	// a slice 插入 b slice 末尾
	ras = append(ras, as...)
	// [0 49]
	fmt.Println(ras)
}

// nil slice
func test44() {
	var a []rune = nil
	// []
	fmt.Println(a)
	if 1 == 1 {
		a = make([]rune, 2)
	}
	// [0 0]
	fmt.Println(a)
}
