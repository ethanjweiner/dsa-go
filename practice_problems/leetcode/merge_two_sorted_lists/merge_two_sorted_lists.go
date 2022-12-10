package practice_problems

type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursive Solution

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	switch {
// 	case list1 == nil:
// 		return list2
// 	case list2 == nil:
// 		return list1
// 	case list1.Val <= list2.Val:
// 		return &ListNode{Val: list1.Val, Next: mergeTwoLists(list1.Next, list2)}
// 	}

// 	return &ListNode{Val: list2.Val, Next: mergeTwoLists(list1, list2.Next)}
// }

// Iterative Solution, destructive

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{} // Keep `dummyHead` as reference for return
	currentNode := dummyHead

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			currentNode.Next = &ListNode{Val: list1.Val, Next: nil}
			list1 = list1.Next
		} else {
			currentNode.Next = &ListNode{Val: list2.Val, Next: nil}
			list2 = list2.Next
		}

		currentNode = currentNode.Next
	}

	if list1 == nil {
		currentNode.Next = list2
	}

	if list2 == nil {
		currentNode.Next = list1
	}

	return dummyHead.Next
}
