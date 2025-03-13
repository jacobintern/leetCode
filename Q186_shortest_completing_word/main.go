package Q186

import (
	"unicode"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	letterCount := make(map[rune]int)
	for _, ch := range licensePlate {
		if unicode.IsLetter(ch) {
			letterCount[unicode.ToLower(ch)]++
		}
	}

	shortestWord := ""
	for _, word := range words {
		if isValidWord(word, letterCount) {
			if shortestWord == "" || len(word) < len(shortestWord) {
				shortestWord = word
			}
		}
	}

	return shortestWord
}

func isValidWord(word string, letterCount map[rune]int) bool {
	wordCount := make(map[rune]int)
	for _, ch := range word {
		wordCount[ch]++
	}

	for char, count := range letterCount {
		if wordCount[char] < count {
			return false
		}
	}

	return true
}
