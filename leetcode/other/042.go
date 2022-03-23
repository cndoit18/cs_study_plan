package other

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	sum := 0

	for i := 1; i < len(height); i++ {
		if height[i] > height[i-1] {
			continue
		}
		next := i
		for j := i; j < len(height); j++ {
			if height[j] >= height[next] {
				next = j
			}
			if height[j] >= height[i-1] {
				break
			}
		}

		second := height[next]
		if height[i-1] < height[next] {
			second = height[i-1]
		}

		for begin := i; begin < next; begin++ {
			sum += second - height[begin]
		}
		i = next
	}
	return sum
}
