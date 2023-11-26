package easy

import (
	"reflect"
	"testing"

	"leetcode/model"
)

func sortedArrayToBST(nums []int) *model.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &model.TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

func TestSortedArrayToBST(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output *model.TreeNode
	}{
		{nums: []int{-10, -3, 0, 5, 9}, output: &model.TreeNode{
			Val: 0,
			Left: &model.TreeNode{
				Val: -3,
				Left: &model.TreeNode{
					Val: -10,
				},
			},
			Right: &model.TreeNode{
				Val: 9,
				Left: &model.TreeNode{
					Val: 5,
				},
			},
		}},
	} {
		res := sortedArrayToBST(test.nums)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
