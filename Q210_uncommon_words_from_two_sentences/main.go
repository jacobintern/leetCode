package Q210

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	appear := make(map[string]int)

	s1Arr := strings.Split(s1, " ")
	s2Arr := strings.Split(s2, " ")

	for _, v := range s1Arr {
		appear[v]++
	}

	for _, v := range s2Arr {
		appear[v]++
	}

	var res []string

	for i, v := range appear {
		if v == 1 {
			res = append(res, i)
		}
	}

	return res
}
