package fizzbuzz

import (
	"strconv"
	"strings"
)

func Get(count int) string {
	result := ""

	for i := 1; i <= count; i++ {
		values := []string{}

		d3 := i%3 == 0
		d5 := i%5 == 0

		if !d3 && !d5 {
			values = append(values, strconv.Itoa(i))
		} else {
			if d3 {
				values = append(values, "Fizz")
			}
			if d5 {
				values = append(values, "Buzz")
			}
		}

		if i > 1 {
			result += ", "
		}
		result += strings.Join(values, " ")
	}

	return result
}
