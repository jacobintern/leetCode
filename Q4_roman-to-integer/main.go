package Q4

import (
	"strings"
)

var numList map[string]int

func init() {
	numList = make(map[string]int)
	numList["I"] = 1
	numList["V"] = 5
	numList["X"] = 10
	numList["L"] = 50
	numList["C"] = 100
	numList["D"] = 500
	numList["M"] = 1000
	numList["IV"] = -2
	numList["IX"] = -2
	numList["XL"] = -20
	numList["XC"] = -20
	numList["CD"] = -200
	numList["CM"] = -200
}

func romanToInt(s string) int {
	var r int
	for i, v := range numList {
		if c := strings.Count(s, i); c > 0 {
			r += c * v
		}
	}
	return r
}
