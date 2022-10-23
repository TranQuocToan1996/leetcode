package medium

import (
	"reflect"
	"sort"
	"testing"
)

// https://leetcode.com/problems/3sum/

func TestThreeSum(t *testing.T) {
	for _, test := range []struct {
		input  []int
		output [][]int
	}{
		{input: []int{-1, 0, 1, 2, -1, -4}, output: [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		}},
		{input: []int{0, 1, 1}, output: [][]int{}},
		{input: []int{0, 0, 0}, output: [][]int{{0, 0, 0}}},
		{input: []int{3, 0, -2, -1, 1, 2}, output: [][]int{{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}},
	} {
		sum := threeSum(test.input)
		if !reflect.DeepEqual(test.output, sum) &&
			len(test.output) > 0 && len(sum) > 0 {
			t.Errorf("expect %v but got %v", test.output, sum)
		}
	}
}

// Runtime: 51 ms, faster than 64.27% of Go online submissions for 3Sum.
// Memory Usage: 7.5 MB, less than 67.92% of Go online submissions for 3Sum.
// Time: O(left* (mid + right)) ~ O(n^3)
// Space: O(1)

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)
	if n < 3 {
		return ans
	}
	sort.Ints(nums)

	for left := 0; left < n-1; left++ {
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}

		for mid, right := left+1, n-1; mid < right; {
			if mid > left+1 && nums[mid] == nums[mid-1] {
				mid++
				continue
			}
			if right < n-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			sum := nums[left] + nums[mid] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[left], nums[mid], nums[right]})
				mid++
				right--
			} else if sum < 0 {
				mid++
				// Move pointer faster for equal val
				for mid < right && nums[mid] == nums[mid-1] {
					mid++
				}
			} else {
				right--
				for mid < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return ans
}

/* Unfortunally, my solution Time Limit Exceeded due to O(n*(a + b)) ~ O(n^2), n is mid, a and b is left and right
func threeSum(nums []int) [][]int {
	leng := len(nums)
	if leng <= 2 {
		return [][]int{}
	}

	sort.Ints(nums)

	cache := map[int][][]int{}

	for mid := 1; mid < leng-1; mid++ {
		left, right := 0, leng-1
		for left < mid && left < right && mid < right {
			currentVal := nums[left] + nums[mid] + nums[right]
			found, ok := cache[currentVal]
			if !ok {
				cache[currentVal] =
					[][]int{{nums[left], nums[mid], nums[right]}}
			} else {
				match := false
				for _, arr := range found {
					if arr[0] == nums[left] && arr[1] == nums[mid] && arr[2] == nums[right] {
						match = true
						break
					}
				}
				if !match {
					found = append(found, []int{nums[left], nums[mid], nums[right]})
					cache[currentVal] = found
				}
			}

			if currentVal >= 0 {
				right--
			} else {
				left++
			}
		}
	}

	return cache[0]
} */
