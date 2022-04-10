package main

import (
	"fmt"
)

func main() {
	type Sum3 struct {
		arr    []int
		target int
	}

	tests := []Sum3{
		{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8},
		{[]int{1, 1, 2, 2, 2, 2}, 5},
	}
	for _, test := range tests {
		fmt.Println(threeSumMulti(test.arr, test.target))
		fmt.Print("\n")
	}
}

// https://leetcode.com/problems/3sum-with-multiplicity/

/*
time complexity: O(),  ms
Space complexity: O(),  MB
*/

// Dynamic Programing: Very good, I'm not yet understand fully this code, but this code make me learning basic concept about Dynamic Programing. I will come back oneday LOL
func threeSumMulti(arr []int, target int) (ans int) {

	//https://www.geeksforgeeks.org/modulo-1097-1000000007/
	const mod = int(1e9 + 7)

	vcnt := make([]int, target+1, target+1)
	twosumcnt := make([]int, target+1, target+1)

	for _, v := range arr {
		if v > target {
			continue
		}

		if ans += twosumcnt[target-v]; ans >= mod {
			ans -= mod
		}

		for oldv, cnt := range vcnt {
			if cnt > 0 {
				if twosum := oldv + v; twosum <= target {
					if twosumcnt[twosum] += cnt; twosumcnt[twosum] >= mod {
						twosumcnt[twosum] -= mod
					}
				}
			}
		}
		vcnt[v]++
	}

	return
}
