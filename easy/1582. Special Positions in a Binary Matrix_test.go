package easy

import (
	"testing"
)

func numSpecial(mat [][]int) int {
	n, m := len(mat), len(mat[0])
	sumX := make([]int, n)
	sumY := make([]int, m)
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			cur := mat[x][y]
			sumX[x] += cur
			sumY[y] += cur
		}
	}
	res := 0
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if mat[x][y] == 1 && sumX[x] == 1 && sumY[y] == 1 {
				res++
			}
		}
	}
	return res
}

func TestNumSpecial(t *testing.T) {
	for _, test := range []struct {
		mat    [][]int
		output int
	}{
		{mat: [][]int{
			{1, 0, 0},
			{0, 0, 1},
			{1, 0, 0},
		}, output: 1},
		{mat: [][]int{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		}, output: 3},
		{mat: [][]int{
			{0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 1},
		}, output: 3},
	} {
		res := numSpecial(test.mat)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
