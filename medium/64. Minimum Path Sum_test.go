package medium

import (
	"testing"

	"leetcode/uti"
)

// time O(rows*cols)
// space O(1): Calculate in the input matrix
func minPathSum(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	for i := 1; i < rows; i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	for i := 1; i < cols; i++ {
		grid[0][i] = grid[0][i-1] + grid[0][i]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			grid[i][j] = grid[i][j] + uti.Min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[rows-1][cols-1]
}

func TestMinPathSum(t *testing.T) {
	for _, test := range []struct {
		input  [][]int
		output int
	}{
		{input: [][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}, output: 7},
		{input: [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}, output: 12},
	} {
		res := minPathSum(test.input)
		if res != test.output {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
