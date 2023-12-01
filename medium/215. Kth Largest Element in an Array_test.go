package medium

import (
	"math/rand"
	"testing"
)

// Quick select:
// Time: Average O(n). Worst O(n^2)
// Space: O(1)
func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for {
		pivotIndex := left + rand.Intn(right-left+1)
		newPivotIndex := partitionFindKthLargest(nums, left, right, pivotIndex)
		if newPivotIndex == len(nums)-k {
			return nums[newPivotIndex]
		} else if newPivotIndex > len(nums)-k {
			right = newPivotIndex - 1
		} else {
			left = newPivotIndex + 1
		}
	}
}

func partitionFindKthLargest(nums []int, left int, right int, pivotIndex int) int {
	pivot := nums[pivotIndex]
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	storedIndex := left
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[i], nums[storedIndex] = nums[storedIndex], nums[i]
			storedIndex++
		}
	}
	nums[right], nums[storedIndex] = nums[storedIndex], nums[right]
	return storedIndex
}

// Method: MinHeap
// Time: O(nlogk)
// Space: O(k)
// func findKthLargest(nums []int, k int) int {
// 	h := model.IntMinHeap{}
// 	for _, num := range nums {
// 		heap.Push(&h, num)
// 		for h.Len() > k {
// 			heap.Pop(&h)
// 		}
// 	}
// 	return h[0]
// }

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
