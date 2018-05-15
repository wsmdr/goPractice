// 判断链表相交
package main

import (
	. "github.com/wsmdr/goPractice/leetcode/ListNode"
	"fmt"
)

func main() {
	list := &LinkedList{}
	list1 := &LinkedList{}
	list2 := &LinkedList{}

	list.AppendNodeElement("c1")
	list.AppendNodeElement("c2")
	list.AppendNodeElement("c3")
	fmt.Println(list.Len())
	list1.AppendNodeElement("a1")
	list1.AppendNodeElement("a2")
	list1.AppendNode(list.Head())
	fmt.Println(list1.Len())

	list2.AppendNodeElement("b1")
	list2.AppendNodeElement("b2")
	list2.AppendNodeElement("b3")
	list2.AppendNode(list.Head())
	fmt.Println(list2.Len())
	fmt.Println("list===")
	PrintListNode(list.Head())
	fmt.Println("list1===")
	PrintListNode(list1.Head())
	fmt.Println("list2===")
	PrintListNode(list2.Head())

	h := getIntersectionNode(list1.Head(), list2.Head())

	fmt.Printf("h.val: %v, h.Next address: %v \n", &h.Val, &h.Next)
}

func getNodeLen(node *ListNode) int {
	var len int
	for node != nil {
		len++
		node = node.Next
	}
	return len
}

func changeHeadNode(m int, head *ListNode) *ListNode {
	for m > 0 {
		head = head.Next
		m--
	}
	return head
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := getNodeLen(headA)
	lenB := getNodeLen(headB)

	if lenA > lenB {
		headA = changeHeadNode(lenA - lenB, headA)
	} else {
		headB = changeHeadNode(lenB - lenA, headB)
	}

	for {
		headA = headA.Next
		headB = headB.Next
		if headA == headB {
			return headB
		}
		if headA == nil || headB == nil {
			return nil
		}
	}
}