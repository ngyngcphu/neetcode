func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s) + 1)
	dp[0], dp[1] = 1, 1
	for i:= 2; i <= len(s); i++ {
		oneDigit := int(s[i - 1] - '0')
		if oneDigit >= 1 && oneDigit <= 9 {
			dp[i] += dp[i - 1]
		}
		twoDigit := 10 * int(s[i - 2] - '0') + oneDigit
		if twoDigit >= 10 && twoDigit <= 26 {
			dp[i] += dp[i - 2]
		}
	}
	return dp[len(s)]
}
