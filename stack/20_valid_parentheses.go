package stack

func isValid(s string) bool {
	stack := New[rune]()
	for _, char := range s {
		if char == '{' || char == '[' || char == '(' {
			stack.Push(char)
			continue
		}
		if char == '}' {
			if parantheses := stack.Pop(); parantheses != '{' {
				return false
			}
		}
		if char == ']' {
			if parantheses := stack.Pop(); parantheses != '[' {
				return false
			}
		}
		if char == ')' {
			if parantheses := stack.Pop(); parantheses != '(' {
				return false
			}
		}
	}
	return len(stack) <= 0
}
