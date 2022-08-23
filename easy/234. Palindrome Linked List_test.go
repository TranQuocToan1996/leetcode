package easy

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestPalindromeLinkList(t *testing.T) {
	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	head2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	if !isPalindrome1(head1) || isPalindrome1(head2) {
		t.Fatal("fail")
	}

}

func isPalindrome1(head *ListNode) bool {
	vals := []int{}
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	left, right := 0, len(vals)-1
	for left < right {
		if vals[left] != vals[right] {
			return false
		}
		left++
		right--
	}
	return true
}
