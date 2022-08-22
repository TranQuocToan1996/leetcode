package easy

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {

	tests := []int{

		121,
		-121,
		10,
	}

	for _, test := range tests {
		fmt.Println(isPalindrome(test))
	}
}

// https://leetcode.com/problems/palindrome-number/

/*
time complexity: O(n), 16 ms
Space complexity: O(1), 4.5 MB
*/

func isPalindrome(x int) bool {
	q := x
	if x < 0 {
		return false
	}
	p := 0
	for ; q != 0; q /= 10 {
		last := q % 10
		p = p*10 + last

	}

	return p == x
}
