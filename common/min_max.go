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
