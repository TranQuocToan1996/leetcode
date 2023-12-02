package hard

import (
	"reflect"
	"testing"

	"leetcode/model"
)

func reverseKGroup(head *model.ListNode, k int) *model.ListNode {
	node, count := head, 0
	for count < k {
		if node == nil {
			return head
		}
		node = node.Next
		count++
	}
	prev := reverseKGroup(node, k)
	for count > 0 {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
		count--
	}

	return prev
}

func TestReverseKGroup(t *testing.T) {
	for _, test := range []struct {
		head   *model.ListNode
		k      int
		output *model.ListNode
	}{
		{
			head: &model.ListNode{
				Val: 1,
				Next: &model.ListNode{
					Val: 2,
					Next: &model.ListNode{
						Val: 3,
						Next: &model.ListNode{
							Val: 4,
							Next: &model.ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			k: 2,
			output: &model.ListNode{
				Val: 2,
				Next: &model.ListNode{
					Val: 1,
					Next: &model.ListNode{
						Val: 4,
						Next: &model.ListNode{
							Val: 3,
							Next: &model.ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		{
			head: &model.ListNode{
				Val: 1,
				Next: &model.ListNode{
					Val: 2,
					Next: &model.ListNode{
						Val: 3,
						Next: &model.ListNode{
							Val: 4,
							Next: &model.ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			k: 3,
			output: &model.ListNode{
				Val: 3,
				Next: &model.ListNode{
					Val: 2,
					Next: &model.ListNode{
						Val: 1,
						Next: &model.ListNode{
							Val: 4,
							Next: &model.ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
	} {
		res := reverseKGroup(test.head, test.k)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
