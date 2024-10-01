package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyQueue struct {
	data []int
}

func ConstructorMyQueue() MyQueue {
	return MyQueue{
		data: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyQueue) Pop() int {
	if !this.Empty() {
		pop := this.Peek()
		this.data = this.data[1:]
		return pop
	}
	return 0
}

func (this *MyQueue) Peek() int {
	if !this.Empty() {
		return this.data[0]
	}
	return 0
}

func (this *MyQueue) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func TestMyQueue(t *testing.T) {
	queue := ConstructorMyQueue()
	queue.Push(1)
	queue.Push(2)
	assert.True(t, queue.Peek() == 1)
	assert.True(t, queue.Pop() == 1)
	assert.True(t, !queue.Empty())
}
