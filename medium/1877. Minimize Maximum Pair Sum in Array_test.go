package medium

import (
	"sort"
	"testing"
)

// Time: O(nLogn)
// Space: O(logn) Base on sort algo. Go using pdqsort.
// O(logn) is came from Quick Sort call stack of pdqsort
func minPairSum(nums []int) int {
	n := len(nums)
	pairs := n / 2
	sort.IntSlice(nums).Sort()
	max := nums[0] + nums[n-1]
	for i := 1; i < pairs; i++ {
		sumPair := nums[i] + nums[n-i-1]
		if max < sumPair {
			max = sumPair
		}
	}
	return max
}

func TestMinPairSum(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{3, 5, 2, 3}, output: 7},
		{nums: []int{3, 5, 4, 2, 4, 6}, output: 8},
	} {
		res := minPairSum(test.nums)
		if res != test.output {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
