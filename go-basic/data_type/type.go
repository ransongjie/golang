package data_type

import (
	"fmt"
	"time"
)

func RunDataType() {
	//整数
	var aint int = 10
	fmt.Println(aint) //int32或int64
	var auint uint = 10
	fmt.Println(auint)  //uint32或uint64
	var abyte byte = 10 //类似uint8
	fmt.Println(abyte)
	var aint8 int8 = 10
	fmt.Println(aint8)
	var auint8 uint8 = 10
	fmt.Println(auint8)
	var aint16 int16 = 10
	fmt.Println(aint16)
	var auint16 uint16 = 10
	fmt.Println(auint16)
	var aint32 uint32 = 10
	fmt.Println(aint32)
	var auint32 uint32 = 10
	fmt.Println(auint32)
	var aint64 uint64 = 10
	fmt.Println(aint64)
	var auint64 uint64 = 10
	fmt.Println(auint64)

	//浮点
	var afloat32 float32 = 10.0
	fmt.Println(afloat32)
	var afloat64 float64 = 10.0
	fmt.Println(afloat64)

	//布尔
	var abool bool = false
	fmt.Println(abool)

	//字符
	var arune rune='a'//rune 等同于uint32
	fmt.Printf("字符类型：%c",arune)
	fmt.Println()

	//复数
	var acomplex64 complex64 = 1 + 2i
	fmt.Println(acomplex64)
	fmt.Println(real(acomplex64))
	fmt.Println(imag(acomplex64))
	var acomplex128 complex128 = 1 + 2i
	fmt.Println(acomplex128)
	fmt.Println(real(acomplex128))
	fmt.Println(imag(acomplex128))

	//null
	var anull *int = nil
	fmt.Println(anull)

	//数组
	var aintArray [3]int = [3]int{}
	fmt.Println(aintArray)
	var bintArray [3]int = [3]int{1, 2, 3}
	fmt.Println(bintArray)
	var cintArray [10]int = [10]int{0: 2, 5: 6}
	fmt.Println(cintArray)
	var dintArray [3][4]int = [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(dintArray)
	var eintArray [3]int = [3]int{1, 2, 3}
	for i, v := range eintArray { //遍历
		fmt.Println(i, ": ", v)
	}

	//指针
	var bint int = 10
	var bintPointer *int = &bint
	fmt.Println(bintPointer)
	fmt.Println(*bintPointer)

	//string
	var astring string = "xcrj"
	fmt.Println(astring)

	//时间
	var atime time.Time = time.Now()
	fmt.Println(atime)
	fmt.Println(atime.Unix())                    //时间戳
	fmt.Println(atime.Format("2006/1/02 15:04")) //时间格式化

	//slice
	var aslice []int = []int{}              //创建
	fmt.Println("aslice len=", len(aslice)) //slice len= 0
	fmt.Println("aslice cap=", cap(aslice)) //slice cap= 0
	var bslice []int = make([]int, 0, 4)    //创建 指定len cap。给定的len将占用空间
	fmt.Println("bslice len=", len(bslice)) //slice len= 0
	fmt.Println("bslice cap=", cap(bslice)) //slice cap= 4
	aslice = append(aslice, 1)              //增
	aslice = append(aslice, 2, 3, 4)        //增多个
	var aaslice []int = aslice[1:]          //删
	fmt.Println(aaslice)
	var cslice []int = []int{} //查
	cslice = append(cslice, 1)
	cslice = append(cslice, 2)
	cslice = append(cslice, 3)
	for i, c := range cslice { //index, value
		fmt.Println(i, ":", c)
	}
	for _, c := range cslice { //忽略index
		fmt.Println(c)
	}

	//map
	var amap map[string]int = map[string]int{"xcrj1": 1, "xcrj2": 2} //创建
	fmt.Println("amap len=", len(amap))                              //amap len= 2。map无cap()函数
	var bmap map[string]int = make(map[string]int, 10)               //创建 make指定cap
	fmt.Println("bmap len=", len(bmap))                              //bmap len= 0
	amap["xcrj3"] = 3                                                //增
	amap["xcrj3"] = 30                                               //改
	for k, v := range amap {                                         //查
		fmt.Println(k, ":", v)
	}
}
