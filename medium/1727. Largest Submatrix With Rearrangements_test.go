package medium

import (
	"sort"
	"testing"

	"leetcode/uti"
)

func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] > 0 && row > 0 {
				matrix[row][col] += matrix[row-1][col]
			}
		}
	}
	res := 0
	for row := 0; row < m; row++ {
		clone := uti.DeepCopySlice(matrix[row])
		sort.Slice(clone, func(i, j int) bool {
			return clone[i] > clone[j]
		})
		for i := range clone {
			res = uti.Max(res, clone[i]*(i+1))
		}
	}
	return res
}

func TestLargestSubmatrix(t *testing.T) {
	for _, test := range []struct {
		matrix [][]int
		output int
	}{
		{matrix: [][]int{
			{0, 0, 1},
			{1, 1, 1},
			{1, 0, 1},
		}, output: 4},
		{matrix: [][]int{
			{1, 0, 1, 0, 1},
		}, output: 3},
		{matrix: [][]int{
			{1, 1, 0},
			{1, 0, 1},
		}, output: 2},
	} {
		res := largestSubmatrix(test.matrix)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
