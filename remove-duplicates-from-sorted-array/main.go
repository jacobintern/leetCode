package Q8

import "fmt"

func removeDuplicates(nums []int) int {
	var newArr []int
	for _, v := range nums {
		if !contains(newArr, v) {
			newArr = append(newArr, v)
		}
	}
	nums = newArr
	return len(nums)
}

func contains(nums []int, num int) bool {
	for _, v := range nums {
		if v == num {
			return true
		}
	}
	return false
}

func removeDuplicates2(nums []int) int {
	length := len(nums)

	j := 0
	for i := 0; i < length; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	fmt.Println(nums)
	return j
}
