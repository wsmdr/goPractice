package main

type MinStack struct {
	s []item
}

type item struct {
	min, data int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (ms *MinStack) Push(x int)  {
	min := x
	if len(ms.s) > 0 && ms.GetMin() < x {
		min = ms.GetMin()
	}
	ms.s = append(ms.s, item{min, x})
}


func (ms *MinStack) Pop() {
	ms.s = ms.s[:len(ms.s)-1]
}

func (ms *MinStack) Top() int {
	return ms.s[len(ms.s)-1].data
}


func (ms *MinStack) GetMin() int {
	return ms.s[len(ms.s)-1].min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
