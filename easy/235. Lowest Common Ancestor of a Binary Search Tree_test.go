package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Left less than right

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root
	v1, v2 := p.Val, q.Val
	for {
		if v1 < cur.Val && v2 < cur.Val {
			cur = cur.Left
		} else if v1 > cur.Val && v2 > cur.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}
}

// recursive
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//      if p.Val < root.Val && q.Val < root.Val {
//          return lowestCommonAncestor(root.Left, p, q)
//      } else if p.Val > root.Val && q.Val > root.Val {
//          return lowestCommonAncestor(root.Right, p, q)
//      } else {
//          return root
//      }
// }
