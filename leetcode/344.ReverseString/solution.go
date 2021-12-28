package leetcode

func reverseString(s []byte) {
	n := len(s)
	for i := 0; i < len(s)/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}
