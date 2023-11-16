package medium

import (
	"reflect"
	"testing"
)

// Time: O(m*n)
// Space: O(1)
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])            // O(1) space
	isFirstRowZero, isFirstColZero := false, false // O(1) space

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					isFirstRowZero = true
				} else {
					matrix[0][j] = 0
				}
				if j == 0 {
					isFirstColZero = true
				} else {
					matrix[i][0] = 0
				}
			}
		}
	}

	// O(m*n)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// O(n)
	if isFirstRowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	// O(m)
	if isFirstColZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

func TestSetZeroes(t *testing.T) {
	for _, test := range []struct {
		matrix [][]int
		expect [][]int
	}{
		{matrix: [][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}, expect: [][]int{
			{1, 0, 1},
			{0, 0, 0},
			{1, 0, 1},
		}},
		{matrix: [][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		}, expect: [][]int{
			{0, 0, 0, 0},
			{0, 4, 5, 0},
			{0, 3, 1, 0},
		}},
	} {
		setZeroes(test.matrix)
		if !reflect.DeepEqual(test.matrix, test.expect) {
			t.Errorf("expect %v but got %v", test.expect, test.matrix)
		}
	}
}
