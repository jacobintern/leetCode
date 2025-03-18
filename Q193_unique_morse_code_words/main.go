package Q193

var tMap map[rune]string = map[rune]string{
	'a': ".-",
	'b': "-...",
	'c': "-.-.",
	'd': "-..",
	'e': ".",
	'f': "..-.",
	'g': "--.",
	'h': "....",
	'i': "..",
	'j': ".---",
	'k': "-.-",
	'l': ".-..",
	'm': "--",
	'n': "-.",
	'o': "---",
	'p': ".--.",
	'q': "--.-",
	'r': ".-.",
	's': "...",
	't': "-",
	'u': "..-",
	'v': "...-",
	'w': ".--",
	'x': "-..-",
	'y': "-.--",
	'z': "--..",
}

func uniqueMorseRepresentations(words []string) int {
	var str string
	res := make(map[string]bool)
	for _, word := range words {
		for _, v := range word {
			if r, exists := tMap[v]; exists {
				str += r
			}
		}

		res[str] = true
		str = ""
	}

	return len(res)
}
