package medium

import (
	"reflect"
	"testing"
)

// Time: O(log(n))
// Time: O(1)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow.Next
	slow.Next = nil
	left, right := sortList(head), sortList(mid)

	return mergeSortList(left, right)
}

func mergeSortList(left, right *ListNode) *ListNode {
	head := &ListNode{}
	temp := head

	for left != nil && right != nil {
		if left.Val < right.Val {
			head.Next = left
			left = left.Next
		} else {
			head.Next = right
			right = right.Next
		}
		head = head.Next
	}

	for left != nil {
		head.Next = left
		head = head.Next
		left = left.Next
	}
	for right != nil {
		head.Next = right
		head = head.Next
		right = right.Next
	}

	return temp.Next
}

func TestSortList(t *testing.T) {
	for _, test := range []struct {
		start  *ListNode
		output *ListNode
	}{
		{
			start: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
		{
			start: &ListNode{
				Val: -1,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 0,
							},
						},
					},
				},
			},
			output: &ListNode{
				Val: -1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		{
			start:  nil,
			output: nil,
		},
	} {
		res := sortList(test.start)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, testInput %v ", test.output, res, test.start)
		}
	}
}
