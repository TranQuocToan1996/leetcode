package medium

import (
	"math/rand"
	"testing"

	"leetcode/uti"
)

type topKFrequentHeapData struct {
	val, frequency int
}
type topKFrequentHeap []topKFrequentHeapData

func (h topKFrequentHeap) Len() int {
	return len(h)
}

func (h topKFrequentHeap) Less(i int, j int) bool {
	return h[i].frequency < h[j].frequency
}

func (h topKFrequentHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *topKFrequentHeap) Push(x any) {
	*h = append(*h, x.(topKFrequentHeapData))
}

func (h *topKFrequentHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Time: O(nlogk)
// Space: O(n)
// Your algorithm's time complexity must be better than O(n log n)
// func topKFrequent(nums []int, k int) []int {
// 	n := len(nums)
// 	if k > n {
// 		return nums
// 	}
// 	seen := map[int]int{}
// 	for _, num := range nums {
// 		seen[num]++
// 	}
// 	topKFrequentHeap := topKFrequentHeap{}
// 	heap.Init(&topKFrequentHeap)
// 	for num, freq := range seen {
// 		heap.Push(&topKFrequentHeap, topKFrequentHeapData{
// 			val:       num,
// 			frequency: freq,
// 		})
// 		for topKFrequentHeap.Len() > k {
// 			heap.Pop(&topKFrequentHeap)
// 		}
// 	}
// 	res := make([]int, 0, k)
// 	for topKFrequentHeap.Len() > 0 {
// 		data := topKFrequentHeap.Pop().(topKFrequentHeapData)
// 		res = append(res, data.val)
// 	}
// 	return res
// }

// Time: Average O(n), Worst O()
// Space: O(n)
func topKFrequent(nums []int, k int) []int {
	seen := map[int]int{}
	for _, num := range nums {
		seen[num]++
	}
	unique := make([]int, 0, len(nums))
	for num := range seen {
		unique = append(unique, num)
	}
	
	partition := func(left, right, pivotIndex int) int {
		freq := seen[unique[pivotIndex]]
		unique[pivotIndex], unique[right] = unique[right], unique[pivotIndex]
		sortFreqIndex := left
		for i := left; i < right; i++ {
			if seen[unique[i]] < freq {
				unique[sortFreqIndex], unique[i] = unique[i], unique[sortFreqIndex]
				sortFreqIndex++
			}
		}
		unique[sortFreqIndex], unique[right] = unique[right], unique[sortFreqIndex]
		return sortFreqIndex
	}

	var quickSelect func(left, right, kSmallest int)
	quickSelect = func(left, right, kSmallest int) {
		if left >= right {
			return
		}

		pivotIndex := left + rand.Intn(right-left+1)
		sortFreqIndex := partition(left, right, pivotIndex)
		if sortFreqIndex < kSmallest {
			quickSelect(sortFreqIndex+1, right, kSmallest)
		} else if sortFreqIndex > kSmallest {
			quickSelect(left, sortFreqIndex-1, kSmallest)
		}
	}
	quickSelect(0, len(unique)-1, len(unique)-k)
	return unique[len(unique)-k:]
}

func TestTopKFrequent(t *testing.T) {
nextTest:
	for _, test := range []struct {
		nums      []int
		k         int
		anyOutput [][]int
	}{
		{nums: []int{1, 1, 1, 2, 2, 3}, k: 2, anyOutput: [][]int{{1, 2}}},
		{nums: []int{1, 1, 2, 2, 3, 3}, k: 2, anyOutput: [][]int{{1, 2}, {1, 3}, {2, 3}}},
		{nums: []int{1}, k: 1, anyOutput: [][]int{{1}}},
		{nums: []int{1, 1, 1, 2, 2, 3}, k: 1, anyOutput: [][]int{{1}}},
		{nums: []int{3, 0, 1, 0}, k: 1, anyOutput: [][]int{{0}}},
	} {
		res := topKFrequent(test.nums, test.k)
		for _, expect := range test.anyOutput {
			if uti.UnorderMatch(res, expect) {
				continue nextTest
			}
		}
		t.Errorf("expect %v but got %v, test %v", test.anyOutput, res, test)
	}
}
