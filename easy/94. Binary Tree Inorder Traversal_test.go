package easy

import (
	"testing"
	"time"

	"leetcode/model"
	"leetcode/uti"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-inorder-traversal/
// https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/
func TestInorderTraversal(t *testing.T) {
	defer uti.MemoryAlloc()
	defer uti.LogRuntime(time.Now())
	tests := []struct {
		tree   *model.TreeNode
		expect []int
	}{
		{
			&model.TreeNode{
				Val: 1,
				Right: &model.TreeNode{
					Val: 2,
					Left: &model.TreeNode{
						Val: 3,
					},
				},
			},
			[]int{1, 3, 2},
		},
		{
			nil,
			[]int{},
		},
		{
			&model.TreeNode{Val: 1},
			[]int{1},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, inorderTraversal(test.tree))
	}
}

func inorderTraversal(root *model.TreeNode) []int {
	res := []int{}

	inorderRecursive(&res, root)
	// inorderStack(&res, root)

	return res
}

// Time complexity: O(n), 2 ms
// Space complexity: O(n), 1.9 MB
func inorderRecursive(res *[]int, cur *model.TreeNode) {
	if cur == nil {
		return
	}

	inorderRecursive(res, cur.Left)
	*res = append(*res, cur.Val) // the append will change the len, cap -> so must use pointer
	inorderRecursive(res, cur.Right)
}

func inorderStack(res *[]int, root *model.TreeNode) {
	stack := []*model.TreeNode{}
	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			// Move to left until point to nil and append cur val to stack
			stack = append(stack, cur)
			cur = cur.Left
		}
		// Pop val from stack to res
		val, ok := pop(&stack)
		if ok {
			cur = val
			*res = append(*res, cur.Val)
		}

		// Continue traversal the right child node -> for loop
		cur = cur.Right
	}
}

func pop(s *[]*model.TreeNode) (*model.TreeNode, bool) {
	if len(*s) == 0 {
		return nil, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
