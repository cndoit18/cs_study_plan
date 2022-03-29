package greedy

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	totalSum := 0
	currentSum := 0
	for i := 0; i < len(gas) && i < len(cost); i++ {
		totalSum += gas[i] - cost[i]
		currentSum += gas[i] - cost[i]
		if currentSum < 0 {
			start = i + 1
			currentSum = 0
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}
