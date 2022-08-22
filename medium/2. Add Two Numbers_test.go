package easy

// https://leetcode.com/problems/add-two-numbers/

/*
Time complexity: O(n), 	21 ms
Space complexity: O(n), 4.6 MB
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carryNum := 0
	l3 := new(ListNode)
	curent := l3

	for l1 != nil || l2 != nil || carryNum != 0 {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		val := val1 + val2 + carryNum
		carryNum = val / 10
		curent.Next = &ListNode{Val: val % 10, Next: nil}
		curent = curent.Next
	}

	return l3.Next
}
