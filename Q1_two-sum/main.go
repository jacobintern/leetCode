package Q1

// solution1
func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// solution2
func twoSum2(nums []int, target int) []int {
	var r []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				r = append(r, i, j)
			}
		}
	}
	return r
}

// solution3
var res []int

func twoSum3(nums []int, target int) []int {
	checkSum(nums, target, nums[0], 0)
	return res
}

func checkSum(nums []int, target int, cur_num int, local int) {
	if len(res) > 0 {
		return
	}

	for i, v := range nums {
		if i != local && i > local && target == v+cur_num {
			res = append(res, local, i)
		}
	}

	checkSum(nums, target, nums[local+1], local+1)
}
