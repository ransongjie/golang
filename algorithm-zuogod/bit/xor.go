package bit

import (
	"fmt"
)

func RunBit() {
	basicRule()
	testSwap()
}

// 异或 具有 交换律 和 结合律
func basicRule() {
	// 0^n=n
	fmt.Println(0 ^ 6) //6
	// n^n=0
	fmt.Println(6 ^ 6) //0
	// 交换律
	fmt.Println((6 ^ 5 ^ 9) == (5 ^ 6 ^ 9)) // true
	// 结合律
	fmt.Println((6 ^ (5 ^ 1) ^ 9) == ((6 ^ 5) ^ 1 ^ 9)) // true
}

func testSwap() {
	as := []int{}
	as = append(as, 1, 2, 3)
	fmt.Println(as)
	swap(&as, 0, 1)
	fmt.Println(as)
}

// 使用异或进行交换
func swap(as *[]int, i int, j int) {
	// a= a^b
	(*as)[i] = (*as)[i] ^ (*as)[j]
	// b= a^b
	(*as)[j] = (*as)[i] ^ (*as)[j]
	// a= a^b
	(*as)[i] = (*as)[i] ^ (*as)[j]
}
