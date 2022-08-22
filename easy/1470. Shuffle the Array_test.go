package easy

import (
	"fmt"
)

func main() {
	type shuffleArray struct {
		nums []int
		n    int
	}

	tests := []shuffleArray{
		{[]int{2, 5, 1, 3, 4, 7}, 3},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4},
		{[]int{1, 1, 2, 2}, 2},
	}
	for _, test := range tests {
		fmt.Println(shuffle(test.nums, test.n))
	}
}

// https://leetcode.com/problems/shuffle-the-array/
/*
time complexity: O(n/2) increase time with the len of arr, 4 ms
Space complexity: O(1), need probably a fix amount of memory for result arr, 3.6 MB
*/
func shuffle(nums []int, n int) []int {
	// Doesn't need to use copy(). But i learnt two ways to deep copy.
	// copy(result, nums)
	// result = append(result, nums...)
	result := make([]int, len(nums))
	for i := 0; i < n; i++ {
		result[2*i] = nums[i]
		result[2*i+1] = nums[i+n]
	}
	return result
}
