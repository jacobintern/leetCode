package Q198

import "strings"

var goatStr []rune = []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

func toGoatLatin(sentence string) string {
	arr := strings.Split(sentence, " ")
	var res []string

	for i, word := range arr {
		if isHeadEqualGoat(word) {
			res = append(res, word+"ma"+strings.Repeat("a", i+1))
		} else {
			res = append(res, word[1:]+string(word[0])+"ma"+strings.Repeat("a", i+1))
		}
	}

	return strings.Join(res, " ")
}

func isHeadEqualGoat(str string) bool {
	target := str[0]
	for _, v := range goatStr {
		if target == byte(v) {
			return true
		}
	}
	return false
}
