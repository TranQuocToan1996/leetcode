package medium

import (
	"testing"
)

// Time: O(n)
// Space: O(1)
func canJump(nums []int) bool {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if i > max {
			return false
		}
		if max < i+nums[i] {
			max = i + nums[i]
		}
	}
	return true
}

func TestCanJump(t *testing.T) {
	for _, test := range []struct {
		input  []int
		output bool
	}{
		{
			input:  []int{2, 3, 1, 1, 4},
			output: true,
		},
		{
			input:  []int{3, 2, 1, 0, 4},
			output: false,
		},
		{
			input:  []int{3, 2, 3, 0, 4},
			output: true,
		},
		{
			input:  []int{0},
			output: true,
		},
		{
			input:  []int{2, 0, 0},
			output: true,
		},
	} {
		res := canJump(test.input)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test.input)
		}
	}
}
