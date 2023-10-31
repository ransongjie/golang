package str

/**
 * https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/
 * 找出字符串中第一个匹配项的下标
 */
func strStr(haystack string, needle string) int {
	next := getNext(needle)
	m, n := len(haystack), len(needle)
	i, j := 0, -1
	for i = 0; i < m; i++ {
		// 不能完全匹配选择部分匹配
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j]
		}

		// h的第i个字符 == n的第j+1个字符
		if haystack[i] == needle[j+1] {
			j++
		}

		// 找到第1个完全匹配
		if j == n-1 {
			return i - n + 1 // 起始位置= h[i] - n长度 +1
		}
	}

	return -1
}

// 获取部分匹配表
func getNext(s string) []int {
	n := len(s)
	i, j := 0, -1
	next := make([]int, n)
	next[0] = -1 // 长度为1的子串，前后缀最大相等长度为-1

	// 长度为i+1的子串，前后缀最大相等长度
	for i = 1; i < n; i++ {
		// 前缀最后1个字符!=后缀最后1个字符
		for j >= 0 && s[j+1] != s[i] {
			// 选择部分匹配
			j = next[j]
		}

		// 前缀最后1个字符==后缀最后1个字符
		if s[j+1] == s[i] {
			j++
		}
		next[i] = j
	}

	return next
}
