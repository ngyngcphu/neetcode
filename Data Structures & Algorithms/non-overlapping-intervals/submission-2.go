func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 1
	end := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if end[1] <= intervals[i][0] {
			count++
			end = intervals[i]
		}
	}
	return len(intervals) - count
}
