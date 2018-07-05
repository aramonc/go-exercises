package multipleSums

import "fmt"

func Execute() {
	var sum = 0

	for n := 1; n < 1000; n++ {
		switch true {
		case n%3 == 0:
			sum += n
			break
		case n%5 == 0:
			sum += n
			break
		}
	}

	fmt.Println("The sum of the multiples of 3 or 5 less than 1000 is ", sum)
}
