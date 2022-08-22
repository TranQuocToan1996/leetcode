package easy

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCallPoint682(t *testing.T) {

	tests := [][]string{
		{"5", "2", "C", "D", "+"},
		{"5", "-2", "4", "C", "D", "9", "+", "+"},
		{"1"},
	}
	for _, test := range tests {
		fmt.Println(calPoints(test))
	}
}

// https://leetcode.com/problems/baseball-game/

/*
time complexity: O(n), 4 ms
Space complexity: O(n), 2.6 MB
*/

func calPoints(ops []string) (sum int) {
	stack := make([]int, len(ops))
	top := 0
	for _, op := range ops {
		switch op {
		case "D":
			stack[top] = 2 * stack[top-1]
			sum += stack[top]
			top++
		case "+":
			stack[top] = stack[top-1] + stack[top-2]
			sum += stack[top]
			top++
		case "C":
			sum -= stack[top-1]
			top--
		default:
			num, _ := strconv.Atoi(op)
			stack[top] = num
			sum += stack[top]
			top++
		}
	}
	return sum
}
