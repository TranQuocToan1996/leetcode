package hard

import (
	"math"
	"sort"
	"testing"
)

// https://leetcode.com/problems/median-of-two-sorted-arrays/

func TestFindMedianSortedArrays(t *testing.T) {
	for _, test := range []struct {
		nums1  []int
		nums2  []int
		expect float64
	}{
		{nums1: []int{1, 2}, nums2: []int{3, 4}, expect: 2.5},
		{nums1: []int{1, 3}, nums2: []int{2}, expect: 2},
		{nums1: []int{}, nums2: []int{1}, expect: 1},
		{nums1: []int{1, 3}, nums2: []int{2, 7}, expect: 2.5},
	} {
		if !almostEqualFloat(findMedianSortedArrays(test.nums1, test.nums2), test.expect) {
			t.Errorf("expect %v but got %v", test.expect, findMedianSortedArrays(test.nums1, test.nums2))
		}
	}
}

// https://stackoverflow.com/questions/47969385/go-float-comparison
func almostEqualFloat(a, b float64) bool {
	const float64EqualityThreshold = 1e-9
	return math.Abs(a-b) <= float64EqualityThreshold
}

// Runtime: 31 ms, faster than 35.60% of Go online submissions for Median of Two Sorted Arrays.
// Memory Usage: 6.1 MB, less than 26.56% of Go online submissions for Median of Two Sorted Arrays.
// Time: mostly time base on the sort.Int (O((n +m)*log(n + m)))
// Space: O(n + m)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	length := len(nums1)
	if length%2 == 0 {
		return float64(nums1[length/2]+nums1[length/2-1]) / 2
	}
	return float64(nums1[length/2])
}

/* Cleaner ways
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    nums1 = append(nums1, nums2...)
    sort.Ints(nums1)
    return float64(nums1[len(nums1)>>1] + nums1[(len(nums1) - 1)>>1]) / 2.0
} */

/*
    Runtime: 8 ms, faster than 97.99% of Go online submissions for Median of Two Sorted Arrays.
    Memory Usage: 5.1 MB, less than 86.49% of Go online submissions for Median of Two Sorted Arrays.
    O(lg(n + m)) : Even better ways. Not yet understand, comment for future review.

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    l := len(nums1) + len(nums2)
    if l % 2 == 0 {
        // find mid and mid - 1
        v1 := findKth(nums1, 0, nums2, 0, l / 2)
        v2 := findKth(nums1, 0, nums2, 0, (l + 2) / 2)
        return (float64(v1) + float64(v2)) / 2.0
    } else {
        // find mid
        return float64(findKth(nums1, 0, nums2, 0, (l + 1) / 2))
    }
}

func findKth(A []int, pa int, B []int, pb int, k int) int {
    // A: [pa..len(A) - 1]
    // B: [pb..len(B) - 1]
    // k: kth [1..n], starting from 1
    m, n := len(A), len(B)
    if pa >= m {return B[pb + k - 1]}
    if pb >= n {return A[pa + k - 1]}

    if k == 1 {return int(math.Min(float64(A[pa]), float64(B[pb])))}
    va := math.MaxInt64
    if pa + k / 2 - 1 < m {
        va = A[pa + k / 2 - 1]
    }
    vb := math.MaxInt64
    if pb + k / 2 - 1 < n {
        vb = B[pb + k / 2 - 1]
    }

    if va <= vb {
        return findKth(A, pa + k / 2, B, pb, k - k / 2)
	}else {
        return findKth(A, pa, B, pb + k / 2, k - k / 2)
    }
} */
