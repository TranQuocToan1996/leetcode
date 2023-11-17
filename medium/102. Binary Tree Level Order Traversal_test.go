package medium

import (
	"reflect"
	"testing"

	"leetcode/model"
)

func levelOrder(root *model.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	mainQueue := []*model.TreeNode{root}
	secondaryQueue := []*model.TreeNode{}
	cur := []int{}
	for len(mainQueue) > 0 || len(secondaryQueue) > 0 {
		if len(mainQueue) == 0 {
			mainQueue = secondaryQueue
			secondaryQueue = nil
		}
		node := mainQueue[0]
		mainQueue = mainQueue[1:]
		cur = append(cur, node.Val)
		if len(mainQueue) == 0 {
			res = append(res, cur)
			cur = nil
		}
		if node.Left != nil {
			secondaryQueue = append(secondaryQueue, node.Left)
		}
		if node.Right != nil {
			secondaryQueue = append(secondaryQueue, node.Right)
		}
	}

	return res
}

func TestLevelOrder(t *testing.T) {
	for _, test := range []struct {
		root   *model.TreeNode
		output [][]int
	}{
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
		}, output: [][]int{
			{3},
			{9, 20},
			{15, 7},
		}},
		{root: &model.TreeNode{
			Val: 1,
		}, output: [][]int{{1}}},
		{root: nil, output: nil},
	} {
		res := levelOrder(test.root)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %+v", test.output, res, test.root)
		}
	}
}
