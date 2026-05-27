type Point struct {
	Distance int
	Index int
}
type MaxHeap []Point
func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Distance > h[j].Distance }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Point))
}
func (h *MaxHeap) Pop() any {
	x := (*h)[len(*h) - 1]
	(*h) = (*h)[:len(*h) - 1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &MaxHeap{}
	for i, point := range points {
		x, y := point[0], point[1]
		distance := x * x + y * y
		if h.Len() < k {
			heap.Push(h, Point{distance, i})
		} else if distance < (*h)[0].Distance {
			heap.Pop(h)
			heap.Push(h, Point{distance, i})
		}
	}
	result := [][]int{}
	for i := range (*h) {
		result = append(result, points[(*h)[i].Index])
	}
	return result
}
