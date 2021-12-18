package hashtable

import "sort"

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for i := range strs {
		str := []byte(strs[i])
		sort.Slice(str, func(i, j int) bool {
			return str[i] < str[j]
		})
		mp[string(str)] = append(mp[string(str)], strs[i])
	}
	result := [][]string{}
	for _, v := range mp {
		result = append(result, v)
	}
	return result
}
