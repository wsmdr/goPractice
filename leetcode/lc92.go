// 交换链表指定位置
package main

import . "github.com/wsmdr/goPractice/leetcode/ListNode"

func main() {
	list := &LinkedList{}
	for i := 1; i < 6; i++ {
		list.AppendNodeElement(i)
	}
	h := reverseBetween(list.Head(), 2, 4)
	PrintListNode(h)
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	var newHead, newNode *ListNode
	result := head

	changeLen := n - m + 1
	for m > 1 {
		newHead = head
		head = head.Next
		m--
	}
	headNode := head
	for changeLen > 0 {
		next := head.Next
		head.Next = newNode
		newNode = head
		head = next
		changeLen--
	}
	headNode.Next = head
	if newHead != nil {
		newHead.Next = newNode
	} else {
		result = newNode
	}
	return result
}
