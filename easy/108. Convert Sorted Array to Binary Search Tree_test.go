package easy

import (
	"reflect"
	"testing"

	"leetcode/model"
)

func sortedArrayToBST(nums []int) *model.TreeNode {
	return nil
}

func TestSortedArrayToBST(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{2, 3, -2, 4}, output: 6},
		{nums: []int{-2, 0, -1}, output: 0},
		{nums: []int{2, -2, -3}, output: 12},
	} {
		res := sortedArrayToBST(test.nums)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
