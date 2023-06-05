package stack

import "strconv"

func evalRPN(tokens []string) int {
	stack := New[int]()
	for _, token := range tokens {
		switch token {
		case "+":
			firstVal := stack.Pop()
			secondVal := stack.Pop()
			stack.Push(firstVal + secondVal)
		case "-":
			firstVal := stack.Pop()
			secondVal := stack.Pop()
			stack.Push(secondVal - firstVal)
		case "*":
			firstVal := stack.Pop()
			secondVal := stack.Pop()
			stack.Push(firstVal * secondVal)
		case "/":
			firstVal := stack.Pop()
			secondVal := stack.Pop()
			stack.Push(secondVal / firstVal)
		default:
			val, _ := strconv.Atoi(token)
			stack.Push(val)

		}
	}
	return stack.Peek()
}
