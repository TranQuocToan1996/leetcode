package main

import (
	"fmt"
	"sort"
)

func main() {
	type Test struct {
		nums []int
		k    int
	}

	tests := []Test{
		{[]int{1, 2, 3, 4}, 5},
		{[]int{3, 1, 3, 4, 3}, 6},
	}
	for _, test := range tests {
		fmt.Println(maxOperations(test.nums, test.k))
	}
}

// https://leetcode.com/problems/two-sum/submissions/

/*
time complexity: O(n), 165 ms
Space complexity: O(n), 10.2 MB
*/

// func maxOperations(nums []int, k int) int {
// 	m := map[int]int{}
// 	var res int
// 	for _, v := range nums {
// 		if m[v] == 0 {
// 			m[k-v]++
// 		} else {
// 			res++
// 			m[v]--
// 		}
// 	}
// 	return res
// }

/*
time complexity: O(n), 152 ms
Space complexity: O(n), 8.9 MB
*/
func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	left, right, res := 0, len(nums)-1, 0
	for left < right {
		if nums[left]+nums[right] == k {
			res++
			left++
			right--
		} else if nums[left]+nums[right] > k {
			right--
		} else {
			left++
		}
	}
	return res
}
