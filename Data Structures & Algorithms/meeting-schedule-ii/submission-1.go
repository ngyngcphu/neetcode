/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

type MinHeap []int
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() any {
	old := *h
	x := old[len(old) - 1]
	*h = old[:len(old) - 1]
	return x
}

func minMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}
	
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, intervals[0].end)

	for i := 1; i < len(intervals); i++ {
		if (*h)[0] <= intervals[i].start {
			heap.Pop(h)
		}
		heap.Push(h, intervals[i].end)
	}
	return h.Len()
}
