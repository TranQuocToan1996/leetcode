package easy

import (
	"fmt"
)

func main() {
	type twoArr struct {
		target []int
		arr    []int
	}

	tests := []twoArr{
		{[]int{1, 2, 3, 4}, []int{2, 4, 1, 3}},
		{[]int{7}, []int{7}},
		{[]int{3, 7, 9}, []int{3, 7, 11}},
	}
	for _, test := range tests {
		fmt.Println(canBeEqual(test.target, test.arr))
	}
}

// https://leetcode.com/problems/make-two-arrays-equal-by-reversing-sub-arrays/

/*
time complexity: I don't know, 3 ms
Space complexity: I don't know neither, 4.3 MB
*/
/* func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	sort.Ints(target)
	sort.Ints(arr)
	return reflect.DeepEqual(target, arr)
}

*/

/*
time complexity: O(n), 4 ms
Space complexity: O(1), 4 MB
*/
func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	count := [1001]int{}
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
		count[target[i]]--
	}
	for i := 0; i < len(count); i++ {
		if count[i] != 0 {
			return false
		}
	}
	return true
}
