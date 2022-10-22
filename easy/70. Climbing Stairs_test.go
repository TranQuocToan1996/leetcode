package easy

import "testing"

// https://leetcode.com/problems/climbing-stairs/

func TestClimbStairs(t *testing.T) {
	for _, test := range []struct {
		steps int
		ways  int
	}{
		{steps: 1, ways: 1},
		{steps: 2, ways: 2},
		{steps: 3, ways: 3},
	} {
		if climbStairs(test.steps) != test.ways {
			t.Errorf("Expect %v but got %v at %v step", test.ways, climbStairs(test.steps), test.steps)
		}
	}
}

// time: O(n) loop through ~ n steps
// memory: O(1): just need a and b. No more memory allocation.

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
// Memory Usage: 1.9 MB, less than 39.33% of Go online submissions for Climbing Stairs.

func climbStairs(n int) int {
	a, b := 1, 1
	for n > 1 {
		n--
		a, b = b, a+b
	}

	return b
}
