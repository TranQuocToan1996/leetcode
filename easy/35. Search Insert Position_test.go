package easy

import "testing"

// https://leetcode.com/problems/search-insert-position/
func TestSearch(t *testing.T) {
	for _, test := range []struct {
		nums     []int
		target   int
		expected int
		name     string
	}{
		{nums: []int{1, 3, 5, 6}, target: 5, expected: 2, name: "found"},
		{nums: []int{1, 3, 5, 6}, target: 2, expected: 1, name: "notfound inside"},
		{nums: []int{1, 3, 5, 6}, target: 7, expected: 4, name: "notfound outside"},
	} {
		if test.expected != searchInsert(test.nums, test.target) {
			t.Errorf("Expect %v, result %v at [%v]", test.expected, searchInsert(test.nums, test.target), test.name)
		}
	}
}

// Runtime: 5 ms, faster than 74.20% of Go online submissions for Search Insert Position.
// Memory Usage: 2.9 MB, less than 100.00% of Go online submissions for Search Insert Position.

// Time: O(log n), binary search
// Space: O(1) no allocation during processing

// You must write an algorithm with O(log n) runtime complexity.
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
