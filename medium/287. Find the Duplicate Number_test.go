package medium

import (
	"testing"
)

// Time: O(n)
// Space: O(n)
// func findDuplicate(nums []int) int {
// 	seen := map[int]bool{}
// 	for _, v := range nums {
// 		if seen[v] {
// 			return v
// 		}
// 		seen[v] = true
// 	}
// 	return 0
// }

/*
You must solve the problem without modifying the array nums and uses only constant extra space.
nums[i] in range [1, n] inclusive
There is only one repeated number in nums.
So we set the negative value if seen. */
// Time: O(n)
// Space: O(1)
// func findDuplicate(nums []int) int {
// 	for _, v := range nums {
// 		abs := int(math.Abs(float64(v)))
// 		if nums[abs] < 0 {
// 			return abs
// 		}
// 		nums[abs] = -nums[abs]
// 	}
// 	return 0
// }

// Imagine the nums[i] is the pointer to the index i.
// Now the array -> link list
// this problem become found the begin of cycle (floyd algorithm)
// Time: O(n)
// Space: O(1)
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// After slow == fast, now the distant slow -> beginCyle == beginLinkList -> beginCyle
	// Keep move same rate until they meet.
	for begin := 0; slow != begin; {
		slow = nums[slow]
		begin = nums[begin]
	}

	return slow
}

func TestFindDuplicate(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{1, 3, 4, 2, 2}, output: 2},
		{nums: []int{3, 1, 3, 4, 2}, output: 3},
	} {
		res := findDuplicate(test.nums)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
