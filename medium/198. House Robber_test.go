package medium

import "testing"

func rob(nums []int) int {
	return 0
}

func TestRob(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{1, 2, 3, 1}, output: 4},
		{nums: []int{2, 7, 9, 3, 1}, output: 12},
	} {
		res := rob(test.nums)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
