package Q157_1

import "container/heap"

// 最小堆（存較大的一半數字）
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] } // 小的優先
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 最大堆（存較小的一半數字）
type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] } // 反向比較
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MedianFinder 結構
type MedianFinder struct {
	left  *MaxHeap // 較小的數
	right *MinHeap // 較大的數
}

// 建構子
func Constructor() MedianFinder {
	left := &MaxHeap{}
	right := &MinHeap{}
	heap.Init(left)
	heap.Init(right)
	return MedianFinder{left, right}
}

// 插入數字
func (mf *MedianFinder) AddNum(num int) {
	if mf.left.Len() == 0 || num <= (*mf.left)[0] {
		heap.Push(mf.left, num) // 插入左側（較小的一半）
	} else {
		heap.Push(mf.right, num) // 插入右側（較大的一半）
	}

	// 平衡兩個 heap
	if mf.left.Len() > mf.right.Len()+1 {
		heap.Push(mf.right, heap.Pop(mf.left)) // 移動元素到右側
	} else if mf.left.Len() < mf.right.Len() {
		heap.Push(mf.left, heap.Pop(mf.right)) // 移動元素到左側
	}
}

// 找中位數
func (mf *MedianFinder) FindMedian() float64 {
	if mf.left.Len() > mf.right.Len() {
		return float64((*mf.left)[0])
	}
	return float64((*mf.left)[0]+(*mf.right)[0]) / 2.0
}
