package common

import "errors"

// 栈类型
// slice必须放入指针
type Stack[T int | int64 | *[]int | *[]int64] []T

func (stack *Stack[T]) Push(t T) {
	*stack = append(*stack, t)
}

func (stack *Stack[T]) Pop() T {
	if stack.IsEmpty() {
		err := errors.New("Stack is empty.")
		panic(err)
	}
	t := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return t
}

func (stack *Stack[T]) Peek() T {
	if stack.IsEmpty() {
		err := errors.New("Stack is empty.")
		panic(err)
	}
	return (*stack)[len(*stack)-1]
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(*stack) == 0
}
