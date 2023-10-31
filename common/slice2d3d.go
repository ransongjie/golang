package common

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
