package arrays

func minSubArrayLen(target int, nums []int) int {
	minLen := 0
	if len(nums) == 0 {
		return 0
	}
	slow, fast, sum := 0, 0, nums[0]
	for fast < len(nums) {
		if sum >= target && (fast-slow < minLen || minLen == 0) {
			minLen = fast - slow + 1
		}
		if sum < target {
			fast++
			if !(fast < len(nums)) {
				break
			}
			sum += nums[fast]
		} else {
			sum -= nums[slow]
			slow++
		}

	}

	return minLen
}
