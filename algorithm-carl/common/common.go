package common

func Max[T int | int64](a T, b T) T {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Min[T int | int64](a T, b T) T {
	if a <= b {
		return a
	} else {
		return b
	}
}

// 空结构体不占用空间
type Void struct{}

// 变量
var void Void

// 泛型map
type Set[K string | int] map[K]Void

// 泛型成员方法
func (set *Set[K]) Add(key K) {
	(*set)[key] = void
}
