package medium

import (
	"testing"
)

// rename searchMatrix when submit to leetcode
// Time: O(m+n)
// Time: O(1).
func searchMatrix2(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	row, col := 0, m-1
	for row < n && col >= 0 {
		found := matrix[row][col]
		if found == target {
			return true
		} else if found > target {
			col--
		} else {
			row++
		}
	}
	return false
}

func TestSearchMatrix2(t *testing.T) {
	for _, test := range []struct {
		matrix [][]int
		target int
		expect bool
	}{
		{
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 5,
			expect: true,
		},
		{
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 20,
			expect: false,
		},
		{
			matrix: [][]int{
				{-1, 3},
			},
			target: 3,
			expect: true,
		},
	} {
		res := searchMatrix2(test.matrix, test.target)
		if res != test.expect {
			t.Errorf("expect %v but got %v, target %v", test.expect, test.matrix, test.target)
		}
	}
}
