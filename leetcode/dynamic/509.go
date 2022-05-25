package dynamic

func fib(n int) int {
	if n < 2 {
		return n
	}

	a, b, c := 0, 1, 0

	for i := 2; i <= n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}
