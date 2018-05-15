package main

type MyStack struct {
	q *Queue
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{q: NewQueue() }
}

/** Push element x onto stack. */
func (ms *MyStack) Push(x int)  {
	if ms.Empty() {
		ms.q.Push(x)
	} else {
		tmp := NewQueue()
		tmp.Push(x)
		for ms.q.Len() > 0 {
			tmp.Push(ms.q.Pop())
		}
		for tmp.Len() > 0 {
			ms.q.Push(tmp.Pop())
		}
	}

}


/** Removes the element on top of the stack and returns that element. */
func (ms *MyStack) Pop() int {
	return ms.q.Pop()
}

/** Get the top element. */
func (ms *MyStack) Top() int {
	return ms.q.Peek()
}


/** Returns whether the stack is empty. */
func (ms *MyStack) Empty() bool {
	return ms.q.IsEmpty()
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func NewQueue() *Queue {
	return &Queue{[]int{}}
}

type Queue struct {
	items []int
}

func (q *Queue) Len() int {
	return len(q.items)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Push(x int) {
	q.items = append(q.items, x)
}

func (q *Queue) Pop() int {
	r := q.Peek()
	q.items = q.items[1:]
	return r
}

func (q *Queue) Peek() int {
	return q.items[0]
}