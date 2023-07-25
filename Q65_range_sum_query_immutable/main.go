package Q65

// type NumArray struct {
// 	nums []int
// }

// func Constructor(nums []int) NumArray {
// 	return NumArray{nums: nums}
// }

// func (this *NumArray) SumRange(left int, right int) int {
// 	res := 0
// 	for _, v := range this.nums[left : right+1] {
// 		res += v
// 	}
// 	return res
// }

type NumArray struct {
	m map[int]int
}

func Constructor(nums []int) NumArray {
	res := NumArray{m: make(map[int]int)}

	tmp := 0
	for i, v := range nums {
		tmp += v
		res.m[i] = tmp
	}

	return res
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.m[right]
	}
	return this.m[right] - this.m[left-1]
}
