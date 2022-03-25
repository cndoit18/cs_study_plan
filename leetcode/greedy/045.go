package greedy

func jump(nums []int) int {
	top := 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		ntop := i
		if top >= len(nums) {
			return sum
		}
		for j := i; j < top; j++ {
			tmp := nums[j] + j + 1
			if tmp > ntop {
				ntop = tmp
			}
		}
		sum++
		top = ntop
	}
	return sum
}
