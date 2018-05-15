package main

import "fmt"

func calculate(s string) int {
	operate := 1
	result := 0
	stack := make([]int, 0, len(s))

	for i := 0; i < len(s);i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num := 0
			for ;i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = 10*num + int(s[i] - '0')
			}
			result += operate*num
			// i 退位
			i--
		case '+':
			operate = 1
		case '-':
			operate = -1
		case '(':
			stack = append(stack, result, operate)
			result = 0
			operate = 1
		case ')':
			operate = stack[len(stack) -1]
			plus := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result = operate*result + plus
		}
	}

	return result
}

func main() {
	res := calculate("(1+(4+5+2)-3)+(6+8)")
	fmt.Println(res)
}


