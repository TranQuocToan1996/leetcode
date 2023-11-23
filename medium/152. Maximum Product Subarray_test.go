package medium

import (
	"testing"

	"leetcode/uti"
)

// Time: O(n)
// Space: O(1)
func maxProduct(nums []int) int {
	max, imax, imin := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		imin = uti.Min(nums[i], nums[i]*imin)
		imax = uti.Max(nums[i], nums[i]*imax)
		max = uti.Max(imax, max)
	}
	return max
}

func TestMaxProduct(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{2, 3, -2, 4}, output: 6},
		{nums: []int{-2, 0, -1}, output: 0},
		{nums: []int{2, -2, -3}, output: 12},
	} {
		res := maxProduct(test.nums)
		if res != test.output {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
