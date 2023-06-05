package stack

import "strings"

func generateParenthesis(n int) []string {
	stack := NewStack[string]()
	res := make([]string, 0)
	backTrackParantheses(0, 0, n, &res, stack)
	return res
}

func backTrackParantheses(open, closed, n int, res *[]string, stack stack[string]) {
	if open == closed && open == n && closed == n {
		*res = append(*res, strings.Join(stack.GetAll(), ""))
		return
	}
	if open < n {
		stack.Push("(")
		backTrackParantheses(open+1, closed, n, res, stack)
		stack.Pop()
	}
	if closed < open {
		stack.Push(")")
		backTrackParantheses(open, closed+1, n, res, stack)
		stack.Pop()
	}
}
