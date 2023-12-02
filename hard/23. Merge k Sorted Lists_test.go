package hard

import (
	"reflect"
	"testing"

	"leetcode/model"
)

// Time: O(n^2 * logn)
// Space: O(n)
// func mergeKLists(lists []*model.ListNode) *model.ListNode {
// 	n := len(lists)
// 	var (
// 		h   = model.IntHeap{}
// 		cur *model.ListNode
// 	)
// 	// Time O(n)
// 	for i := 0; i < n; i++ {
// 		cur = lists[i]
// 		for cur != nil {
// 			// Time: O(nlogn)
// 			heap.Push(&h, cur.Val)
// 			cur = cur.Next
// 		}
// 	}
// 	return buildSortedListFromHeap(&h)
// }

// func buildSortedListFromHeap(h heap.Interface) *model.ListNode {
// 	root := &model.ListNode{}
// 	cur := root
// 	for h.Len() > 0 {
// 		node := &model.ListNode{Val: heap.Pop(h).(int)}
// 		cur.Next = node
// 		cur = cur.Next
// 	}
// 	return root.Next
// }

// Time: O(nlogn)
// Space: O(logn) for calltack. No extra space for node.
func mergeKLists(lists []*model.ListNode) *model.ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) >= 2 {
		// pop 2 lists
		l1, l2 := lists[0], lists[1]
		lists = lists[2:]

		merged := mergeTwoListNode(l1, l2)
		lists = append(lists, merged)
	}
	return lists[0]
}

func mergeTwoListNode(node1, node2 *model.ListNode) *model.ListNode {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	if node1.Val < node2.Val {
		node1.Next = mergeTwoListNode(node1.Next, node2)
		return node1
	} else {
		node2.Next = mergeTwoListNode(node1, node2.Next)
		return node2
	}
}

func TestMergeKLists(t *testing.T) {
	for _, test := range []struct {
		name   string
		lists  []*model.ListNode
		expect *model.ListNode
	}{
		{
			name: "Normal",
			lists: []*model.ListNode{
				{Val: 1, Next: &model.ListNode{Val: 4, Next: &model.ListNode{Val: 5}}},
				{Val: 1, Next: &model.ListNode{Val: 3, Next: &model.ListNode{Val: 4}}},
				{Val: 2, Next: &model.ListNode{Val: 6}},
			},
			expect: &model.ListNode{Val: 1, Next: &model.ListNode{Val: 1, Next: &model.ListNode{Val: 2, Next: &model.ListNode{Val: 3, Next: &model.ListNode{
				Val:  4,
				Next: &model.ListNode{Val: 4, Next: &model.ListNode{Val: 5, Next: &model.ListNode{Val: 6}}},
			}}}}},
		},
		{
			name:   "A nil node",
			lists:  []*model.ListNode{nil},
			expect: nil,
		},
		{
			name:   "No node",
			lists:  []*model.ListNode{},
			expect: nil,
		},
		{
			name:   "Two nil nodes",
			lists:  []*model.ListNode{nil, nil},
			expect: nil,
		},
		{
			name:   "First node nil, Second node normal",
			lists:  []*model.ListNode{nil, {Val: 1}},
			expect: &model.ListNode{Val: 1},
		},
	} {
		res := mergeKLists(test.lists)
		if !reflect.DeepEqual(res, test.expect) {
			t.Errorf("[%v] expect %v but got %v, test %v", test.name, test.expect, res, test)
		}
	}
}
