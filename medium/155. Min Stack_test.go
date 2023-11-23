package medium

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type MinStack struct {
	data   []int
	minVal int
}

// Need rename Constructor
func NewMinStack() MinStack {
	const cap = 1 << 10
	return MinStack{data: make([]int, 0, cap), minVal: math.MaxInt}
}

func (this *MinStack) Push(val int) {
	if this.minVal > val {
		this.minVal = val
	}
	this.data = append(this.data, val)
}

func (this *MinStack) Pop() {
	if len(this.data) > 0 {
		popVal := this.data[len(this.data)-1]
		this.data = this.data[:len(this.data)-1]
		if popVal == this.minVal {
			this.minVal = math.MaxInt
			for _, val := range this.data {
				if this.minVal > val {
					this.minVal = val
				}
			}
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.data) > 0 {
		return this.minVal
	}
	return 0
}

func TestNewMinStack(t *testing.T) {
	minStack := NewMinStack()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	require.Equal(t, -3, minStack.GetMin()) // return -3
	minStack.Pop()
	require.Equal(t, minStack.Top(), 0)     // return 0
	require.Equal(t, -2, minStack.GetMin()) // return -2
}
