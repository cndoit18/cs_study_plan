package dynamic

func wordBreak(s string, wordDict []string) bool {
	workMap := make(map[string]struct{})
	for _, v := range wordDict {
		workMap[v] = struct{}{}
	}

	findWord := func(s string) bool {
		_, ok := workMap[s]
		return ok
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if findWord(s[j:i]) && dp[j] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
