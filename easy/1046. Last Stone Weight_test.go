package easy

import (
	"fmt"
	"sort"
)

func main() {

	tests := [][]int{
		{2, 7, 4, 1, 8, 1},
		{1},
	}
	for _, test := range tests {
		fmt.Println(lastStoneWeight(test))
	}
}

// https://leetcode.com/problems/make-two-arrays-equal-by-reversing-sub-arrays/

/*
time complexity: O(),  ms
Space complexity: O(),  MB
*/
func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	for len(stones) > 1 {
		stones[len(stones)-1] -= stones[len(stones)-2]
		stones[len(stones)-2] = 0
		sort.Ints(stones)
		for stones[0] == 0 && len(stones) > 1 {
			stones = stones[1:]
		}
	}
	return stones[0]
}

/*
// Using heap
func lastStoneWeight(stones []int) int {
	pq := IntHeap(stones)
	heap.Init(&pq)
	for pq.Len() > 1 {
		heap.Push(&pq, heap.Pop(&pq).(int)-heap.Pop(&pq).(int))

	}
	return heap.Pop(&pq).(int)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
} */
