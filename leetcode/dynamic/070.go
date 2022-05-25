package dynamic

func climbStairs(n int) int {
	if n < 4 {
		return n
	}
	a, b, c := 2, 3, 5
	for i := 4; i <= n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}
