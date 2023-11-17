package medium

import (
	"testing"
)

// Time: O(m*n*4^len(word))
// Space: O(len(word)) call stack
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	var dfs func(x, y int, word string) bool
	dfs = func(x, y int, word string) bool {
		if len(word) == 0 {
			return true
		}
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[0] {
			return false
		}

		temp := board[x][y]
		const checked byte = '#'
		board[x][y] = checked
		next := word[1:]
		found := dfs(x+1, y, next) || dfs(x-1, y, next) ||
			dfs(x, y+1, next) || dfs(x, y-1, next)
		board[x][y] = temp
		return found
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, word) {
				return true
			}
		}
	}

	return false
}

func TestWordSearch(t *testing.T) {
	for _, test := range []struct {
		board  [][]byte
		word   string
		output bool
	}{
		{board: [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, word: "ABCCED", output: true},
		{board: [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, word: "SEE", output: true},
		{board: [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, word: "ABCB", output: false},
	} {
		res := exist(test.board, test.word)
		if res != test.output {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
