package main

import . "github.com/wsmdr/goPractice/leetcode/ListNode"

func main() {
	a1 := []int{1,2,4}
	a2 := []int{1,3,4}

	list1 := &LinkedList{}
	list2 := &LinkedList{}
	for _, v := range a1 {
		list1.AppendNodeElement(v)
	}
	for _, v := range a2 {
		list2.AppendNodeElement(v)
	}

	node := mergeTwoLists(list1.Head(), list2.Head())

	PrintListNode(node)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	head := node
	for l1 != nil && l2 != nil {
		var val1 int
		switch v := l1.Val.(type) {
		case int:
			val1 = v
		}
		var val2 int
		switch v := l2.Val.(type) {
		case int:
			val2 = v
		}
		if val1 < val2 {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	return head.Next
}
