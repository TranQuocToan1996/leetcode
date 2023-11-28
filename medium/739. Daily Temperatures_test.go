package medium

import (
	"reflect"
	"testing"
)

// Time: O(n)
// Space: O(n)
func dailyTemperatures(temperatures []int) []int {
	var (
		n     = len(temperatures)
		res   = make([]int, n)
		stack = make([]int, 0, n)
	)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			popIndex := stack[len(stack)-1]
			res[popIndex] = i - popIndex
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

func TestDailyTemperatures(t *testing.T) {
	for _, test := range []struct {
		temperatures []int
		output       []int
	}{
		{temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73}, output: []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{temperatures: []int{30, 40, 50, 60}, output: []int{1, 1, 1, 0}},
		{temperatures: []int{30, 60, 90}, output: []int{1, 1, 0}},
	} {
		res := dailyTemperatures(test.temperatures)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
