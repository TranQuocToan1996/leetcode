package uti

import "cmp"

func Max[T cmp.Ordered](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

type Minus interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func AbsMinus[T Minus](x1, x2 T) T {
	if x1 > x2 {
		return x1 - x2
	}
	return x2 - x1
}
