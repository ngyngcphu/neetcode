func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	
	prev2, prev1 := 1, 1
	for i := 2; i <= len(s); i++ {
		cur := 0
		if s[i - 1] != '0' {
			cur = prev1
		}
		twoDigit, _ := strconv.Atoi(s[i - 2:i])
		if twoDigit >= 10 && twoDigit <= 26 {
			cur += prev2
		}
		prev2, prev1 = prev1, cur
	}
	return prev1
}
