package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for i := 1; i <= 20; i++ {
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
			fmt.Print(", ")
		}
		fmt.Print(strings.Join(values, " "))
	}
}
