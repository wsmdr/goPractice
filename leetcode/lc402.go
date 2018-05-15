package main

func removeKdigits(num string, k int) string {
	resNum := len(num) - k
	top := 0
	s := make([]byte, len(num))

	for v := range num {
		for k >0 && top > 0 && s[top-1] > num[v] {
			k--
			top--
		}
		s[top] = num[v]
		top++
	}
	i := 0
	for i < resNum && s[i] == '0' {
		i++
	}
	if i == resNum {
		return "0"
	}
	return string(s[i:resNum])
}
