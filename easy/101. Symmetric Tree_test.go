package easy

import (
	"leetcode/uti"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSymetric(t *testing.T) {
	defer uti.LogRuntime(time.Now())
	tests := []struct {
		tree   *TreeNode
		expect bool
	}{
		{
			&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			true,
		},
		{
			&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, isSymmetric(test.tree))
	}
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// return isMirrorTreeRecursive(root.Left, root.Right)
	return isMirrorTreeBFS(root)
}

// uti.LogRuntime = 12917 nsec
// Letcode runtime na ms
// Leetcode memory 2.9 MB (beat 70%)
// Time complexity :O(n). Because we traverse the entire input tree once, the total run time is O(n), where n is the total number of nodes in the tree.
// Space complexity :O(n) The number of recursive calls is bound by the height of the tree. In the worst case, the tree is linear and the height is in O(n). Therefore, space complexity due to recursive calls on the stack is O(n) in the worst case.
func isMirrorTreeRecursive(tree1, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}

	if tree1 == nil || tree2 == nil {
		return false
	}

	return tree1.Val == tree2.Val && isMirrorTreeRecursive(tree1.Left, tree2.Right) && isMirrorTreeRecursive(tree1.Right, tree2.Left)
}

const (
	maxChildsPerNode = 2
)

type treeQueue [][maxChildsPerNode]*TreeNode

// pop change the len, cap so that must use pointer receiver
func (q *treeQueue) pop() (left, right *TreeNode) {
	if len(*q) == 0 {
		return nil, nil
	}

	left, right = (*q)[0][0], (*q)[0][1]

	*q = (*q)[1:]
	return left, right
}

func (q treeQueue) isEmpty() bool {
	return len(q) == 0
}

// uti.LogRuntime = 12333 nsec
// Letcode runtime na ms
// Leetcode memory 2.9 MB (beat 70%)
// Time complexity :O(n). Because we traverse the entire input tree once, the total run time is O(n), where n is the total number of nodes in the tree.
// Space complexity :O(n) There is additional space required for the search queue. In the worst case, we have to insert O(n) nodes in the queue. Therefore, space complexity is O(n).

func isMirrorTreeBFS(root *TreeNode) bool {
	if root == nil {
		return true
	}

	q := treeQueue{
		{root.Left, root.Right},
	}

	for q.isEmpty() {
		left, right := q.pop()

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		q = append(q, treeQueue{{left.Left, right.Right}, {left.Right, right.Left}}...)
	}

	return true
}
