package easy

import (
	"leetcode/model"
	"reflect"
	"testing"
)

// https://leetcode.com/problems/merge-two-sorted-lists/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
// Memory Usage: 2.5 MB, less than 50.40% of Go online submissions for Merge Two Sorted Lists.

// Time: O(n + m): with n and m is the size if each linklist
// Memory: O(n + m): Have to allocation new memory to hold val from both linklist

func TestMergeTwoList(t *testing.T) {
	for _, test := range []struct {
		name     string
		list1    *model.ListNode
		list2    *model.ListNode
		expected *model.ListNode
	}{
		// {
		// 	name: "two list not empty",
		// 	list1: &model.ListNode{
		// 		Val: 1,
		// 		Next: &model.ListNode{
		// 			Val: 2,
		// 			Next: &model.ListNode{
		// 				Val: 4,
		// 			},
		// 		},
		// 	},
		// 	list2: &model.ListNode{
		// 		Val: 1,
		// 		Next: &model.ListNode{
		// 			Val: 3,
		// 			Next: &model.ListNode{
		// 				Val: 4,
		// 			},
		// 		},
		// 	},
		// 	expected: &model.ListNode{
		// 		Val: 1,
		// 		Next: &model.ListNode{
		// 			Val: 1,
		// 			Next: &model.ListNode{
		// 				Val: 2,
		// 				Next: &model.ListNode{
		// 					Val: 3,
		// 					Next: &model.ListNode{
		// 						Val: 4,
		// 						Next: &model.ListNode{
		// 							Val: 4,
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// },
		{
			name:     "two list empty",
			list1:    nil,
			list2:    nil,
			expected: nil,
		},
		{
			name:  "one list not empty",
			list1: nil,
			list2: &model.ListNode{
				Val: 0,
			},
			expected: &model.ListNode{
				Val: 0,
			},
		},
	} {
		expectList := listNodeToSlice(test.expected)
		mergeList := listNodeToSlice(mergeTwoLists(test.list1, test.list2))

		if !reflect.DeepEqual(expectList, mergeList) {
			t.Errorf("[%v] expect %v, merge %v", test.name, expectList, mergeList)
		}
	}
}

func listNodeToSlice(sll *model.ListNode) []int {
	list := []int{}
	cur := sll

	for cur != nil {
		list = append(list, cur.Val)
		cur = cur.Next
	}

	return list
}

/**
 * Definition for singly-linked list. I already move to model package.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func linkListSetVal(from *model.ListNode, to *model.ListNode) (fromNext *model.ListNode) {
	to.Val = from.Val
	return from.Next
}

func allocationListNode(list *model.ListNode) *model.ListNode {
	new := &model.ListNode{}
	if list == nil {
		list = new
		return list
	} else {
		list.Next = new
		return list.Next
	}
}

func mergeTwoLists(list1 *model.ListNode, list2 *model.ListNode) *model.ListNode {

	mergeList := allocationListNode(nil)
	cur := mergeList
	cur1 := list1
	cur2 := list2

	for cur1 != nil || cur2 != nil {
		cur = allocationListNode(cur)
		if cur1 != nil && cur2 != nil {
			if cur1.Val < cur2.Val {
				cur1 = linkListSetVal(cur1, cur)
			} else {
				cur2 = linkListSetVal(cur2, cur)
			}
		} else if cur1 != nil {
			cur1 = linkListSetVal(cur1, cur)
		} else {
			cur2 = linkListSetVal(cur2, cur)
		}
	}

	return mergeList.Next
}
