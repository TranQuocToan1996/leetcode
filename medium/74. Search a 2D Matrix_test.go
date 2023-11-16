package medium

import (
	"reflect"
	"testing"
)

// You must solve this problem without using the library's sort function.
// nums[i] only 0,1,2
// Time: O(n)
// Space: O(1), due to limit key of the map
func sortColors(nums []int) {
	n := len(nums)
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	numWrite := 0
	for i := 0; i < n && len(count) > 0; i++ {
		for count[numWrite] <= 0 {
			numWrite++
		}
		nums[i] = numWrite
		count[numWrite]--
	}
}

func TestSortColors(t *testing.T) {
	for _, test := range []struct {
		input  []int
		expect []int
	}{
		{
			input:  []int{2, 0, 2, 1, 1, 0},
			expect: []int{0, 0, 1, 1, 2, 2},
		},
		{
			input:  []int{2, 0, 1},
			expect: []int{0, 1, 2},
		},
		{
			input:  []int{0},
			expect: []int{0},
		},
	} {
		sortColors(test.input)
		if !reflect.DeepEqual(test.input, test.expect) {
			t.Errorf("expect %v but got %v", test.expect, test.input)
		}
	}
}
