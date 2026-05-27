type Node struct {
	Distance float64
	Val int
}
type MaxHeap []Node
func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Distance > h[j].Distance }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Node))
}
func (h *MaxHeap) Pop() any {
	x := (*h)[len(*h) - 1]
	(*h) = (*h)[:len(*h) - 1]
	return x
}

func findClosestElements(arr []int, k int, x int) []int {
	h := &MaxHeap{}
	for	_, num := range arr {
		distance := math.Abs(float64(num - x))
		if h.Len() < k {
			heap.Push(h, Node{distance, num})
		} else if distance < (*h)[0].Distance {
			heap.Pop(h)
			heap.Push(h, Node{distance, num})
		}
	}
	result := []int{}
	for i := range *h {
		result = append(result, (*h)[i].Val)
	}
	sort.Ints(result)
	return result
}
