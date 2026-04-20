func minWindow(s string, t string) string {
	chTs := []rune(t)
	chSs := []rune(s)
	countT := make(map[rune]int)
	res := []int{-1, -1}
	window := make(map[rune]int)
	have := 0
	start := 0
	minLen := math.MaxInt
	
	for _, chT := range chTs {
		countT[chT]++
	}

	for end, chS := range chSs {
		window[chS]++
		if countT[chS] > 0 && window[chS] == countT[chS] {
			have++
		}
		for have == len(countT) {
			if minLen > end - start + 1 {
				minLen = end - start + 1
				res = []int{start, end}
			}
			window[chSs[start]]--
			if countT[chSs[start]] > 0 && window[chSs[start]] < countT[chSs[start]] {
				have--
			}
			start++
		}
	}
	if minLen == math.MaxInt {
		return ""
	}
	return s[res[0]:res[1] + 1]
}
