package Q16

import (
	"strings"
)

func addBinary(a string, b string) string {
	if a == "0" && b == "0" {
		return "0"
	}
	var res string
	aLen := len(a)
	bLen := len(b)
	if aLen > bLen {
		addlen := aLen - bLen
		for addlen > 0 {
			b = "0" + b
			addlen--
		}
	} else {
		addlen := bLen - aLen
		for addlen > 0 {
			a = "0" + a
			addlen--
		}
	}
	aArr := strings.Split(a, "")
	bArr := strings.Split(b, "")
	i := len(aArr) - 1
	plus := "0"
	for i >= 0 {
		if i != 0 {
			if plus == "0" && aArr[i] == "0" && bArr[i] == "0" {
				res = "0" + res
				plus = "0"
			} else if (plus == "0" && aArr[i] == "0" && bArr[i] == "1") ||
				(plus == "0" && aArr[i] == "1" && bArr[i] == "0") ||
				(plus == "1" && aArr[i] == "0" && bArr[i] == "0") {
				res = "1" + res
				plus = "0"
			} else if (plus == "0" && aArr[i] == "1" && bArr[i] == "1") ||
				(plus == "1" && aArr[i] == "0" && bArr[i] == "1") ||
				(plus == "1" && aArr[i] == "1" && bArr[i] == "0") {
				res = "0" + res
				plus = "1"
			} else if plus == "1" && aArr[i] == "1" && bArr[i] == "1" {
				res = "1" + res
				plus = "1"
			}
		} else {
			if (plus == "0" && aArr[i] == "0" && bArr[i] == "1") ||
				(plus == "0" && aArr[i] == "1" && bArr[i] == "0") ||
				(plus == "1" && aArr[i] == "0" && bArr[i] == "0") {
				res = "1" + res
			} else if (plus == "0" && aArr[i] == "1" && bArr[i] == "1") ||
				(plus == "1" && aArr[i] == "0" && bArr[i] == "1") ||
				(plus == "1" && aArr[i] == "1" && bArr[i] == "0") {
				res = "10" + res
			} else if plus == "1" && aArr[i] == "1" && bArr[i] == "1" {
				res = "11" + res
			}
		}
		i--
	}
	return res
}
