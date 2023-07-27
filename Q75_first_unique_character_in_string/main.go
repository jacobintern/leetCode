package Q75

type info struct {
	key string
	p   int
	c   int
}

func firstUniqChar(s string) int {
	arr := []info{}
	for i, v := range s {
		cur := string(v)
		index, isExist := findElement(arr, cur)
		if isExist {
			arr[index].c++
		} else {
			arr = append(arr, info{cur, i, 1})
		}
	}

	for _, v := range arr {
		if v.c == 1 {
			return v.p
		}
	}

	return -1
}

func findElement(arr []info, target string) (int, bool) {
	for i, v := range arr {
		if v.key == target {
			return i, true
		}
	}
	return -1, false
}

func firstUniqChar2(s string) int {
	var list = make([]int, 26)
	for _, v := range s {
		list[v-'a']++
	}

	for i, v := range s {
		if list[v-'a'] == 1 {
			return i
		}
	}

	return -1
}
