package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {

}

func reverse(x int) int {
	var res string
	s := strings.Split(strconv.Itoa(x), "")
	for _, v := range s {
		if v != "-" {
			res = v + res
		}
	}
	if x < 0 {
		res = "-" + res
	}
	r, err := strconv.Atoi(res)
	if err != nil {
		log.Fatal(err)
	}

	if math.MinInt32 < r && r < math.MaxInt32 {
		return r
	} else {
		return 0
	}
}
