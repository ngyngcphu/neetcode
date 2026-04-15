func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}

	freqhm := make([][2]int, 0, len(count))
	for k, v := range count {
		freqhm = append(freqhm, [2]int{k, v})
	}
	sort.Slice(freqhm, func(i, j int) bool {
		return freqhm[i][1] > freqhm[j][1]
	})

	var res []int
	for i := 0; i < k; i++ {
		res = append(res, freqhm[i][0])
	}
	return res
}
