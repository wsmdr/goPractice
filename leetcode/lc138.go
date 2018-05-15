package main

import "fmt"

type RandomListNode struct {
	Val int
	Random *RandomListNode
	Next *RandomListNode
}

func main() {
	a := &RandomListNode{1, nil, nil}
	b := &RandomListNode{2, nil, nil}
	c := &RandomListNode{3, nil, nil}
	d := &RandomListNode{4, nil, nil}
	e := &RandomListNode{5, nil, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	a.Random = c
	b.Random = d
	c.Random = c
	e.Random = d
	fmt.Println("old RandomListNode:")
	node := a
	for a != nil {
		fmt.Printf("%v\n", a)
		a = a.Next
	}

	fmt.Println("new RandomListNode:")
	n := copyRandomList(node)
	for n != nil {
		fmt.Printf("%v\n", n)
		n = n.Next
	}
}

func copyRandomList2(head *RandomListNode) *RandomListNode {

}


func copyRandomList(head *RandomListNode) *RandomListNode {
	nodeMap := make(map[*RandomListNode]int)
	newNodeMap := make(map[int]*RandomListNode)
	node := head
	i := 0
	for node != nil {
		nodeMap[node] = i
		newNode := &RandomListNode{node.Val, nil, nil}
		newNodeMap[i] = newNode
		node = node.Next
		i++
	}

	node = head
	i = 0
	for node != nil {
		newNodeMap[i].Next = newNodeMap[i+1]
		if node.Random != nil {
			id := nodeMap[node.Random]
			newNodeMap[i].Random = newNodeMap[id]
		}
		node = node.Next
		i++
	}

	return newNodeMap[0]
}