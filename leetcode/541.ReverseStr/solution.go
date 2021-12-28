package leetcode

func reverseStr(s string, k int) string {
	ss := []byte(s)
	n := len(s)
	for i := 0; i < n; i += 2 * k {
		if i+k <= n {
			reverse(ss[i : i+k])
		} else {
			reverse(ss[i:n])
		}
	}
	return string(ss)
}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
