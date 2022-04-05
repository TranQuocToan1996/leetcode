package main

import (
	"fmt"
)

func main() {

	tests := [][]int{
		{3, 4, 5, 2},
		{1, 5, 4, 5},
		{3, 7},
	}
	for _, test := range tests {
		fmt.Println(maxProduct(test))
	}
}

// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/

/*
time complexity: O(n) increases time with the len of arr, 4 ms
Space complexity: O(1), needs probably a fix amount of memory for max1 and max2, 3 MB
*/
func maxProduct(nums []int) int {
	max1 := 0
	max2 := 0

	for i := range nums {
		if nums[i] >= max1 {
			max2 = max1
			max1 = nums[i]
			continue
		}

		if nums[i] > max2 {
			max2 = nums[i]
		}
	}

	return (max1 - 1) * (max2 - 1)
}

/*
time complexity: O(n^2) 2 loops, 5 ms
Space complexity: O(1), 3 MB
*/
/* func maxProduct(nums []int) int {
	var maxMultiply int
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			multiply := (nums[i] - 1) * (nums[j] - 1)
			if maxMultiply < multiply {
				maxMultiply = multiply
			}
		}
	}
	return maxMultiply
} */

/* Looks like using merge sort and sort.Ints() not so good.
func maxProduct(nums []int) int {
	nums = mergeSort(nums)
	fmt.Println(nums)
	length := len(nums)
	leftMultiply := (nums[0] - 1) * (nums[1] - 1)
	rightMultiply := (nums[length-1] - 1) * (nums[length-2] - 1)
	if leftMultiply < rightMultiply {
		return rightMultiply
	}
	return leftMultiply
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
*/
