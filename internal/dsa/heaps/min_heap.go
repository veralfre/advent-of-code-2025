package heaps

type IntMinHeap []int

func NewIntMinHeap() *IntMinHeap {
	h := &IntMinHeap{}
	return h
}

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] } // Note: < for min heap
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }


func (h *IntMinHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}