package easy

import (
	"fmt"
	"sort"
)

func main() {
	// first checking
	fmt.Println(canMakeArithmeticProgression([]int{3, 5, 1}))
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4}))

}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/submissions/
/* I'm not sure about the sort.Ints() complexity, but about the for loop below that:
time complexity: O(n)
Space complexity O(1)
*/
func canMakeArithmeticProgression(arr []int) bool {

	length := len(arr)
	if length < 3 {
		return true
	}
	// Sort array
	sort.Ints(arr)
	difference := arr[1] - arr[0]

	for i := 1; i < length; i++ {
		if difference != arr[i]-arr[i-1] {
			return false
		}
	}
	return true
}
