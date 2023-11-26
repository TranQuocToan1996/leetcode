package uti

func DeepCopySlice[T any](orig []T) []T {
	cpy := make([]T, len(orig))
	copy(cpy, orig)
	return cpy
}

func ReverseSlice[T any](nums []T, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func RemoveElementSlice[T any](slice []T, s int) []T {
	if len(slice) > 0 {
		return append(slice[:s], slice[s+1:]...)
	}
	return slice
}

func InsertToSlice[T any](a []T, index int, value T) []T {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}
