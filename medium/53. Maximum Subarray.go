package medium

// Constraints: Length nums > 0
// Time: O(n)
// Space: O(1)
func maxSubArray(nums []int) int {
	var (
		max = nums[0]
		cur = nums[0]
	)

	for i := 1; i < len(nums); i++ {
		if cur < 0 {
			cur = 0
		}
		cur += nums[i]
		if cur > max {
			max = cur
		}
	}

	return max
}
