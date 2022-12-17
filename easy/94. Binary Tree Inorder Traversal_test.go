package easy

import (
	"leetcode/uti"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-inorder-traversal/
// https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/
func TestInorderTraversal(t *testing.T) {
	defer uti.LogRuntime(time.Now())
	tests := []struct {
		tree   *TreeNode
		expect []int
	}{
		{
			&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
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
			&TreeNode{Val: 1},
			[]int{1},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, inorderTraversal(test.tree))
	}
}

// func inorderTraversal(root *TreeNode) []int {
// 	res := []int{}

// 	inorderRecursive(&res, root)

// 	return res
// }

func inorderRecursive(res *[]int, cur *TreeNode) {
	if cur == nil {
		return
	}

	inorderRecursive(res, cur.Left)
	*res = append(*res, cur.Val) // the append will change the len, cap -> so must use pointer
	inorderRecursive(res, cur.Right)
}
func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	inorderStack(&res, root)

	return res
}

func inorderStack(res *[]int, root *TreeNode) {
	stack := []*TreeNode{}
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

func pop(s *[]*TreeNode) (*TreeNode, bool) {
	if len(*s) == 0 {
		return nil, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
