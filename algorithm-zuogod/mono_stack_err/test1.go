package mono_stack

// 错在 receiver 的使用上应当使用 指针 receiver

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

func Max[T int | int64](a T, b T) T {
	if a >= b {
		return a
	} else {
		return b
	}
}

// 栈类型
// slice必须放入指针
type Stack[T int | int64 | *[]int | *[]int64] []T

// attention 值receiver 传入副本，在函数中操作副本
// 统一用指针receiver。当然对象不需要被修改时可以使用值receiver
func (stack Stack[T]) push(t T) {
	stack = append(stack, t)
}

func (stack Stack[T]) pop() T {
	if stack.isEmpty() {
		err := errors.New("Stack is empty.")
		panic(err)
	}
	t := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return t
}

func (stack Stack[T]) peek() T {
	if stack.isEmpty() {
		err := errors.New("Stack is empty.")
		panic(err)
	}
	return stack[len(stack)-1]
}

func (stack Stack[T]) isEmpty() bool {
	return len(stack) == 0
}

// 目标：子数组累加和与子数组中最小值的乘积A的最大值
// [2,1,3] (2+1+3)*1=6
// as数组可以包含相等的数
func RunMonoStack() {
	times := 1000
	maxLen := 10
	maxVal := 10

	for i := 0; i < times; i++ {
		as := getAs(maxLen, maxVal)
		a := cp(as)
		b := fp(as)
		if a != b {
			fmt.Println(a)
			fmt.Println(b)
			panic(errors.New("a is not equal to b."))
		}
	}
}

// 生成as[]
func getAs(maxLen, maxVal int) []int {
	//attension 未设定种子数而产生的随机数是固定数
	rand.Seed(100)
	len := rand.Intn(maxLen) + 1 // 1~maxLen
	as := make([]int, len)
	for i := 0; i < len; i++ {
		as[i] = rand.Intn(maxVal + 1) // 0~maxLen+1
	}
	return as
}

// 单调栈求目标
func fp(as []int) int {
	rs := fpg(as)
	max := math.MinInt
	for _, e := range rs {
		max = Max[int](max, getA(e, as))
	}
	return max
}

// 获取目标 slice，即获得每种结果
func fpg(as []int) (rs []Result) {
	//处理特殊情况
	n := len(as) // nil slice的len=0
	if as == nil || n == 0 {
		return nil
	}
	//构建结果
	//可以存储相等元素的单调栈
	stack := Stack[*[]int]{}
	//遍历as
	for i := 0; i < n; i++ {
		//单调栈为空，直接放入
		if stack.isEmpty() {
			ss := []int{}
			ss = append(ss, i)
			stack.push(&ss)
		}
		//单调栈顶元素>当前元素，直接放入
		if as[(*stack.peek())[0]] > as[i] {
			ss := []int{}
			ss = append(ss, i)
			stack.push(&ss)
		}
		//单调栈顶元素=当前元素，放入slice尾部
		if as[(*stack.peek())[0]] == as[i] {
			ss := stack.peek()
			*ss = append(*ss, i)
		}
		//单调栈顶元素<当前元素，循环出栈，直到栈空||栈顶元素!<当前元素
		//出栈元素 in (新栈顶元素,当前元素)，左边界可能不存在，即可能栈空导致无新栈顶元素
		for !stack.isEmpty() && as[(*stack.peek())[0]] < as[i] {
			//当前元素as[i]是右边界
			rightIdx := i
			rightVal := as[rightIdx]

			// 出栈元素是子数组中最小值
			ss := stack.pop()

			// 新栈顶元素 是左边界
			// 默认是ss的第1个元素
			leftIdx := (*ss)[0]
			leftVal := as[leftIdx]
			if !stack.isEmpty() {
				leftIdx = (*stack.peek())[len(*stack.peek())-1]
				leftVal = as[leftIdx]
			}

			// 遍历ss构建目标
			for _, idx := range *ss {
				result := Result{}
				result.leftIdx = leftIdx
				result.leftVal = leftVal
				result.rightIdx = rightIdx
				result.rightVal = rightVal
				result.curIdx = idx
				result.curVal = as[idx]
				rs = append(rs, result)
			}
		}
		//入栈当前元素
		i--
	}

	//处理单调栈中剩余元素，右边界一定不存在，即没有当前元素
	for !stack.isEmpty() {
		ss := stack.pop()

		//默认是ss的最后1个元素
		rightIdx := (*ss)[len(*ss)-1]
		rightVal := as[rightIdx]

		//默认是ss的第1个元素
		leftIdx := (*ss)[0]
		leftVal := as[leftIdx]
		if !stack.isEmpty() {
			leftIdx = (*stack.peek())[len(*stack.peek())-1]
			leftVal = as[leftIdx]
		}

		// 遍历ss构建目标
		for _, idx := range *ss {
			result := Result{}
			result.leftIdx = leftIdx
			result.leftVal = leftVal
			result.rightIdx = rightIdx
			result.rightVal = rightVal
			result.curIdx = idx
			result.curVal = as[idx]
			rs = append(rs, result)
		}
	}
	return
}

// 对数器求目标
func cp(as []int) int {
	rs := cpg(as)
	max := math.MinInt
	for _, e := range rs {
		max = Max[int](max, getA(e, as))
	}
	return max
}

// 求子数组累加和*子数组最小值的乘积A
func getA(e Result, as []int) (a int) {
	sum := 0
	for i := e.leftIdx; i <= e.rightIdx; i++ {
		sum += as[i]
	}

	return sum * e.curVal
}

// 获取目标 slice，即获得每种结果
func cpg(as []int) (rs []Result) {
	n := len(as) // nil slice的len=0
	if as == nil || n == 0 {
		return nil
	}
	//将as中每个元素当做子数组的最小值，求子数组的左右边界
	for i := 0; i < n; i++ {
		e := Result{}
		e.curIdx = i
		e.curVal = as[i]

		e.leftIdx = i     //可能没有左边界索引
		e.leftVal = as[i] //可能没有左边界值
		// 求左边界
		for j := i - 1; j > -1; j-- {
			if as[j] > as[i] {
				e.leftIdx = j
				e.leftVal = as[j]
				break
			}
		}
		// 求右边界
		e.rightIdx = i     //可能没有右边界索引
		e.rightVal = as[i] //可能没有右边界值
		for j := i + 1; j < n; j++ {
			if as[j] > as[i] {
				e.rightIdx = j
				e.rightVal = as[j]
				break
			}
		}

		rs = append(rs, e)
	}

	return
}

// 获取目标 slice，即获得每种结果
func cpg_bk(as []int) (es []Element) {
	n := len(as) // nil slice的len=0
	if as == nil || n == 0 {
		return nil
	}

	//将as中每个元素当做子数组的最小值，求子数组的左右边界
	for i := 0; i < n; i++ {
		e := Element{}
		e.curIdxP = &i
		e.curValP = &as[i]

		e.leftIdxP = nil //可能没有左边界索引
		e.leftValP = nil //可能没有左边界值
		// 求左边界
		for j := i - 1; j > -1; j-- {
			if as[j] > as[i] {
				e.leftIdxP = &j
				e.leftValP = &as[j]
				break
			}
		}
		// 求右边界
		e.rightIdxP = nil //可能没有右边界索引
		e.rightValP = nil //可能没有右边界值
		for j := i + 1; j < n; j++ {
			if as[j] > as[i] {
				e.rightIdxP = &j
				e.rightValP = &as[j]
				break
			}
		}

		es = append(es, e)
	}

	return
}

// 记录子数组左边界索引和值，右边界索引和值，子数组最小值索引和值
type Element struct {
	leftIdxP  *int // 使用指针的目的 nil代表不存在
	leftValP  *int
	rightIdxP *int
	rightValP *int
	curIdxP   *int
	curValP   *int
}

type Result struct {
	leftIdx  int
	leftVal  int
	rightIdx int
	rightVal int
	curIdx   int
	curVal   int
}
