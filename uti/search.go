package uti

import "cmp"

func BinarySearch[T cmp.Ordered](s []T, target T) bool {
	found := false
	low := 0
	high := len(s) - 1
	for low <= high {
		mid := (low + high) / 2
		if s[mid] == target {
			found = true
			break
		}
		if s[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return found
}
