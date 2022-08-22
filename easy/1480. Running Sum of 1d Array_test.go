package easy

import (
	"fmt"
	"testing"
)

func TestSumArrayOneD(t *testing.T) {

	tests := [][]int{
		{1, 2, 3, 4},
		{1, 1, 1, 1, 1},
		{3, 1, 2, 10, 1},
	}
	for _, test := range tests {
		fmt.Println(runningSum(test))
	}
}

// https://leetcode.com/problems/running-sum-of-1d-array/
/*
time complexity: O(n - 1),0 ms
Space complexity: O(1),2.7 MB
*/
func runningSum(nums []int) []int {
	length := len(nums)
	for i := 1; i < length; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
