func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{')': '(', '}': '{', ']': '['}
	chrs := []rune(s)

	for _, c := range chrs {
		if open, ok := mapping[c]; ok {
			if len(stack) == 0 || open != stack[len(stack) - 1] {
				return false
			}
			stack = stack[:len(stack) - 1]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
