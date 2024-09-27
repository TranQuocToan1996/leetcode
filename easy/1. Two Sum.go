package easy

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for index, num := range nums {
		if v, ok := m[num]; ok {
			return []int{v, index}
		}
		m[target-num] = index
	}
	return []int{0, 1}
}
