package Q81

import "strconv"

// when 3 then Fizz
// when 5 then Buzz
// when 15 then FizzBuzz
func fizzBuzz(n int) []string {
	res := []string{}
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}
