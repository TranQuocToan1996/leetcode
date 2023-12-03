package easy

import (
	"testing"

	"leetcode/uti"
)

func minTimeToVisitAllPoints(points [][]int) int {
	n := len(points)
	res := 0
	for i := 1; i < n; i++ {
		res += calculateTimeMove(points[i-1], points[i])
	}

	return res
}

func calculateTimeMove(start, end []int) int {
	xMax := start[0] - end[0]
	if xMax < 0 {
		xMax = -xMax
	}
	yMax := start[1] - end[1]
	if yMax < 0 {
		yMax = -yMax
	}
	return uti.Max(xMax, yMax)
}

func TestMinTimeToVisitAllPoints(t *testing.T) {
	for _, test := range []struct {
		points [][]int
		output int
	}{
		{points: [][]int{
			{1, 1},
			{3, 4},
			{-1, 0},
		}, output: 7},
		{points: [][]int{
			{3, 2},
			{-2, 2},
		}, output: 5},
	} {
		res := minTimeToVisitAllPoints(test.points)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
