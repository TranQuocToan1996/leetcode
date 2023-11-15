package medium

import (
	"reflect"
	"testing"
)

func nextPermutation(nums []int) {
	j := len(nums) - 1
	for ; j > 0 && nums[j-1] >= nums[j]; j-- {
	}
	if j != 0 {
		l := len(nums) - 1
		for ; l > j-1 && nums[j-1] >= nums[l]; l-- {
		}
		nums[j-1], nums[l] = nums[l], nums[j-1]
	}
	for k := len(nums) - 1; j < k; j, k = j+1, k-1 {
		nums[j], nums[k] = nums[k], nums[j]
	}
}

func TestNextPermutation(t *testing.T) {
	for _, test := range []struct {
		input  []int
		expect []int
	}{
		{input: []int{1, 2, 3}, expect: []int{1, 3, 2}},
		{input: []int{3, 2, 1}, expect: []int{1, 2, 3}},
		{input: []int{1, 1, 5}, expect: []int{1, 5, 1}},
	} {
		nextPermutation(test.input)
		if !reflect.DeepEqual(test.input, test.expect) {
			t.Errorf("expect %v but got %v", test.expect, test.input)
		}
	}
}
