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

// 二维切片
func Slice2D[T int | int64](row int, col int) [][]T {
	s2d := make([][]T, row)
	for i := 0; i < row; i++ {
		s2d[i] = make([]T, col)
	}
	return s2d
}

// 三维切片
func Slice3D[T int | int64](row int, col int, high int) [][][]T {
	s3d := make([][][]T, row)
	for i := 0; i < row; i++ {
		s3d[i] = make([][]T, col)
		for j := 0; j < col; j++ {
			s3d[i][j] = make([]T, high)
		}
	}
	return s3d
}
