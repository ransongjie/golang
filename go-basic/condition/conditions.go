package condition

import "fmt"

//条件，循环
func RunCondition(){
	// 条件
	var a int =10;
	if a>10{
		fmt.Println("a>10")
	}else if a==10{
		fmt.Println("a=10")
	}else{
		fmt.Println("a<10")
	}

	// 循环
	///数组循环
	var as [3]int=[3]int{1,2,3}
	for i:=0;i<len(as);i++{
		fmt.Println(as[i])
		///跳出所在循环
		for true{
			break
		}
	}
	for _,a:=range as{
		fmt.Println(a)
		///跳过后面的语句，执行下一次循环
		if a>1{
			continue
		}
		fmt.Println(a)
		fmt.Println(a)
	}
	///切片循环
	var aslice []int=[]int{}
	aslice=append(aslice,1)
	aslice=append(aslice,2)
	aslice=append(aslice,3)
	for i,a:=range aslice{
		fmt.Println("idx=",i,", value=",a)
	}
	for _,a:=range aslice{
		fmt.Println(a)
	}
	///map循环
	var amap map[string]int=map[string]int{}
	amap["xcrj1"]=1
	amap["xcrj2"]=2
	amap["xcrj3"]=3
	for k,v:=range amap{
		fmt.Println(k,": ",v)
	}

	// 分支
	///默认无case穿透
	var d int=1
	switch d{
	case 1:
		fmt.Println("d=1")
	case 2:
		fmt.Println("d=2")
	case 3:
		fmt.Println("d=3")
	default:
		fmt.Println("d=?")
	}
	///case穿透，直到无fallthrough
	fmt.Println("case穿透")
	switch d{
	case 1:
		fmt.Println("d=1")
		fallthrough
	case 2:
		fmt.Println("d=2")
		fallthrough
	case 3:
		fmt.Println("d=3")
	default:
		fmt.Println("d=?")
	}
}