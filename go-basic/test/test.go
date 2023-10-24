package test
import (
	"fmt"
	"strings"
)

func RunTest(){
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
	test12()
}

// func()(sum int) 返回值变量被默认初始化
func test1(){
	s:=[]int{1,2,3}
	sum:=sum1(s)
	fmt.Println(sum)
}
// 返回值变量默认已经初始化, int返回值变量默认初始化为0
func sum1(nums []int) (sum int){
	fmt.Println(sum)
	return
}

//全局变量
var(
	ans []string=[]string{}
)
func test2(){
	ans=append(ans,"xcrj")
	//[xcrj]
	fmt.Println(ans)
}

//string 获取 字符
func test3(){
	//string 获取 字符
	var str string="xcrj"
	fmt.Println(str[1])
	//print format
	fmt.Printf("%c",str[1])
}

//slice string 拼接 join
func test4(){
	// string join 字符串join
	var ls []string=[]string{}
	ls=append(ls,"128")
	ls=append(ls,"0")
	ls=append(ls,"0")
	ls=append(ls,"1")
	var str string=strings.Join(ls,".")
	fmt.Println(str)
}

//golang 变量 常量
func test5(){
	var a int=10
	const B int=10
	fmt.Println(a)
}

//slice append 链式操作
func test6(){
	var lss [][]int=[][]int{}
	var ls []int=[]int{}
	ls=append(ls,1,2,3)
	//append(x,append)
	lss=append(lss,append([]int{},ls...))
	fmt.Println(lss)
}

//移除slice指定位置元素，移除第i个元素
func test7(){
	var s []int=[]int{}
	s=append(s,1)
	s=append(s,2)
	s=append(s,3)
	//删除2
	i:=1
	s=append(s[:i],s[i+1:]...)
	fmt.Println(s)
}

// slice copy
func test8(){
	// 在切片指定位置插入值
	var s []int=[]int{}
	s=append(s,1)
	s=append(s,2)
	s=append(s,3)
	i:=1
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

//slice 中元素交换
func test9(){
	var as []int=[]int{}
	as=append(as,10)
	as=append(as,20)
	// 类似python
	as[0],as[1]=as[1],as[0]
	fmt.Println(as)
}

//map v,ok:=amap["key"]
func test10(){
	var am map[int]int=map[int]int{}
	am[1]=1
	// 值在前，TRUE FALSE在后
	a,ok:=am[1]
	fmt.Println(a," ",ok)
}

// slice 值传递
func test11(){
	// 若函数外部需要和函数内部同时看到变化，则不进行引用传递
	var as []int=[]int{}
	as=append(as,1)
	// 函数调用，值传递，函数中的as和函数外的as不是同一个
	test11a(as)
	// [1]
	fmt.Println(as)
}
func test11a(as []int){
	// [1]
	fmt.Println(as)
	// 根本原因：append时，slice超出cap的情况下会分配一个新的底层array
	// 造成函数内外的as不是同一个
	as=append(as,2)
	// [1 2]
	fmt.Println(as)
}

// slice 引用传递
func test12(){
	var as []int=[]int{}
	as=append(as,1)
	//函数调用，引用传递，函数中的as和函数外的as是同一个
	test12a(&as)
	// [1]
	fmt.Println(as)
}
func test12a(as *[]int){
	// [1]
	fmt.Println(*as)
	*as=append(*as,2)
	// [1 2]
	fmt.Println(*as)
}