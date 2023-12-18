package io

import (
	"fmt"
	"io"
)

func RunIO() {
	test6()
}

// 4
// 1 2 3 4
// 按空格输入整数数组
func test6() {
	var n int
	fmt.Scanf("%d", &n)
	var as []int
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a) // 分隔符 回车 或 换行
		as = append(as, a)
	}
	fmt.Println(as)
}

// xcrj 19
// 输入字符串和数字
func test5() {
	// 获取键盘输入
	var name string
	var age int
	fmt.Print("按空格隔开输入字符串和数字:")
	_, _ = fmt.Scanln(&name, &age) // name age 必须写在同一行，以空格为分割符
	// _, _ = fmt.Scan(&name, &age)// name age 可以写在不同行，以换行 和 空格为分割符
	fmt.Printf("我叫 %s, 今年 %d 岁！", name, age)
}

// 1 2
// 两个数字按空格读取
func test4() {
	var a, b int
	fmt.Println("请输入2个数字以空格隔开:")
	for {
		_, err := fmt.Scan(&a, &b)
		if err == io.EOF {
			break
		}
		fmt.Println(a + b)
	}
}

// xcrj
// 输入1个字符串
func test1() {
	var input string
	fmt.Println("请输入一个字符串：")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("读取输入时发生错误：", err)
		return
	}

	fmt.Println("您输入的字符串是：", input)
}

// 输入1个整数
func test2() {
	var input int
	fmt.Println("请输入一个整数：")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("读取输入时发生错误：", err)
		return
	}
	fmt.Println("您输出的整数是：", input)
}
