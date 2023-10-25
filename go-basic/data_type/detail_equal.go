package data_type

import (
	"fmt"
	"time"
	"reflect"
)

func RunEqual(){
	/*基本类型*/
	//整数
	a:=10
	b:=10
	fmt.Println(a==b)// true
	//浮点
	c:=10.0
	d:=10.000000000009
	fmt.Println(c==d)// false, 浮点数使用差值比较，满足精度要求即可
	//复数
	e:=2+1i
	f:=2+1i
	fmt.Println(e==f)// true
	//布尔
	h:=true
	g:=true
	fmt.Println(h==g)// true
	//字符串
	i:="xcrj"
	j:="xcrj"
	fmt.Println(i==j)// true
	ua:="xc"+"rj"
	fmt.Println(i==ua)// true
	u1:="xc"
	u2:="rj"
	ub:=u1+u2
	fmt.Println("字符串：",i==ub)// true
	uc:=u1+"rj"
	fmt.Println("字符串：",i==uc)// true
	//数组
	as:=[3]int{1,2,3}
	bs:=[3]int{1,2,3}
	fmt.Println(as==bs)// true
	//地址指针
	var ap uintptr=20
	var bp uintptr=20
	fmt.Println(ap==bp)// true

	/*引用类型*/
	//时间
	var atime time.Time = time.Now()
	var btime time.Time = time.Now()
	fmt.Println(atime==btime)// false
	fmt.Println(atime.Equal(btime))// false 
	//slice
	a1:=[]int{1,2,3}
	b1:=[]int{1,2,3}
	fmt.Println(a1==nil)// slice==只能是nil，不能slice==slice，编译不通过
	fmt.Println(reflect.DeepEqual(a1,b1))// true
	//map
	m1:=map[string]int{"ab":1}
	m2:=map[string]int{"ab":1}
	fmt.Println(m1==nil)// map==只能是nil，不能map==map，编译不通过
	fmt.Println(reflect.DeepEqual(m1,m2))// true
	//interface
	//nil
	var p *int=nil;
    var q interface{}=p//interface{}变量被赋值为*int类型nil
    fmt.Println(p==q)//true
    fmt.Println(p==nil)//true
    fmt.Println(q==nil)//false interface{}运行时类型是*int, nil是Type类型，类型不一致
	var o interface{}=nil //interface{}变量被赋值为nil
	fmt.Println(o==nil)//true
	
	/*结构体类型*/
	//struct
	s1:=MyA{18,"xcrj"}
	s2:=MyA{18,"xcrj"}
	fmt.Println(s1==s2)// true
	fmt.Println(reflect.DeepEqual(s1,s2))// true
	//普通指针
	p1:=new(MyA)
	p1.age=18
	p1.name="xcrj"
	p2:=new(MyA)
	p2.age=18
	p2.name="xcrj"
	fmt.Println("p1==p2: ", p1==p2)// false
	fmt.Println("reflect.DeepEqual(p1,p2): ", reflect.DeepEqual(p1,p2))// true

	/*自定义类型*/
	//type
	q1:=10
	var q2 MyInt=10
	// fmt.Println(q1==q2)// 编译错误类型不匹配
	fmt.Println(reflect.DeepEqual(q1,q2))// false。强类型语言
}
type MyA struct{
	age int
	name string
}
type MyInt int