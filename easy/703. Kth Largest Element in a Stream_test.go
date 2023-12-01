package easy

import (
	"container/heap"
	"testing"

	"leetcode/model"
)

func TestXxx(t *testing.T) {
	// tests := []{

	// }
	// for _, test := range tests {
	// 	fmt.Println(canBeEqual(test.target, test.arr))
	// }
}

// https://leetcode.com/problems/make-two-arrays-equal-by-reversing-sub-arrays/

/*
time complexity: O(),  ms
Space complexity: O(),  MB
*/

// real thing starts here
type KthLargest struct {
	Nums *model.IntHeap
	K    int
}

func Constructor(k int, nums []int) KthLargest {
	h := &model.IntHeap{}
	heap.Init(h)
	// push all elements to the heap
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
	}
	// remove the smaller elements from the heap such that only the k largest elements are in the heap
	for len(*h) > k {
		heap.Pop(h)
	}
	return KthLargest{h, k}
}

func (this *KthLargest) Add(val int) int {
	if len(*this.Nums) < this.K {
		heap.Push(this.Nums, val)
	} else if val > (*this.Nums)[0] {
		heap.Push(this.Nums, val)
		heap.Pop(this.Nums)
	}
	return (*this.Nums)[0]
}
