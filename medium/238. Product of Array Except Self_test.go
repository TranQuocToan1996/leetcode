package medium

import (
	"testing"

	"leetcode/uti"
)

// You must write an algorithm that runs in O(n) time and without using the division operation.
// Time: O(n)
// Space: O(n)
// func productExceptSelf(nums []int) []int {
// 	var (
// 		n            = len(nums)
// 		leftProduct  = make([]int, n)
// 		rightProduct = make([]int, n)
// 	)
// 	leftProduct[0] = nums[0]
// 	rightProduct[n-1] = nums[n-1]

// 	for i := 1; i < n; i++ {
// 		leftProduct[i] = leftProduct[i-1] * nums[i]
// 		rightProduct[n-1-i] = rightProduct[n-i] * nums[n-1-i]
// 	}
// 	nums[0], nums[n-1] = rightProduct[1], leftProduct[n-2]
// 	for i := 1; i < n-1; i++ {
// 		nums[i] = leftProduct[i-1] * rightProduct[i+1]
// 	}

// 	return nums
// }

// Time: O(n)
// Space: O(1)
// func productExceptSelf(nums []int) []int {
// 	product := 1
// 	numZero := 0
// 	for i := 0; i < len(nums); i++ {
// 		if numZero == 2 {
// 			break
// 		}
// 		if nums[i] == 0 {
// 			numZero++
// 			continue
// 		}
// 		product *= nums[i]
// 	}
// 	if numZero >= 2 {
// 		for i := range nums {
// 			nums[i] = 0
// 		}
// 	} else if numZero == 1 {
// 		for i := range nums {
// 			if nums[i] == 0 {
// 				nums[i] = product
// 			} else {
// 				nums[i] = 0
// 			}
// 		}
// 	} else {
// 		for i := range nums {
// 			nums[i] = product / nums[i]
// 		}
// 	}

// 	return nums
// }

// Time: O(n)
// Space: O(1)
func productExceptSelf(nums []int) []int {
	var (
		n   = len(nums)
		res = make([]int, n)
		pre = 1
		suf = 1
	)
	for i := range res {
		res[i] = 1
	}
	for i := range res {
		res[i] *= pre
		pre *= nums[i]
		res[n-1-i] *= suf
		suf *= nums[n-1-i]
	}
	return res
}

func TestProductExceptSelf(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output []int
	}{
		{nums: []int{1, 2, 3, 4}, output: []int{24, 12, 8, 6}},
		{nums: []int{-1, 1, 0, -3, 3}, output: []int{0, 0, 9, 0, 0}},
		{nums: []int{4, 3, 2, 1, 2}, output: []int{12, 16, 24, 48, 24}},
	} {
		res := productExceptSelf(test.nums)
		if !uti.UnorderMatch(test.output, res) {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
