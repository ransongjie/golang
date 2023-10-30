package random

import (
	"fmt"
	"math/rand"
)

// 未设定种子数而产生的随机数是固定数
func RunRandom() {
	testRandom1()
	testRandom2()
}

func testRandom1() {
	r := rand.Int()
	fmt.Println(r) // 533070578455567690
}

func testRandom2() {
	rand.Seed(100) // ！！！设置1次到处生效，全局变量
	r := rand.Int()
	fmt.Println(r) // 7530908113823513298
}
