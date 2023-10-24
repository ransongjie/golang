package data_type

import "fmt"

func RunSliceDetail() {
	testSlice1()
}

//slice扩容
//扩容操作可能会导致底层数据的复制，因此可能会引起性能开销。如果需要频繁地添加元素到slice中，
//最好预先指定一个较大的容量，以减少扩容的次数。
func testSlice1() {
	q := []int{}
	fmt.Println("q len:", len(q), "cap:", cap(q)) //len: 0 cap: 0

	//创建时给定len=0=cap
	y := make([]int, 0)
	fmt.Println("len:", len(y), "cap:", cap(y)) //len: 0 cap: 0

	//创建时给定len=20=cap
	z := make([]int, 20)
	fmt.Println("len:", len(z), "cap:", cap(z)) //len: 20 cap: 20

	//创建时给定len和cap
	a := make([]int, 1, 2)
	a = append(a, 1)
	fmt.Println("len:", len(a), "cap:", cap(a)) //len: 2 cap: 2

	//cap在1024以下，以2倍增长
	b := make([]int, 0, 2)
	b = append(b, 1, 2, 3)
	fmt.Println("b len:", len(b), "cap:", cap(b)) //len: 3 cap: 4
	b = append(b, 1, 2, 3)
	fmt.Println("b len:", len(b), "cap:", cap(b)) //len: 6 cap: 8
	b = append(b, 1, 2, 3)
	fmt.Println("b len:", len(b), "cap:", cap(b)) //len: 9 cap: 16

	//cap增长到1024后，以1.25倍增长
	c := []int{}
	for i := 0; i < 1024; i++ {
		c = append(c, i)
	}
	fmt.Println("c len:", len(c), "cap:", cap(c)) //c len: 1024 cap: 1280
}
