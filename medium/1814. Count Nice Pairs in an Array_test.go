package medium

import "testing"

func countNicePairs(nums []int) int {
	mem := make([][]int, len(nums)+1)
	for k := range mem {
		mem[k] = make([]int, len(nums)+1)
	}

	for i := 1; i < len(mem); i++ {
		mem[0][i] = nums[i-1]
		mem[i][0] = revInt(nums[i-1])
	}
	for i := 1; i < len(mem); i++ {
		for j := 1; j < len(mem[0]); j++ {
			mem[i][j] = mem[0][j] + mem[i][0]
		}
	}
	return 0
}

func revInt(num int) int {
	reversed := 0
	for num != 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}
	return reversed
}

func TestCountNicePairs(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{42, 11, 1, 97}, output: 3},
		{nums: []int{13, 10, 35, 24, 76}, output: 4},
	} {
		res := countNicePairs(test.nums)
		if res != test.output {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
