package err
import "fmt"
import "errors"

/*
 什么情况下退出程序，异常退出（panic退出），正常退出（return退出）

 errors.New("err msg") //定义异常
 panic("err msg") //抛出异常

 recover() 
 - 捕获panic的异常
 - recover需要放到defer中

 defer 函数
 - defer不能对未指定返回值变量名的变量赋值，类似java的finally。
 - 退出之前执行，包括异常退出和正常退出。
 - 多个defer压栈
 */

func RunErr(){
	testErr0()
	testErr1()
	testErr2()
	testErr3()
	testErr4()
	testErr5()
	testErr6()
	testErr7()
}

// defer 用来在退出时执行一个函数，常用于释放资源和异常处理
func testErr0(){
	defer fmt.Println("defer被执行")
	
}

// panic defer recover
// 系统panic
func testErr1(){
	defer func(){
		if err:=recover();err!=nil{
			//runtime error: index out of range [100] with length 3
			fmt.Println("ops err: ",err)
		}
	}()

	al:=[]int{}
	al=append(al,1)
	al=append(al,2)
	al=append(al,3)
	fmt.Println(al[100])
}

// panic defer recover
// 自行panic
func testErr2(){
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println("err: ",err)
		}
	}()//

	panic("zero err")
}

// errors包 自定义err
func testErr3(){
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()

	err:=errors.New("self defined err")
	panic(err)
}

// return error, err!=nil
func testErr4(){
	a,err:=runErr4()
	if err!=nil{
		fmt.Println(a,", ",err)
	}
}

func runErr4()(int,error){
	err:=errors.New("my err")
	return 0,err
}

// defer
func testErr5(){
	c:=runErr5()
	//c= 30
	fmt.Println("c= ",c)
} 

func runErr5()(c int){
	defer func(){
		fmt.Println("defer is executed when return")
	}()

	a:=10
	b:=20
	return a+b
}

// defer压栈, 写在后面的先执行
func testErr6(){
	defer func(){
		fmt.Println("I'am first defer")
	}()
	defer func(){
		fmt.Println("I'am second defer")
	}()
	defer func(){
		fmt.Println("I'am third defer")
	}()

	return
} 

// defer与return
func testErr7(){
	r1:=runErr6a()
	//r1= 1
	fmt.Println("r1=",r1)

	r2:=runErr6b()
	//r2= 11
	fmt.Println("r2=",r2)
}

//defer不能对未指定返回值变量名的变量赋值
//未指定返回值变量名，由golang创建临时变量存储返回值，defer不可见 系统创建的临时返回变量
//同于java
func runErr6a()(int){
	r:=1

	defer func(){
		r+=10
		//runErr6a defer a= 11
		fmt.Println("runErr6a defer a=",r)
	}()

	return r
}

//defer能对指定返回值变量名的变量赋值
func runErr6b()(r int){
	defer func(){
		r+=10
		//runErr6b defer a= 11
		fmt.Println("runErr6b defer a=",r)
	}()
	
	r=1
	return r
}