package fizzBuzz

import (
	"fmt"
	"strconv"
)

func Execute() {
	output := ""

	for n := 1; n <= 100; n++ {
		if n%3 == 0 {
			output += "Fizz"
		}

		if n%5 == 0 {
			output += "Buzz"
		}

		if len(output) == 0 {
			output = strconv.Itoa(n)
		}

		fmt.Println(output)

		output = ""
	}
}
