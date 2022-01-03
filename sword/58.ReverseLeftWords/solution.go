package sword

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

func reverseLeftWords2(s string, n int) string {
	length := len(s)
	ss := []byte(s)
	var tmp []byte

	for i := 0; i < n; i++ {
		tmp = append(tmp, s[i])
	}

	ss = append(ss[n:length], tmp...)
	return string(ss)
}

// 先翻转前n个字符，在翻转n后字符，然后全部翻转
func reverseLeftWords3(s string, n int) string {
	ss := []byte(s)
	var left, right int
	length := len(s)

	left, right = 0, n-1
	for left < right {
		ss[left], ss[right] = ss[right], ss[left]
		left++
		right--
	}

	left, right = n, length-1
	for left < right {
		ss[left], ss[right] = ss[right], ss[left]
		left++
		right--
	}

	left, right = 0, length-1
	for left < right {
		ss[left], ss[right] = ss[right], ss[left]
		left++
		right--
	}

	return string(ss)

}
