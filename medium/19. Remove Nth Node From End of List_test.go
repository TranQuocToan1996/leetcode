package medium

import (
	"fmt"
	"leetcode/uti"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func TestRemoveNthFromEnd(t *testing.T) {
	defer uti.LogRuntime(time.Now())
	for _, test := range []struct {
		head   *ListNode
		n      int
		output *ListNode
	}{
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
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
			n: 2,
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
		{
			head:   &ListNode{Val: 1},
			n:      1,
			output: nil,
		},
		{
			head:   &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			n:      1,
			output: &ListNode{Val: 1},
		},
	} {
		res := removeNthFromEnd(test.head, test.n)
		assert.Equal(t, test.output, res)
	}
}

// Change the return
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// uti.LogRuntime base on concurrency, so that best and worst big difference.
	// Runtime: O(n) 2ms (53.83%): Need to traveral all singly linked list once, faster than normal because the channel send signal complete before all callstack done
	// Space: O(1) 2.2 MB (30.28%): only allocate for right, left, isEnd, channel. Callstack O(height(head))
	return removeNthFromEnd_recursiveWithChannel(head, n)

	// uti.LogRuntime 35958 nsec
	// Runtime: O(n) 2ms (53.83%): Need to traveral all singly linked list once
	// Space: O(1) 2.2 MB (30.28%): only allocate for right, left, isEnd
	// return removeNthFromEnd_twoPointersSameRate(head, n)

	// uti.LogRuntime 37292 nsec
	// Runtime: O(n) 2ms (53.83%): Need to traveral all singly linked list once
	// Space: O(1) 2.2 MB (30.28%): only allocate for right, left, isEnd. Callstack O(height(head))
	// return removeNthFromEnd_recursive(head, n)

}

func removeNthFromEnd_twoPointersSameRate(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}
	left, right := head, head

	var isEnd = false
	for i := 0; i < n; i++ {
		if right.Next == nil {
			isEnd = true
			break
		}
		right = right.Next
	}

	if isEnd {
		return head.Next
	} else {
		for right.Next != nil {
			right = right.Next
			left = left.Next
		}
		left.Next = left.Next.Next
	}

	return head
}

func removeNthFromEnd_recursiveWithChannel(head *ListNode, n int) *ListNode {
	prehead := &ListNode{
		Next: head,
	}
	done := make(chan struct{})
	traversalWithChannel(prehead, n, done)
	<-done
	return prehead.Next
}

func traversalWithChannel(node *ListNode, n int, res chan struct{}) int {
	if node == nil {
		return 1
	}
	place := traversalWithChannel(node.Next, n, res)
	if place == n+1 {
		node.Next = node.Next.Next
		close(res)
	}
	fmt.Println(place)
	return place + 1
}

func removeNthFromEnd_recursive(head *ListNode, n int) *ListNode {
	prehead := &ListNode{
		Next: head,
	}
	traversal(prehead, n)
	return prehead.Next
}

func traversal(node *ListNode, n int) int {
	if node == nil {
		return 1
	}
	place := traversal(node.Next, n)
	if place == n+1 {
		node.Next = node.Next.Next
	}
	return place + 1
}
