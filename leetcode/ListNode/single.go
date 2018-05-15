package ListNode

import "fmt"

type ListNode struct {
	Val interface{}
	Next *ListNode
}

type LinkedList struct {
	length int
	head *ListNode
}

func (list *LinkedList) init() {
	list.length = 0
	list.head = nil
}

// append node to the end of LinkedList
func (list *LinkedList) AppendNode(node *ListNode) bool {
	if node == nil {
		return false
	}
	if list.head == nil {
		list.head = node
	} else {
		current := list.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
	for node != nil {
		list.length++
		node = node.Next
	}
	return true
}

// append node's val to the end of LinkedList
func (list *LinkedList) AppendNodeElement(element interface{}) bool {
	node := &ListNode{element, nil}
	if list.head == nil {
		list.head = node
	} else {
		current := list.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
	list.length++

	return true
}

// insert node to the position
func (list *LinkedList) Insert(element interface{}, position int) bool {
	if position > list.length || position < 0 {
		return false
	}
	var node = &ListNode{element, nil}
	current := list.head
	if position == 0 {
		node.Next = current
		list.head = node
	} else {
		var pre *ListNode
		for i := 1; i <= position; i++ {
			pre = current
			current = current.Next
		}
		node.Next = current
		pre.Next = node
	}
	list.length++

	return true
}

func (list *LinkedList) Index(position int) *ListNode {
	if position > list.length || position < 1 {
		return nil
	}
	current := list.head
	for i := 2; i <= position; i++ {
		current = current.Next
	}
	return current
}

// remove the ListNode
func (list *LinkedList) Remove(position int) bool {
	if position >= list.length || position < 0 {
		return false
	}
	current := list.head
	if position == 0 {
		list.head = current.Next
	} else {
		var pre, next *ListNode
		for i := 1; i <= position; i++ {
			pre = current
			next = pre.Next
			current = next
		}
		pre.Next = next.Next
	}
	return true
}

// LinkedList is empty
func (list LinkedList) Empty() bool {
	if list.length == 0 {
		return true
	}
	return false
}

// get the LinkedList's length
func (list LinkedList) Len() int {
	return list.length
}

// get the LinkedList's first node
func (list LinkedList) Head() *ListNode {
	return list.head
}

// print ListNode
func PrintListNode(h *ListNode) {
	for h != nil {
		fmt.Printf("node.val: %v - %v, node.Next: %v \n", h.Val, &h.Val, &h.Next)
		h = h.Next
	}
}