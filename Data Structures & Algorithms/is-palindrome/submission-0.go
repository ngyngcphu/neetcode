func isPalindrome(s string) bool {
	newStr := ""
	characters := []rune(s)
	for _, s := range characters {
		if (s >= 'a' && s <= 'z') || (s >= '0' && s <= '9') {
			newStr += string(s)
		} else if s >= 'A' && s <= 'Z' {
			newStr += string(s - 'A' + 'a')
		}
	}
	reversedStr := reverse(newStr)
	return reversedStr == newStr
}

func reverse(s string) string {
	newStr := []rune(s)
	n := len(s)
	for i := 0; i < n / 2; i++ {
		newStr[i], newStr[n - 1 - i] = newStr[n - 1 - i], newStr[i]
	}
	return string(newStr)
}
