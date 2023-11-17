package medium

import (
	"reflect"
	"testing"

	"leetcode/model"
)

func buildTree(preorder []int, inorder []int) *model.TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// preorder = [3,9,20,15,7] inorder = [9,3,15,20,7]
	// [3,9,20,null,null,15,7]
	mid := indexOf(inorder, preorder[0])
	root := &model.TreeNode{Val: inorder[mid]}
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return root
}

func indexOf(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

func TestBuildTree(t *testing.T) {
	for _, test := range []struct {
		preorder []int
		inorder  []int
		output   *model.TreeNode
	}{
		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			output: &model.TreeNode{
				Val: 3,
				Left: &model.TreeNode{
					Val: 9,
				},
				Right: &model.TreeNode{
					Val: 20,
					Left: &model.TreeNode{
						Val: 15,
					},
					Right: &model.TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			preorder: []int{-1},
			inorder:  []int{-1},
			output: &model.TreeNode{
				Val: -1,
			},
		},
	} {
		res := buildTree(test.preorder, test.inorder)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %+v", test.output, res, test)
		}
	}
}
