package medium

import (
	"reflect"
	"testing"

	"leetcode/model"
)

// Time: O(n)
// Space: O(logn) call stack
func flatten(root *model.TreeNode) {
	var priv *model.TreeNode
	var helper func(*model.TreeNode)
	helper = func(node *model.TreeNode) {
		if node == nil {
			return
		}
		helper(node.Right)
		helper(node.Left)
		node.Right = priv
		node.Left = nil
		priv = node
	}

	helper(root)
}

func TestFlattenTree(t *testing.T) {
	for _, test := range []struct {
		root   *model.TreeNode
		expect *model.TreeNode
	}{
		{root: &model.TreeNode{
			Val: 1,
			Left: &model.TreeNode{
				Val: 2,
				Left: &model.TreeNode{
					Val: 3,
				},
				Right: &model.TreeNode{
					Val: 4,
				},
			},
			Right: &model.TreeNode{
				Val: 5,
				Right: &model.TreeNode{
					Val: 6,
				},
			},
		}, expect: &model.TreeNode{
			Val: 1,
			Right: &model.TreeNode{
				Val: 2,
				Right: &model.TreeNode{
					Val: 3,
					Right: &model.TreeNode{
						Val: 4,
						Right: &model.TreeNode{
							Val: 5,
							Right: &model.TreeNode{
								Val: 6,
							},
						},
					},
				},
			},
		}},
		{root: nil, expect: nil},
		{root: &model.TreeNode{Val: 0}, expect: &model.TreeNode{Val: 0}},
	} {
		flatten(test.root)
		if !reflect.DeepEqual(test.root, test.expect) {
			t.Errorf("expect %v but got %v", test.expect, test.root)
		}
	}
}
