package easy

import "testing"

func findSpecialInteger(arr []int) int {
	n := len(arr)
	freqTarget := n / 4 // 25%
	count := map[int]int{}
	for _, num := range arr {
		count[num]++
		if count[num] > freqTarget {
			return num
		}
	}
	return arr[0]
}

func TestFindSpecialInteger(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{1, 2, 2, 6, 6, 6, 6, 7, 10}, output: 6},
		{nums: []int{1, 1}, output: 1},
	} {
		res := findSpecialInteger(test.nums)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
