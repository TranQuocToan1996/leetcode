package medium

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// type NodeIndex struct {
// 	node  *Node
// 	index int
// }

// Time: O(n)
// Space: O(n)
// func copyRandomList(head *Node) *Node {
// 	if head == nil {
// 		return nil
// 	}
// 	m := make(map[*Node]*NodeIndex)
// 	var newList []*Node

// 	cur := head
// 	index := 0
// 	for cur != nil {
// 		m[cur] = &NodeIndex{node: cur, index: index}
// 		newNode := &Node{
// 			Val: cur.Val,
// 		}
// 		if index != 0 {
// 			newList[len(newList)-1].Next = newNode
// 		}
// 		newList = append(newList, newNode)
// 		index++
// 		cur = cur.Next
// 	}

// 	cur = head
// 	index = 0
// 	for cur != nil {
// 		if cur.Random != nil {
// 			randomIndex := m[cur.Random].index
// 			newList[index].Random = newList[randomIndex]
// 		}

// 		index++
// 		cur = cur.Next
// 	}

// 	return newList[0]
// }

// Time: O(n)
// Space: O(1) if not count output
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	oldToNew := make(map[*Node]*Node)

	cur := head
	for cur != nil {
		oldToNew[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		oldToNew[cur].Next = oldToNew[cur.Next]
		oldToNew[cur].Random = oldToNew[cur.Random]
		cur = cur.Next
	}

	return oldToNew[head]
}
