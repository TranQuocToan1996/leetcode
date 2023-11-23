package medium

import (
	"reflect"
	"testing"

	"leetcode/uti"
)

// Time: O(n)
// Space: O(1)
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	uti.ReverseSlice(nums, 0, n-1)
	uti.ReverseSlice(nums, 0, k-1)
	uti.ReverseSlice(nums, k, n-1)
}

// Time: O(n)
// Space: O(n)
// func rotate(nums []int, k int) {
// 	n := len(nums)
// 	if k == 0 || n <= 1 {
// 		return
// 	}

// 	k = k % n

// 	left, right := uti.DeepCopySlice(nums[n-k:]), uti.DeepCopySlice(nums[:n-k])
// 	for i := 0; i < len(nums); i++ {
// 		if i < k {
// 			nums[i] = left[i]
// 			continue
// 		}
// 		nums[i] = right[i-k]
// 	}
// }

// Time: O(n^2)
// Space: O(1)
// func rotate(nums []int, k int) {
// 	n := len(nums)
// 	k %= n
// 	temp := 0
// 	for k > 0 {
// 		for i := 0; i < n; i++ {
// 			if i == 0 {
// 				temp = nums[0]
// 				nums[0] = nums[n-1]
// 				continue
// 			}
// 			if i == n-1 {
// 				nums[i] = temp
// 				continue
// 			}
// 			nums[i], temp = temp, nums[i]
// 		}
// 		k--
// 	}
// }

func TestRotate(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		k      int
		output []int
	}{
		{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, output: []int{5, 6, 7, 1, 2, 3, 4}},
		{nums: []int{-1, -100, 3, 99}, k: 2, output: []int{3, 99, -1, -100}},
		{nums: []int{1, 2}, k: 3, output: []int{2, 1}},
	} {
		rotate(test.nums, test.k)
		if !reflect.DeepEqual(test.nums, test.output) {
			t.Errorf("expect %v but got %v", test.output, test.nums)
		}
	}
}
