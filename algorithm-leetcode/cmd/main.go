package main

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"unicode"

	"com.xcrj/golang/algorithm-leetcode/internal/offersa/pass1"
)

func main() {
	test54()
}
func test54() {
	s := []int{1, 2, 3}

	sum := sum54(s)
	fmt.Println(sum)
}

// 返回值变量默认已经初始化
// int返回值变量默认初始化为0
func sum54(nums []int) (sum int) {
	fmt.Println(sum)
	return
}

var (
	ans []string = []string{}
)

func test53() {
	ans = append(ans, "xcrj")
	fmt.Println(ans)
}

func test52() {
	//string 获取 字符
	var str string = "xcrj"
	fmt.Println(str[1])
	//print format
	fmt.Printf("%c", str[1])
}
func test51() {
	// string join 字符串join
	var ls []string = []string{}
	ls = append(ls, "128")
	ls = append(ls, "0")
	ls = append(ls, "0")
	ls = append(ls, "1")
	var str string = strings.Join(ls, ".")
	fmt.Println(str)
}
func test50() {
	//golang  变量 常量
	var a int = 10
	const B int = 10
	fmt.Println(a)
}
func test49() {
	// slice深拷贝，append方法，将产生临时切片
	var lss [][]int = [][]int{}
	var ls []int = []int{}
	ls = append(ls, 1)
	ls = append(ls, 2)
	ls = append(ls, 3)
	// golang链式操作
	lss = append(lss, append([]int{}, ls...))
	fmt.Println(lss)
}
func test48() {
	//在切片移除指定位置元素
	var s []int = []int{}
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	//删除2
	i := 1
	s = append(s[:i], s[i+1:]...)
	fmt.Println(s)
}
func test47() {
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
	//[3,0], [2,3]
	copy(s[i+1:], s[i:])
	s[i] = 10
	fmt.Println(s)
}
func test46() {
	//slice 中元素交换
	var as []int = []int{}
	as = append(as, 10)
	as = append(as, 20)
	as[0], as[1] = as[1], as[0]
	fmt.Println(as)
}
func test45() {
	var am map[int]int = map[int]int{}
	am[1] = 1
	// 值在前，TRUE FALSE灾后
	a, ok := am[1]
	fmt.Println(a, " ", ok)
}
func test43() {
	// xcrj 若函数外部需要和函数内部同时看到变化，则不进行引用传递
	// 值传递，函数中的as和函数外的as不是同一个
	var as []int = []int{}
	as = append(as, 1)
	test44(as)
	// [1]
	fmt.Println(as)
}
func test44(as []int) {
	fmt.Println(as)
	as = append(as, 2)
	// [1 2]
	fmt.Println(as)
}
func test42() {
	//深拷贝 slice
	//slice append slice
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
	//
	ls[0] = 10
	fmt.Println(lss)
}
func test41() {
	//slice append slice
	var lss [][]int = [][]int{}
	var ls []int = []int{}
	ls = append(ls, 1)
	ls = append(ls, 2)
	ls = append(ls, 3)
	// 是append 引用
	lss = append(lss, ls)
	fmt.Println(lss)
	//
	ls[0] = 10
	fmt.Println(lss)
}
func test40() {
	//切片元素赋值, 必须使用 索引，不能使用range
	var fs []int = make([]int, 3)
	// for _,f:=range fs{
	// 	f=10;
	// }
	for i := 0; i < 3; i++ {
		fs[i] = 10
	}
	fmt.Println(fs)
}
func test39() {
	var mc2 pass1.MyCalendar2 = pass1.Constructor58b()
	var ins [][]int = [][]int{}
	ins = append(ins, []int{10, 20})
	ins = append(ins, []int{15, 25})
	ins = append(ins, []int{20, 30})
	for _, se := range ins {
		var r bool = mc2.Book(se[0], se[1])
		fmt.Println(r)
	}
}
func test38() {
	// 测试二维 slice 按照索引取值
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
func test37() {
	// golang 函数变量
	// 见Solution54.go
}
func test36() {
	// golang值传递和引用传递
	// * Type, *变量名
	// 见Solution52.go
}
func test35() {
	//结论 是顺序遍历的
	//slice range是否顺序遍历
	var as []int = []int{}
	as = append(as, 1)
	as = append(as, 2)
	as = append(as, 3)
	for _, a := range as {
		fmt.Println(a)
	}
}
func test34() {
	//fmt
	//Printf()
	//Sprintf()
}
func test33() {
	//方法变量 默认值
	var c int
	var d float32
	//0
	fmt.Printf("%d", c)
	//0.0
	fmt.Printf("%0.1f", d)
}

var a int
var b float32

func test32() {
	//文件变量 默认值
	//0
	fmt.Println(a)
	//0.0
	fmt.Println(b)
}
func test31() {
	var a rune = '2'
	var b rune = '0'
	var i int = int(a - b)
	fmt.Println(i)
}
func test30() {
	// rune a - rune b 》 int
	var a rune = 'a'
	var b rune = 'b'
	// 类型转换
	var i int = int(b - a)
	fmt.Println(i)
}
func test29() {
	//slice -1
	var s []int = []int{}
	s = append(s, 1)
	s = append(s, 2)
	// slice 的索引值不能为 -数
	// fmt.Println(s[-1])
}
func test28() {
	//无返回值函数 java返回void
}
func test27() {
	//slice 实现 队列
	var queue []int = []int{}
	//入队
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	//对头元素值
	var peak int = queue[0]
	fmt.Println(peak)
	//出队
	queue = queue[1:]
}
func test26() {
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
func test25() {
	//判断字母 数字
	var a rune = 'a'
	var ab bool = unicode.IsLetter(a)
	fmt.Println(ab)

	var a1 rune = '2'
	var ab1 bool = unicode.IsDigit(a1)
	fmt.Println(ab1)
}
func test24() {
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
func test23() {
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

func test22() {
	//获取变量类型
	fmt.Println(reflect.TypeOf(pass1.MinWindow("a", "a")).String())
}
func test21() {
	// []rune切片转string
	var str1 string = "xcrj"
	// 写法 强制类型转换
	var cs []rune = []rune(str1)
	// 写法 强制类型转换
	var str2 string = string(cs)
	fmt.Println(str2)
}

func test20() {
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

func test19() {
	//创建set
	// map[string]类型
	var set map[string]void = make(map[string]void)
	// set["xcrj"]=值
	set["xcrj"] = vv
	set["xcrj1"] = vv
	fmt.Println(set)
}

func test18() {
	// array to slice
	var as [3]int = [3]int{1, 2, 3}
	var s1 []int = as[:]
	fmt.Println(s1)
	// slice to array
	var bs [3]int = [3]int{}
	copy(bs[:], s1) // copy(目的,原)
	fmt.Println(bs)
}

func test17() {
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
func test16() {
	// string to []rune
	var str string = "xcrj"
	var cs []rune = []rune(str)
	var c rune = cs[0]
	// 输出字符
	fmt.Printf("%c", c)
}
func test15() {
	var astr string = "xcrj"
	fmt.Println(len(astr))
}
func test14() {
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

func test13() {
	var as []int = []int{}
	// xcrj go 只有有序列表 没有 无序列表 set
	// 重复值 后者 不会 覆盖前者 有序列表
	as = append(as, 1)
	as = append(as, 1)
	// [1,1]
	fmt.Println(as)
}

func test12() {
	var as []int = []int{}
	as = append(as, 1)
	as = append(as, 3)
	as = append(as, 2)
	// 不需要 as=sort.Ints(as)
	sort.Ints(as)
	fmt.Println(as)
}
func test11() {
	var c rune = 'b' - 'a'
	var s []int = make([]int, 2)
	s[c] = 1
	fmt.Println(s)
}

func test10() {
	var str string = "abc"
	for _, c := range str {
		fmt.Printf("%c", c)
	}
}

func test9() {

	var a int = 1
	fmt.Println(a << 1)

	// 最好使用精确类型
	var b int32 = 1
	fmt.Println(b << 1)
}

func test8() {
	// 创建空slice 必须写0
	var as []int = make([]int, 0)
	as = append(as, 4)
	as = append(as, 5)
	as = append(as, 6)

	// 必须 for _,a:=range as
	for a := range as {
		// 0 1 2, 默认输出索引值
		fmt.Println(a)
	}
}

func test7() {
	var amap map[string]int = make(map[string]int)
	amap["xcrj"] = 1
	v, ok := amap["xcrj1"]
	fmt.Println(v, " ", ok)
	// get or default
	var r int = amap["xcrj1"]
	fmt.Println(r)
	amap["xcrj1"] = r
	fmt.Println(amap)
}

func test6() {
	var amap map[string]int = make(map[string]int, 2)
	// 0, 实际长度
	fmt.Println(len(amap))
	amap["xcrj"] = 1
	amap["xcrj2"] = 2
	amap["xcrj3"] = 3
	// 3, 实际长度
	fmt.Println(len(amap))
	// no cap func for map
	// fmt.Println(cap(amap))
}

func test5() {
	var a string = "a"
	var b string = "b"
	fmt.Println(a + b)
}
func test4() {
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

func test3() {
	var as []rune = make([]rune, 0)
	as = append(as, '1')
	var bs []rune = make([]rune, 0)
	bs = append(bs, '1')
	bs = append(bs, '2')

	var ras []rune = make([]rune, len(bs))
	// dest src
	copy(ras, as)
	// [49 0]
	fmt.Println(ras)
	for _, a := range ras {
		// 格式化为字符后输出
		fmt.Printf("%c", a)
	}
}

func test2() {
	var a []rune = nil
	// []
	fmt.Println(a)
	if 1 == 1 {
		a = make([]rune, 2)
	}
	// [0 0]
	fmt.Println(a)
}

func test1() {
	// slice len 和 cap
	// 1=len 2=cap
	// a[0]默认值为0，赋值len长度的值为0
	var a []int = make([]int, 1, 2)
	// 1,
	fmt.Println(len(a))
	a = append(a, 1)
	a = append(a, 2)
	// 3, 实际长度
	fmt.Println(len(a))
	// 4, todo(slice cap 如何变化)
	fmt.Println(cap(a))
}
