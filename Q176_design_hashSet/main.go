package Q176

import (
	"container/heap"
)

type hashSet []int

func (h hashSet) Len() int           { return len(h) }
func (h hashSet) Less(i, j int) bool { return h[i] < h[j] }
func (h *hashSet) Pop() any          { return nil }
func (h *hashSet) Push(x any)        { *h = append(*h, x.(int)) }
func (h hashSet) Swap(i, j int)      {}

type MyHashSet struct {
	stack *hashSet
	idx   map[int][]int
}

func Constructor() MyHashSet {
	h := &hashSet{}
	heap.Init(h)

	return MyHashSet{stack: h, idx: make(map[int][]int)}
}

func (this *MyHashSet) Add(key int) {
	if l := len(*this.stack); l == 0 {
		this.idx[key] = []int{0}
	} else {
		if this.idx[key] == nil {
			this.idx[key] = []int{l}
		} else {
			arr := this.idx[key]
			arr = append(arr, l)
			this.idx[key] = arr
		}
	}
	this.stack.Push(key)
}

func (this *MyHashSet) Remove(key int) {
	if idxs := this.idx[key]; idxs != nil {
		for _, v := range idxs {
			heap.Remove(this.stack, v)
		}
	}

	delete(this.idx, key)
}

func (this *MyHashSet) Contains(key int) bool {
	return this.idx[key] != nil
}

// type MyHashSet struct {
// 	stack  []int
// }

// func (this *MyHashSet) Add(key int) {
// 	if slices.Index(this.stack, key) < 0 {
// 		this.stack = append(this.stack, key)
// 	}
// 	fmt.Println(this.stack)
// }

// func (this *MyHashSet) Remove(key int) {
// 	if index := slices.Index(this.stack, key); index > -1 {
// 		this.stack = append(this.stack[:index], this.stack[index+1:]...)
// 	}
// 	fmt.Println(this.stack)
// }

// func (this *MyHashSet) Contains(key int) bool {
// 	fmt.Println(this.stack)
// 	return slices.Index(this.stack, key) > -1
// }

// type MyHashSet struct {
//     List []bool
// }

// func Constructor() MyHashSet {
//     return MyHashSet{make([]bool, 1000001)}
// }

// func (this *MyHashSet) Add(key int)  {
//     this.List[key] = true
// }

// func (this *MyHashSet) Remove(key int)  {
//     this.List[key] = false
// }

// func (this *MyHashSet) Contains(key int) bool {
//     return this.List[key]
// }
