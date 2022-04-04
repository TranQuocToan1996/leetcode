package main

import (
	"fmt"
)

func main() {

	fmt.Println(xorOperation(5, 0))
	fmt.Println(xorOperation(4, 3))
}

// https://leetcode.com/problems/xor-operation-in-an-array/
/*
time complexity: O(n) , 0 ms
Space complexity: O(1) , 1.9 MB
*/
func xorOperation(n int, start int) int {
	xor := start
	for i := 1; i < n; i++ {
		xor ^= (start + 2*i)
	}

	return xor
}
