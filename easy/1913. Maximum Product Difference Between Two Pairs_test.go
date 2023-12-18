package easy

import (
	"math"
	"testing"

	"leetcode/uti"
)

// Track 2-small, 2-big
// Space: O(1)
// Time: O(n)
func maxProductDifference(nums []int) int {
	biggest, secondBiggest := math.MinInt, math.MinInt
	smallest, secondsmallest := math.MaxInt, math.MaxInt
	for _, num := range nums {
		if biggest < num {
			secondBiggest = biggest
			biggest = num
		} else {
			secondBiggest = uti.Max(secondBiggest, num)
		}
		if smallest > num {
			secondsmallest = smallest
			smallest = num
		} else {
			secondsmallest = uti.Min(secondsmallest, num)
		}
	}
	return biggest*secondBiggest - smallest*secondsmallest
}

// Method: Sort array
// Space: logn
// Time: nLogn
// func maxProductDifference(nums []int) int {
// 	n := len(nums)
// 	sort.IntSlice(nums).Sort()
// 	return nums[n-1]*nums[n-2] - nums[0]*nums[1]
// }

// Space: 1
// Time: nLogn
// func maxProductDifference(nums []int) int {
// 	maxHeap := model.IntMaxHeap{}
// 	minHeap := model.IntHeap{}
// 	for i := range nums {
// 		heap.Push(&maxHeap, nums[i])
// 		if maxHeap.Len() > 2 {
// 			heap.Pop(&maxHeap)
// 		}
// 		heap.Push(&minHeap, nums[i])
// 		if minHeap.Len() > 2 {
// 			heap.Pop(&minHeap)
// 		}
// 	}
// 	return -maxHeap.IntHeap[0]*maxHeap.IntHeap[1] + minHeap[0]*minHeap[1]
// }

func TestMaxProductDifference(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{5, 6, 2, 7, 4}, output: 34},
		{nums: []int{4, 2, 5, 9, 7, 4, 8}, output: 64},
		{nums: []int{5, 9, 4, 6}, output: 34},
	} {
		res := maxProductDifference(test.nums)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
