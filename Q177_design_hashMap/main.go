package Q177

// type hashMap map[int]int

// type MyHashMap struct {
// 	mp *hashMap
// }

// func (h hashMap) Len() int           { return len(h) }
// func (h hashMap) Less(i, j int) bool { return h[i] < h[j] }
// func (h hashMap) Pop() any           { return nil }
// func (h hashMap) Push(x any)         {}
// func (h hashMap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func Constructor() MyHashMap {
// 	mp := &hashMap{}
// 	heap.Init(mp)

// 	return MyHashMap{mp: mp}
// }

// func (this *MyHashMap) Put(key int, value int) {
// 	(*this.mp)[key] = value
// }

// func (this *MyHashMap) Get(key int) int {
// 	v, exists := (*this.mp)[key]
// 	if exists {
// 		return v
// 	}

// 	return -1
// }

// func (this *MyHashMap) Remove(key int) {
// 	delete(*this.mp, key)
// }

type MyHashMap struct {
	mp map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{mp: make(map[int]int)}
}

func (this *MyHashMap) Put(key int, value int) {
	this.mp[key] = value
}

func (this *MyHashMap) Get(key int) int {
	v, exists := this.mp[key]
	if exists {
		return v
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	delete(this.mp, key)
}
