package stackqueue

func isValid(s string) bool {
	stack := Stack[rune]{}

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack.Push(c)
			continue
		}
		pop, ok := stack.Pop()
		if !ok {
			return false
		}

		if c == ')' && pop != '(' {
			return false
		}
		if c == ']' && pop != '[' {
			return false
		}

		if c == '}' && pop != '{' {
			return false
		}
	}

	return stack.IsEmpty()
}

func isValid2(s string) bool {
	stack := []rune{}
	closers := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	for _, c := range s {
		open, ok := closers[c]
		if ok {
			stack = append(stack, open)
			continue
		}
		if len(stack) == 0 {
			return false
		}

		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop != c {
			return false
		}
	}

	return len(stack) == 0
}
