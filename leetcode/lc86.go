package main

import (
	. "github.com/wsmdr/goPractice/leetcode/ListNode"
	"fmt"
)

func main() {
	arr := []int{1,4,3,2,5,2}
	list := &LinkedList{}
	for _,v := range arr {
		list.AppendNodeElement(v)
	}
	PrintListNode(list.Head())
	fmt.Println("=============")
	node := partition(list.Head(), 3)
	PrintListNode(node)
}

func partition(head *ListNode, x int) *ListNode {
	leftHead := &ListNode{}
	rightHead := &ListNode{}
	left := leftHead
	right := rightHead

	for head != nil {
		var val int
		switch v := head.Val.(type) {
		case int:
			val = v
		}
		if val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = head.Next
	}
	left.Next = rightHead.Next
	right.Next = nil
	return leftHead.Next
}
