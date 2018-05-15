// 链表是否有环
package main

import (
	. "github.com/wsmdr/goPractice/leetcode/ListNode"
	"fmt"
)

func main() {
	list := &LinkedList{}

	for i := 1; i < 8; i++ {
		list.AppendNodeElement(i)
	}

	node1 := list.Index(3)

	node2 := list.Index(list.Len())
	node2.Next = node1

	cycle, ok := detectCycle(list.Head())
	if ok {
		fmt.Printf("h.val: %v - %v, h.Next address: %v \n", cycle.Val, &cycle.Val, &cycle.Next)
	} else {
		fmt.Println("链表无环")
	}
}

// 套圈
func detectCycle(head *ListNode) (*ListNode, bool) {
	fast := head
	slow := head
	var meet *ListNode
	for {
		fast = fast.Next
		slow = slow.Next
		if fast == nil {
			return nil,false
		}
		fast = fast.Next
		if fast == slow {
			meet = fast
			break
		}
	}
	if meet == nil {
		return nil, false
	}
	for {
		if head == meet {
			break
		}
		head = head.Next
		meet = meet.Next
	}
	return meet, true
}