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
