package medium

import (
	"testing"

	"leetcode/uti"
)

/*
IDEA:
A robber has 2 options: a) rob current house i; b) don't rob current house.
If an option "a" is selected it means she can't rob previous i-1 house but can safely proceed to the one before previous i-2 and gets all cumulative loot that follows.
If an option "b" is selected the robber gets all the possible loot from robbery of i-1 and all the following buildings.
So it boils down to calculating what is more profitable:

robbery of current house + loot from houses before the previous
loot from the previous house robbery and any loot captured before that
rob(i) = max( rob(i - 2) + currentHouseValue, rob(i - 1) )
*/

func rob(nums []int) int {
	return robIterativeConstantMem(nums)
}

// Time: O(2^n)
// Time: O(log(n))
func robRecursive(nums []int, i int) int {
	if i < 0 {
		return 0
	}
	return uti.Max(robRecursive(nums, i-2)+nums[i], robRecursive(nums, i-1))
}

// Time: O(n)
// Time: O(n)
func robRecursiveMemo(m map[int]int, nums []int, i int) int {
	if i < 0 {
		return 0
	}
	v, ok := m[i]
	if ok {
		return v
	}

	priv := robRecursiveMemo(m, nums, i-1)
	m[i-1] = priv
	beforePriv := robRecursiveMemo(m, nums, i-2)
	m[i-2] = beforePriv

	return uti.Max(beforePriv+nums[i], priv)
}

// Time: O(n)
// Space: O(1)
func robIterativeConstantMem(nums []int) int {
	left, right, temp := 0, 0, 0
	for _, num := range nums {
		temp = right
		right = uti.Max(left+num, right)
		left = temp
	}
	return right
}

func TestRob(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{1, 2, 3, 1}, output: 4},
		{nums: []int{2, 7, 9, 3, 1}, output: 12},
		{nums: []int{100, 2, 3, 200}, output: 300},
		{
			// Time complexity check:
			nums:   []int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240},
			output: 4173,
		},
	} {
		res := rob(test.nums)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
