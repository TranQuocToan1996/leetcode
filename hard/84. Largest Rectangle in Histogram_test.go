package hard

import "testing"

type largestRectangleAreaData struct {
	index, height int
}

// Monotonic increase (meet smaller, pop and append keep index)
// Time: O(n)
// Space: O(n)
func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := []largestRectangleAreaData{{index: 0, height: heights[0]}}
	max := heights[0]
	for i := 1; i < n; i++ {
		lastIndex := i
		for len(stack) > 0 && stack[len(stack)-1].height > heights[i] {
			pop := stack[len(stack)-1]
			lastIndex = pop.index
			area := pop.height * (i - pop.index)
			if max < area {
				max = area
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, largestRectangleAreaData{index: lastIndex, height: heights[i]})
	}

	// Calculate area of remain element in stack
	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		area := pop.height * (n - pop.index)
		if max < area {
			max = area
		}
		stack = stack[:len(stack)-1]
	}
	return max
}

func TestLargestRectangleArea(t *testing.T) {
	for _, test := range []struct {
		heights []int
		output  int
	}{
		{heights: []int{2, 1, 5, 6, 2, 3}, output: 10},
		{heights: []int{2, 4}, output: 4},
	} {
		res := largestRectangleArea(test.heights)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
