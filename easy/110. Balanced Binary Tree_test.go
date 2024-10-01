package easy

import (
	"math"
	"testing"

	"leetcode/model"
	"leetcode/uti"
)

func TestLevelOrder(t *testing.T) {
	for _, test := range []struct {
		root   *model.TreeNode
		output bool
	}{
		{root: nil, output: true},
		{root: &model.TreeNode{
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
		}, output: true},
	} {
		res := isBalanced(test.root)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %+v", test.output, res, test.root)
		}
	}
}

func isBalanced(root *model.TreeNode) bool {
	if root == nil {
		return true
	}
	heightLeft := maxDepth(root.Left)
	heightRight := maxDepth(root.Right)
	if math.Abs(float64(heightLeft-heightRight)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(node *model.TreeNode) int {
	if node == nil {
		return 0
	}
	return uti.Max(maxDepth(node.Left), maxDepth(node.Right)) + 1
}
