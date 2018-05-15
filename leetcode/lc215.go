package main

import "container/heap"

func findKthLargest(nums []int, k int) int {
	tmp := highHeap(nums)
	h := &tmp
	heap.Init(h)
	if k == 1 {
		return (*h)[0]
	}
	for i := 1; i < k; i++ {
		heap.Remove(h, 0)
	}
	return (*h)[0]
}

type highHeap []int

func (h highHeap) Len() int {
	return len(h)
}

func (h highHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h highHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *highHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *highHeap) Pop() interface{} {
	r := (*h)[len(*h)-1]
	*h = (*h)[0:len(*h) -1]
	return r
}
