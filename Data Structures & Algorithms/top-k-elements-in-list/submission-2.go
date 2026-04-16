type Pair struct {
	freq int
	num int
}
type MinHeap []Pair

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n - 1]
	*h = old[:n-1]
	return val
}

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for num, freq := range count {
		heap.Push(h, Pair{freq, num})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	var res []int
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(Pair).num)
	}
	return res
}