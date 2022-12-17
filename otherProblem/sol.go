package sol

import "sort"

/*

Write a function:

func Solution(A []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000]. */

func Solution(A []int) int {
	response := 1

	sort.Ints(A)
	for _, num := range A {
		if num <= 0 {
			continue
		}
		if num == response {
			response++
		} else if num < response {
			continue
		} else {
			return response
		}
	}

	return response
}
