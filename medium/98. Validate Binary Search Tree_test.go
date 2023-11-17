package medium

import (
	"math"
	"testing"

	"leetcode/model"
)

// Time: O(n)
// Space: O(logn) call stack validateBST_DFSfunc
func isValidBST(root *model.TreeNode) bool {
	// return validateBST_DFSfunc(root, math.MinInt, math.MaxInt)
	return validateBST_BFSfunc(root)
}

func validateBST_DFSfunc(node *model.TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if min < node.Val && node.Val < max {
		return validateBST_DFSfunc(node.Left, min, node.Val) &&
			validateBST_DFSfunc(node.Right, node.Val, max)
	}
	return false
}

type checkValidBST struct {
	node     *model.TreeNode
	min, max int
}

// Time: O(n)
// Space: O(logn) queue length
func validateBST_BFSfunc(root *model.TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []checkValidBST{{node: root, min: math.MinInt, max: math.MaxInt}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.min >= cur.node.Val || cur.node.Val >= cur.max {
			return false
		}
		if cur.node.Left != nil {
			queue = append(queue, checkValidBST{
				node: cur.node.Left,
				min:  cur.min,
				max:  cur.node.Val,
			})
		}
		if cur.node.Right != nil {
			queue = append(queue, checkValidBST{
				node: cur.node.Right,
				min:  cur.node.Val,
				max:  cur.max,
			})
		}
	}
	return true
}

func TestIsValidBST(t *testing.T) {
	for _, test := range []struct {
		root   *model.TreeNode
		output bool
	}{
		{root: &model.TreeNode{
			Val: 2,
			Left: &model.TreeNode{
				Val: 1,
			},
			Right: &model.TreeNode{
				Val: 3,
			},
		}, output: true},
		{root: &model.TreeNode{
			Val: 5,
			Left: &model.TreeNode{
				Val: 1,
			},
			Right: &model.TreeNode{
				Val: 4,
				Left: &model.TreeNode{
					Val: 3,
				},
				Right: &model.TreeNode{
					Val: 6,
				},
			},
		}, output: false},
		{root: &model.TreeNode{
			Val: 5,
			Left: &model.TreeNode{
				Val: 4,
			},
			Right: &model.TreeNode{
				Val: 6,
				Left: &model.TreeNode{
					Val: 3,
				},
				Right: &model.TreeNode{
					Val: 7,
				},
			},
		}, output: false},
	} {
		res := isValidBST(test.root)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %+v", test.output, res, test.root)
		}
	}
}
