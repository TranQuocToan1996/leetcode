package medium

import (
	"container/heap"
	"testing"

	"leetcode/model"
)

func findKthLargest(nums []int, k int) int {
	h := model.IntMinHeap{}
	for _, num := range nums {
		heap.Push(&h, num)
		for h.Len() > k {
			heap.Pop(&h)
		}
	}
	return h[0]
}

func TestFindKthLargest(t *testing.T) {
nextTest:
	for _, test := range []struct {
		nums      []int
		k         int
		anyOutput []int
	}{
		{nums: []int{3, 2, 1, 5, 6, 4}, k: 2, anyOutput: []int{5}},
		{nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4, anyOutput: []int{4}},
	} {
		res := findKthLargest(test.nums, test.k)
		for _, expect := range test.anyOutput {
			if res == expect {
				continue nextTest
			}
		}
		t.Errorf("expect %v but got %v, test %v", test.anyOutput, res, test)
	}
}
