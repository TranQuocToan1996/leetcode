package medium

import (
	"reflect"
	"sort"
	"testing"

	"leetcode/uti"
)

// Time: O(nlogn + n) = O(nlogn)
// Space: n (output)
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0, len(intervals))
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if isOverLap(res[len(res)-1], intervals[i]) {
			mergeInterval(res[len(res)-1], intervals[i])
			continue
		}
		res = append(res, intervals[i])
		continue
	}

	return res
}

func isOverLap(left, right []int) bool {
	return left[1] >= right[0]
}

func mergeInterval(left, right []int) {
	if isOverLap(left, right) {
		left[0] = uti.Min(left[0], right[0])
		left[1] = uti.Max(left[1], right[1])
	}
}

func TestMergeInterval(t *testing.T) {
	for _, test := range []struct {
		input  [][]int
		output [][]int
	}{
		{
			input: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			output: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			input: [][]int{
				{1, 4},
				{4, 5},
			},
			output: [][]int{
				{1, 5},
			},
		},
	} {
		res := merge(test.input)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test.input)
		}
	}
}
