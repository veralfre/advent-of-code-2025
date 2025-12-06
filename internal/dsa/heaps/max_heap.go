package heaps

type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Note: > for max heap
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }



func NewIntMaxHeap() *IntMaxHeap {
	h := &IntMaxHeap{}
	return h
}

func (h *IntMaxHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}