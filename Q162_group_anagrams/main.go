package Q162

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _, v := range strs {
		sortedStr := sortString(v)
		t, exist := strMap[sortedStr]
		if exist {
			strMap[sortedStr] = append(t, v)
		} else {
			strMap[sortedStr] = []string{v}
		}
	}

	res := make([][]string, len(strMap))
	i := 0
	for _, arr := range strMap {
		res[i] = make([]string, len(arr))
		res[i] = arr
		i++
	}

	return res
}

func sortString(str string) string {
	var res string
	var arr []string
	for _, v := range str {
		arr = append(arr, string(v))
	}

	sort.Strings(arr)

	for _, v := range arr {
		res += v
	}

	return res
}
