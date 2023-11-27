package medium

import (
	"fmt"
	"testing"
)

type orangesRottingAddr struct {
	i, j int
}

// Time: O(n*m)
// Space: O(n*m)
func orangesRotting(grid [][]int) int {
	const (
		empty  = 0
		fresh  = 1
		rotten = 2
	)
	var (
		n, m     = len(grid), len(grid[0])
		noRotten = 0
		freshPos = map[string]orangesRottingAddr{}
	)

	isAloneFresh := func(i, j int) bool {
		if i-1 >= 0 && grid[i-1][j] != empty {
			return false
		}
		if i+1 < n && grid[i+1][j] != empty {
			return false
		}
		if j-1 >= 0 && grid[i][j-1] != empty {
			return false
		}
		if j+1 < m && grid[i][j+1] != empty {
			return false
		}
		return true
	}

	isNearRotten := func(i, j int) bool {
		if i-1 >= 0 && grid[i-1][j] == rotten {
			return true
		}
		if i+1 < n && grid[i+1][j] == rotten {
			return true
		}
		if j-1 >= 0 && grid[i][j-1] == rotten {
			return true
		}
		if j+1 < m && grid[i][j+1] == rotten {
			return true
		}
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == fresh {
				if isAloneFresh(i, j) {
					return -1
				}
				freshPos[fmt.Sprintf("%v_%v", i, j)] = orangesRottingAddr{i, j}
			}
			if grid[i][j] == rotten {
				noRotten++
			}
		}
	}
	if noRotten == 0 {
		if len(freshPos) == 0 {
			return 0
		}
		return -1
	}
	res := 0
	for len(freshPos) > 0 {
		beginLen := len(freshPos)
		var willRotten []orangesRottingAddr
		for key, addr := range freshPos {
			if isNearRotten(addr.i, addr.j) {
				delete(freshPos, key)
				willRotten = append(willRotten, addr)
			}
		}
		for _, addr := range willRotten {
			grid[addr.i][addr.j] = rotten
		}
		res++
		if beginLen == len(freshPos) {
			return -1
		}
	}

	return res
}

func TestOrangesRotting(t *testing.T) {
	for _, test := range []struct {
		grid   [][]int
		output int
	}{
		{grid: [][]int{
			{2, 1, 1},
			{1, 1, 0},
			{0, 1, 1},
		}, output: 4},
		{grid: [][]int{
			{2, 1, 1},
			{0, 1, 1},
			{1, 0, 1},
		}, output: -1},
		{grid: [][]int{
			{0, 2},
		}, output: 0},
		{grid: [][]int{
			{0, 1},
		}, output: -1},
		{grid: [][]int{
			{1},
			{1},
			{1},
			{1},
		}, output: -1},
		{grid: [][]int{
			{0},
		}, output: 0},
		{grid: [][]int{
			{2},
			{2},
			{1},
			{0},
			{1},
			{1},
		}, output: -1},
	} {
		res := orangesRotting(test.grid)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
