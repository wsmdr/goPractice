// 反转单链表
package main

import (
	. "github.com/wsmdr/goPractice/leetcode/ListNode"
	"fmt"
)

func main() {
	list := &LinkedList{}
	for i := 1; i < 6; i++ {
		list.AppendNodeElement(i)
	}
	PrintListNode(list.Head())
	fmt.Println("=========")
	newNode := reverseList(list.Head())

	PrintListNode(newNode)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newNode *ListNode
	for {
		next := head.Next
		head.Next = newNode
		newNode = head
		head = next
		if head == nil {
			break
		}
	}
	return newNode
}