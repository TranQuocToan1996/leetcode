package medium

import (
	"math"
	"reflect"
	"testing"
)

// Time: O(m * n*logn) m is len(l), n is len(nums)
// Space: O(n)
// func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
// 	res := make([]bool, len(l))
// 	for i := 0; i < len(l); i++ {
// 		res[i] = checkArithmetic(uti.DeepCopySlice(nums[l[i] : r[i]+1]))
// 	}
// 	return res
// }

// func checkArithmetic(nums []int) bool {
// 	if len(nums) <= 2 {
// 		return true
// 	}
// 	sort.IntSlice(nums).Sort()
// 	diff := nums[1] - nums[0]
// 	for i := 2; i < len(nums); i++ {
// 		if nums[i]-nums[i-1] != diff {
// 			return false
// 		}
// 	}
// 	return true
// }

// Time: O(m * n) m is len(l), n is len(nums)
// Space: O(n)
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	res := make([]bool, len(l))
	isArithmetic := func(nums []int) bool {
		if len(nums) <= 2 {
			return true
		}
		seen := make(map[int]bool, len(nums))
		min, max := math.MaxInt, math.MinInt
		for i := 0; i < len(nums); i++ {
			seen[nums[i]] = true
			if min > nums[i] {
				min = nums[i]
			}
			if max < nums[i] {
				max = nums[i]
			}
		}
		if (max-min)%(len(nums)-1) != 0 {
			return false
		}
		diff := (max - min) / (len(nums) - 1)
		cur := min
		for i := 0; i < len(nums); i++ {
			if !seen[cur] {
				return false
			}
			cur += diff
		}
		return true
	}

	for i := 0; i < len(res); i++ {
		res[i] = isArithmetic(nums[l[i] : r[i]+1])
	}

	return res
}

func TestCheckArithmeticSubarrays(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		l      []int
		r      []int
		output []bool
	}{
		{
			nums: []int{4, 6, 5, 9, 3, 7}, l: []int{0, 0, 2}, r: []int{2, 3, 5},
			output: []bool{true, false, true},
		},
		{
			nums: []int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}, l: []int{0, 1, 6, 4, 8, 7}, r: []int{4, 4, 9, 7, 9, 10},
			output: []bool{false, true, false, false, true, true},
		},
		{
			nums: []int{-3, -6, -8, -4, -2, -8, -6, 0, 0, 0, 0}, l: []int{5, 4, 3, 2, 4, 7, 6, 1, 7}, r: []int{6, 5, 6, 3, 7, 10, 7, 4, 10},
			output: []bool{true, true, true, true, false, true, true, true, true},
		},
	} {
		res := checkArithmeticSubarrays(test.nums, test.l, test.r)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
