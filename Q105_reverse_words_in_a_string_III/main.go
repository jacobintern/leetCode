package Q105

import (
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	reversedWords := make([]string, len(words))

	for i, word := range words {
		reversedWords[i] = reverseString(word)
	}

	return strings.Join(reversedWords, " ")
}

func reverseString(input string) string {
	runes := []rune(input)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func reverseWords2(s string) string {
	res := ""
	for k, v := range strings.Split(s, " ") {
		tmp := []rune(v)
		l := len(v)
		for i := 0; i < l/2; i++ {
			tmp[i], tmp[l-i-1] = tmp[l-i-1], tmp[i]
		}

		if k != len(strings.Split(s, " "))-1 {
			res += string(tmp) + " "
		} else {
			res += string(tmp)
		}
	}

	return res
}
