package easy

import "fmt"

// Method 1: runtime 0ms, memory 2.1 MB
func plus(num int) int {
	if num <= 0 {
		return num
	}
	return num + plus(num-1)
}

func numIdenticalPairs(nums []int) int {
	pair := make(map[int]int)
	for _, v := range nums {
		pair[v]++
	}
	fmt.Println(pair)
	var result int = 0
	for _, v := range pair {
		result += plus(v - 1)
	}
	return result
}

/* // Method 2: runtime 3ms, memory: 2.1 MB
func numIdenticalPairs(nums []int) int {
	result := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				result++
			}
		}
	}
	return result
} */

// Method 3: runtime 0ms, memory 2.1MB
// Do phuc tap thoi gian do duyet qua mang 1 lan O(n), do phuc tap khong gian trong truong hop n so khac nhau O(n)
/* func numIdenticalPairs(nums []int) int {
	result := 0
	countMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		count, isExist := countMap[nums[i]]
		if isExist {
			result += count
			countMap[nums[i]] = count + 1
		} else {
			countMap[nums[i]] = 1
		}
	}
	return result
} */
