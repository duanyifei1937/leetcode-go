package leetcode

// 根据模式串 ==> next[]
func getNext(next []int, s string) {
	// 初始化
	j := 0
	next[0] = j

	// j: 字符首；
	// i:字符尾；
	for i := 1; i < len(s); i++ {
		// 首尾不相同：
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		// 首尾相同：
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}

func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := make([]int, n)
	getNext(next, needle)

	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}

	}
	return -1

}
