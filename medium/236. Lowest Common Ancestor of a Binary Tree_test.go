package medium

import (
	"testing"

	"leetcode/model"
)

// func lowestCommonAncestor(root, p, q *model.TreeNode) *model.TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	var pAncestor, qAncestor, ancestor []*model.TreeNode
// 	var dfs func(node *model.TreeNode)
// 	dfs = func(node *model.TreeNode) {
// 		if pAncestor != nil && qAncestor != nil {
// 			return
// 		}
// 		if node == nil {
// 			return
// 		}
// 		ancestor = append(ancestor, node)
// 		if node == q {
// 			qAncestor = uti.DeepCopySlice(ancestor)
// 		}
// 		if node == p {
// 			pAncestor = uti.DeepCopySlice(ancestor)
// 		}
// 		dfs(node.Left)
// 		dfs(node.Right)
// 		ancestor = ancestor[:len(ancestor)-1]
// 	}
// 	dfs(root)
// 	for i := len(qAncestor) - 1; i >= 0; i-- {
// 		for j := len(pAncestor) - 1; j >= 0; j-- {
// 			if qAncestor[i] == pAncestor[j] {
// 				return qAncestor[i]
// 			}
// 		}
// 	}
// 	return nil
// }

// Time: O(n)
// Space: O(logn)
func lowestCommonAncestor(root, p, q *model.TreeNode) *model.TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

func TestLowestCommonAncestor(t *testing.T) {
	for _, test := range []struct {
		root     *model.TreeNode
		getNodes func(node *model.TreeNode) (*model.TreeNode, *model.TreeNode, *model.TreeNode)
	}{
		{
			root: &model.TreeNode{
				Val: 3,
				Left: &model.TreeNode{
					Val: 5,
					Left: &model.TreeNode{
						Val: 6,
					},
					Right: &model.TreeNode{
						Val: 2,
						Left: &model.TreeNode{
							Val: 7,
						},
						Right: &model.TreeNode{
							Val: 4,
						},
					},
				},
				Right: &model.TreeNode{
					Val: 1,
					Left: &model.TreeNode{
						Val: 0,
					},
					Right: &model.TreeNode{
						Val: 8,
					},
				},
			}, getNodes: func(node *model.TreeNode) (*model.TreeNode, *model.TreeNode, *model.TreeNode) {
				return node.Left, node.Right, node
			},
		},
		{
			root: &model.TreeNode{
				Val: 3,
				Left: &model.TreeNode{
					Val: 5,
					Left: &model.TreeNode{
						Val: 6,
					},
					Right: &model.TreeNode{
						Val: 2,
						Left: &model.TreeNode{
							Val: 7,
						},
						Right: &model.TreeNode{
							Val: 4,
						},
					},
				},
				Right: &model.TreeNode{
					Val: 1,
					Left: &model.TreeNode{
						Val: 0,
					},
					Right: &model.TreeNode{
						Val: 8,
					},
				},
			}, getNodes: func(node *model.TreeNode) (*model.TreeNode, *model.TreeNode, *model.TreeNode) {
				return node.Left, node.Left.Right.Right, node.Left
			},
		},
	} {
		p, q, output := test.getNodes(test.root)
		res := lowestCommonAncestor(test.root, p, q)
		if res != output {
			t.Errorf("expect %v but got %v, test %v", output, res, test)
		}
	}
}
