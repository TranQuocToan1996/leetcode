package medium

import (
	"testing"

	"leetcode/model"
)

// func kthSmallest(root *model.TreeNode, k int) int {
// 	res := make([]int, k)
// 	for k := range res {
// 		res[k] = math.MaxInt
// 	}
// 	var dfs func(node *model.TreeNode)
// 	dfs = func(node *model.TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		res = insertLSmallest(res, node.Val)
// 		dfs(node.Left)
// 		dfs(node.Right)
// 	}
// 	dfs(root)
// 	return res[len(res)-1]
// }

// func insertLSmallest(kSmallest []int, value int) []int {
// 	if kSmallest[len(kSmallest)-1] <= value {
// 		return kSmallest
// 	}
// 	for i := 0; i < len(kSmallest); i++ {
// 		if value < kSmallest[i] {
// 			kSmallest = uti.InsertToSlice(kSmallest, i, value)
// 			kSmallest = kSmallest[:len(kSmallest)-1]
// 			return kSmallest
// 		}
// 	}
// 	return kSmallest
// }

// Time: O(n)
// Space: O(1)
func kthSmallest(root *model.TreeNode, k int) int {
	c := make(chan int)
	go inorderTravelChannel(root, c)
	for i := 0; i < k-1; i++ {
		<-c
	}
	return <-c
}

func inorderTravelChannel(node *model.TreeNode, c chan int) {
	if node == nil {
		return
	}
	inorderTravelChannel(node.Left, c)
	c <- node.Val
	inorderTravelChannel(node.Right, c)
}

func TestKthSmallest(t *testing.T) {
	for _, test := range []struct {
		root   *model.TreeNode
		k      int
		output int
	}{
		{root: &model.TreeNode{
			Val: 3,
			Left: &model.TreeNode{
				Val: 1,
				Right: &model.TreeNode{
					Val: 2,
				},
			},
			Right: &model.TreeNode{
				Val: 4,
			},
		}, k: 1, output: 1},
	} {
		res := kthSmallest(test.root, test.k)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
