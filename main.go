package main

import (
	"fmt"
	"github.com/aramonc/go-exercises/fizzBuzz"
	"github.com/aramonc/go-exercises/mathUtil"
	"github.com/aramonc/go-exercises/multipleSums"
)

func main() {

	// Run FizzBuzz exercise
	fizzBuzz.Execute()

	// Sum the multiples of 3 or 5 below 1000
	multipleSums.Execute()

	// Print the divided integer & whether it was event
	fmt.Println(mathUtil.Half(1))
	fmt.Println(mathUtil.Half(2))
	fmt.Println(mathUtil.Half(3))
	fmt.Println(mathUtil.Half(4))
}
