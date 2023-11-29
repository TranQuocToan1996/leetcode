package easy

import "testing"

// Time: O(1)
// Space: O(1)
func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		if num%2 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}

func TestMaxProduct(t *testing.T) {
	for _, test := range []struct {
		nums   uint32
		output int
	}{
		{nums: 0b0000000000000000000000000001011, output: 3},
		{nums: 0b0000000000000000000000010000000, output: 1},
		{nums: 0b11111111111111111111111111111101, output: 31},
	} {
		res := hammingWeight(test.nums)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
