package medium

import "testing"

// Time: O(n)
// Space: O(n)
func subarraySum(nums []int, k int) int {
	var (
		n         = len(nums)
		res       = 0
		prefixSum = make(map[int]int, n)
		curSum    = 0
	)
	prefixSum[0] = 1 // Imagine there is an empty slice before nums
	for i := 0; i < n; i++ {
		curSum += nums[i]
		oldSum := curSum - k
		res += prefixSum[oldSum]
		prefixSum[curSum]++
	}
	return res
}

// Time: O(n^3)
// Space: O(1)
// Brute force
// func subarraySum(nums []int, k int) int {
// 	n := len(nums)
// 	res := 0
// 	for start := 0; start < n; start++ {
// 		for end := start + 1; end <= n; end++ {
// 			sum := 0
// 			for i := start; i < end; i++ {
// 				sum += nums[i]
// 			}
// 			if sum == k {
// 				res++
// 			}
// 		}
// 	}
// 	return res
// }

func TestSubarraySum(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		k      int
		output int
	}{
		{nums: []int{1, 1, 1, 2, 1, 1}, k: 2, output: 4},
		{nums: []int{1, 2, 3}, k: 3, output: 2},
	} {
		res := subarraySum(test.nums, test.k)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
