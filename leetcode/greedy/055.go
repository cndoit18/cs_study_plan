package greedy

func canJump(nums []int) bool {
	cover := 0
	for i, v := range nums {
		t := i + v
		if cover < i {
			break
		}
		if cover < t {
			cover = t
		}
	}
	return cover >= len(nums)-1
}
