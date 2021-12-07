package arrays

// 在一排树中，第 i 棵树产生 tree[i] 型的水果。
// 你可以从你选择的任何树开始，然后重复执行以下步骤：

// 1. 把这棵树上的水果放进你的篮子里。如果你做不到，就停下来。
// 2. 移动到当前树右侧的下一棵树。如果右边没有树，就停下来。

// 请注意，在选择一颗树后，你没有任何选择：你必须执行步骤 1，然后执行步骤 2，然后返回步骤 1，然后执行步骤 2，依此类推，直至停止。

func totalFruit(fruits []int) int {
	basket := map[int]int{}
	slow, fast, maxLen := 0, 0, 0
	for fast < len(fruits) {
		if _, ok := basket[fruits[fast]]; (!ok && len(basket) < 2) || ok {
			basket[fruits[fast]]++
			fast++
			if (fast - slow) > maxLen {
				maxLen = fast - slow
			}
			continue
		}

		key := fruits[slow]
		if nums := basket[key]; nums < 2 {
			delete(basket, key)
		} else {
			basket[key]--
		}
		slow++
	}
	return maxLen
}
