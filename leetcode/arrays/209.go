package arrays

// 给定一个含有 n 个正整数的数组和一个正整数 target 。
// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

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
