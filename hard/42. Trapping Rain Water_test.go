package hard

import (
	"testing"
)

// Time: O(n * level) where level is max height
// Space: O(1)
// func trap(height []int) int {
// 	var (
// 		n     = len(height)
// 		water = 0
// 		level = math.MinInt
// 	)
// 	for _, h := range height {
// 		level = uti.Max(level, h)
// 	}
// 	for ; level >= 0; level-- {
// 		candidateWater := 0
// 		canTrap := false
// 		for i := 0; i < n; i++ {
// 			if height[i] >= level {
// 				if canTrap {
// 					water += candidateWater
// 					candidateWater = 0
// 				}
// 				canTrap = true
// 				continue
// 			}
// 			if canTrap {
// 				candidateWater++
// 			}
// 		}
// 	}
// 	return water
// }

// Dynamic programming
// Intuition: Check the left/right max height.
// At i position, candidateWater = min(maxLeft, maxRight)
// Only add to water response if we got candidateWater > 0.
// Time: O(n)
// Space: O(n)
// func trap(height []int) int {
// 	var (
// 		water    = 0
// 		n        = len(height)
// 		maxLeft  = make([]int, n)
// 		maxRight = make([]int, n)
// 	)
// 	for i := 1; i < n; i++ {
// 		maxLeft[i] = uti.Max(height[i-1], maxLeft[i-1])
// 		maxRight[n-i-1] = uti.Max(height[n-i], maxRight[n-i])
// 	}
// 	for i := 0; i < n; i++ {
// 		water += uti.Max(0, uti.Min(maxLeft[i], maxRight[i])-height[i])
// 	}
// 	return water
// }

// Two pointer
/*
Intuition: Let's check {2, 1, 3, 1, 2}
we would fill up to the 3, and then see that we cannot fill over to the 2 because we would overflow,
so we instead mirror the algorithm and bring from the 2 backward
2, 1, 3 -> left
3, 1 ,2 -> right
*/
// Time: O(n)
// Space: O(1)
func trap(height []int) int {
	var (
		water             = 0
		n                 = len(height)
		left, right       = 0, n - 1
		maxLeft, maxRight = height[0], height[n-1]
	)

	for left < right {
		if height[left] < height[right] {
			if maxLeft <= height[left] {
				maxLeft = height[left]
			} else {
				water += (maxLeft - height[left])
			}
			left++
		} else {
			if maxRight <= height[right] {
				maxRight = height[right]
			} else {
				water += (maxRight - height[right])
			}
			right--
		}
	}

	return water
}

func TestTrap(t *testing.T) {
	for _, test := range []struct {
		height []int
		output int
	}{
		{height: []int{2, 1, 3, 1, 4}, output: 3},
		{height: []int{2, 1, 3, 1, 2}, output: 2},
		{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, output: 6},
		{height: []int{4, 2, 0, 3, 2, 5}, output: 9},
	} {
		res := trap(test.height)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
