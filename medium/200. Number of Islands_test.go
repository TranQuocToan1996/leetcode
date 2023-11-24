package medium

import "testing"

// Time: O(m * n)
// Space: O(1)
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				markAdjacentIslands(grid, i, j)
			}
		}
	}
	return count
}

func markAdjacentIslands(grid [][]byte, i, j int) {
	const visited byte = '#'
	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = visited
	// Top
	if j-1 >= 0 {
		markAdjacentIslands(grid, i, j-1)
	}
	// Bot
	if j+1 < len(grid[0]) {
		markAdjacentIslands(grid, i, j+1)
	}
	// Left
	if i-1 >= 0 {
		markAdjacentIslands(grid, i-1, j)
	}
	// Right
	if i+1 < len(grid) {
		markAdjacentIslands(grid, i+1, j)
	}
}

func TestNumIslands(t *testing.T) {
	for _, test := range []struct {
		grid   [][]byte
		output int
	}{
		{grid: [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}, output: 1},
		{grid: [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}, output: 3},
	} {
		res := numIslands(test.grid)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
