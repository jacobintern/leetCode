package Q196

import (
	"slices"
	"strings"
)

var ignoreList = []rune{'!', '?', '\'', ';', '.', ',', ' '}

func mostCommonWord(paragraph string, banned []string) string {
	words := make(map[string]int)
	var word strings.Builder
	paragraph += " "

	for _, v := range paragraph {
		if containsRune(ignoreList, v) {
			tmp := strings.ToLower(word.String())
			if !slices.Contains(banned, tmp) && tmp != "" {
				words[tmp]++
				word.Reset()
			} else {
				word.Reset()
			}
			continue
		}

		word.WriteRune(v)
	}
	max, maxStr := 0, ""
	for w, v := range words {
		if v > max {
			max = v
			maxStr = w
		}
	}
	return maxStr
}

func containsRune(list []rune, r rune) bool {
	for _, v := range list {
		if v == r {
			return true
		}
	}
	return false
}
