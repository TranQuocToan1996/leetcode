package medium

import (
	"sort"
	"testing"

	"leetcode/uti"
)

// Time: O(nlogn)
// Space: O(logn)
func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] > points[j][0]
	})
	res := 0
	for i := 1; i < len(points); i++ {
		width := uti.AbsMinus(points[i][0], points[i-1][0])
		if res < width {
			res = width
		}
	}
	return res
}

func TestMaxWidthOfVerticalArea(t *testing.T) {
	for _, test := range []struct {
		points [][]int
		output int
	}{
		{points: [][]int{
			{8, 7},
			{9, 9},
			{7, 4},
			{9, 7},
		}, output: 1},
	} {
		res := maxWidthOfVerticalArea(test.points)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
