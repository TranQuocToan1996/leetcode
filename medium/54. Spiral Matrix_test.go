package medium

import (
	"reflect"
	"testing"
)

// Time: O(m*n)
// Space: O(m*n)
func spiralOrder(matrix [][]int) []int {
	var (
		left   = 0
		right  = len(matrix[0]) - 1
		top    = 0
		bot    = len(matrix) - 1
		result = make([]int, 0, len(matrix)*len(matrix[0]))
	)

	done := func() bool {
		return !(left <= right && top <= bot)
	}

	for {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		if done() {
			return result
		}

		for i := top; i <= bot; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if done() {
			return result
		}

		for i := right; i >= left; i-- {
			result = append(result, matrix[bot][i])
		}
		bot--
		if done() {
			return result
		}

		for i := bot; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		if done() {
			return result
		}
	}
}

func TestSpiralOrder(t *testing.T) {
	for _, test := range []struct {
		input  [][]int
		output []int
	}{
		{input: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}, output: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{input: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}, output: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	} {
		res := spiralOrder(test.input)
		if !reflect.DeepEqual(test.output, res) {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
