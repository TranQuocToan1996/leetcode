package easy

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"leetcode/model"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func TestAverageLevelBinaryTree(t *testing.T) {
	tests := []struct {
		tree   *model.TreeNode
		expect []float64
	}{
		{
			&model.TreeNode{
				Val: 1,
				Left: &model.TreeNode{
					Val: 2,
					Left: &model.TreeNode{
						Val: 3,
					},
				},
				Right: &model.TreeNode{
					Val: 3,
				},
			},
			[]float64{1.00000, 2.5000, 3.0},
		},
		{
			&model.TreeNode{Val: 1, Left: &model.TreeNode{Val: 2}},
			[]float64{1.00000, 2.0000},
		},
	}

	for no, test := range tests {
		beginDFS := time.Now()
		resDFS := averageOfLevels(test.tree)
		fmt.Println("DFS nanosec", time.Since(beginDFS).Nanoseconds())
		if !reflect.DeepEqual(resDFS, test.expect) {
			t.Errorf("[%v]: Fail test DFS", no)
		}

		beginBFS := time.Now()
		resBFS := averageOfLevelsBFS(test.tree)
		fmt.Println("DFS nanosec", time.Since(beginBFS).Nanoseconds())
		if !reflect.DeepEqual(resBFS, test.expect) {
			t.Errorf("[%v]: Fail test BFS", no)
		}
	}
}

// Using Depth First Search (DFS)
// Time complexity : O(n). The whole tree is traversed once only. Here, n refers to the total number of nodes in the given binary tree.

// Space complexity : O(h). res and count array of size h are used. Here, h refers to the height(maximum number of levels) of the given binary tree. Further, the depth of the recursive tree could go upto h only.

type totalHeight struct {
	totalByHeight  map[int]int
	numberByHeight map[int]int
}

func averageOfLevels(root *model.TreeNode) []float64 {
	if root == nil {
		return []float64{0}
	}

	// outside of tree
	curHeight := 0
	total := totalHeight{
		map[int]int{},
		map[int]int{},
	}

	sumByHeight(root, total, curHeight)
	lengthArr := 0
	for key := range total.totalByHeight {
		if lengthArr < key {
			lengthArr = key
		}
	}

	var res []float64 = make([]float64, lengthArr+1)

	for key, val := range total.totalByHeight {
		res[key] = float64(val) / float64(total.numberByHeight[key])
	}

	return res
}

func sumByHeight(node *model.TreeNode, total totalHeight, height int) {
	if node == nil {
		return
	}

	total.totalByHeight[height] += node.Val
	total.numberByHeight[height] += 1
	height++

	if node.Left != nil {
		sumByHeight(node.Left, total, height)
	}
	if node.Right != nil {
		sumByHeight(node.Right, total, height)
	}
}

// Approach #2 Breadth First Search (BFS)

// Time complexity : O(n). The whole tree is traversed atmost once. Here, n refers to the number of nodes in the given binary tree.

// Space complexity : O(m). The size of queue or temp can grow upto atmost the maximum number of nodes at any level in the given binary tree. Here, m refers to the maximum mumber of nodes at any level in the input tree.
func averageOfLevelsBFS(root *model.TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	queue := []*model.TreeNode{root}
	ans := []float64{}
	for len(queue) != 0 {
		n, sum := len(queue), 0
		for i := 0; i < n; i++ {
			pop := queue[0]
			queue = queue[1:]
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
			sum += pop.Val
		}
		ans = append(ans, float64(sum)/float64(n))
	}
	return ans
}
