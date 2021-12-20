package hashtable

func isHappy(n int) bool {
	m := map[int]struct{}{
		1: {},
	}

	f := func(n int) int {
		sum := 0
		for n != 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}
		return sum
	}

	for {
		if _, ok := m[n]; ok {
			return n == 1
		}
		m[n] = struct{}{}
		n = f(n)
	}
}
