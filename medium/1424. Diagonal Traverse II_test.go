package medium

import (
	"reflect"
	"testing"
)

// Time: O(n)
// Space: O(n)
// func findDiagonalOrder(nums [][]int) []int {
// 	n := len(nums)
// 	maxLenRow := 0
// 	for i := 0; i < n; i++ {
// 		if len(nums[i]) > maxLenRow {
// 			maxLenRow = len(nums[i])
// 		}
// 	}
// 	res := make([]int, 0)
// 	list := make([][]int, n+maxLenRow-1)

// 	for i := n - 1; i >= 0; i-- {
// 		for j := maxLenRow - 1; j >= 0; j-- {
// 			if len(nums[i]) <= j {
// 				continue
// 			}
// 			list[i+j] = append(list[i+j], nums[i][j])
// 		}
// 	}
// 	for i := 0; i < len(list); i++ {
// 		res = append(res, list[i]...)
// 	}
// 	return res
// }

// Time: O(n)
// Space: O(n)
// DFS
type findDiagonalOrderDFS struct {
	x, y int
}

func findDiagonalOrder(nums [][]int) []int {
	const cap = 1 << 13
	res := make([]int, 0, cap)
	queue := []findDiagonalOrderDFS{{0, 0}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.x >= len(nums) || node.y >= len(nums[node.x]) {
			continue
		}
		res = append(res, nums[node.x][node.y])
		if node.y == 0 {
			queue = append(queue, findDiagonalOrderDFS{node.x + 1, node.y})
		}
		queue = append(queue, findDiagonalOrderDFS{node.x, node.y + 1})
	}

	return res
}

func TestFindDiagonalOrder(t *testing.T) {
	for _, test := range []struct {
		nums   [][]int
		output []int
	}{
		{nums: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}, output: []int{1, 4, 2, 7, 5, 3, 8, 6, 9}},
		{nums: [][]int{
			{1, 2, 3, 4, 5},
			{6, 7},
			{8},
			{9, 10, 11},
			{12, 13, 14, 15, 16},
		}, output: []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16}},
	} {
		res := findDiagonalOrder(test.nums)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
