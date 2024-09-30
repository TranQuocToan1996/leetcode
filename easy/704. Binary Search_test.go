package easy

import "testing"

func TestBinarySearch(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		target int
		output int
	}{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			output: 4,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: -2,
			output: -1,
		},
		{
			nums:   []int{2, 5},
			target: 5,
			output: 1,
		},
	} {
		res := search(test.nums, test.target)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
