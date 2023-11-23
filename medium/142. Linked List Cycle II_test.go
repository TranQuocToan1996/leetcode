package medium

import "testing"

func detectCycle(head *ListNode) *ListNode {
	return nil
}

func TestDetectCycle(t *testing.T) {
	for _, test := range []struct {
		start             *ListNode
		setupAndGetOutput func(start *ListNode) *ListNode
	}{
		{
			start: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			setupAndGetOutput: func(start *ListNode) *ListNode {
				begin := start.Next
				start.Next.Next.Next.Next = begin
				return begin
			},
		},
		{
			start: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			setupAndGetOutput: func(start *ListNode) *ListNode {
				start.Next.Next = start
				return start.Next.Next
			},
		},
		{
			start: &ListNode{
				Val: 1,
			},
			setupAndGetOutput: func(start *ListNode) *ListNode {
				return nil
			},
		},
	} {
		output := test.setupAndGetOutput(test.start)
		res := detectCycle(test.start)
		if res != output {
			t.Errorf("expect %v but got %v", output, res)
		}
	}
}
