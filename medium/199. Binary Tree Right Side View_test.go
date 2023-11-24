package medium

import (
	"reflect"
	"testing"

	"leetcode/model"
)

func rightSideView(root *model.TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		res         = []int{root.Val}
		mainQueue   = []*model.TreeNode{root}
		secondQueue []*model.TreeNode
		level       = 1
	)
	for len(mainQueue) > 0 {
		node := mainQueue[0]
		mainQueue = mainQueue[1:]
		if node.Right != nil {
			secondQueue = append(secondQueue, node.Right)
		}
		if node.Left != nil {
			secondQueue = append(secondQueue, node.Left)
		}
		if len(res) < level {
			res = append(res, node.Val)
		}
		if len(mainQueue) == 0 {
			mainQueue = secondQueue
			secondQueue = nil
			level++
		}
	}
	return res
}

func TestRightSideView(t *testing.T) {
	for _, test := range []struct {
		root   *model.TreeNode
		output []int
	}{
		{root: &model.TreeNode{
			Val: 1,
			Right: &model.TreeNode{
				Val: 3,
				Right: &model.TreeNode{
					Val: 4,
				},
			},
			Left: &model.TreeNode{
				Val: 2,
				Right: &model.TreeNode{
					Val: 5,
				},
			},
		}, output: []int{1, 3, 4}},
		{root: nil, output: nil},
	} {
		res := rightSideView(test.root)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
