package uti

func DeepCopySlice[T any](orig []T) []T {
	cpy := make([]T, len(orig))
	copy(cpy, orig)
	return cpy
}
