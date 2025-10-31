package array

//最长公共前缀

//两两比较
func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, str := range strs[1:] {
		i := 0
		for i < len(prefix) && i < len(str) && prefix[i] == str[i] {
			i++
		}

		prefix = prefix[:i]
		if prefix == "" {
			break
		}
	}

	return prefix
}

//leetcode做法
func longestCommonPrefix2(strs []string) string {
	s0 := strs[0]
	for j, c := range s0 {
		for _, v := range strs {
			if j == len(v) || v[j] != byte(c) {
				return s0[:j]
			}
		}
	}
	return s0
}
