package Q45

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {

	if len(nums) == 0 {
		return []string{}
	}

	res := []string{}
	// start
	preNum, rangeS := nums[0], nums[0]

	// loop check
	for _, v := range nums {
		if v == preNum {
			continue
		}

		if v-preNum != 1 {
			res = append(res, insert(preNum, rangeS))
			rangeS = v
		}

		preNum = v
	}

	return append(res, insert(preNum, rangeS))
}

func insert(v, s int) string {
	if v == s {
		return fmt.Sprint(v)
	}

	return fmt.Sprint(s, "->", v)
}

func summaryRanges2(nums []int) []string {

	if len(nums) == 0 {
		return []string{}
	}

	res := []string{}
	// start
	curNum := nums[0]
	res = append(res, strconv.Itoa(nums[0]))

	// loop check
	for i := 0; i < len(nums); i++ {
		if curNum != nums[i] && res[len(res)-1] != strconv.Itoa(nums[i-1]) {
			res[len(res)-1] += "->" + strconv.Itoa(nums[i-1])
			res = append(res, strconv.Itoa(nums[i]))
			curNum = nums[i]
		} else if curNum != nums[i] && res[len(res)-1] == strconv.Itoa(nums[i-1]) {
			res = append(res, strconv.Itoa(nums[i]))
			curNum = nums[i]
		} else if i == len(nums)-1 && res[len(res)-1] != strconv.Itoa(nums[i]) {
			res[len(res)-1] += "->" + strconv.Itoa(nums[i])
		}
		curNum += 1
	}
	return res
}
