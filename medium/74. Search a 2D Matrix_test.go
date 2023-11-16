package medium

import (
	"testing"

	"leetcode/uti"
)

// Time: O(log(m) * log(n))
// Space: O(1)
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	if target > matrix[m-1][n-1] {
		return false
	}
	row := 0
	top, bot := 0, m-1
	for top <= bot {
		mid := (bot + top) / 2
		maxRowVal := matrix[mid][n-1]
		minRowVal := matrix[mid][0]
		if maxRowVal >= target && minRowVal <= target {
			row = mid
			break
		}
		if maxRowVal < target {
			top = mid + 1
		} else {
			bot = mid - 1
		}
	}

	return uti.BinarySearch(matrix[row], target)
}

// Time: O(log(n*m))
// Space: O(1)
// func searchMatrix(matrix [][]int, target int) bool {
//     m := len(matrix)
//     n := len(matrix[0])
//     left, right := 0, m*n-1

//     for left <= right {
//         mid := left + (right-left)/2
//         mid_val := matrix[mid/n][mid%n]

//         if mid_val == target {
//             return true
//         } else if mid_val < target {
//             left = mid + 1
//         } else {
//             right = mid - 1
//         }
//     }
//     return false
// }

func TestSearchMatrix(t *testing.T) {
	for _, test := range []struct {
		matrix [][]int
		target int
		expect bool
	}{
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 3,
			expect: true,
		},
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 13,
			expect: false,
		},
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 34,
			expect: true,
		},
	} {
		res := searchMatrix(test.matrix, test.target)
		if res != test.expect {
			t.Errorf("expect %v but got %v", test.expect, test.matrix)
		}
	}
}
