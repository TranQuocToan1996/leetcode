package easy

import "leetcode/model"

func invertTree(root *model.TreeNode) *model.TreeNode {
	res := root
	swapChild(root)
	return res
}

func swapChild(node *model.TreeNode) {
	if node == nil {
		return
	}
	swapChild(node.Left)
	swapChild(node.Right)
	temp := node.Left
	node.Left = node.Right
	node.Right = temp
}
