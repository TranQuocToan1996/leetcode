package medium

import "testing"

// Time: O(n^2)
// Space O(n^2)
// func countNicePairs(nums []int) int {
// 	mem := make([][]int, len(nums)+1)
// 	for k := range mem {
// 		mem[k] = make([]int, len(nums)+1)
// 	}

// 	for i := 1; i < len(mem); i++ {
// 		mem[0][i] = nums[i-1]
// 		mem[i][0] = revInt(nums[i-1])
// 	}
// 	for i := 1; i < len(mem); i++ {
// 		for j := 1; j < len(mem[0]); j++ {
// 			mem[i][j] = mem[0][j] + mem[i][0]
// 		}
// 	}
// 	res := 0
// 	mod := int(1e9 + 7)
// 	for i := 1; i < len(mem); i++ {
// 		for j := 1; j < len(mem[0]); j++ {
// 			if i == j {
// 				continue
// 			}
// 			if mem[i][j] == mem[j][i] {
// 				res++
// 				res = res % mod
// 			}
// 		}
// 	}

// 	return res / 2
// }

// Time: O(n)
// Space O(n)
func countNicePairs(nums []int) int {
	const mod = int(1e9 + 7)
	var (
		n   = len(nums)
		res = 0
		mem = make(map[int]int)
		rev = make([]int, len(nums))
	)
	for i := 0; i < n; i++ {
		rev[i] = nums[i] - revInt(nums[i])
	}

	for i := 0; i < n; i++ {
		res = (res + mem[rev[i]]) % mod
		mem[rev[i]]++
	}

	return res
}

func revInt(num int) int {
	reversed := 0
	for num != 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}
	return reversed
}

func TestCountNicePairs(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{42, 11, 1, 97}, output: 2},
		{nums: []int{13, 10, 35, 24, 76}, output: 4},
	} {
		res := countNicePairs(test.nums)
		if res != test.output {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
