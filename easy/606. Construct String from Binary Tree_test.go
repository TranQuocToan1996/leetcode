package easy

import (
	"fmt"
	"testing"

	"leetcode/model"
)

func tree2str(root *model.TreeNode) string {
	res := ""
	if root != nil {
		res += fmt.Sprint(root.Val)
	}
	if root.Left != nil {
		res += "("
		res += tree2str(root.Left)
		res += ")"
	}
	if root.Right != nil {
		if root.Left == nil {
			res += "()"
		}
		res += "("
		res += tree2str(root.Right)
		res += ")"
	}
	return res
}

func TestTreeToString(t *testing.T) {
	for _, test := range []struct {
		root   *model.TreeNode
		output string
	}{
		{root: &model.TreeNode{
			Val: 1,
			Left: &model.TreeNode{
				Val: 2,
				Left: &model.TreeNode{
					Val: 4,
				},
			},
			Right: &model.TreeNode{
				Val: 3,
			},
		}, output: "1(2(4))(3)"},
		{root: &model.TreeNode{
			Val: 1,
			Left: &model.TreeNode{
				Val: 2,
				Right: &model.TreeNode{
					Val: 4,
				},
			},
			Right: &model.TreeNode{
				Val: 3,
			},
		}, output: "1(2()(4))(3)"},
	} {
		res := tree2str(test.root)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
