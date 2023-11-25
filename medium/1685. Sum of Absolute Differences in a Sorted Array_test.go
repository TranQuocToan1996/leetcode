package medium

import (
	"reflect"
	"testing"
)

// func getSumAbsoluteDifferences(nums []int) []int {
// 	var (
// 		n         = len(nums)
// 		prefixSum = make([]int, n)
// 		res       = make([]int, n)
// 	)

// 	prefixSum[0] = nums[0]
// 	for i := 1; i < n; i++ {
// 		prefixSum[i] = prefixSum[i-1] + nums[i]
// 	}

// 	for i := 0; i < n; i++ {
// 		res[i] = nums[i]*(-n+2*i+2) - 2*prefixSum[i] + prefixSum[n-1]
// 	}

// 	return res
// }

func getSumAbsoluteDifferences(nums []int) []int {
	var (
		n   = len(nums)
		res = make([]int, n)
	)

	total := nums[0]
	for i := 1; i < n; i++ {
		total += nums[i]
	}

	cur := 0
	for i := 0; i < n; i++ {
		cur += nums[i]
		// leftSum := nums[i]*(i+1) - cur
		// rightSum := prefixSum[n-1] - cur - total
		// res[i] = rightSum + leftSum
		res[i] = nums[i]*(-n+2*i+2) - 2*cur + total
	}

	return res
}

func TestGetSumAbsoluteDifferences(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output []int
	}{
		{nums: []int{2, 3, 5}, output: []int{4, 3, 5}},
		{nums: []int{1, 4, 6, 8, 10}, output: []int{24, 15, 13, 15, 21}},
	} {
		res := getSumAbsoluteDifferences(test.nums)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
