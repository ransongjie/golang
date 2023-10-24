package funcs
import "fmt"
/*
 文件函数
 结构体函数
 闭包-匿名函数
 多值返回
 函数作入参
 访问控制
 */

//函数 文件函数
func RunFuncs1(){
	fmt.Println("函数")
}

//方法 结构体函数
func RunFuncs2(){
	var employee Employee=Employee{}
	employee.toSay()
}
type Employee struct{}
func (employee Employee) toSay(){
	fmt.Println("我是打工人")
}

//闭包 匿名函数
func RunFuncs3(){
	var anonymous func()//声明
	anonymous=func(){//赋值
		fmt.Println("我是一个匿名函数")
	}
	anonymous()//调用
}

//函数入参
func RunFuncs4(){
	funcfunc(paramFunc)
}
func funcfunc(paramFunc func()){
	paramFunc()
}
func paramFunc(){
	fmt.Println("我是一个入参函数")
}

//访问控制
//首字母大写 整个包唯一 在其它包中可访问
func RunFuncs5(){}
//首字母小写 整个包唯一 在其它包中不可访问
func runFuncs6(){
	fmt.Println("hello")
}