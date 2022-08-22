package easy

import (
	"fmt"
)

func main() {
	// first checking
	salaries := [][]int{
		{4000, 3000, 1000, 2000},
		{1000, 2000, 3000},
	}
	for _, salary := range salaries {
		fmt.Println(average(salary))
	}
}

// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
/*
time complexity: I don't really know due to the sort.Ints()
Space complexity: I don't really know due to the sort.Ints()
*/
/* func average(salary []int) float64 {
	sort.Ints(salary)
	// Remove the min, max salary
	salary = salary[1 : len(salary)-1]
	var sum float64
	for _, money := range salary {
		sum += float64(money)
	}

	return sum / float64(len(salary))
} */

/*
time complexity: O(n)
Space complexity: O(1)
*/
func average(salary []int) float64 {
	length := len(salary)
	min, max, total := 0, 0, 0
	for i := 0; i < length; i++ {
		if salary[i] > max {
			max = salary[i]
		}
		if salary[i] < min {
			min = salary[i]
		}
		total += salary[i]
	}
	total -= max + min
	return float64(total) / float64(length)
}
