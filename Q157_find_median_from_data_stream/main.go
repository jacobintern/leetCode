package Q157

import "sort"

type MedianFinder struct {
	left  []int
	right []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		left:  []int{},
		right: []int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.left) == 0 || num <= this.left[0] {
		this.left = append(this.left, num)                // 插入 left
		sort.Sort(sort.Reverse(sort.IntSlice(this.left))) // 使 left 仍為降序
	} else {
		this.right = append(this.right, num) // 插入 right
		sort.Ints(this.right)                // 使 right 仍為升序
	}

	// 保持平衡（left 大小 >= right）
	if len(this.left) > len(this.right)+1 {
		this.right = append(this.right, this.left[0])
		this.left = this.left[1:] // 移除 left[0]
		sort.Ints(this.right)     // 重新排序 right
	} else if len(this.left) < len(this.right) {
		this.left = append(this.left, this.right[0])
		this.right = this.right[1:]                       // 移除 right[0]
		sort.Sort(sort.Reverse(sort.IntSlice(this.left))) // 重新排序 left
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.left) > len(this.right) {
		return float64(this.left[0]) // left 比 right 多 1 個
	}

	return float64(this.left[0]+this.right[0]) / 2.0
}
