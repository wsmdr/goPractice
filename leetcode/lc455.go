package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var child,cookie int
	for child < len(g) && cookie < len(s) {
		if g[child] <= s[cookie] {
			child++
		}
		cookie++
	}
	return child
}
