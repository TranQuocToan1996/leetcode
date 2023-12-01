package medium

import (
	"sort"
	"testing"

	"leetcode/uti"
)

// Your algorithm's time complexity must be better than O(n log n)
func topKFrequent(nums []int, k int) []int {
	sort.IntSlice(nums).Sort()
	

	return nil
}

func TestTopKFrequent(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		k      int
		output []int
	}{
		{nums: []int{1, 1, 1, 2, 2, 3}, k: 2, output: []int{1, 2}},
		{nums: []int{1}, k: 1, output: []int{1}},
	} {
		res := topKFrequent(test.nums, test.k)
		if !uti.UnorderMatch(res, test.output) {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
