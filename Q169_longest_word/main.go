package Q169

import (
	"regexp"
	"strings"
	"unicode"
)

func longestWord(sen string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]`)
	cleanSen := re.ReplaceAllString(sen, " ")

	words := strings.Fields(cleanSen)

	longest := words[0]

	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

func LongestWord2(sen string) string {
	var curStr strings.Builder
	var maxStr string

	for _, c := range sen {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			curStr.WriteRune(c)
		} else {
			if curStr.Len() > len(maxStr) {
				maxStr = curStr.String()
			}
			curStr.Reset()
		}
	}

	if curStr.Len() > len(maxStr) {
		maxStr = curStr.String()
	}

	return maxStr
}
