package medium

import (
	"reflect"
	"testing"
)

// Time: O(2*n*m)
// Space: O(max(n,m)). Need update grid input
func onesMinusZeros(grid [][]int) [][]int {
	var (
		n, m    = len(grid), len(grid[0])
		onesRow = make([]int, n)
		onesCol = make([]int, m)
	)

	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if grid[x][y] == 1 {
				onesRow[x]++
				onesCol[y]++
			}
		}
	}

	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			grid[x][y] = 2*onesRow[x] + 2*onesCol[y] - n - m
		}
	}

	return grid
}

func TestOnesMinusZeros(t *testing.T) {
	for _, test := range []struct {
		grid   [][]int
		output [][]int
	}{
		{
			grid: [][]int{
				{0, 1, 1},
				{1, 0, 1},
				{0, 0, 1},
			},
			output: [][]int{
				{0, 0, 4},
				{0, 0, 4},
				{-2, -2, 2},
			},
		},
	} {
		res := onesMinusZeros(test.grid)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
