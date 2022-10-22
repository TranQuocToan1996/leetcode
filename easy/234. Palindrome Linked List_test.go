package easy

import (
	"leetcode/model"
	"testing"
)

func TestPalindromeLinkList(t *testing.T) {
	head1 := &model.ListNode{
		Val: 1,
		Next: &model.ListNode{
			Val: 2,
			Next: &model.ListNode{
				Val: 2,
				Next: &model.ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	head2 := &model.ListNode{
		Val: 1,
		Next: &model.ListNode{
			Val:  2,
			Next: nil,
		},
	}

	if !isPalindrome1(head1) || isPalindrome1(head2) {
		t.Fatal("fail")
	}

}

func isPalindrome1(head *model.ListNode) bool {
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
