package medium

import (
	"testing"

	"leetcode/model"
)

// Time: O(n)
// space: O(n)
func pathSum(root *model.TreeNode, targetSum int) int {
	var (
		res   = 0
		cache = map[int]int{0: 1}
		dfs   func(root *model.TreeNode, curSum int)
	)
	dfs = func(root *model.TreeNode, curSum int) {
		if root == nil {
			return
		}
		curSum += root.Val
		oldSum := curSum - targetSum
		res += cache[oldSum]
		cache[curSum]++
		dfs(root.Left, curSum)
		dfs(root.Right, curSum)
		cache[curSum]--
	}
	dfs(root, 0)
	return res
}

func TestPathSum(t *testing.T) {
	for _, test := range []struct {
		root      *model.TreeNode
		targetSum int
		output    int
	}{
		{root: &model.TreeNode{
			Val: 10,
			Left: &model.TreeNode{
				Val: 5,
				Left: &model.TreeNode{
					Val: 3,
					Left: &model.TreeNode{
						Val: 3,
					},
					Right: &model.TreeNode{
						Val: -2,
					},
				},
				Right: &model.TreeNode{
					Val: 2,
					Right: &model.TreeNode{
						Val: 1,
					},
				},
			},
			Right: &model.TreeNode{
				Val: -3,
				Right: &model.TreeNode{
					Val: 11,
				},
			},
		}, targetSum: 8, output: 3},
	} {
		res := pathSum(test.root, test.targetSum)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
