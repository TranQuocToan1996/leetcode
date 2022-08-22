package easy

import (
	"fmt"
	"math"
)

func main() {

	tests := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
	}
	for _, test := range tests {
		fmt.Println(maxArea(test))
	}
}

// https://leetcode.com/problems/container-with-most-water/

/*
time complexity: O(n^2),  76 ms
Space complexity: O(1), 8.7 MB
*/
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0

	for left < right {
		max = int(math.Max(float64(max), float64((right-left)*
			int(math.Min(float64(height[left]), float64(height[right]))))))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

/*
time complexity: O(n^2),  Out of time ms
Space complexity: O(1),  MB
*/
/* func maxArea(height []int) int {
	leng := len(height)
	result := 0
	for i := 0; i < leng; i++ {
		for j := i + 1; j < leng; j++ {
			area := minOf(height[i], height[j]) * (j - i)
			if result < area {
				result = area
			}
		}
	}
	return result
}

func minOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
} */
