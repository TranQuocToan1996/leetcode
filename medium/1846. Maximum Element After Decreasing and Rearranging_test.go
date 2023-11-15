package medium

import (
	"sort"
	"testing"
)

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	max := 1
	if len(arr) == 1 {
		return max
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max++
		}
	}
	return max
}

func TestMaximumElementAfterDecrementingAndRearranging(t *testing.T) {
	for _, test := range []struct {
		input  []int
		output int
	}{
		{input: []int{2, 2, 1, 2, 1}, output: 2},
		{input: []int{100, 1, 1000}, output: 3},
		{input: []int{1, 2, 3, 4, 5}, output: 5},
	} {
		res := maximumElementAfterDecrementingAndRearranging(test.input)
		if res != test.output {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
