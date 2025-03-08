package Q174

import "container/heap"

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(v any)        { *h = append(*h, v.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	*h = old[:n-1]
	return nil
}

type KthLargest struct {
	k    int
	heap *minHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := &minHeap{}
	heap.Init(h)

	kl := KthLargest{k, h}
	for _, n := range nums {
		kl.Add(n)
	}

	return kl
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.heap, val)

	if this.heap.Len() > this.k {
		heap.Pop(this.heap)
	}

	return (*this.heap)[0]
}
