package mathUtil

import "sort"

func Max(numbers ...int) int {
	sort.Ints(numbers)

	return numbers[len(numbers)-1]
}
