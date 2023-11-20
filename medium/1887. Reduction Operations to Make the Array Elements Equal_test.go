package medium

import (
	"sort"
	"testing"
)

func reductionOperations(nums []int) int {
	sort.IntSlice(nums).Sort()
	res := 0
	cur := 0
	nextMax := nums[len(nums)-1]

	for i := len(nums) - 1; i >= 0; i-- {
		if nextMax > nums[i] {
			nextMax = nums[i]
			res += cur
		}
		cur++
	}

	return res
}

func TestReductionOperations(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{5, 1, 3}, output: 3},
		{nums: []int{1, 1, 1}, output: 0},
		{nums: []int{1, 1, 2, 2, 3}, output: 4},
	} {
		res := reductionOperations(test.nums)
		if res != test.output {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
