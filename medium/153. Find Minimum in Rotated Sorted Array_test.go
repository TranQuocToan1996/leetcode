package medium

import (
	"testing"
)

// You must write an algorithm that runs in O(log n) time.
// Space: O(1)
func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func TestFindMin(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{3, 4, 5, 1, 2}, output: 1},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, output: 0},
		{nums: []int{11, 13, 15, 17}, output: 11},
		{nums: []int{3, 1, 2}, output: 1},
	} {
		res := findMin(test.nums)
		if res != test.output {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}

/*
4 5 6 7 8 9 0 1 2 3
l       m         r
m > r => l = m + 1
4 5 6 7 8 9 0 1 2 3
          l   m   r
m < r => r = m
4 5 6 7 8 9 0 1 2 3
          l m r
m < r => r = m
4 5 6 7 8 9 0 1 2 3
          l r
m > r => l = m + 1
7 8 0 1 2 3 4 5 6
            lr
r >= l => return l

7 8 0 1 2 3 4 5 6
l       m       r
m < r => r = m
7 8 0 1 2 3 4 5 6
l m     m
m > r => l = m + 1
7 8 0 1 2 3 4 5 6
    l m r
m < l => r = m
7 8 0 1 2 3 4 5 6
    l r
m < r => r = m
7 8 0 1 2 3 4 5 6
    lr
r >= l => return l

8 0 1 2 3 4 5 6 7
l       m       r
m < r => r = m
8 0 1 2 3 4 5 6 7
l   m   r
m < r => r = m
8 0 1 2 3 4 5 6 7
l m r
m < r => r = m
8 0 1 2 3 4 5 6 7
l r
m > r => l = m + 1
8 0 1 2 3 4 5 6 7
  lr
r >=l => return l

0 1 2 3 4
l   m   r
m < r => r = m
0 1 2 3 4
l m r
m < r => r = m
0 1 2 3 4
l r
m < r => r = m
0 1 2 3 4
lr
r >= l => return l
*/
