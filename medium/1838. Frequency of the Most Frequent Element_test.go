package medium

import (
	"sort"
	"testing"
)

// Time: O(nlogn) for sort
// Space: O(logn) for sort
// Sliding Window
func maxFrequency(nums []int, k int) int {
	sort.IntSlice(nums).Sort()
	left := 0
	sum := 0
	for right := 0; right < len(nums); right++ {
		target := nums[right]
		sum += target
		if (right-left+1)*target-sum > k {
			sum -= nums[left]
			left++
		}
	}
	return len(nums) - left
}

func TestMaxFrequency(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		k      int
		output int
	}{
		{nums: []int{1, 2, 4}, k: 5, output: 3},
		{nums: []int{1, 4, 8, 13}, k: 5, output: 2},
		{nums: []int{3, 9, 6}, k: 2, output: 1},
	} {
		res := maxFrequency(test.nums, test.k)
		if res != test.output {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
