package sword

// ' ' ==> '%20'
func replaceSpace3(s string) string {
	ss := []byte(s)
	var count int
	for i := range ss {
		if ss[i] == ' ' {
			count++
		}
	}
	add_s := make([]byte, 2 * count)
	ns := append(ss, add_s...)

	right, left := len(ns)-1, len(ss)-1
	for left >= 0 {
		if ns[left] != ' ' {
			ns[right] = ns[left]
			left--
			right--
		} else {
			ns[right] = '0'
			ns[right-1] = '2'
			ns[right-2] = '%'
			right-=3
			left--
		}
	}
	return string(ns)
}