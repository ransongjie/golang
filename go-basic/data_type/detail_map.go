package data_type

import "fmt"

func RunMapDetail() {
	testMap1()
}

// map可以指定cap
func testMap1() {
	a := map[string]int{}
	fmt.Println("a len:", len(a)) //len: 0。map无cap()函数

	b := make(map[string]int, 2)  //make指定len
	fmt.Println("b len:", len(b)) //len: 0
}
