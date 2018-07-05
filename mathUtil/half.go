package mathUtil

func Half(n int) (int, bool) {
	return n / 2, n%2 == 0
}
