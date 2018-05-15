package main

import (
	"sort"
	"fmt"
)

func subsets(nums []int) [][]int {
	res := make([][]int, 0)

	rescursion(nums, []int{}, &res)

	return res
}

var i int

func rescursion(nums, temp []int, res *[][]int) {
	fmt.Println(i)
	l := len(nums)
	i++
	if l == 0 {
		sort.Ints(temp)
		*res = append(*res, temp)
		return
	}
	rescursion(nums[:l-1], temp, res)

	rescursion(nums[:l-1], append([]int{nums[l-1]}, temp...), res)
}

func main() {
	num := []int{1,2,3}
	res := subsets(num)

	fmt.Println(res)
}
