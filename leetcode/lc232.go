package main

import "fmt"

type MyQueue struct {
	a,b *Stack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{a: NewStack(), b: NewStack()}
}


/** Push element x to the back of queue. */
func (mq *MyQueue) Push(x int)  {
	mq.b.Push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (mq *MyQueue) Pop() int {
	if mq.a.Len() == 0 {
		for mq.b.Len() >0 {
			mq.a.Push(mq.b.Pop())
		}
	}
	return mq.a.Pop()
}


/** Get the front element. */
func (mq *MyQueue) Peek() int {
	r := mq.Pop()
	mq.a.Push(r)
	return r
}


/** Returns whether the queue is empty. */
func (mq *MyQueue) Empty() bool {
	return mq.a.Len() + mq.b.Len() == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)

	a := obj.Peek()
	b := obj.Peek()
	fmt.Println(a,b)

	fmt.Printf("a:%v=b:%v", obj.a, obj.b)
}

func NewStack() *Stack {
	return &Stack{[]int{}}
}

type Stack struct {
	items []int
}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) Pop() int {
	r := s.Top()
	s.items = s.items[0:s.Len()-1]
	return r
}

func (s *Stack) Top() int {
	return s.items[s.Len()-1]
}