package hard

import (
	"container/heap"
	"testing"

	"leetcode/model"

	"github.com/stretchr/testify/require"
)

type MedianFinder struct {
	big   model.IntHeap
	small model.IntMaxHeap
}

// Space: O(n)
func Constructor() MedianFinder {
	return MedianFinder{
		big:   model.IntHeap{},
		small: model.IntMaxHeap{},
	}
}

// Time: O(logn)
func (this *MedianFinder) AddNum(num int) {
	if this.small.Len() < this.big.Len() {
		heap.Push(&this.big, num)
		heap.Push(&this.small, heap.Pop(&this.big))
		return
	}

	heap.Push(&this.small, num)
	heap.Push(&this.big, heap.Pop(&this.small))
}

// Time: O(1)
func (this *MedianFinder) FindMedian() float64 {
	if this.small.Len() > this.big.Len() {
		return float64(this.small.IntHeap[0])
	}
	if this.small.Len() < this.big.Len() {
		return float64(this.big[0])
	}
	return (float64(this.small.IntHeap[0]) + float64(this.big[0])) / 2
}

func TestMedianFinder(t *testing.T) {
	medianFinder := Constructor()
	medianFinder.AddNum(1)
	medianFinder.AddNum(2)
	require.Equal(t, 1.5, medianFinder.FindMedian())
	medianFinder.AddNum(3)
	require.Equal(t, 2.0, medianFinder.FindMedian())
}
